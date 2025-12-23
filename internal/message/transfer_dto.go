package message

type AmountInfo struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type DomesticPayoutRequest struct {
	PartnerReferenceNo     string             `json:"partnerReferenceNo"`
	Amount                 AmountInfo         `json:"amount"`
	BeneficiaryAccountName string             `json:"beneficiaryAccountName"`
	BeneficiaryAccountNo   string             `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode    string             `json:"beneficiaryBankCode"`
	BeneficiaryEmail       string             `json:"beneficiaryEmail"`
	SourceAccountNo        string             `json:"sourceAccountNo"`
	TransactionDate        string             `json:"transactionDate"`
	AdditionalInfo         DomesticPayoutInfo `json:"additionalInfo"`
}

type DomesticPayoutInfo struct {
	TransferType string `json:"transferType"`
	PurposeCode  string `json:"purposeCode"`
}

type DomesticPayoutResponse struct {
	ResponseCode         string            `json:"responseCode"`
	ResponseMessage      string            `json:"responseMessage"`
	PartnerReferenceNo   string            `json:"partnerReferenceNo"`
	ReferenceNo          string            `json:"referenceNo"`
	Amount               AmountInfo        `json:"amount"`
	BeneficiaryAccountNo string            `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode  string            `json:"beneficiaryBankCode"`
	SourceAccountNo      string            `json:"sourceAccountNo"`
	AdditionalInfo       map[string]string `json:"additionalInfo"`
}

type InhousePayoutRequest struct {
	PartnerReferenceNo   string            `json:"partnerReferenceNo"`
	BeneficiaryEmail     string            `json:"beneficiaryEmail"`
	Amount               AmountInfo        `json:"amount"`
	BeneficiaryAccountNo string            `json:"beneficiaryAccountNo"`
	Remark               string            `json:"remark"`
	SourceAccountNo      string            `json:"sourceAccountNo"`
	TransactionDate      string            `json:"transactionDate"`
	AdditionalInfo       InhousePayoutInfo `json:"additionalInfo"`
}

type InhousePayoutInfo struct {
	EconomicActivity   string `json:"economicActivity"`
	TransactionPurpose string `json:"transactionPurpose"`
}

type InhousePayoutResponse struct {
	ResponseCode         string            `json:"responseCode"`
	ResponseMessage      string            `json:"responseMessage"`
	PartnerReferenceNo   string            `json:"partnerReferenceNo"`
	ReferenceNo          string            `json:"referenceNo"`
	Amount               AmountInfo        `json:"amount"`
	BeneficiaryAccountNo string            `json:"beneficiaryAccountNo"`
	SourceAccountNo      string            `json:"sourceAccountNo"`
	TransactionDate      string            `json:"transactionDate"`
	AdditionalInfo       InhousePayoutInfo `json:"additionalInfo"`
}

type BillPayoutRequest struct {
	VirtualAccountNo    string `json:"virtualAccountNo"`
	VirtualAccountEmail string `json:"virtualAccountEmail"`
	SourceAccountNo     string `json:"sourceAccountNo"`
	PartnerReferenceNo  string `json:"partnerReferenceNo"`
	PaidAmount          string `json:"paidAmount"`
	TrxDateTime         string `json:"trxDateTime"`
}

type BillPayoutResponse struct {
	ResponseCode       string         `json:"responseCode"`
	ResponseMessage    string         `json:"responseMessage"`
	VirtualAccountData BillPayoutData `json:"virtualAccountData"`
}

type BillPayoutData struct {
	VirtualAccountNo    string                  `json:"virtualAccountNo"`
	VirtualAccountName  string                  `json:"virtualAccountName"`
	VirtualAccountEmail string                  `json:"virtualAccountEmail"`
	SourceAccountNo     string                  `json:"sourceAccountNo"`
	PartnerReferenceNo  string                  `json:"partnerReferenceNo"`
	ReferenceNo         string                  `json:"referenceNo"`
	PaidAmount          AmountInfo              `json:"paidAmount"`
	TotalAmount         AmountInfo              `json:"TotalAmount"`
	TrxDateTime         string                  `json:"trxDateTime"`
	BillDetails         []BillPayoutDetails     `json:"billDetails"`
	FreeTexts           []BillPayoutDescription `json:"freeTexts"`
	FeeAmount           AmountInfo              `json:"feeAmount"`
	ProductName         string                  `json:"productName"`
}

type BillPayoutDescription struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

type BillPayoutDetails struct {
	BillDescription BillPayoutDescription `json:"billDescription"`
	BillAmount      []AmountInfo          `json:"billAmount"`
}

type TokenResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AccessToken     string `json:"accessToken"`
	TokenType       string `json:"tokenType"`
	ExpiresIn       int16  `json:"expiresIn"`
}
