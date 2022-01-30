package responses

type ToHeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesSchedulingAgreement string `json:"SalesSchedulingAgreement"`
			PartnerFunction          string `json:"PartnerFunction"`
			Customer                 string `json:"Customer"`
			Supplier                 string `json:"Supplier"`
			AddressID                string `json:"AddressID"`
		} `json:"results"`
	} `json:"d"`
}
