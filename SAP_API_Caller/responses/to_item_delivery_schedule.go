package responses

type ToItemDeliverySchedule struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesSchedulingAgreement       string      `json:"SalesSchedulingAgreement"`
			SalesSchedulingAgreementItem   string      `json:"SalesSchedulingAgreementItem"`
			IntDeliveryScheduleNumber      string      `json:"IntDeliveryScheduleNumber"`
			SchedulingAgreementReleaseType string      `json:"SchedulingAgreementReleaseType"`
			DeliveryScheduleStartDate      string      `json:"DeliveryScheduleStartDate"`
			DeliveryScheduleEndDate        string      `json:"DeliveryScheduleEndDate"`
			CustomerDeliveryScheduleNumber string      `json:"CustomerDeliveryScheduleNumber"`
			DeliveryScheduleDate           string      `json:"DeliveryScheduleDate"`
			LastDeliveryDocPostingDate     string      `json:"LastDeliveryDocPostingDate"`
			LastDeliveryDocument           string      `json:"LastDeliveryDocument"`
			LastIntDeliveryScheduleNumber  string      `json:"LastIntDeliveryScheduleNumber"`
			CreationDate                   string      `json:"CreationDate"`
			DeliveryScheduleCreationDate   string      `json:"DeliveryScheduleCreationDate"`
			LastReceiptQuantity            string      `json:"LastReceiptQuantity"`
			CumulativeReceiptQuantity      string      `json:"CumulativeReceiptQuantity"`
			CumulativeIssuedQuantity       string      `json:"CumulativeIssuedQuantity"`
			CumulativeDeliveredQuantity    string      `json:"CumulativeDeliveredQuantity"`
			OrderQuantityUnit              string      `json:"OrderQuantityUnit"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			ToItemScheduleLine struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SalesSchedgAgrmtSchedLine"`
		} `json:"results"`
	} `json:"d"`
}
