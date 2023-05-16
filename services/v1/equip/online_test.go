package equip

import (
	"context"
	"testing"

	"github.com/Kotodian/gokit/id"
	"github.com/stretchr/testify/assert"

	"gitee.com/csms/jxeu-ocpp/internal/config"
	"gitee.com/csms/jxeu-ocpp/internal/log"
	"gitee.com/csms/jxeu-ocpp/pkg/api"
	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
)

func TestEquipOnlineRequest(t *testing.T) {
	config.TestConfig()
	api.Init()
	log.InitNopLogger()

	ctx := context.TODO()
	p := services.OCPP16()

	remoteAddress := "127.0.0.1"
	req := []*equipOnlineRequest{
		newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String()),
		newTestEquipOnlineRequestWithRemoteAddress(services.TestSN, p, services.TestAccessPod, id.Next().String(), remoteAddress),
	}

	for _, v := range req {
		err := OnlineRequest(ctx, v)
		assert.Nil(t, err)
	}
}

func newTestEquipOnlineRequest(sn string, p *services.Protocol, pod, msgID string) *equipOnlineRequest {
	return NewEquipOnlineRequest(sn, p, pod, msgID)
}

func newTestEquipOnlineRequestWithRemoteAddress(sn string, p *services.Protocol, pod, msgID, remoteAddress string) *equipOnlineRequest {
	req := NewEquipOnlineRequest(sn, p, pod, msgID)
	req.Data.RemoteAddress = &remoteAddress
	return req
}
