package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewTwiplaSSRWebsiteAPI(client *TwiplaSSRApiClient) *TwiplaSSRWebsiteAPI {
	return &TwiplaSSRWebsiteAPI{client: client}
}

type TwiplaSSRWebsiteAPI struct {
	client *TwiplaSSRApiClient
}

func (t *TwiplaSSRWebsiteAPI) New(websiteID string, visaCustomerID string) error {
	jsonData, err := json.Marshal(
		SSRWebsiteSettings{
			Paused:          false,
			AnyPage:         true,
			ClickAndScroll:  false,
			TextObfuscation: false,
			MinDuration:     -1,
		})
	if err != nil {
		return err
	}

	url := t.client.apiGateway + fmt.Sprintf("/api/v2/aaas/users/%s/websites/%s/ssr-settings", visaCustomerID, websiteID)
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
		return fmt.Errorf("can't create new ssr website. %s", string(payload))
	}

	return nil
}
