package test

import (
	`testing`

	`github.com/ttlv/kdniao`
	`github.com/ttlv/kdniao/sdk`
)

func TestExpressQuery(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	expressQuerySdk := sdk.NewExpressQuery(config, logger)
	req := expressQuerySdk.GetRequest("550000609031770")
	resp, err := expressQuerySdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessID)
	t.Log("resp.LogisticCode", resp.LogisticCode)
}

