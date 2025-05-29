package equip

// import (
// 	"context"
// 	"encoding/json"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// )

// func TestResetRequestUnmarshalJSON(t *testing.T) {
// 	r := EquipResetRequest{}
// 	err := json.Unmarshal([]byte(`{
//    "msgId":"415682532126725",
//    "protocol":{
//       "name":"OCPP",
//       "version":"1.6"
//    },
//    "data":{
//       "type":1
//    }
// }`), &r)
// 	assert.Nil(t, err)

// }

// func TestResetRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	req := []*equipResetCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 			//Data: nil,
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Failed),
// 			//Data: nil,
// 		},
// 	}
// 	for _, v := range req {
// 		err := ResetCallbackRequest(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }

// func TestResetRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	ctx := context.TODO()

// 	p := services.OCPP16()
// 	req := []*equipResetCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 			//Data: nil,
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Failed),
// 			//Data: nil,
// 		},
// 	}
// 	for _, v := range req {
// 		err := ResetCallbackRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }