package auth

import (
	"context"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"hechuangfil-admin/config"
	"hechuangfil-admin/handler"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg/lotus"
	"hechuangfil-admin/redisclient"
	"hechuangfil-admin/utils"
	"log"
	"time"
)

func GetBlock() {
	go func() {
		var ctx context.Context
		rheight, err := redisclient.GetLotusHeight()
		if err != nil {
			rheight = 0
			//log.Fatal("无法获取高度：",err.Error())
			//continue
		}
		if rheight < config.ApplicationConfig.SetHeight {
			rheight = config.ApplicationConfig.SetHeight
		}
		for {
			stime := time.Second * 30
			chain, err := lotus.FullAPI.ChainHead(ctx)
			if err != nil {
				log.Fatal("无法获取链上信息：", err.Error())
				//time.Sleep(stime)
				//continue
			}
			if int64(chain.Height())-rheight > 2 {
				stime = time.Second
			}
			rheight++
			if rheight >= int64(chain.Height()) {
				rheight = int64(chain.Height()) - 2
			}
			FilAddress := redisclient.GetFilAddress()
			go func(height abi.ChainEpoch) {
				log.Println("获取高度：", height)
				s1, err := lotus.FullAPI.ChainGetTipSetByHeight(ctx, height, types.TipSetKey{})
				if err != nil {
					log.Println("无法获取链上信息：", err.Error(), "，height：", height)
					//time.Sleep(time.Second*30)
					return
				}
				for _, v := range s1.Cids() {
					ms, err := lotus.FullAPI.ChainGetBlockMessages(ctx, v)
					if err != nil {
						log.Println("无法获取链上信息：", err.Error())
						//time.Sleep(time.Second*30)
						return
					}
					for _, val := range ms.SecpkMessages {
						//log.Println(val.Message.From,"------",val.Message.Cid(),"------",val.Message.To,"------",val.Message.Method,"------",height)
						if val.Message.Method == handler.SEND {
							//log.Println(val.Message.From,"------",val.Cid(),"------",val.Message.To)
							RechargeRecord := redisclient.GetRechargeRecord()
							if FilAddress[val.Message.To.String()] != "" && RechargeRecord[val.Cid().String()] == "" {
								var re models.Recharge
								param := make(map[string]string)
								param["to_address"] = val.Message.To.String()
								param["height"] = height.String()
								param["amount"] = utils.NanoOrAttoToFILstr(val.Message.Value.String(), utils.AttoFIL)
								param["from_address"] = val.Message.From.String()
								param["cid"] = val.Cid().String()
								param["customer_id"] = FilAddress[val.Message.To.String()]
								if err := re.Insert(param); err != nil {
									log.Println("保存数据失败1：", err.Error(), "，CID：", val.Cid().String())
									continue
								}
								record := make(map[string]string)
								record[param["cid"]] = param["customer_id"]
								if _, err := redisclient.UpdateRechargeRecord(record); err != nil {
									log.Println("保存数据失败2：", err.Error(), "，CID：", val.Cid().String())
									continue
								}
							}

							//修改提现状态
							if models.CidWithdraw[val.Cid().String()] {
								var wi models.Withdraw
								if err = wi.UpdateByCid(val.Cid().String(), height.String()); err != nil {
									log.Println("修改提现状态失败！", err.Error(), val.Cid().String(), val.Message.Cid())
									continue
								}
								delete(models.CidWithdraw, val.Cid().String())
							}
						}
					}
					//log.Println("--------------------------------------------------------------------------")
					for _, val := range ms.BlsMessages {
						//log.Println(val.From,"------",val.Cid(),"------",val.To,"------",val.Method,"------",height)
						if val.Method == handler.SEND {
							//log.Println(val.From,"------",val.Cid(),"------",val.To)
							RechargeRecord := redisclient.GetRechargeRecord()
							if FilAddress[val.To.String()] != "" && RechargeRecord[val.Cid().String()] == "" {
								var re models.Recharge
								param := make(map[string]string)
								param["to_address"] = val.To.String()
								param["height"] = height.String()
								param["amount"] = utils.NanoOrAttoToFILstr(val.Value.String(), utils.AttoFIL)
								param["from_address"] = val.From.String()
								param["cid"] = val.Cid().String()
								param["customer_id"] = FilAddress[val.To.String()]
								if err := re.Insert(param); err != nil {
									log.Println("保存数据失败1：", err.Error(), "，CID：", val.Cid().String())
									continue
								}
								record := make(map[string]string)
								record[param["cid"]] = param["customer_id"]
								if _, err := redisclient.UpdateRechargeRecord(record); err != nil {
									log.Println("保存数据失败2：", err.Error(), "，CID：", val.Cid().String())
									continue
								}
							}

							//修改提现状态
							if models.CidWithdraw[val.Cid().String()] {
								var wi models.Withdraw
								if err = wi.UpdateByCid(val.Cid().String(), height.String()); err != nil {
									log.Println("修改提现状态失败！", err.Error(), val.Cid().String(), val.Cid())
									continue
								}
								delete(models.CidWithdraw, val.Cid().String())
							}
						}
					}
				}

				//默认不会出现问题,直接保存高度
				err = redisclient.SetLotusHeight(int64(height))
				if err != nil {
					log.Println("保存高度失败：", err.Error(), ",height：", rheight)
				}
			}(abi.ChainEpoch(rheight))
			time.Sleep(stime)
		}
	}()
}

const (
	INTERVAL               = 5 * time.Minute //归集周期
	SETTLEMENT_AMOUNT      = 0.1             //归集最小余额
	RETAIN_SERVICE_CHARAGE = 0.01            //保留手续费
)

func CashSweep(ctx context.Context) {
	go func() {
		for {
			var conf models.Config
			conf, err := conf.GetConfig(models.COLLECTION_ADDRESS)
			if err != nil {
				log.Println("未设置归集钱包!")
			}
			select {
			case <-time.After(INTERVAL):
				FilAddress := redisclient.GetFilAddress()
				for k, _ := range FilAddress {
					amount, _ := lotus.Balance(ctx, k)
					if amount >= SETTLEMENT_AMOUNT {
						if sm, err := lotus.Send(ctx, k, conf.Value, amount-RETAIN_SERVICE_CHARAGE); err != nil {
							log.Println("归集失败：", k, err.Error())
						} else {
							log.Println("归集成功：", sm.Cid(), k, sm.Message.Value)
						}
					}
				}
			}
		}
	}()
}
