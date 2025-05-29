package equip

// func TestSetChargingTimerRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	req := []*equipSetChargingTimerCallbackRequest{
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

// 	for _, value := range req {
// 		err := SetChargingTimerCallbackRequest(ctx, value)
// 		assert.Nil(t, err)
// 	}
// }

// func TestSetChargingTimerRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	req := []*equipSetChargingTimerCallbackRequest{
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

// 	for _, value := range req {
// 		err := SetChargingTimerCallbackRequestWithGeneric(ctx, value)
// 		assert.Nil(t, err)
// 	}
// }