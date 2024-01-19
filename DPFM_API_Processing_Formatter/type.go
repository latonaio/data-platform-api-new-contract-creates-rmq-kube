package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Contract                         int      `json:"Contract"`
	ContractDate                     *string  `json:"ContractDate"`
	ContractStatus                   *string  `json:"ContractStatus"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	ContractValidityStartDate        *string  `json:"ContractValidityStartDate"`
	ContractValidityEndDate          *string  `json:"ContractValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount"`
	TransactionCurrency              *string  `json:"TransactionCurrency"`
	PricingDate                      *string  `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
}

type ItemUpdates struct {
	Contract                         		 int      `json:"Contract"`
	ContractItem                             int      `json:"ContractItem"`
	ContractStatus                           *string  `json:"ContractStatus"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	ContractItemText                         *string  `json:"ContractItemText"`
	ContractItemTextByBuyer                  *string  `json:"ContractItemTextByBuyer"`
	ContractItemTextBySeller                 *string  `json:"ContractItemTextBySeller"`
	Product                                  *string  `json:"Product"`
	ProductionVersion				 		 *int     `json:"ProductionVersion"`
	ProductionVersionItem			 		 *int     `json:"ProductionVersionItem"`
	BillOfMaterial                           *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                       *int     `json:"BillOfMaterialItem"`
	ServicesRenderingDate                    *string  `json:"ServicesRenderingDate"`
	ContractQuantityInBaseUnit               *float32 `json:"ContractQuantityInBaseUnit"`
	ContractQuantityInDeliveryUnit           *float32 `json:"ContractQuantityInDeliveryUnit"`
	QuantityPerPackage                       *float32 `json:"QuantityPerPackage"`
	TaxAmount                                *float32 `json:"TaxAmount"`
	GrossAmount                              *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                      *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	InspectionPlan                           *int     `json:"InspectionPlan"`
	InspectionPlant                          *string  `json:"InspectionPlant"`
	InspectionOrder                          *int     `json:"InspectionOrder"`
	TransactionTaxClassification             *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry    *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry  *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                 *string  `json:"DefinedTaxClassification"`
	PaymentTerms                             *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                   *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                           *string  `json:"PaymentDueDate"`
	NetPaymentDays                           *int     `json:"NetPaymentDays"`
	PaymentMethod                            *string  `json:"PaymentMethod"`
	Project                                  *int     `json:"Project"`
	WBSElement                               *int     `json:"WBSElement"`
	Equipment                                *int     `json:"Equipment"`
	ItemBlockStatus                          *bool    `json:"ItemBlockStatus"`
	ExternalReferenceDocument                *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem            *string  `json:"ExternalReferenceDocumentItem"`
}

type ItemPricingElementUpdates struct {
	Contract                  int      `json:"Contract"`
	ContractItem              int      `json:"ContractItem"`
	SupplyChainRelationshipID *int     `json:"SupplyChainRelationshipID"`
	Buyer                     *int     `json:"Buyer"`
	Seller                    *int     `json:"Seller"`
	PricingProcedureCounter   int      `json:"PricingProcedureCounter"`
	ConditionRateValue        *float32 `json:"ConditionRateValue"`
	ConditionAmount           *float32 `json:"ConditionAmount"`
}

type PartnerUpdates struct {
	Contract                int     `json:"Contract"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
	Contract    int     `json:"Contract"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
