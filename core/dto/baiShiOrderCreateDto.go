package dto

import "encoding/xml"

type OrderExpressDto struct {
	BizData     string `json:"bizData"`
	ServiceType string `json:"serviceType"`
	PartnerID   string `json:"partnerID"`
	Sign        string `json:"sign"`
}

type BaiShiRequestDto struct {
	XMLName       xml.Name `xml:"request"`
	CustomerCode  string   `xml:"customerCode"`
	CustomerName  string   `xml:"customerName"`
	ProjectCode   string   `xml:"projectCode"`
	TransportMode string   `xml:"transportMode"`
	VehicleModel  string   `xml:"vehicleModel"`

	LogisticsCode    string `xml:"logisticsCode"`
	Remark           string `xml:"remark"`
	OrderDescription string `xml:"orderDescription"`
	ShippingOrderNo  string `xml:"shippingOrderNo"`
	OrderCode        string `xml:"orderCode"`

	GoodsValue  string `xml:"goodsValue"`
	CheapAmount string `xml:"cheapAmount"`
	NumberType  string `xml:"type"`
	Value       string `xml:"value"`

	Sender      BaiShiSenderDto           `xml:"sender"`
	Receiver    BaiShiReceiverDto         `xml:"receiver"`
	Items       []*BaiShiItemDtoNew       `xml:"items"`
	Additionals []*BaiShiAdditionalDtoNew `xml:"additionals"`
}

type BaiShiSenderDto struct {
	XMLName  xml.Name `xml:"sender"`
	Name     string   `xml:"name"`
	Province string   `xml:"province"`
	City     string   `xml:"city"`
	District string   `xml:"district"`
	Address  string   `xml:"address"`

	ContactName     string `xml:"contactName"`
	Phone           string `xml:"phone"`
	EarlyPickupTime string `xml:"earlyPickupTime"`
	LatePickupTime  string `xml:"latePickupTime"`
}

type BaiShiReceiverDto struct {
	XMLName  xml.Name `xml:"receiver"`
	Name     string   `xml:"name"`
	Province string   `xml:"province"`
	City     string   `xml:"city"`
	District string   `xml:"district"`
	Address  string   `xml:"address"`

	ContactName       string `xml:"contactName"`
	Phone             string `xml:"phone"`
	EarlyDeliveryTime string `xml:"earlyDeliveryTime"`
	LateDeliveryTime  string `xml:"lateDeliveryTime"`
}

type BaiShiItemDtoNew struct {
	Item BaiShiItemDto `xml:"item"`
}

type BaiShiItemDto struct {
	XMLName      xml.Name `xml:"item"`
	LineNo       string   `xml:"lineNo"`
	ItemCode     string   `xml:"itemCode"`
	ItemName     string   `xml:"itemName"`
	ItemCount    string   `xml:"itemCount"`
	PackageCount string   `xml:"packageCount"`

	PackageUomCode string `xml:"packageUomCode"`
	Weight         string `xml:"weight"`
	Volume         string `xml:"volume"`
	UnitPrice      string `xml:"unitPrice"`
	DeclaredValue  string `xml:"declaredValue"`

	VolumeWeight string `xml:"volumeWeight"`
	Remark       string `xml:"remark"`
}

type BaiShiAdditionalDtoNew struct {
	Additional BaiShiAdditionalDto `xml:"additional"`
}
type BaiShiAdditionalDto struct {
	AddType string `xml:"type"`
	Value   string `xml:"value"`
}

type BaiShiResponseDto struct {
	XMLName          xml.Name `xml:"response"`
	Result           bool     `xml:"result"`
	Note             string   `xml:"note"`
	ErrorCode        string   `xml:"errorCode"`
	ErrorDescription string   `xml:"errorDescription"`
}
