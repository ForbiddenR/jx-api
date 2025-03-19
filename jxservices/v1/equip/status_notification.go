package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipStatusNotificationRequest{}

type equipStatusNotificationRequest struct {
	services.Base
	Data *equipStatusNotificationRequestDetail `json:"data"`
}

type equipStatusNotificationRequestDetail struct {
	EvseSerial      *int    `json:"evseSerial,omitempty"`
	ConnectorSerial string  `json:"connectorSerial"`
	Status          int     `json:"status"`
	ErrorCode       *string `json:"errorCode,omitempty"`
	VendorErrorCode *string `json:"vendorErrorCode,omitempty"`
	Timestamp       int64   `json:"timestamp"`
}

type StatusNotificationRequestConfig struct {
	ConnectorSerial string
	Status          int
	Timestamp       int64
}

func NewEquipStatusNotificationRequestOCPP16(sn, id, pod, msgId string, connectorId string, status int, errorCode StatusNotificationErrorCodeEnum, timestamp int64) *equipStatusNotificationRequest {
	ec := string(errorCode)
	return &equipStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    services.OCPP16(),
			Category:    services.StatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: connectorId,
			Status:          status,
			ErrorCode:       &ec,
			Timestamp:       timestamp,
		},
	}
}

func NewStatusNotification(base services.Base, config *StatusNotificationRequestConfig) *equipStatusNotificationRequest {
	req := &equipStatusNotificationRequest{
		Base: base,
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: config.ConnectorSerial,
			Status:          config.Status,
			Timestamp:       config.Timestamp,
		},
	}
	return req
}

func NewEquipStatusNotificationRequest(sn, id, pod, msgId string, p *services.Protocol, connectorId string, status int, timestamp int64) *equipStatusNotificationRequest {
	return &equipStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.StatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipStatusNotificationRequestDetail{
			ConnectorSerial: connectorId,
			Status:          status,
			Timestamp:       timestamp,
		},
	}
}

func (equipStatusNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.StatusNotification
}

func (e *equipStatusNotificationRequest) TraceId() string {
	return e.MsgID
}

func (equipStatusNotificationRequest) IsCallback() bool {
	return false
}

type ConnectorStatusTypeEnum int

const (
	ConnectorStatusUnavailable ConnectorStatusTypeEnum = iota
	ConnectorStatusAvailable
	ConnectorStatusPreparing
	ConnectorStatusCharging
	ConnectorStatusSuspendedEV
	ConnectorStatusSuspendedEVSE
	ConnectorStatusFinishing
	ConnectorStatusFaulted
	ConnectorStatusReserved
	ConnectorStatusOccupied
)

type ConnectorStatus201TypeEnum int

const (
	ConnectorStatus201Unavailable   ConnectorStatus201TypeEnum = iota // 不可用
	ConnectorStatus201Available                                       // 空闲可用
	ConnectorStatus201Occupied                                        // 占用
	ConnectorStatus201Reserved                                        // 预约
	ConnectorStatus201Faulted                                         // 故障
	ConnectorStatus201Preparing                                       // 准备中
	ConnectorStatus201Charging                                        // 充电中
	ConnectorStatus201SuspendedEV                                     // 车端挂起，不输入电能
	ConnectorStatus201SuspendedEVSE                                   // 桩端挂起，不输出电能
	ConnectorStatus201Finishing                                       // 结束中
)

type StatusNotificationErrorCodeEnum string

const (
	StatusNotificationErrorCodeNoError              StatusNotificationErrorCodeEnum = "NoError"
	StatusNotificationErrorCodeConnectorLockFailure StatusNotificationErrorCodeEnum = "ConnectorLockFailure"
	StatusNotificationErrorCodeEVCommunicationError StatusNotificationErrorCodeEnum = "EVCommunicationError"
	StatusNotificationErrorCodeGroundFailure        StatusNotificationErrorCodeEnum = "GroundFailure"
	StatusNotificationErrorCodeHighTemperature      StatusNotificationErrorCodeEnum = "HighTemperature"
	StatusNotificationErrorCodeInternalError        StatusNotificationErrorCodeEnum = "InternalError"
	StatusNotificationErrorCodeLocalListConflict    StatusNotificationErrorCodeEnum = "LocalListConflict"
	StatusNotificationErrorCodeOverCurrentFailure   StatusNotificationErrorCodeEnum = "OverCurrentFailure"
	StatusNotificationErrorCodeOverVoltage          StatusNotificationErrorCodeEnum = "OverVoltage"
	StatusNotificationErrorCodePowerMeterFailure    StatusNotificationErrorCodeEnum = "PowerMeterFailure"
	StatusNotificationErrorCodePowerSwitchFailure   StatusNotificationErrorCodeEnum = "PowerSwitchFailure"
	StatusNotificationErrorCodeReaderFailure        StatusNotificationErrorCodeEnum = "ReaderFailure"
	StatusNotificationErrorCodeResetFailure         StatusNotificationErrorCodeEnum = "ResetFailure"
	StatusNotificationErrorCodeUnderVoltage         StatusNotificationErrorCodeEnum = "UnderVoltage"
	StatusNotificationErrorCodeWeakSignal           StatusNotificationErrorCodeEnum = "WeakSignal"
	StatusNotificationErrorCodeOtherError           StatusNotificationErrorCodeEnum = "OtherError"
)

var _ services.Response = &equipStatusNotificationResponse{}

type equipStatusNotificationResponse struct {
	api.Response
}

func (resp *equipStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

func StatusNotificationRequest(ctx context.Context, req *equipStatusNotificationRequest) error {
	// // header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, &equipStatusNotificationResponse{})
}
