package equip

// import (
// 	"context"
// 	"testing"

// 	// "gitee.com/csms/jxeu-ocpp/internal/config"
// 	// "gitee.com/csms/jxeu-ocpp/internal/errors"
// 	// "gitee.com/csms/jxeu-ocpp/pkg/api"
// 	// "gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	// "gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
// 	"github.com/samber/lo"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/ForbiddenR/jx-api/services"
// 	"github.com/ForbiddenR/jx-api/apierrors"
// )

// // func TestSetVariablesRequest(t *testing.T) {
// // 	config.TestConfig()
// // 	api.Init()
// // 	ctx := context.TODO()

// // 	p := services.OCPP16()
// // 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
// // 	req := []*equipSetVariablesCallbackRequest{
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				Protocol:    p,
// // 				AccessPod:   services.TestAccessPod,
// // 				MsgID:       "1",
// // 			},
// // 			Callback: services.NewCBError(err),
// // 		},
// // 		{
// // 			Base: services.Base{
// // 				EquipmentSn: services.TestSN,
// // 				Protocol:    p,
// // 				AccessPod:   services.TestAccessPod,
// // 				MsgID:       "1",
// // 			},
// // 			Callback: services.NewCB(services.Successful),
// // 		},
// // 	}

// // 	lo.ForEach(req, func(v *equipSetVariablesCallbackRequest, _ int) {
// // 		err := SetVariablesCallbackRequest(ctx, v)
// // 		assert.Nil(t, err)
// // 	})
// // }

// func TestSetVariablesRequestWithGeneric(t *testing.T) {
// 	// config.TestConfig()
// 	// api.Init()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)
// 	ctx := context.TODO()
// 	p := services.OCPP16()

// 	err := apierrors.NewCallbackErrorGenericError(services.TestSN, "protocol.RemoteStopTransactionFeatureName", "test err")
// 	req := []*equipSetVariablesCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCBError(err),
// 		},
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

// 	lo.ForEach(req, func(v *equipSetVariablesCallbackRequest, _ int) {
// 		err := SetVariablesRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	})
// }