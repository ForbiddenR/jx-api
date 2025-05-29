package equip

// func TestRemoteStopTransactionRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	p := services.OCPP16()
// 	ctx := context.TODO()
// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
// 	req := []*equipRemoteStopTransactionCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCBError(err),
// 		},
// 	}
// 	for _, v := range req {
// 		generalErr := RemoteStopTransactionRequest(ctx, v)
// 		assert.Nil(t, generalErr)
// 	}
// }

// func TestRemoteStopTransactionRequestWithGeneric(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.RemoteStopTransactionFeatureName, "test err")
// 	req := []*equipRemoteStopTransactionCallbackRequest{
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCB(services.Successful),
// 		},
// 		{
// 			Base: services.Base{
// 				EquipmentSn: services.TestSN,
// 				Protocol:    p,
// 				AccessPod:   services.TestAccessPod,
// 				MsgID:       "1",
// 			},
// 			Callback: services.NewCBError(err),
// 		},
// 	}
// 	for _, v := range req {
// 		generalErr := RemoteStopTransactionRequestWithGeneric(ctx, v)
// 		assert.Nil(t, generalErr)
// 	}
// }