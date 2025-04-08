package jxservices

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeCheck(t *testing.T) {
	t.Log(typeCheckOr("t", func(r string) int {
		return len(r)
	}, "t"))
}

func TestProtocol(t *testing.T) {
	p := NewProtocol("IEC104", "0.1")
	t.Log(p.Equal(iec001))
	t.Log(p.Equal(iec002))
}

func TestMarshal(t *testing.T) {
	proto := iec001
	bytes, err := json.Marshal(proto)
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", bytes)

	pt := &Protocol{}
	err = json.Unmarshal(bytes, pt)
	if err != nil {
		panic(err)
	}
	t.Log(pt.String())
} 

func TestServiceRequestNameServices(t *testing.T) {
	assert.Equal(t, Register.FirstUpper(), "EquipRegister")
	// assert.Equal(t, Register.Split(), []string{"register"})

	assert.Equal(t, Online.FirstUpper(), "EquipOnline")
	// assert.Equal(t, Online.Split(), []string{"online"})

	assert.Equal(t, BootNotification.FirstUpper(), "BootNotification")
	// assert.Equal(t, BootNotification.Split(), []string{"boot", "notification"})

	assert.Equal(t, StatusNotification.FirstUpper(), "StatusNotification")
	// assert.Equal(t, StatusNotification.Split(), []string{"status", "notification"})

	assert.Equal(t, Authorize.FirstUpper(), "Authorize")
	// assert.Equal(t, Authorize.Split(), []string{"authorize"})

	// assert.Equal(t, RemoteStartTransaction.Split(), []string{"remote", "start", "transaction"})

	// assert.Equal(t, RemoteStopTransaction.Split(), []string{"remote", "stop", "transaction"})

	assert.Equal(t, RemoteStartTransaction.FirstUpper(), "RemoteStartTransaction")

	// assert.Equal(t, UpdateFirmware.Split(), []string{"push", "firmware", "equipment"})

	// assert.Equal(t, FirmwareStatusNotification.Split(), []string{"push", "firmware", "notification"})
}

func TestURL(t *testing.T) {
	assert.Equal(t, "ac/firmwareStatusNotification",
		Equip+"/"+"firmwareStatusNotification")

	assert.Equal(t, "ac/callback/pushFirmwareCallback",
		Equip+"/"+Callback+"/"+UpdateFirmware.String()+"Callback")

	assert.Equal(t, "ac/callback/sendQRCodeCallback", Equip+"/"+Callback+"/"+SendQRCode.String()+"Callback")
}
