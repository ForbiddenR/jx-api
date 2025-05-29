package equip

// func TestGetOfflineReason(t *testing.T) {
// 	str := GetOfflineReason(io.EOF)
// 	assert.Equal(t, "eof", str)
// 	str = GetOfflineReason(errors.New("read timeout"))
// 	assert.Equal(t, "overTime", str)
// }

// func TestEquipOnlineAndOfflineRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()
// 	log.InitNopLogger()

// 	ctx := context.TODO()
// 	p := services.OCPP16()
// 	req := newTestEquipOnlineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String())

// 	equipId, err := OnlineRequestWithGeneric(ctx, req)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, equipId)

// 	offreq := newTestEquipOfflineRequest(services.TestSN, p, services.TestAccessPod, id.Next().String(), EOF)

// 	err = OfflineRequestWithGeneric(ctx, offreq)

// 	assert.Nil(t, err)

// }

// func newTestEquipOfflineRequest(sn string, p *services.Protocol, pod, msgID, reason string) *equipOfflineRequest {
// 	return NewEquipOfflineRequest(sn, p, pod, msgID, reason)
// }