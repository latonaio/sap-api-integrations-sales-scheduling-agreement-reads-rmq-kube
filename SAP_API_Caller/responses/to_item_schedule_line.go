package responses

type ToItemScheduleLine struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesSchedulingAgreement       string `json:"SalesSchedulingAgreement"`
			SalesSchedulingAgreementItem   string `json:"SalesSchedulingAgreementItem"`
			ScheduleLine                   string `json:"ScheduleLine"`
			SchedulingAgreementReleaseType string `json:"SchedulingAgreementReleaseType"`
			ScheduleLineCategory           string `json:"ScheduleLineCategory"`
			OrderQuantityUnit              string `json:"OrderQuantityUnit"`
			SalesDelivDateCategory         string `json:"SalesDelivDateCategory"`
			RequestedDeliveryDate          string `json:"RequestedDeliveryDate"`
			RequestedDeliveryTime          string `json:"RequestedDeliveryTime"`
			ScheduleLineOrderQuantity      string `json:"ScheduleLineOrderQuantity"`
			CorrectedQtyInOrderQtyUnit     string `json:"CorrectedQtyInOrderQtyUnit"`
			ScheduleLineOpenQuantity       string `json:"ScheduleLineOpenQuantity"`
			ConfdOrderQtyByMatlAvailCheck  string `json:"ConfdOrderQtyByMatlAvailCheck"`
			ProductAvailabilityDate        string `json:"ProductAvailabilityDate"`
			ProductAvailabilityTime        string `json:"ProductAvailabilityTime"`
			DelivBlockReasonForSchedLine   string `json:"DelivBlockReasonForSchedLine"`
			TransportationPlanningDate     string `json:"TransportationPlanningDate"`
			TransportationPlanningTime     string `json:"TransportationPlanningTime"`
			GoodsIssueDate                 string `json:"GoodsIssueDate"`
			LoadingDate                    string `json:"LoadingDate"`
			GoodsIssueTime                 string `json:"GoodsIssueTime"`
			LoadingTime                    string `json:"LoadingTime"`
			GoodsMovementType              string `json:"GoodsMovementType"`
		} `json:"results"`
	} `json:"d"`
}
