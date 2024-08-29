package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TwiplaIntpcAPI struct {
	client        *TwiplaApiClient
	websiteAPI    *TwiplaWebsiteAPI
	ssrWebsiteAPI *TwiplaSSRWebsiteAPI
}

func (t *TwiplaIntpcAPI) New(args NewIntpcArgs) error {
	err := t.newTwiplaCustomer(args)
	if err != nil {
		return err
	}

	// check if recordings should work without onboarding
	// create ssr website/settings
	if t.ssrWebsiteAPI == nil {
		return nil
	}

	website, err := t.websiteAPI.GetByID(args.Website.ExtID)
	if err != nil {
		return err
	}

	return t.ssrWebsiteAPI.New(website.ID, website.VisaCustomerID)
}

func (t *TwiplaIntpcAPI) newTwiplaCustomer(args NewIntpcArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/customers"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't create new customer. %d, %s", res.StatusCode, string(payload))
	}

	return nil
}

func (t *TwiplaIntpcAPI) List(pag PagArgs) (*[]Intpc, error) {
	url := fmt.Sprintf("%s/v2/3as/customers?page=%d&pageSize=%d", t.client.apiGateway, pag.Page, pag.PageSize)
	r, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		payload, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("can't get intp customers. %d, %s", res.StatusCode, string(payload))
	}

	return NewTwiplaJSON[[]Intpc](res.Body).Unmarshal()
}

func (t *TwiplaIntpcAPI) GetByID(ID string) (*Intpc, error) {
	url := t.client.apiGateway + "/v2/3as/customers/" + ID
	r, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("not found")
	}

	if res.StatusCode != http.StatusOK {
		payload, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("can't get intp customer. %d, %s", res.StatusCode, string(payload))
	}

	return NewTwiplaJSON[Intpc](res.Body).Unmarshal()
}

func NewTwiplaIntpcAPI(
	client *TwiplaApiClient,
	websiteAPI *TwiplaWebsiteAPI,
	ssrWebsiteAPI *TwiplaSSRWebsiteAPI,
) *TwiplaIntpcAPI {
	return &TwiplaIntpcAPI{
		client:        client,
		websiteAPI:    websiteAPI,
		ssrWebsiteAPI: ssrWebsiteAPI,
	}
}
