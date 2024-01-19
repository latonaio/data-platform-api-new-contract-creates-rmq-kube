package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-orders-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Contract:                  			data.Contract,
		ContractDate:                 		data.ContractDate,
		ContractStatus:              		data.ContractStatus,
		SupplyChainRelationshipBillingID:   data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID:   data.SupplyChainRelationshipPaymentID,
		BillToParty:               			data.BillToParty,
		BillFromParty:             			data.BillFromParty,
		Payer:                     			data.Payer,
		Payee:                     			data.Payee,
		ContractValidityStartDate:    		data.ContractValidityStartDate,
		ContractValidityEndDate:      		data.ContractValidityEndDate,
		InvoicePeriodStartDate:    			data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:      			data.InvoicePeriodEndDate,
		TotalNetAmount:            			data.TotalNetAmount,
		TotalTaxAmount:            			data.TotalTaxAmount,
		TotalGrossAmount:          			data.TotalGrossAmount,
		TransactionCurrency:       			data.TransactionCurrency,
		PricingDate:               			data.PricingDate,
		PriceDetnExchangeRate:     			data.PriceDetnExchangeRate,
		PaymentTerms:              			data.PaymentTerms,
		PaymentMethod:             			data.PaymentMethod,
		AccountingExchangeRate:    			data.AccountingExchangeRate,
		InvoiceDocumentDate:       			data.InvoiceDocumentDate,
		HeaderText:                			data.HeaderText,
		HeaderBlockStatus:         			data.HeaderBlockStatus,
		ExternalReferenceDocument:		  	data.ExternalReferenceDocument,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item
	return &ItemUpdates{
		Contract:                                 dataHeader.Contract,
		ContractItem:                             data.ContractItem,
		ContractStatus:              			  data.ContractStatus,
		SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
		DeliverToParty:                           data.DeliverToParty,
		DeliverFromParty:                         data.DeliverFromParty,
		DeliverToPlant:                           data.DeliverToPlant,
		DeliverFromPlant:                         data.DeliverFromPlant,
		ContractItemText:                         data.ContractItemText,
		ContractItemTextByBuyer:                  data.ContractItemTextByBuyer,
		ContractItemTextBySeller:                 data.ContractItemTextBySeller,
		Product:                                  data.Product,
		ProductionVersion:                        data.ProductionVersion,
		ProductionVersionItem:                    data.ProductionVersionItem,
		BillOfMaterial:							  data.BillOfMaterial,
		BillOfMaterialItem:						  data.BillOfMaterialItem,
		ServicesRenderingDate:					  data.ServicesRenderingDate,
		ContractQuantityInBaseUnit:               data.ContractQuantityInBaseUnit,
		ContractQuantityInDeliveryUnit:           data.ContractQuantityInDeliveryUnit,
		QuantityPerPackage:                       data.QuantityPerPackage,
		TaxAmount:                                data.TaxAmount,
		GrossAmount:                              data.GrossAmount,
		InvoiceDocumentDate:                      data.InvoiceDocumentDate,
		ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
		ProductionPlant:                          data.ProductionPlant,
		InspectionPlan:                           data.InspectionPlan,
		InspectionPlant:                          data.InspectionPlant,
		InspectionOrder:                          data.InspectionOrder,
		TransactionTaxClassification:             data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:    data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry:  data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                 data.DefinedTaxClassification,
		PaymentTerms:                             data.PaymentTerms,
		DueCalculationBaseDate:                   data.DueCalculationBaseDate,
		PaymentDueDate:                           data.PaymentDueDate,
		NetPaymentDays:                           data.NetPaymentDays,
		PaymentMethod:							  data.PaymentMethod,
		Project:                                  data.Project,
		WBSElement:                               data.WBSElement,
		Equipment:                                data.Equipment,
		ItemBlockStatus:                          data.ItemBlockStatus,
		ExternalReferenceDocument:				  data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem:	  		  data.ExternalReferenceDocumentItem,
	}
}

func ConvertToItemPricingElementUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	dataHeader := header
	dataItem := item
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		Contract:                  dataHeader.Contract,
		ContractItem:              dataItem.ContractItem,
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		PricingProcedureCounter:   data.PricingProcedureCounter,
		ConditionRateValue:        data.ConditionRateValue,
		ConditionAmount:           data.ConditionAmount,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		Contract:                dataHeader.Contract,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}

func ConvertToAddressUpdates(header dpfm_api_input_reader.Header, address dpfm_api_input_reader.Address) *AddressUpdates {
	dataHeader := header
	data := address

	return &AddressUpdates{
		Contract:    dataHeader.Contract,
		AddressID:   data.AddressID,
		PostalCode:  data.PostalCode,
		LocalRegion: data.LocalRegion,
		Country:     data.Country,
		District:    data.District,
		StreetName:  data.StreetName,
		CityName:    data.CityName,
		Building:    data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}
