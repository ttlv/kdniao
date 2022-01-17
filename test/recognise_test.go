package test

import (
	"testing"

	"github.com/ttlv/kdniao"
	"github.com/ttlv/kdniao/sdk"
)

func TestRecognise(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniao.NewKdniaoLogger()

	apiRecogniseSdk := sdk.NewApiRecognise(config, logger)
	req := apiRecogniseSdk.GetRequest("JDVB10958632214")
	resp, err := apiRecogniseSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.LogisticCode", resp.LogisticCode)
	for _, shipper := range resp.Shippers {
		t.Log(shipper.ShipperCode, shipper.ShipperName)
	}
}
