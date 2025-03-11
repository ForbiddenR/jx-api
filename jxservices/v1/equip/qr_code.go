package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipQRCodeRequest{}

type equipQRCodeRequest struct {
	services.Base
	Data *equipQRCoeRequestDetail `json:"data"`
}

type equipQRCoeRequestDetail struct {
	EvseSerial  *string `json:"evseSerial"`
	ConnectorId string  `json:"connectorSerial"`
}

func NewEquipQRCodeRequest(sn, id, pod, msgID string, protocol *services.Protocol, connectorId string) *equipQRCodeRequest {
	return &equipQRCodeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			Category:    services.QRCode.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipQRCoeRequestDetail{
			ConnectorId: connectorId,
		},
	}
}

func (q *equipQRCodeRequest) GetName() services.Request2ServicesNameType {
	return services.QRCode
}

func (q *equipQRCodeRequest) TraceId() string {
	return q.MsgID
}

func (q *equipQRCodeRequest) IsCallback() bool {
	return false
}

var _ services.Response = &equipQRCodeResponse{}

type equipQRCodeResponse struct {
	api.Response
	Data *equipQRCodeResponseDetail `json:"data"`
}

type equipQRCodeResponseDetail struct {
	Qrcode string `json:"qrcode"`
}

func (q *equipQRCodeResponse) GetStatus() int {
	return q.Status
}

func (q *equipQRCodeResponse) GetMsg() string {
	return q.Msg
}

func QRCodeRequest(ctx context.Context, req services.Request) (*equipQRCodeResponse, error) {
	header := services.GetSimpleHeaderValue(services.QRCode)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, header, &equipQRCodeResponse{})
}
