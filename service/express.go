package service

import (
	"tp-api-go/config"
	. "tp-api-go/core"
	. "tp-api-go/core/dto"
	"tp-api-go/core/helper"

	"encoding/xml"

	"time"
)

type ExpressOrderService struct{}

func (f *ExpressOrderService) CreateExpressOrder(dto *ExpressInfoDto) (response bool, apiError *APIError) {
	//1.build bizData
	//2.sign
	//3.build request params
	//4.request
	//5.response

	currentDateTime := time.Now()
	bizData := new(BaiShiRequestDto)

	//1.common params
	bizData.ShippingOrderNo = dto.ShippingOrderNo
	bizData.ProjectCode = config.Config.BaiShi.ProjectCode
	bizData.CustomerCode = config.Config.BaiShi.CustomerCode
	bizData.TransportMode = config.Config.BaiShi.TransportMode
	bizData.OrderCode = dto.CustomerOrderCode

	//2.sernder
	bizData.Sender.Name = dto.SenderName
	bizData.Sender.Province = dto.SenderProvince
	bizData.Sender.City = dto.SenderCity
	bizData.Sender.District = dto.SenderDistrict
	bizData.Sender.Address = dto.SenderAddress

	bizData.Sender.ContactName = dto.SenderContactName
	bizData.Sender.Phone = dto.SenderMobile
	bizData.Sender.EarlyPickupTime = currentDateTime.Format("2006-01-02 15:04")
	bizData.Sender.LatePickupTime = currentDateTime.Add(time.Hour * 24).Format("2006-01-02 15:04")

	//3.receiver
	bizData.Receiver.Name = dto.Name
	bizData.Receiver.Province = dto.Province
	bizData.Receiver.City = dto.City
	bizData.Receiver.District = dto.District
	bizData.Receiver.Address = dto.ShippingAddress

	bizData.Receiver.ContactName = dto.Name
	bizData.Receiver.Phone = dto.MobileNumber
	bizData.Receiver.EarlyDeliveryTime = currentDateTime.Add(time.Hour * 24 * 2).Format("2006-01-02 15:04")
	bizData.Receiver.LateDeliveryTime = currentDateTime.Add(time.Hour * 24 * 3).Format("2006-01-02 15:04")

	//4.items
	bizData.Items = []*BaiShiItemDtoNew{&BaiShiItemDtoNew{Item: BaiShiItemDto{
		ItemCode:       dto.ItemCode,
		ItemName:       dto.ItemName,
		PackageCount:   dto.ItemCount,
		PackageUomCode: "件"}}}

	//5.additional
	if len(dto.Value) != 0 {
		bizData.Additionals = []*BaiShiAdditionalDtoNew{&BaiShiAdditionalDtoNew{Additional: BaiShiAdditionalDto{
			AddType: "IDENTIFICATION_CODE",
			Value:   dto.Value}}}
	}

	var bizDataXml string
	var sign string
	if output, err := xml.MarshalIndent(bizData, "", " "); err != nil {
		response = false
		apiError = helper.NewApiError(10001, err.Error())
		return
	} else {
		bizDataXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		` + string(output)
		apiError = nil

		if sign, err = helper.BaiShiSign(bizDataXml, config.Config.BaiShi.PartnerKey); err != nil {
			response = false
			apiError = helper.NewApiError(10001, err.Error())
			return
		}
	}

	reqMap := make(map[string]string, 0)
	reqMap["serviceType"] = config.Config.BaiShi.ServiceType
	reqMap["partnerID"] = config.Config.BaiShi.PartnerId
	reqMap["bizData"] = bizDataXml
	reqMap["sign"] = sign
	var result string
	if result, apiError = helper.GetUrlRespHtml(config.Config.BaiShi.OrderExpressUrl, reqMap); apiError != nil {
		response = false
		return
	}
	var resp BaiShiResponseDto
	if e := xml.Unmarshal([]byte(result), &resp); e != nil {
		response = false
		apiError = helper.NewApiError(10001, e.Error())
		return
	}

	if resp.Result == false {
		response = false
		apiError = helper.NewApiError(10001, resp.ErrorCode+resp.ErrorDescription)
		return
	} else {
		response = resp.Result
		apiError = nil
		return
	}
	return
}
