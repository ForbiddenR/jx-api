package jxservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	assert.Equal(t, "ac/firmwareStatusNotification",
		Equip+"/"+"firmwareStatusNotification")

	assert.Equal(t, "ac/callback/pushFirmwareCallback",
		Equip+"/"+Callback+"/"+UpdateFirmware.String()+"Callback")

	assert.Equal(t, "ac/callback/sendQRCodeCallback", Equip+"/"+Callback+"/"+SendQRCode.String()+"Callback")
}

func TestFeatureCollection(t *testing.T) {
	result, ok := UnsupportedFeatures.Get("test")
	t.Log(result, ok)
}
