package equip

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/internal/errors"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetVariablesRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	// If sending a callback error, we don't need to fill the field 'data'.
// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
// 	req := []*equipGetVariablesCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCBError(err),
// 			},
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCB(services.Failed),
// 			},
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCB(services.Successful),
// 			},
// 		},
// 	}

// 	for _, v := range req {
// 		err := GetVariablesCallbackRequest(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }

// func TestGetVariablesRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	// If sending a callback error, we don't need to fill the field 'data'.
// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
// 	req := []*equipGetVariablesCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCBError(err),
// 			},
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCB(services.Failed),
// 			},
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: &equipGetVariablesCallbackRequestDetail{
// 				CB: services.NewCB(services.Successful),
// 			},
// 		},
// 	}

// 	for _, v := range req {
// 		err := GetVariablesCallbackRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }