package sdk

import (
	"fmt"
	"io"
	"net/http"
)

type NewPackageArgs struct {
	Name     string
	ExtID    *string
	STPs     int
	Price    float32
	Period   TwiplaPeriod
	Currency TwiplaCurrency
}

type Package struct {
	ID          string         `json:"id"`
	ExtID       *string        `json:"externalId"`
	Name        string         `json:"name"`
	Price       float32        `json:"price"`
	Currency    TwiplaCurrency `json:"currency"`
	Period      string         `json:"period"`
	Recommended bool           `json:"recommended"`
	Touchpoints float64        `json:"touchpoints"`
	CreatedAt   string         `json:"createdAt"`
}

type TwiplaPackageAPI struct {
	client *TwiplaAPIClient
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

func NewTwiplaPackageAPI(
	twiplaAPIClient *TwiplaAPIClient,
) *TwiplaPackageAPI {
	return &TwiplaPackageAPI{
		client: twiplaAPIClient,
	}
}
