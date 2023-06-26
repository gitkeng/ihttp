package sysutil

import (
	"encoding/base64"
	"fmt"
	. "github.com/klauspost/cpuid/v2"
)

func hwUniqID(appId string) (string, error) {
	id, err := ID()
	if err != nil {
		return "", err
	}

	uniqId := fmt.Sprint(id, "|", CPU.BrandName, "|", CPU.VendorString, "|", CPU.Family, "|", CPU.PhysicalCores, "|", CPU.ThreadsPerCore, "|", CPU.LogicalCores)
	return protect(appId, uniqId), nil
}

func MachineHash(appId string) (string, error) {
	if uniqId, err := hwUniqID(appId); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString([]byte(uniqId)), nil
	}
}

func ValidMachineHash(appId, machineId string) bool {
	checkId, err := base64.StdEncoding.DecodeString(machineId)
	if err != nil {
		return false
	}
	if machineId, err := hwUniqID(appId); err != nil {
		return false
	} else {
		return machineId == string(checkId)
	}
}
