package sysutil_test

import (
	"github.com/gitkeng/ihttp/util/sysutil"
	"testing"
)

func TestID(t *testing.T) {
	id, err := sysutil.ID()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("ID: %s", id)
}

func TestValidMachineHash(t *testing.T) {
	machineId, err := sysutil.MachineHash("superrich-exchange")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("machineId: %s", machineId)

	if !sysutil.ValidMachineHash("superrich-exchange", machineId) {
		t.Error("machineId is invalid")
		return
	}

	t.Logf("machineId is valid")
}

func TestMachineId(t *testing.T) {
	uniqId, err := sysutil.MachineHash("superrich-exchange")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("uniqId: %s", uniqId)
}
