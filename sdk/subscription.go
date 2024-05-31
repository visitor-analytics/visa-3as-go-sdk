package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TwiplaSubscriptionAPI struct {
	client *TwiplaAPIClient
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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("can't upgrade subscription")
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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("can't downgrade subscription")
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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("can't cancel subscription")
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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("can't resume subscription")
	}

	return nil
}

func NewTwiplaSubscriptionAPI(
	twiplaAPIClient *TwiplaAPIClient,
) *TwiplaSubscriptionAPI {
	return &TwiplaSubscriptionAPI{
		client: twiplaAPIClient,
	}
}
