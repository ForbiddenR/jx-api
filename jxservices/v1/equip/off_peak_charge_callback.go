package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipOffPeakChargeCallbackRequest{}

type equipOffPeakChargeCallbackRequest struct {
	services.Base
	Callback services.CB `json:"callback"`
}

func (*equipOffPeakChargeCallbackRequest) GetName() services.Request2ServicesNameType {
	return services.OffPeakCharge
}

func (e *equipOffPeakChargeCallbackRequest) TraceId() string {
	return e.MsgID
}

func (equipOffPeakChargeCallbackRequest) IsCallback() bool {
	return true
}

func NewEquipOffPeakChargeCallbackRequest(base services.Base, status int) *equipOffPeakChargeCallbackRequest {
	req := &equipOffPeakChargeCallbackRequest{
		Base:     base,
		Callback: services.NewCB(status),
	}
	return req
}

func NewEquipOffPeakChargeCallbackRequestError(base services.Base, err *apierrors.CallbackError) *equipOffPeakChargeCallbackRequest {
	return &equipOffPeakChargeCallbackRequest{
		Base:     base,
		Callback: services.NewCBError(err),
	}
}

var _ services.Response = &equipOffPeakChargeCallbackResponse{}

type equipOffPeakChargeCallbackResponse struct {
	api.Response
}

func OffPeakChargeCallbackRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}