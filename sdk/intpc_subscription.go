package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TwiplaIntpcSubscriptionAPI struct {
	client *TwiplaApiClient
}

func (t *TwiplaIntpcSubscriptionAPI) Upgrade(args UpgradeIntpcArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/upgrade"
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

func (t *TwiplaIntpcSubscriptionAPI) Downgrade(args DowngradeIntpcArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/downgrade"
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

func (t *TwiplaIntpcSubscriptionAPI) Cancel(intpcID string) error {
	jsonData, err := json.Marshal(struct {
		IntpcID string `json:"intpcId"`
	}{
		IntpcID: intpcID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/cancel"
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

func (t *TwiplaIntpcSubscriptionAPI) Resume(intpcID string) error {
	jsonData, err := json.Marshal(struct {
		IntpcID string `json:"intpcId"`
	}{
		IntpcID: intpcID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/resume"
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

func (t *TwiplaIntpcSubscriptionAPI) Deactivate(intpcID string) error {
	jsonData, err := json.Marshal(struct {
		IntpcID string `json:"intpcId"`
	}{
		IntpcID: intpcID,
	})

	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/deactivate"
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

func (t *TwiplaIntpcSubscriptionAPI) Reactivate(args ReactivateIntpcArgs) error {
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	url := t.client.apiGateway + "/v3/3as/intpc-subscriptions/reactivate"
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

func NewTwiplaIntpcSubscriptionAPI(
	twiplaAPIClient *TwiplaApiClient,
) *TwiplaIntpcSubscriptionAPI {
	return &TwiplaIntpcSubscriptionAPI{
		client: twiplaAPIClient,
	}
}
