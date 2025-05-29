package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUploadParkingLockDataRequest struct {
	services.Base
	Data *equipUploadParkingLockDataRequestDetail `json:"data"`
}

type equipUploadParkingLockDataRequestDetail struct {
	ConnectorId   string `json:"connectorSerial"`
	LockStatus    uint8  `json:"lockStatus"`
	VehicleStatus uint8  `json:"vehicleStatus"`
	Power         uint8  `json:"power"`
	AlarmStatus   uint8  `json:"alarmStatus"`
}

func (r equipUploadParkingLockDataRequest) GetName() services.Request2ServicesNameType {
	return services.UploadParkingLockData
}

func (r *equipUploadParkingLockDataRequest) TraceId() string {
	return r.MsgID
}

func (equipUploadParkingLockDataRequest) IsCallback() bool {
	return false
}

func NewEquipUploadParkingLockDataRequest(sn, pod, msgId string, p *services.Protocol, connectorId string, lockStatus, vehicleStatus, power, alarmStatus uint8) *equipUploadParkingLockDataRequest {
	return &equipUploadParkingLockDataRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipUploadParkingLockDataRequestDetail{
			ConnectorId:   connectorId,
			LockStatus:    lockStatus,
			VehicleStatus: vehicleStatus,
			Power:         power,
			AlarmStatus:   alarmStatus,
		},
	}
}

type equipUploadParkingLockDataResponse struct {
	api.Response
	Data *equipUploadParkingLockDataResponseDetail `json:"data"`
}

type equipUploadParkingLockDataResponseDetail struct {
	Authenticated bool `json:"authenticated"`
}

func (resp *equipUploadParkingLockDataResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipUploadParkingLockDataResponse) GetMsg() string {
	return resp.Msg
}

func UploadParkingLockDataRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
