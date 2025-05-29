package equip

import (
	"context"
	"strings"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
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

// var _ services.Request = &equipOfflineRequest{}

type equipOfflineRequest struct {
	services.Base
	Data *equipOfflineRequestDetail `json:"data"`
}

type equipOfflineRequestDetail struct {
	OfflineReason string `json:"offlineReason"`
}

func NewEquipOfflineRequest(sn string, protocol *services.Protocol, pod, msgID string, reason string) *equipOfflineRequest {
	return &equipOfflineRequest{
		Base: services.Base{
			EquipmentSn: sn,
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

func OfflineRequest(ctx context.Context, req *equipOfflineRequest) error {
	return services.Transport(ctx, req)
}
