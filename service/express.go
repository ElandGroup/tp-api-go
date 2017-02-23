package service

import (
	"fmt"
	"tp-api-go/config"
	. "tp-api-go/core"
	. "tp-api-go/core/dto"
	"tp-api-go/core/helper"

	"encoding/xml"

	"github.com/ddliu/go-httpclient"
)

type ExpressOrderService struct{}

func (ExpressOrderService) CreateExpressOrder(expressInfoDto *ExpressInfoDto) (response string, apiError *APIError) {
	//1.build bizData
	//2.sign
	//3.build request params
	//4.request
	//5.response

	//request := new(OrderExpressDto)

	bizData := new(BaiShiRequestDto)
	if output, err := xml.MarshalIndent(bizData, " ", " "); err != nil {
		response = ""
		apiError = helper.NewApiError(10001, err.Error())
		return
	} else {
		response = string(output)
		fmt.Println(response)
		apiError = nil
	}

	httpclient.Defaults(httpclient.Map{
		"Content-Type": "application/x-www-form-urlencoded;charsets=utf-8"})
	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = config.Config.BaiShi.ServiceType
	reqMap["partnerID"] = config.Config.BaiShi.PartnerId
	reqMap["bizData"] = ""
	reqMap["sign"] = ""
	httpclient.Post(config.Config.BaiShi.OrderExpressUrl, reqMap)
	return
}
