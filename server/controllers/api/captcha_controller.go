package api

import (
	"fmt"
	"bbs-go/pkg/bbsurls"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/web"
	"github.com/sirupsen/logrus"

	"bbs-go/pkg/config"
)

type CaptchaController struct {
	Ctx iris.Context
}

func (c *CaptchaController) GetRequest() *web.JsonResult {
	captchaId := c.Ctx.FormValue("captchaId")
	fmt.Println("captchaId:",captchaId)
	if strs.IsNotBlank(captchaId) { // reload
		fmt.Println("captchaId:",captchaId)
		if !captcha.Reload(captchaId) {
			fmt.Println("Reload:",captchaId)
			// reload 失败，重新加载验证码
			captchaId = captcha.NewLen(4)
		}
	} else {
		fmt.Println("captchaId2:",captchaId)
		captchaId = captcha.NewLen(4)
	}
	captchaUrl := bbsurls.AbsUrl(":" + config.Instance.Port + "/api/captcha/show?captchaId=" + captchaId + "&r=" + strs.UUID())
	fmt.Println("captchaUrl:",captchaUrl)
	return web.NewEmptyRspBuilder().
		Put("captchaId", captchaId).
		Put("captchaUrl", captchaUrl).
		JsonResult()
}

func (c *CaptchaController) GetShow() {
	fmt.Println("0")
	captchaId := c.Ctx.URLParam("captchaId")

	if captchaId == "" {
		fmt.Println("1")
		c.Ctx.StatusCode(404)
		return
	}
	
	if !captcha.Reload(captchaId) {
		fmt.Println("2")
		c.Ctx.StatusCode(404)
		return
	}
	
	fmt.Println("3")
	c.Ctx.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(c.Ctx.ResponseWriter(), captchaId, captcha.StdWidth, captcha.StdHeight); err != nil {
		fmt.Println("4")
		logrus.Error(err)
	}
}

func (c *CaptchaController) GetVerify() *web.JsonResult {
	captchaId := c.Ctx.URLParam("captchaId")
	captchaCode := c.Ctx.URLParam("captchaCode")
	success := captcha.VerifyString(captchaId, captchaCode)
	return web.NewEmptyRspBuilder().Put("success", success).JsonResult()
}
