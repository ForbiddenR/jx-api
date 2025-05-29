package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type DiagnosticsStatusNotificationType int

const (
	DiagnosticsStatusNotificationTypeIdle         DiagnosticsStatusNotificationType = 1
	DiagnosticsStatusNotificationTypeUploadFailed DiagnosticsStatusNotificationType = 2
	DiagnosticsStatusNotificationTypeUploaded     DiagnosticsStatusNotificationType = 3
	DiagnosticsStatusNotificationTypeUploading    DiagnosticsStatusNotificationType = 4
)

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

func NewEquipDiagnosticsStatusNotificationRequestOCPP16(sn, pod, msgID string, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			Status: status,
		},
	}
	return req
}

func NewEquipDiagnosticsStatusNotificationRequest(sn, pod, msgID string, p *services.Protocol, requestId int64, status DiagnosticsStatusNotificationType) *equipDiagnosticsStatusNotificationRequest {
	req := &equipDiagnosticsStatusNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipDiagnosticsStatusNotificationRequestDetail{
			RequestId: &requestId,
			Status:    status,
		},
	}
	return req
}

// var _ services.Response = &equipDiagnosticsStatusNotificationResponse{}

// type equipDiagnosticsStatusNotificationResponse struct {
// 	api.Response
// 	Data *equipDiagnosticsStatusNotificationResponseDetail `json:"data"`
// }

// func (resp *equipDiagnosticsStatusNotificationResponse) GetStatus() int {
// 	return resp.Status
// }

// func (resp *equipDiagnosticsStatusNotificationResponse) GetMsg() string {
// 	return resp.Msg
// }

// type equipDiagnosticsStatusNotificationResponseDetail struct {
// }

func DiagnosticsStatusNotificationRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
