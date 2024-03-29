package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

type equipDataTransferRequest struct {
	services.Base
	Callback *equipDataTransferRequestDetail `json:"data"`
}

type equipDataTransferRequestDetail struct {
}

func (r equipDataTransferRequest) GetName() string {
	return services.DataTransfer.String()
}

func NewEquipDataTransferRequest() *equipDataTransferRequest {
	req := &equipDataTransferRequest{}
	return req
}

type equipDataTransferResponse struct {
	api.Response
	Data *equipDataTransferResponseDetail `json:"data"`
}

type equipDataTransferResponseDetail struct {
}

func (resp *equipDataTransferResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipDataTransferResponse) GetMsg() string {
	return resp.Msg
}

func DataTransferRequest(ctx context.Context, req services.CallbackRequest) error {
	header := services.GetCallbackHeaderValue(services.DataTransfer)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipDataTransferResponse{})
}
