package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesSchedulingAgreement       string `json:"SalesSchedulingAgreement"`
			SalesSchedulingAgreementItem   string `json:"SalesSchedulingAgreementItem"`
			SalesSchedgAgrmtItemCategory   string `json:"SalesSchedgAgrmtItemCategory"`
			LastChangeDate                 string `json:"LastChangeDate"`
			Material                       string `json:"Material"`
			MaterialByCustomer             string `json:"MaterialByCustomer"`
			MaterialGroup                  string `json:"MaterialGroup"`
			MaterialPricingGroup           string `json:"MaterialPricingGroup"`
			ProductionPlant                string `json:"ProductionPlant"`
			StorageLocation                string `json:"StorageLocation"`
			SalesSchedgAgrmtItemText       string `json:"SalesSchedgAgrmtItemText"`
			PurchaseOrderByCustomer        string `json:"PurchaseOrderByCustomer"`
			OrderQuantity                  string `json:"OrderQuantity"`
			OrderQuantityUnit              string `json:"OrderQuantityUnit"`
			TargetQuantity                 string `json:"TargetQuantity"`
			TargetQuantityUnit             string `json:"TargetQuantityUnit"`
			SalesDocumentRjcnReason        string `json:"SalesDocumentRjcnReason"`
			NetAmount                      string `json:"NetAmount"`
			NetPriceAmount                 string `json:"NetPriceAmount"`
			NetPriceQuantity               string `json:"NetPriceQuantity"`
			NetPriceQuantityUnit           string `json:"NetPriceQuantityUnit"`
			TransactionCurrency            string `json:"TransactionCurrency"`
			PricingDate                    string `json:"PricingDate"`
			ShippingPoint                  string `json:"ShippingPoint"`
			ShippingType                   string `json:"ShippingType"`
			DeliveryPriority               string `json:"DeliveryPriority"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
			ItemBillingBlockReason         string `json:"ItemBillingBlockReason"`
			SDProcessStatus                string `json:"SDProcessStatus"`
			DeliveryStatus                 string `json:"DeliveryStatus"`
			ItemGeneralIncompletionStatus  string `json:"ItemGeneralIncompletionStatus"`
			ItemBillingIncompletionStatus  string `json:"ItemBillingIncompletionStatus"`
			ItemDeliveryIncompletionStatus string `json:"ItemDeliveryIncompletionStatus"`
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
			ToItemDeliverySchedule struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SalesSchedgAgrmtDelivSched"`
		} `json:"results"`
	} `json:"d"`
}
