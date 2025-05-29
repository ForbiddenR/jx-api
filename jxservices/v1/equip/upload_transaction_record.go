package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipUploadTransactionRecordRequest struct {
	services.Base
	Data *equipUploadTransactionRecordRequestData `json:"data"`
}

type equipUploadTransactionRecordRequestData struct {
}

func (r *equipUploadTransactionRecordRequest) GetName() services.Request2ServicesNameType {
	return services.UploadTransactionRecord
}

func (r *equipUploadTransactionRecordRequest) TraceId() string {
	return r.MsgID
}

func (r *equipUploadTransactionRecordRequest) IsCallback() bool {
	return false
}

func NewUploadTransactionRecordRequest(sn, pod, msgId string, p *services.Protocol) *equipUploadTransactionRecordRequest {
	return &equipUploadTransactionRecordRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipUploadTransactionRecordRequestData{},
	}
}

// type equipUploadTransactionRecordResponse struct {
// 	api.Response
// }

// func (r *equipUploadTransactionRecordResponse) GetMsg() string {
// 	return r.Msg
// }

// func (r *equipUploadTransactionRecordResponse) GetStatus() int {
// 	return r.Status
// }

func UplaodTransactionRecordRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
