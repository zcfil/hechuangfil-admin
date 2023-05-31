package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"net/http"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// @Summary 初始化图片验证码
// @Description 获取JSON 登录UID 和Base64 图片验证码
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string	"{"code": 200,"data":"base64","id":"uuid", "msg": "success"}"
// @Success 200 {string} string	"{"code": 0, "message": "err"}"
// @Router /api/v1/getCaptcha [get]
// @Security
func GenerateCaptchaHandler(c *gin.Context) {
	var param configJsonBody
	param.Id = uuid.New().String()
	param.CaptchaType = "string"
	param.DriverString = base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := param.DriverString.ConvertFonts()

	cap := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cap.Generate()
	body := map[string]interface{}{"code": 200, "data": b64s, "id": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	c.JSON(http.StatusOK, body)
}
