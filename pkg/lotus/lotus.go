package lotus

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"
	"hechuangfil-admin/utils"
)

var FullAPI v0api.FullNode

// 转账
func Send(ctx context.Context, from, addrTo string, amount float64) (*types.SignedMessage, error) {
	value := utils.Float64ToString(amount)
	fr, err := address.NewFromString(from)
	if err != nil {
		return nil, err
	}
	add, err := address.NewFromString(addrTo)
	if err != nil {
		return nil, err
	}
	val, err := types.ParseFIL(value)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}
	sm, err := FullAPI.MpoolPushMessage(ctx, &types.Message{
		To:    add,
		From:  fr,
		Value: types.BigInt(val),
	}, nil)
	if err != nil {
		return nil, err
	}
	return sm, err
}

// 获取钱包余额
func Balance(ctx context.Context, addr string) (float64, error) {
	add, err := address.NewFromString(addr)
	if err != nil {
		return 0, err
	}
	bal, err := FullAPI.WalletBalance(ctx, add)
	if err != nil {
		return 0, err
	}
	bls := utils.NanoOrAttoToFIL(bal.String(), utils.AttoFIL)
	return bls, err
}
