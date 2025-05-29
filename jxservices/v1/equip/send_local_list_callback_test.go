package equip

// func TestSendLocalListRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	req := []*equipSendLocalListCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(1),
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(2),
// 		},
// 	}
// 	for _, v := range req {
// 		err := SendLocalListCallbackRequest(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }

// func TestSendLocalListRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	req := []*equipSendLocalListCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 			},
// 			Callback: services.NewCB(1),
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 			},
// 			Callback: services.NewCB(2),
// 		},
// 	}
// 	for _, v := range req {
// 		err := SendLocalListCallbackRequestWithGeneric(ctx, v)
// 		assert.Nil(t, err)
// 	}
// }