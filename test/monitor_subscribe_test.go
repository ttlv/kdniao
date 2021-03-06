package test

import (
	"testing"

	"github.com/ttlv/kdniao"
	"github.com/ttlv/kdniao/sdk"
)

func TestMonitorSubscribe(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	apiMonitorSubscribeSdk := sdk.NewApiMonitorSubscribe(config, logger)
	req := apiMonitorSubscribeSdk.GetRequest("4303618027892", "YD")
	resp, err := apiMonitorSubscribeSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.UpdateTime", resp.UpdateTime)
}
