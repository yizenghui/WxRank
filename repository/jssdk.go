package repository

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
)

//GetSign 获取排名
func GetSign(url string) {

	// npm install weixin-js-sdk
	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)

	jsClt := jssdk.NewDefaultTicketServer(clt)

	jsStr, _ := jsClt.Ticket()

	signature := jssdk.WXConfigSign(jsStr, "ADSDXSXSAXSAX", "1507867736", "wb.readfollow.com")
	fmt.Println(signature)

	panic(signature)
}
