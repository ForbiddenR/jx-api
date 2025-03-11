package v1

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	api "github.com/ForbiddenR/jxapi"
	esam "github.com/ForbiddenR/jxapi/jxesam"
	utils "github.com/ForbiddenR/jxapi/jxutils"
)

const readWait = 80 * time.Second

type accessVerifyRequest struct {
	EquipmentSn      string  `json:"equipmentSn"`
	Protocol         string  `json:"deviceProtocol,omitempty"`
	ProtocolVersion  string  `json:"deviceProtocolVersion,omitempty"`
	RequestPort      string  `json:"requestPort"`
	RemoteAddress    *string `json:"remoteAddress"`
	CertSerialNumber *string `json:"certSerialNumber"`
	AccountPassword  *string `json:"accountPassword"`
}

func NewAccessVerifyRequest(sn, requestPort, protocol, protocolVersion string) *accessVerifyRequest {
	return &accessVerifyRequest{
		EquipmentSn:     sn,
		Protocol:        protocol,
		ProtocolVersion: protocolVersion,
		RequestPort:     requestPort,
	}
}

type accessVerifyResponse struct {
	api.Response
	Data *accessVerifyResponseData `json:"data"`
}

type accessVerifyResponseData struct {
	ID                  string         `json:"id"`
	BaseUrl             string         `json:"baseUrl"`
	HearbeatInterval    int            `json:"keepalive"`
	TransactionInterval int            `json:"transactionInterval"`
	Registered          bool           `json:"registered"`
	ReadWait            utils.Duration `json:"readWait"`
}

func AccessVerifyRequest(ctx context.Context, ticket string, traceId string, request *accessVerifyRequest) (*accessVerifyResponse, error) {
	header := map[string]string{api.Perms: esam.Access.Perm(), esam.TicketKey: ticket, api.TraceId: traceId}
	url := api.EsamUrl + esam.Equip + "/verify"
	resp, err := api.SendRequest(ctx, url, request, header)
	if err != nil {
		return nil, err
	}
	accessResponse := new(accessVerifyResponse)
	err = json.Unmarshal(resp, accessResponse)
	if err != nil {
		return nil, err
	}
	if accessResponse.Status == 1 {
		return accessResponse, errors.New(accessResponse.Msg)
	}
	if accessResponse.Data == nil {
		return accessResponse, errors.New("data is empty")
	}

	if accessResponse.Data.BaseUrl == "" {
		return accessResponse, errors.New("baseUrl is empty")
	}

	if accessResponse.Data.ReadWait == 0 {
		accessResponse.Data.ReadWait = utils.Duration(readWait)
	}

	return accessResponse, nil
}
