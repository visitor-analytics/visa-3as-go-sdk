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

func (t *TwiplaSSRWebsiteAPI) New(args NewSSRWebsiteArgs) error {
	jsonData, err := json.Marshal(struct {
		ExtID    string             `json:"externalWebsiteId"`
		Name     string             `json:"name"`
		Platform string             `json:"platform"`
		Settings SSRWebsiteSettings `json:"recordingsSettings"`
	}{
		ExtID:    args.ExtID,
		Name:     args.Name,
		Platform: "AAAS",
		Settings: SSRWebsiteSettings{
			Paused:          false,
			AnyPage:         true,
			ClickAndScroll:  false,
			TextObfuscation: false,
			MinDuration:     0,
		},
	})
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/api/websites"
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
