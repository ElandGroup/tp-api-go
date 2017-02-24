package dto

type ExpressInfoDto struct {
	City              string `city`
	CustomerOrderCode string `customerOrderCode`
	District          string `district`
	ItemCode          string `itemCode`
	ItemCount         string `itemCount`

	ItemName     string `itemName`
	MobileNumber string `mobileNumber`
	Name         string `name`
	Note         string `note`
	PhoneNumber  string `phoneNumber`
	Province     string `province`

	SenderAddress     string `senderAddress`
	SenderCity        string `senderCity`
	SenderContactName string `senderContactName`
	SenderDistrict    string `senderDistrict`
	SenderMobile      string `senderMobile`

	SenderName     string `senderName`
	SenderPhone    string `senderPhone`
	SenderProvince string `senderProvince`

	ShippingAddress string `shippingAddress`
	ShippingOrderNo string `shippingOrderNo`
	Value           string `value`
}
