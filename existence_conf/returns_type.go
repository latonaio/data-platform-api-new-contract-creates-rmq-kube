package existence_conf

type Returns struct {
	ConnectionKey                                string                                        `json:"connection_key"`
	Result                                       bool                                          `json:"result"`
	RedisKey                                     string                                        `json:"redis_key"`
	RuntimeSessionID                             string                                        `json:"runtime_session_id"`
	BusinessPartnerID                            *int                                          `json:"business_partner"`
	Filepath                                     string                                        `json:"filepath"`
	ServiceLabel                                 string                                        `json:"service_label"`
	ProductMasterReturn                          ProductMasterReturn                           `json:"ProductMaster"`
	BPGeneralReturn                              BPGeneralReturn                               `json:"BusinessPartnerGeneral"`
	PlantGeneralReturn                           PlantGeneralReturn                            `json:"PlantGeneral"`
	SupplyChainRelationshipGeneralReturn         *SupplyChainRelationshipGeneralReturn         `json:"SupplyChainRelationshipGeneral,omitempty"`
	SupplyChainRelationshipBillingRelationReturn *SupplyChainRelationshipBillingRelationReturn `json:"SupplyChainRelationshipBillingRelation,omitempty"`
	SupplyChainRelationshipPaymentRelationReturn *SupplyChainRelationshipPaymentRelationReturn `json:"SupplyChainRelationshipPaymentRelation,omitempty"`
	CurrencyReturn                               CurrencyReturn                                `json:"Currency"`
	CountryReturn                                CountryReturn                                 `json:"Country"`
	PaymentTermsReturn                           PaymentTermsReturn                            `json:"PaymentTerms"`
	PaymentMethodReturn                          PaymentMethodReturn                           `json:"PaymentMethod"`
	AddressReturn                                AddressReturn                                 `json:"Address"`
	QuantityUnitReturn                           QuantityUnitReturn                            `json:"QuantityUnit"`
	ProjectReturn                                ProjectReturn                                 `json:"Project"`
	APISchema                                    string                                        `json:"api_schema"`
	Accepter                                     []string                                      `json:"accepter"`
	Deleted                                      bool                                          `json:"deleted"`
}

type ProductMasterReturn struct {
	General struct {
		Product       string `json:"Product"`
		ExistenceConf bool   `json:"ExistenceConf"`
	} `json:"General"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type SupplyChainRelationshipGeneralReturn struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type SupplyChainRelationshipBillingRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
}

type SupplyChainRelationshipPaymentRelationReturn struct {
	SupplyChainRelationshipID        int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int `json:"Buyer"`
	Seller                           int `json:"Seller"`
	BillToParty                      int `json:"BillToParty"`
	BillFromParty                    int `json:"BillFromParty"`
	Payer                            int `json:"Payer"`
	Payee                            int `json:"Payee"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}

type PaymentTermsReturn struct {
	PaymentTerms string `json:"PaymentTerms"`
}

type PaymentMethodReturn struct {
	PaymentMethod string `json:"PaymentMethod"`
}

type AddressReturn struct {
	AddressID       int    `json:"AddressID"`
	ValidityEndDate string `json:"ValidityEndDate"`
}

type QuantityUnitReturn struct {
	QuantityUnit string `json:"QuantityUnit"`
}

type ProjectReturn struct {
	Project string `json:"Project"`
}
