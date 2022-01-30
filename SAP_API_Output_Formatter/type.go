package sap_api_output_formatter

type SalesSchedulingAgreement struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	APISchema           string `json:"api_schema"`
	SchedulingAgreement string `json:"sales_scheduling_agreement"`
	Deleted             bool   `json:"deleted"`
}    
    
type Header struct {
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
    ToHeaderPartner               string      `json:"to_SchAgrmtPartner"`
    ToItem                        string      `json:"to_SchedgAgrmtItm"`
}

type Item struct {
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
	ToItemPricingElement           string `json:"to_PricingElement"`
	ToItemDeliverySchedule         string `json:"to_SalesSchedgAgrmtDelivSched"`
}

type ToHeaderPartner struct {
	SalesSchedulingAgreement string `json:"SalesSchedulingAgreement"`
	PartnerFunction          string `json:"PartnerFunction"`
	Customer                 string `json:"Customer"`
	Supplier                 string `json:"Supplier"`
	AddressID                string `json:"AddressID"`
}

type ToItem struct {
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
	ToItemPricingElement           string `json:"to_PricingElement"`
	ToItemDeliverySchedule         string `json:"to_SalesSchedgAgrmtDelivSched"`
}

type ToItemDeliverySchedule struct {
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
	ToItemScheduleLine             string      `json:"to_SalesSchedgAgrmtSchedLine"`
}

type ToItemScheduleLine struct {
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
}

type ToItemPricingElement struct {
	SalesSchedulingAgreement      string `json:"SalesSchedulingAgreement"`
	SalesSchedulingAgreementItem  string `json:"SalesSchedulingAgreementItem"`
	PricingProcedureStep          string `json:"PricingProcedureStep"`
	PricingProcedureCounter       string `json:"PricingProcedureCounter"`
	ConditionApplication          string `json:"ConditionApplication"`
	ConditionType                 string `json:"ConditionType"`
	PricingDateTime               string `json:"PricingDateTime"`
	ConditionCalculationType      string `json:"ConditionCalculationType"`
	ConditionBaseValue            string `json:"ConditionBaseValue"`
	ConditionRateValue            string `json:"ConditionRateValue"`
	ConditionCurrency             string `json:"ConditionCurrency"`
	ConditionQuantity             string `json:"ConditionQuantity"`
	ConditionQuantityUnit         string `json:"ConditionQuantityUnit"`
	ConditionCategory             string `json:"ConditionCategory"`
	PricingScaleType              string `json:"PricingScaleType"`
	ConditionRecord               string `json:"ConditionRecord"`
	ConditionSequentialNumber     string `json:"ConditionSequentialNumber"`
	TaxCode                       string `json:"TaxCode"`
	ConditionAmount               string `json:"ConditionAmount"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	PricingScaleBasis             string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue      string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit       string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency   string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged    bool   `json:"ConditionIsManuallyChanged"`
}
