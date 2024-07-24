package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TwiplaSubscriptionAPI struct {
	client *TwiplaApiClient
}

func (t *TwiplaSubscriptionAPI) Upgrade(args UpgradeArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/upgrade"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't upgrade subscription. %s", string(payload))
	}

	return nil
}

func (t *TwiplaSubscriptionAPI) Downgrade(args DowngradeArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/downgrade"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't downgrade subscription. %s", string(payload))
	}

	return nil
}

func (t *TwiplaSubscriptionAPI) Cancel(websiteExtID string) error {
	jsonData, err := json.Marshal(struct {
		IntpWebsiteId string `json:"intpWebsiteId"`
	}{
		IntpWebsiteId: websiteExtID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/cancel"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't cancel subscription. %s", string(payload))
	}

	return nil
}

func (t *TwiplaSubscriptionAPI) Resume(websiteExtID string) error {
	jsonData, err := json.Marshal(struct {
		IntpWebsiteId string `json:"intpWebsiteId"`
	}{
		IntpWebsiteId: websiteExtID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/resume"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't resume subscription. %s", string(payload))
	}

	return nil
}

func (t *TwiplaSubscriptionAPI) Deactivate(websiteExtID string) error {
	jsonData, err := json.Marshal(struct {
		IntpWebsiteId string `json:"intpWebsiteId"`
	}{
		IntpWebsiteId: websiteExtID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/deactivate"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't deactivate subscription. %s", string(payload))
	}

	return nil
}

func (t *TwiplaSubscriptionAPI) Reactivate(args ReactivateArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v2/3as/notifications/subscriptions/reactivate"
	r, err := t.client.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		payload, _ := io.ReadAll(res.Body)
		return fmt.Errorf("can't reactivate subscription. %s", string(payload))
	}

	return nil
}

func NewTwiplaSubscriptionAPI(
	twiplaAPIClient *TwiplaApiClient,
) *TwiplaSubscriptionAPI {
	return &TwiplaSubscriptionAPI{
		client: twiplaAPIClient,
	}
}
