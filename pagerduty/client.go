package pagerduty

import (
	"os"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/wtfutil/wtf/wtf"
)

// GetOnCalls returns a list of people currently on call
func GetOnCalls() ([]pagerduty.OnCall, error) {
	client := pagerduty.NewClient(apiKey())

	var results []pagerduty.OnCall

	var queryOpts pagerduty.ListOnCallOptions
	queryOpts.Since = time.Now().Format("2006-01-02T15:04:05Z07:00")
	queryOpts.Until = time.Now().Format("2006-01-02T15:04:05Z07:00")

	oncalls, err := client.ListOnCalls(queryOpts)
	if err != nil {
		return nil, err
	}

	results = append(results, oncalls.OnCalls...)

	for oncalls.APIListObject.More == true {
		queryOpts.APIListObject.Offset = oncalls.APIListObject.Offset
		oncalls, err = client.ListOnCalls(queryOpts)
		if err != nil {
			return nil, err
		}
		results = append(results, oncalls.OnCalls...)
	}

	return results, nil
}

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.pagerduty.apiKey",
		os.Getenv("WTF_PAGERDUTY_API_KEY"),
	)
}
