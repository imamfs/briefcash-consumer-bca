package entity

import "time"

type Transaction struct {
	ID                      int64      `gorm:"column:id;primaryKey"`
	Sender                  int64      `gorm:"column:data_sender_id"`
	Recipient               int64      `gorm:"column:data_recipient_id"`
	MerchantCode            string     `gorm:"column:merchant_code"`
	PartnerReferenceNo      string     `gorm:"column:partner_reference_no"`
	BankReferenceNo         *string    `gorm:"column:bank_reference_no"`
	SystemReferenceNo       *string    `gorm:"column:system_reference_no"`
	Amount                  float64    `gorm:"column:amount"`
	Currency                string     `gorm:"column:currency"`
	Remark                  string     `gorm:"column:remark"`
	TransactionType         string     `gorm:"column:transaction_type"`
	TransactionDate         time.Time  `gorm:"column:transaction_date"`
	Status                  string     `gorm:"column:status"`
	IsReversal              bool       `gorm:"column:is_reversal"`
	CompanyCharge           float32    `gorm:"column:company_charge"`
	PartnerCharge           float32    `gorm:"column:partner_charge"`
	AdditionalPartnerCharge float32    `gorm:"column:additional_partner_charge"`
	TaxCharge               float32    `gorm:"column:tax_charge"`
	IsReconcile             bool       `gorm:"column:is_reconcile"`
	ReconcileDate           *time.Time `gorm:"column:reconcile_date"`
}

type BankPartner struct {
	BankCode     string `gorm:"column:bank_code"`
	BankName     string `gorm:"column:bank_name"`
	DomesticPath string `gorm:"column:external_transfer_url"`
	InhousePath  string `gorm:"column:internal_transfer_url"`
	BillPath     string `gorm:"column:va_transfer_url"`
	SknPath      string `gorm:"column:sknbi_transfer_url"`
	RtgsPath     string `gorm:"column:rtgs_transfer_url"`
	TokenPath    string `gorm:"column:access_token_url"`
	BasePath     string `gorm:"column:base_url"`
	ClientKey    string `gorm:"column:client_key"`
	ClientSecret string `gorm:"column:client_secret"`
	PartnerId    string `gorm:"column:partner_id"`
	ChannelId    string `gorm:"column:channel_id"`
}
