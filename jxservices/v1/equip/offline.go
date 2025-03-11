package equip

import (
	"context"
	"strings"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type OfflineReason string

const (
	// Cannot received any message in the stated period.
	Timeout = "Timeout disconnected"
	// Be closed by charger station.
	EOF = "Active disconnected"
	// Closing the ws connection.
	Initiative = "Passive disconnected"
)

func GetOfflineReason(err error) string {
	if err == nil {
		return Initiative
	}
	if strings.Contains(strings.ToLower(err.Error()), "eof") {
		return EOF
	}
	return Timeout
}

var _ services.Request = &equipOfflineRequest{}

type equipOfflineRequest struct {
	services.Base
	Data *equipOfflineRequestDetail `json:"data"`
}

type equipOfflineRequestDetail struct {
	OfflineReason string `json:"offlineReason"`
}

func NewEquipOfflineRequest(sn, id, pod, msgID string, protocol *services.Protocol, reason string) *equipOfflineRequest {
	return &equipOfflineRequest{
		Base: services.Base{
			Category:    services.Offline.FirstUpper(),
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipOfflineRequestDetail{
			OfflineReason: reason,
		},
	}
}

func (equipOfflineRequest) GetName() services.Request2ServicesNameType {
	return services.Offline
}

func (e *equipOfflineRequest) TraceId() string {
	return e.MsgID
}

func (equipOfflineRequest) IsCallback() bool {
	return false
}

var _ services.Response = &equipOfflineResponse{}

type equipOfflineResponse struct {
	api.Response
}

func (resp *equipOfflineResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipOfflineResponse) GetMsg() string {
	return resp.Msg
}

func OfflineRequest(ctx context.Context, req *equipOfflineRequest) error {
	header := services.GetSimpleHeaderValue(services.Offline)
	url := services.GetSimpleURL(req)
	return services.RequestWithoutResponse(ctx, req, url, header, &equipOfflineResponse{})
}
