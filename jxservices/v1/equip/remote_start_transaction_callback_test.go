package equip

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"github.com/samber/lo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRemoteStartTransactionRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	req := []*equipRemoteStartTransactionCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 		},
// 	}

// 	lo.ForEach(req, func(v *equipRemoteStartTransactionCallbackRequest, _ int) {
// 		err := RemoteStartTransactionCallbackRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	})
// }

// func TestRemoteStartTransactionRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)
// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	req := []*equipRemoteStartTransactionCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 		},
// 	}

// 	lo.ForEach(req, func(v *equipRemoteStartTransactionCallbackRequest, _ int) {
// 		err := RemoteStartTransactionCallbackRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	})
// }