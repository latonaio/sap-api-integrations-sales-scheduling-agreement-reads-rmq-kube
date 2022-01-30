package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-scheduling-agreement-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			SalesSchedulingAgreement:      data.SalesSchedulingAgreement,
			SalesSchedgAgrmtType:          data.SalesSchedgAgrmtType,
			CreationDate:                  data.CreationDate,
			LastChangeDate:                data.LastChangeDate,
			SalesOrganization:             data.SalesOrganization,
			DistributionChannel:           data.DistributionChannel,
			OrganizationDivision:          data.OrganizationDivision,
			SalesGroup:                    data.SalesGroup,
			SalesOffice:                   data.SalesOffice,
			SoldToParty:                   data.SoldToParty,
			SalesSchedgAgrmtDate:          data.SalesSchedgAgrmtDate,
			SDDocumentReason:              data.SDDocumentReason,
			PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
			CustomerPurchaseOrderType:     data.CustomerPurchaseOrderType,
			CustomerPurchaseOrderDate:     data.CustomerPurchaseOrderDate,
			SalesDistrict:                 data.SalesDistrict,
			TotalNetAmount:                data.TotalNetAmount,
			TransactionCurrency:           data.TransactionCurrency,
			PricingDate:                   data.PricingDate,
			ShippingType:                  data.ShippingType,
			ShippingCondition:             data.ShippingCondition,
			DeliveryBlockReason:           data.DeliveryBlockReason,
			DelivSchedTypeMRPRlvnceCode:   data.DelivSchedTypeMRPRlvnceCode,
			AgrmtValdtyStartDate:          data.AgrmtValdtyStartDate,
			AgrmtValdtyEndDate:            data.AgrmtValdtyEndDate,
			HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
			CustomerPaymentTerms:          data.CustomerPaymentTerms,
			PaymentMethod:                 data.PaymentMethod,
			OverallSDProcessStatus:        data.OverallSDProcessStatus,
			OverallSDDocumentRejectionSts: data.OverallSDDocumentRejectionSts,
			TotalBlockStatus:              data.TotalBlockStatus,
			OverallDeliveryStatus:         data.OverallDeliveryStatus,
			OverallDeliveryBlockStatus:    data.OverallDeliveryBlockStatus,
			OverallBillingBlockStatus:     data.OverallBillingBlockStatus,
			TotalCreditCheckStatus:        data.TotalCreditCheckStatus,
			ToHeaderPartner:               data.ToHeaderPartner.Deferred.URI,
			ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			SalesSchedulingAgreement:       data.SalesSchedulingAgreement,
			SalesSchedulingAgreementItem:   data.SalesSchedulingAgreementItem,
			SalesSchedgAgrmtItemCategory:   data.SalesSchedgAgrmtItemCategory,
			LastChangeDate:                 data.LastChangeDate,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialGroup:                  data.MaterialGroup,
			MaterialPricingGroup:           data.MaterialPricingGroup,
			ProductionPlant:                data.ProductionPlant,
			StorageLocation:                data.StorageLocation,
			SalesSchedgAgrmtItemText:       data.SalesSchedgAgrmtItemText,
			PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
			OrderQuantity:                  data.OrderQuantity,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			TargetQuantityUnit:             data.TargetQuantityUnit,
			SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
			NetAmount:                      data.NetAmount,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			NetPriceQuantityUnit:           data.NetPriceQuantityUnit,
			TransactionCurrency:            data.TransactionCurrency,
			PricingDate:                    data.PricingDate,
			ShippingPoint:                  data.ShippingPoint,
			ShippingType:                   data.ShippingType,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			SDProcessStatus:                data.SDProcessStatus,
			DeliveryStatus:                 data.DeliveryStatus,
			ItemGeneralIncompletionStatus:  data.ItemGeneralIncompletionStatus,
			ItemBillingIncompletionStatus:  data.ItemBillingIncompletionStatus,
			ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
			ToItemDeliverySchedule:         data.ToItemDeliverySchedule.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
			SalesSchedulingAgreement: data.SalesSchedulingAgreement,
			PartnerFunction:          data.PartnerFunction,
			Customer:                 data.Customer,
			Supplier:                 data.Supplier,
			AddressID:                data.AddressID,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			SalesSchedulingAgreement:       data.SalesSchedulingAgreement,
			SalesSchedulingAgreementItem:   data.SalesSchedulingAgreementItem,
			SalesSchedgAgrmtItemCategory:   data.SalesSchedgAgrmtItemCategory,
			LastChangeDate:                 data.LastChangeDate,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialGroup:                  data.MaterialGroup,
			MaterialPricingGroup:           data.MaterialPricingGroup,
			ProductionPlant:                data.ProductionPlant,
			StorageLocation:                data.StorageLocation,
			SalesSchedgAgrmtItemText:       data.SalesSchedgAgrmtItemText,
			PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
			OrderQuantity:                  data.OrderQuantity,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			TargetQuantity:                 data.TargetQuantity,
			TargetQuantityUnit:             data.TargetQuantityUnit,
			SalesDocumentRjcnReason:        data.SalesDocumentRjcnReason,
			NetAmount:                      data.NetAmount,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			NetPriceQuantityUnit:           data.NetPriceQuantityUnit,
			TransactionCurrency:            data.TransactionCurrency,
			PricingDate:                    data.PricingDate,
			ShippingPoint:                  data.ShippingPoint,
			ShippingType:                   data.ShippingType,
			DeliveryPriority:               data.DeliveryPriority,
			IncotermsClassification:        data.IncotermsClassification,
			CustomerPaymentTerms:           data.CustomerPaymentTerms,
			ItemBillingBlockReason:         data.ItemBillingBlockReason,
			SDProcessStatus:                data.SDProcessStatus,
			DeliveryStatus:                 data.DeliveryStatus,
			ItemGeneralIncompletionStatus:  data.ItemGeneralIncompletionStatus,
			ItemBillingIncompletionStatus:  data.ItemBillingIncompletionStatus,
			ItemDeliveryIncompletionStatus: data.ItemDeliveryIncompletionStatus,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
			ToItemDeliverySchedule:         data.ToItemDeliverySchedule.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemDeliverySchedule(raw []byte, l *logger.Logger) ([]ToItemDeliverySchedule, error) {
	pm := &responses.ToItemDeliverySchedule{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemDeliverySchedule. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemDeliverySchedule := make([]ToItemDeliverySchedule, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemDeliverySchedule = append(toItemDeliverySchedule, ToItemDeliverySchedule{
			SalesSchedulingAgreement:       data.SalesSchedulingAgreement,
			SalesSchedulingAgreementItem:   data.SalesSchedulingAgreementItem,
			IntDeliveryScheduleNumber:      data.IntDeliveryScheduleNumber,
			SchedulingAgreementReleaseType: data.SchedulingAgreementReleaseType,
			DeliveryScheduleStartDate:      data.DeliveryScheduleStartDate,
			DeliveryScheduleEndDate:        data.DeliveryScheduleEndDate,
			CustomerDeliveryScheduleNumber: data.CustomerDeliveryScheduleNumber,
			DeliveryScheduleDate:           data.DeliveryScheduleDate,
			LastDeliveryDocPostingDate:     data.LastDeliveryDocPostingDate,
			LastDeliveryDocument:           data.LastDeliveryDocument,
			LastIntDeliveryScheduleNumber:  data.LastIntDeliveryScheduleNumber,
			CreationDate:                   data.CreationDate,
			DeliveryScheduleCreationDate:   data.DeliveryScheduleCreationDate,
			LastReceiptQuantity:            data.LastReceiptQuantity,
			CumulativeReceiptQuantity:      data.CumulativeReceiptQuantity,
			CumulativeIssuedQuantity:       data.CumulativeIssuedQuantity,
			CumulativeDeliveredQuantity:    data.CumulativeDeliveredQuantity,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			LastChangeDate:                 data.LastChangeDate,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
		})
	}

	return toItemDeliverySchedule, nil
}

func ConvertToToItemScheduleLine(raw []byte, l *logger.Logger) ([]ToItemScheduleLine, error) {
	pm := &responses.ToItemScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemScheduleLine := make([]ToItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemScheduleLine = append(toItemScheduleLine, ToItemScheduleLine{
			SalesSchedulingAgreement:       data.SalesSchedulingAgreement,
			SalesSchedulingAgreementItem:   data.SalesSchedulingAgreementItem,
			ScheduleLine:                   data.ScheduleLine,
			SchedulingAgreementReleaseType: data.SchedulingAgreementReleaseType,
			ScheduleLineCategory:           data.ScheduleLineCategory,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			SalesDelivDateCategory:         data.SalesDelivDateCategory,
			RequestedDeliveryDate:          data.RequestedDeliveryDate,
			RequestedDeliveryTime:          data.RequestedDeliveryTime,
			ScheduleLineOrderQuantity:      data.ScheduleLineOrderQuantity,
			CorrectedQtyInOrderQtyUnit:     data.CorrectedQtyInOrderQtyUnit,
			ScheduleLineOpenQuantity:       data.ScheduleLineOpenQuantity,
			ConfdOrderQtyByMatlAvailCheck:  data.ConfdOrderQtyByMatlAvailCheck,
			ProductAvailabilityDate:        data.ProductAvailabilityDate,
			ProductAvailabilityTime:        data.ProductAvailabilityTime,
			DelivBlockReasonForSchedLine:   data.DelivBlockReasonForSchedLine,
			TransportationPlanningDate:     data.TransportationPlanningDate,
			TransportationPlanningTime:     data.TransportationPlanningTime,
			GoodsIssueDate:                 data.GoodsIssueDate,
			LoadingDate:                    data.LoadingDate,
			GoodsIssueTime:                 data.GoodsIssueTime,
			LoadingTime:                    data.LoadingTime,
			GoodsMovementType:              data.GoodsMovementType,
		})
	}

	return toItemScheduleLine, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			SalesSchedulingAgreement:     data.SalesSchedulingAgreement,
			SalesSchedulingAgreementItem: data.SalesSchedulingAgreementItem,
			PricingProcedureStep:         data.PricingProcedureStep,
			PricingProcedureCounter:      data.PricingProcedureCounter,
			ConditionApplication:         data.ConditionApplication,
			ConditionType:                data.ConditionType,
			PricingDateTime:              data.PricingDateTime,
			ConditionCalculationType:     data.ConditionCalculationType,
			ConditionBaseValue:           data.ConditionBaseValue,
			ConditionRateValue:           data.ConditionRateValue,
			ConditionCurrency:            data.ConditionCurrency,
			ConditionQuantity:            data.ConditionQuantity,
			ConditionQuantityUnit:        data.ConditionQuantityUnit,
			ConditionCategory:            data.ConditionCategory,
			PricingScaleType:             data.PricingScaleType,
			ConditionRecord:              data.ConditionRecord,
			ConditionSequentialNumber:    data.ConditionSequentialNumber,
			TaxCode:                      data.TaxCode,
			ConditionAmount:              data.ConditionAmount,
			TransactionCurrency:          data.TransactionCurrency,
			PricingScaleBasis:            data.PricingScaleBasis,
			ConditionScaleBasisValue:     data.ConditionScaleBasisValue,
			ConditionScaleBasisUnit:      data.ConditionScaleBasisUnit,
			ConditionScaleBasisCurrency:  data.ConditionScaleBasisCurrency,
			ConditionIsManuallyChanged:   data.ConditionIsManuallyChanged,
		})
	}

	return toItemPricingElement, nil
}
