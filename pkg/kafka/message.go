package kafka

// ThirdPaymentUpdatePayStatusNotifyMessage Third payment update pay status notify kafka
type ThirdPaymentUpdatePayStatusNotifyMessage struct {
	PayStatus int64  `json:"payStatus"`
	Sn        string `json:"Sn"`
}
