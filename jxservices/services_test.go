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

func TestURL(t *testing.T) {
	assert.Equal(t, "ac/firmwareStatusNotification",
		Equip+"/"+"firmwareStatusNotification")

	assert.Equal(t, "ac/callback/pushFirmwareCallback",
		Equip+"/"+Callback+"/"+UpdateFirmware.String()+"Callback")

	assert.Equal(t, "ac/callback/sendQRCodeCallback", Equip+"/"+Callback+"/"+SendQRCode.String()+"Callback")
}
