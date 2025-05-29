package equip

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/csms/jxeu-ocpp/internal/config"
// 	"gitee.com/csms/jxeu-ocpp/internal/errors"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api"
// 	"gitee.com/csms/jxeu-ocpp/pkg/api/services"
// 	"gitee.com/csms/jxeu-ocpp/pkg/ocpp1.6/protocol"
// 	"github.com/samber/lo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetDiagnosticsRequest(t *testing.T) {
// 	config.TestConfig()
// 	api.Init()

// 	ctx := context.TODO()
// 	// logOpt := &log.Options{}
// 	// logOpt.Development = true
// 	// log.InitLogger(logOpt)
// 	p := services.OCPP16()

// 	err := errors.NewCallbackErrorGenericError(services.TestSN, protocol.GetDiagnosticsFeatureName,
// 		"unsupported")

// 	req := []*equipGetDiagnosticsCallbackRequest{
// 		NewEquipGetDiagnosticsCallbackRequest(services.TestSN, services.TestAccessPod, "1", p, services.Successful),
// 		NewEquipGetDiagnosticsCallbackRequestError(services.TestSN, services.TestAccessPod, "1", p, err),
// 	}

// 	lo.ForEach(req, func(item *equipGetDiagnosticsCallbackRequest, _ int) {
// 		err := GetDiagnosticsCallbackRequestG(ctx, item)
// 		assert.Nil(t, err)
// 	})
// }