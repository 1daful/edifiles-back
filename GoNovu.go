package cron

import (
	"context"
	"fmt"
	"log"
	"time"

	novu "github.com/novuhq/go-novu/lib"
)

func Notify(eventId string) {
	subscriberID := "<<REPLACE_WITH_YOUR_SUBSCRIBER>"
	apiKey := "<REPLACE_WITH_YOUR_API_KEY>"

	ctx := context.Background()
	to := map[string]interface{}{
		"lastName":     "Doe",
		"firstName":    "John",
		"subscriberId": subscriberID,
		"email":        "john@doemail.com",
	}

	payload := map[string]interface{}{
		"name": "Hello World",
		"organization": map[string]interface{}{
			"logo": "https://happycorp.com/logo.png",
		},
	}


	data := novu.ITriggerPayloadOptions{To: to, Payload: payload}
	novuClient := novu.NewAPIClient(apiKey, &novu.Config{RetryConfig: &novu.RetryConfigType{RetryMax: 3, WaitMin: 1 * time.Second, WaitMax: 5 * time.Second, InitialDelay: 0 * time.Second}})

	resp, err := novuClient.EventApi.Trigger(ctx, eventId, data)
	if err != nil {
		log.Fatal("novu error", err.Error())
		return
	}

	fmt.Println(resp)

	// get integrations
	integrations, err := novuClient.IntegrationsApi.GetAll(ctx)
	if err != nil {
		log.Fatal("Get all integrations error: ", err.Error())
	}
	fmt.Println(integrations)
}
