package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewPackageArgs struct {
	Name     string  `json:"name"`
	ExtID    *string `json:"externalId,omitempty"`
	STPs     int     `json:"touchpoints"`
	Price    float32 `json:"price"`
	Period   string  `json:"period"`
	Currency string  `json:"currency"`
}

type Package struct {
	ID          string  `json:"id"`
	ExtID       *string `json:"externalId"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Currency    string  `json:"currency"`
	Period      string  `json:"period"`
	Recommended bool    `json:"recommended"`
	Touchpoints float64 `json:"touchpoints"`
	CreatedAt   string  `json:"createdAt"`
}

type TwiplaPackageAPI struct {
	client *TwiplaApiClient
}

func (t *TwiplaPackageAPI) List() (*[]Package, error) {
	url := t.client.apiGateway + "/v2/3as/packages/"
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
		return nil, fmt.Errorf("can't get intp packages. %s", string(payload))
	}

	return NewTwiplaJSON[[]Package](res.Body).Unmarshal()
}

func (t *TwiplaPackageAPI) GetByID(ID string) (*Package, error) {
	url := t.client.apiGateway + "/v2/3as/packages/" + ID
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
		return nil, fmt.Errorf("can't get intp package. %s", string(payload))
	}

	return NewTwiplaJSON[Package](res.Body).Unmarshal()
}

func (t *TwiplaPackageAPI) Create(args NewPackageArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/packages"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't create new package. %d, %s", res.StatusCode, string(payload))
	}

	return nil
}

func NewTwiplaPackageAPI(
	twiplaAPIClient *TwiplaApiClient,
) *TwiplaPackageAPI {
	return &TwiplaPackageAPI{
		client: twiplaAPIClient,
	}
}
