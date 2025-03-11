package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type DiagnosticsStatusNotificationType int

const (
	DiagnosticsStatusNotificationTypeIdle         DiagnosticsStatusNotificationType = 1
	DiagnosticsStatusNotificationTypeUploadFailed DiagnosticsStatusNotificationType = 2
	DiagnosticsStatusNotificationTypeUploaded     DiagnosticsStatusNotificationType = 3
	DiagnosticsStatusNotificationTypeUploading    DiagnosticsStatusNotificationType = 4
)

// func OCPP16GetDiagnosticsStatus(status protocol.DiagnosticsStatusNotificationJsonStatus) DiagnosticsStatusNotificationType {
// 	var result DiagnosticsStatusNotificationType
// 	switch status {
// 	case protocol.DiagnosticsStatusNotificationJsonStatusIdle:
// 		result = DiagnosticsStatusNotificationTypeIdle
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploadFailed:
// 		result = DiagnosticsStatusNotificationTypeUploadFailed
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploaded:
// 		result = DiagnosticsStatusNotificationTypeUploaded
// 	case protocol.DiagnosticsStatusNotificationJsonStatusUploading:
// 		result = DiagnosticsStatusNotificationTypeUploading
// 	}
// 	return result
// }

var _ services.Request = &equipDiagnosticsStatusNotificationRequest{}

type equipDiagnosticsStatusNotificationRequest struct {
	services.Base
	Data *equipDiagnosticsStatusNotificationRequestDetail `json:"data"`
}

type equipDiagnosticsStatusNotificationRequestDetail struct {
	RequestId *int64                            `json:"requestId,omitempty"`
	Status    DiagnosticsStatusNotificationType `json:"status"`
}

func (equipDiagnosticsStatusNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.DiagnosticsStatusNotification
}

func (e *equipDiagnosticsStatusNotificationRequest) TraceId() string {
	return e.MsgID
}

func (equipDiagnosticsStatusNotificationRequest) IsCallback() bool {
	return false
}

func NewEquipDiagnosticsStatusNotificationRequestOCPP16(sn, id, pod, msgId string, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			Category:    services.DiagnosticsStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			Status: status,
		},
	}
	return req
}

func NewEquipDiagnosticsStatusNotificationRequest(sn, id, pod, msgId string, p *services.Protocol, requestId int64, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.DiagnosticsStatusNotification.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			RequestId: &requestId,
			Status:    status,
		},
	}
	return req
}

var _ services.Response = &equipDiagnosticsStatusNotificationResponse{}

type equipDiagnosticsStatusNotificationResponse struct {
	api.Response
	Data *equipDiagnosticsStatusNotificationResponseDetail `json:"data"`
}

func (resp *equipDiagnosticsStatusNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipDiagnosticsStatusNotificationResponse) GetMsg() string {
	return resp.Msg
}

type equipDiagnosticsStatusNotificationResponseDetail struct {
}

func DiagnosticsStatusNotificationRequest(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.DiagnosticsStatusNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipDiagnosticsStatusNotificationResponse{})
}
