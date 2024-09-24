package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TwiplaWebsiteAPI struct {
	client              *TwiplaApiClient
	twiplaSSRWebsiteAPI *TwiplaSSRWebsiteAPI
}

func (t *TwiplaWebsiteAPI) newTwiplaWebsite(intpcID string, args NewWebsiteArgs) (*Website, error) {
	jsonData, err := json.Marshal(struct {
		NewWebsiteArgs
		IntpcID string `json:"intpCustomerId"`
	}{
		IntpcID:        intpcID,
		NewWebsiteArgs: args,
	})
	if err != nil {
		return nil, err
	}

	url := t.client.apiGateway + "/v2/3as/websites"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("status code: %d; payload: %s", res.StatusCode, string(payload))
	}

	return NewTwiplaJSON[Website](res.Body).Unmarshal()
}

func (t *TwiplaWebsiteAPI) New(intpcID string, args NewWebsiteArgs) error {
	twiplaWebsite, err := t.newTwiplaWebsite(intpcID, args)
	if err != nil {
		return err
	}

	if t.twiplaSSRWebsiteAPI != nil {
		return t.twiplaSSRWebsiteAPI.New(twiplaWebsite.ID, twiplaWebsite.VisaCustomerID)
	}

	return nil
}

func (t *TwiplaWebsiteAPI) List(pag PagArgs) (*[]Website, error) {
	url := fmt.Sprintf("%s/v2/3as/websites?page=%d&pageSize=%d", t.client.apiGateway, pag.Page, pag.PageSize)
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
		return nil, fmt.Errorf("status code: %d; payload: %s", res.StatusCode, string(payload))
	}

	return NewTwiplaJSON[[]Website](res.Body).Unmarshal()
}

func (t *TwiplaWebsiteAPI) GetByID(ID string) (*Website, error) {
	url := t.client.apiGateway + "/v2/3as/websites/" + ID
	r, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return NewTwiplaJSON[Website](res.Body).Unmarshal()
	}

	if res.StatusCode == http.StatusNotFound {
		//return nil, &NotFoundError{
		//	Resource: fmt.Sprintf("3as external id: %s", ID),
		//	Err:      WebsiteNotFoundError,
		//}
		return nil, fmt.Errorf("not found")
	}

	payload, _ := io.ReadAll(res.Body)
	return nil, fmt.Errorf("status code: %d; payload: %s", res.StatusCode, string(payload))
}

func (t *TwiplaWebsiteAPI) GetByIntID(ID string) (*Website, error) {
	url := t.client.apiGateway + "/v2/3as/internal/websites/" + ID
	r, err := t.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return NewTwiplaJSON[Website](res.Body).Unmarshal()
	}

	if res.StatusCode == http.StatusNotFound {
		//return nil, &NotFoundError{
		//	Resource: fmt.Sprintf("3as internal id: %s", ID),
		//	Err:      WebsiteNotFoundError,
		//}
		return nil, fmt.Errorf("not found")
	}

	payload, _ := io.ReadAll(res.Body)
	return nil, fmt.Errorf("status code: %d; payload: %s", res.StatusCode, string(payload))
}

func NewTwiplaWebsiteAPI(
	client *TwiplaApiClient,
	ssrAPI *TwiplaSSRWebsiteAPI,
) *TwiplaWebsiteAPI {
	return &TwiplaWebsiteAPI{
		client:              client,
		twiplaSSRWebsiteAPI: ssrAPI,
	}
}
