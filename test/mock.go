package test

import (
	"errors"
	"strings"

	"github.com/ttlv/kdniao"
	"github.com/ttlv/kdniao/util"
)

func getConfig() (kdniao.KdniaoConfig, error) {
	eBusinessId, appKey, err := getConfigValue()
	return kdniao.NewKdniaoConfig(eBusinessId, appKey), err
}

func getConfigValue() (string, string, error) {
	configStr, err := util.FileGetContents("./config.txt")
	if err != nil {
		return "", "", err
	}
	configs := strings.Split(configStr, ",")
	if 2 != len(configs) {
		return "", "", errors.New("eBusinessId or appKey is empty")
	}
	return strings.TrimSpace(configs[0]), strings.TrimSpace(configs[1]), nil
}
