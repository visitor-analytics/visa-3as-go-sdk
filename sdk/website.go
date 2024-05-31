package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TwiplaWebsiteAPI struct {
	client *TwiplaAPIClient
}

func (t *TwiplaWebsiteAPI) New(intpcID string, args NewWebsiteArgs) error {
	jsonData, err := json.Marshal(struct {
		NewWebsiteArgs
		IntpcID string `json:"intpCustomerId"`
	}{
		IntpcID:        intpcID,
		NewWebsiteArgs: args,
	})
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/websites"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("can't create new website")
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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get intp websites")
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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't get intp website")
	}

	return NewTwiplaJSON[Website](res.Body).Unmarshal()
}

func NewTwiplaWebsiteAPI(client *TwiplaAPIClient) *TwiplaWebsiteAPI {
	return &TwiplaWebsiteAPI{client: client}
}
