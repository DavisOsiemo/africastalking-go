package payment

import (
	"africastalking/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Service is a service
type Service struct {
	Username string
	APIKey   string
	Env      string
}

// NewService creates a new Service
func NewService() Service {
	return Service{}
}

// RequestB2C sends a B2C request
func (service Service) RequestB2C(body B2CRequest) (*B2CResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/b2c/request"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal b2c req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var b2cResponse B2CResponse
	json.NewDecoder(response.Body).Decode(&b2cResponse)
	defer response.Body.Close()
	return &b2cResponse, nil
}

// RequestB2B sends a B2B request
func (service Service) RequestB2B(body B2BRequest) (*B2BResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/b2b/request"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal b2b req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var b2bResponse B2BResponse
	json.NewDecoder(response.Body).Decode(&b2bResponse)
	defer response.Body.Close()
	return &b2bResponse, nil
}

// MobileCheckout requests
func (service Service) MobileCheckout(body MobileCheckoutRequest) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/mobile/checkout/request"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal mobile checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// CardCheckoutCharge requests
func (service Service) CardCheckoutCharge(body CardCheckoutRequest) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/card/checkout/charge"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal card checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// CardCheckoutValidate requests
func (service Service) CardCheckoutValidate(body CardValidateCheckoutRequest) (*CheckoutValidateResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/card/checkout/validate"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal card validate checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	defer response.Body.Close()
	return &cvr, nil
}

// BankCheckoutCharge requests
func (service Service) BankCheckoutCharge(body BankCheckoutRequest) (*CheckoutResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/checkout/charge"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank checkout charge req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var checkoutResponse CheckoutResponse
	json.NewDecoder(response.Body).Decode(&checkoutResponse)
	defer response.Body.Close()
	return &checkoutResponse, nil
}

// BankCheckoutValidate requests
func (service Service) BankCheckoutValidate(body BankValidateCheckoutRequest) (*CheckoutValidateResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/checkout/validate"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank validate checkout req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var cvr CheckoutValidateResponse
	json.NewDecoder(response.Body).Decode(&cvr)
	defer response.Body.Close()
	return &cvr, nil
}

// BankTransfer requests
func (service Service) BankTransfer(body BankTransferRequest) (*BankTransferResponse, error) {
	host := util.GetAPIHost(service.Env)
	url := host + "/bank/transfer"

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("could not marshal bank transfer req body %v: ", err)
	}

	response, err := service.newRequest(url, reqBody)
	if err != nil {
		return nil, err
	}

	var btr BankTransferResponse
	json.NewDecoder(response.Body).Decode(&btr)
	defer response.Body.Close()
	return &btr, nil
}

func (service Service) newRequest(url string, body []byte) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("apiKey", service.APIKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(request)
}