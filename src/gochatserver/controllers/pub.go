package controllers

import (
	"github.com/astaxie/beego"
	"gochatserver/models"
	"fmt"
)

type PubController struct  {
	beego.Controller
}
//这里填上在企业号回调模式中设置的AESkey




//TODO 这里负责回调模式的验证
func (c * 	PubController) Get() {
	var signature models.Signature
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	fmt.Println(signature.Echostr)
	c.Ctx.WriteString(signature.Echostr)

}




func (c * 	PubController) Post(){

}