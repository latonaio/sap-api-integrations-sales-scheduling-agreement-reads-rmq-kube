package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-sales-scheduling-agreement-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesSchedulingAgreement(salesSchedulingAgreement, salesSchedulingAgreementItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(salesSchedulingAgreement)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(salesSchedulingAgreement, salesSchedulingAgreementItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(salesSchedulingAgreement string) {
	headerData, err := c.callSalesSchedulingAgreementSrvAPIRequirementHeader("A_SalesSchedgAgrmt", salesSchedulingAgreement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerData, "function": "SalesSchedulingAgreementHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerPartnerData, "function": "SalesSchedulingAgreementPartner"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerPartnerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "SalesSchedulingAgreementItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemDeliveryScheduleData, err := c.callToItemDeliverySchedule(itemData[0].ToItemDeliverySchedule)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemDeliveryScheduleData, "function": "SalesSchedulingAgreementItemDeliverySchedule"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemDeliveryScheduleData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemDeliveryScheduleData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemScheduleLineData, "function": "SalesSchedulingAgreementItemScheduleLine"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemPricingElementData, "function": "SalesSchedulingAgreementItemPricingElement"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

}

func (c *SAPAPICaller) callSalesSchedulingAgreementSrvAPIRequirementHeader(api, salesSchedulingAgreement string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_SALES_SCHEDULING_AGREEMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, salesSchedulingAgreement)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDeliverySchedule(url string) ([]sap_api_output_formatter.ToItemDeliverySchedule, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDeliverySchedule(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(salesSchedulingAgreement, salesSchedulingAgreementItem string) {
	itemData, err := c.callSalesSchedulingAgreementSrvAPIRequirementItem("A_SalesSchedgAgrmtItem", salesSchedulingAgreement, salesSchedulingAgreementItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "SalesSchedulingAgreementItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemDeliveryScheduleData, err := c.callToItemDeliverySchedule(itemData[0].ToItemDeliverySchedule)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemDeliveryScheduleData, "function": "SalesSchedulingAgreementItemDeliverySchedule"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemDeliveryScheduleData)

	itemScheduleLineData, err := c.callToItemScheduleLine(itemDeliveryScheduleData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemScheduleLineData, "function": "SalesSchedulingAgreementItemScheduleLine"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemScheduleLineData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemPricingElementData, "function": "SalesSchedulingAgreementItemPricingElement"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)

}

func (c *SAPAPICaller) callSalesSchedulingAgreementSrvAPIRequirementItem(api, salesSchedulingAgreement, salesSchedulingAgreementItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_SALES_SCHEDULING_AGREEMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, salesSchedulingAgreement, salesSchedulingAgreementItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, salesSchedulingAgreement string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SalesSchedulingAgreement eq '%s'", salesSchedulingAgreement))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, salesSchedulingAgreement, salesSchedulingAgreementItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SalesSchedulingAgreement eq '%s' and SalesSchedulingAgreementItem eq '%s'", salesSchedulingAgreement, salesSchedulingAgreementItem))
	req.URL.RawQuery = params.Encode()
}
