package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			SalesSchedulingAgreement      string      `json:"SalesSchedulingAgreement"`
			SalesSchedgAgrmtType          string      `json:"SalesSchedgAgrmtType"`
			CreationDate                  string      `json:"CreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			SalesOrganization             string      `json:"SalesOrganization"`
			DistributionChannel           string      `json:"DistributionChannel"`
			OrganizationDivision          string      `json:"OrganizationDivision"`
			SalesGroup                    string      `json:"SalesGroup"`
			SalesOffice                   string      `json:"SalesOffice"`
			SoldToParty                   string      `json:"SoldToParty"`
			SalesSchedgAgrmtDate          string      `json:"SalesSchedgAgrmtDate"`
			SDDocumentReason              string      `json:"SDDocumentReason"`
			PurchaseOrderByCustomer       string      `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderType     string      `json:"CustomerPurchaseOrderType"`
			CustomerPurchaseOrderDate     string      `json:"CustomerPurchaseOrderDate"`
			SalesDistrict                 string      `json:"SalesDistrict"`
			TotalNetAmount                string      `json:"TotalNetAmount"`
			TransactionCurrency           string      `json:"TransactionCurrency"`
			PricingDate                   string      `json:"PricingDate"`
			ShippingType                  string      `json:"ShippingType"`
			ShippingCondition             string      `json:"ShippingCondition"`
			DeliveryBlockReason           string      `json:"DeliveryBlockReason"`
			DelivSchedTypeMRPRlvnceCode   string      `json:"DelivSchedTypeMRPRlvnceCode"`
			AgrmtValdtyStartDate          string      `json:"AgrmtValdtyStartDate"`
			AgrmtValdtyEndDate            string      `json:"AgrmtValdtyEndDate"`
			HeaderBillingBlockReason      string      `json:"HeaderBillingBlockReason"`
			CustomerPaymentTerms          string      `json:"CustomerPaymentTerms"`
			PaymentMethod                 string      `json:"PaymentMethod"`
			OverallSDProcessStatus        string      `json:"OverallSDProcessStatus"`
			OverallSDDocumentRejectionSts string      `json:"OverallSDDocumentRejectionSts"`
			TotalBlockStatus              string      `json:"TotalBlockStatus"`
			OverallDeliveryStatus         string      `json:"OverallDeliveryStatus"`
			OverallDeliveryBlockStatus    string      `json:"OverallDeliveryBlockStatus"`
			OverallBillingBlockStatus     string      `json:"OverallBillingBlockStatus"`
			TotalCreditCheckStatus        string      `json:"TotalCreditCheckStatus"`
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem                        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
