package jxservices

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	"github.com/ForbiddenR/jxapi/jxutils/store"
)

type callbackGenerator func(base Base, err *apierrors.CallbackError) Request

var UnsupportedFeatures = store.NewReceptacle[callbackGenerator]()

// These constants are usually used in services package many times.
const (
	Equip           = "ac"
	Equipment       = "equip"
	Callback        = "callback"
	QueuePrefix     = "mq_services_"
	TestConnectorId = "1"
	TestSN          = "JK000000006"
	TestAccessPod   = "jx-acos-0"
	CallbackSuffix  = "Callback"
)

// Define some default status fields in callback.
const (
	CallbackError   = -1
	Successful      = 0
	Failed          = 1
	Unsupported     = 2
	PriceSchemeLost = 3
)

// These constants are associated with the topics of mqtt.
// todo Some customize feature related to DataTransfer and sending to charging point will be actually announced.
const (
	// Core
	DataTransferFeatureName           = "dataTransfer"
	ChangeAvailabilityFeatureName     = "changeAvailability"
	ChangeConfigurationFeatureName    = "setVariables"
	GetConfigurationFeatureName       = "getVariables"
	GetBaseReportFeatureName          = "getBaseReport"
	ClearCacheFeatureName             = "clearCache"
	RemoteStartTransactionFeatureName = "remoteStartTransaction"
	RemoteStopTransactionFeatureName  = "remoteStopTransaction"
	ResetFeatureName                  = "reset"
	UnlockConnectorFeatureName        = "unlockConnector"
	// Firmware
	GetDiagnosticsFeatureName = "getDiagnostics"
	UpdateFirmwareFeatureName = "pushFirmware"
	// LocalAuth
	GetLocalListVersionFeatureName = "getLocalListVersion"
	SendLocalListFeatureName       = "setLocalAuthorizeList"
	// RemoteTrigger
	TriggerMessageFeatureName         = "triggerMessage"
	CallStatusNotificationFeatureName = "callStatusNotification"
	// Reservation
	CancelReservationFeatureName = "cancelReservation"
	ReserveNowFeatureName        = "reserveNow"
	// SmartCharging
	ClearChargingProfileFeatureName = "clearChargingProfile"
	GetCompositeScheduleFeatureName = "detCompositeSchedule"
	SetChargingProfileFeatureName   = "setChargingProfile"
	// Customize Features
	SetChargingTimerFeatureName      = "setChargingTimer"
	SetLoadBalanceFeatureName        = "setLoadBalance"
	SetFactoryResetFeatureName       = "setFactoryReset"
	CloseFeatureName                 = "close"
	SetIntellectChargeFeatureName    = "setIntellectCharge"
	CancelIntellectChargeFeatureName = "cancelIntellectCharge"
	SetPriceSchemeFeatureName        = "setPriceScheme"
	// Customized Features about 104 protocol
	SendQRCodeFeatureName              = "sendQRCode"
	GetIntellectChargeFeatureName      = "getIntellectCharge"
	OffPeakChargeFeatureName           = "offPeakCharging"
	RequestStartDischargingFeatureName = "remoteStartDischarging"
	RequestStopDischargingFeatureName  = "remoteStopDischarging"
	// New for Yunkuaichong 104 protocol
	TemplateFetchFeatureName                  = "templateFetch"
	ReadCurrentMonitorFeatureName             = "readCurrentMonitor"
	RemoteUpdateAccountBalanceFeatureName     = "remoteUpdateAccountBalance"
	ClearOfflineCardsFeatureName              = "clearOfflineCards"
	QueryOfflineCardsFeatureName              = "queryOfflineCards"
	SetParametersFeatureName                  = "setParameters"
	RemoteControlParkingLockFeatureName       = "remoteControlParkingLock"
	RemoteStartParallelTransactionFeatureName = "startParallelTransaction"
)

type Request2ServicesNameType string

const (
	Authorize                     Request2ServicesNameType = "authorize"
	BootNotification              Request2ServicesNameType = "bootNotification"
	ClearCache                    Request2ServicesNameType = "clearCache"
	DataTransfer                  Request2ServicesNameType = "dataTransfer"
	GetBaseReport                 Request2ServicesNameType = "getBaseReport"
	GetConfiguration              Request2ServicesNameType = "getVariables"
	MeterValues                   Request2ServicesNameType = "meterValues"
	UpdateTransaction             Request2ServicesNameType = "updateTransaction"
	Online                        Request2ServicesNameType = "equipOnline"
	Offline                       Request2ServicesNameType = "equipOffline"
	Register                      Request2ServicesNameType = "equipRegister"
	StatusNotification            Request2ServicesNameType = "statusNotification"
	StartTransaction              Request2ServicesNameType = "startTransaction"
	StopTransaction               Request2ServicesNameType = "stopTransaction"
	RemoteStartTransaction        Request2ServicesNameType = "remoteStartTransaction"
	RemoteStopTransaction         Request2ServicesNameType = "remoteStopTransaction"
	Reset                         Request2ServicesNameType = "reset"
	ReservationStatusNotification Request2ServicesNameType = "reservationStatusNotification"
	ChangeConfiguration           Request2ServicesNameType = "setVariables"
	SendLocalList                 Request2ServicesNameType = "setLocalAuthorizeList"
	SetChargingTimer              Request2ServicesNameType = "setChargingTimer"
	ChargingTimerNotification     Request2ServicesNameType = "timingStatusNotification"
	// ExpiredChargingTimer          Request2ServicesNameType = "expiredChargingTimer"
	UpdateFirmware                Request2ServicesNameType = "pushFirmware"
	FirmwareStatusNotification    Request2ServicesNameType = "firmwareStatusNotification"
	CallStatusNotification        Request2ServicesNameType = "callStatusNotification"
	GetDiagnostics                Request2ServicesNameType = "getDiagnostics"
	DiagnosticsStatusNotification Request2ServicesNameType = "diagnosticsStatusNotification"
	SetLoadBalance                Request2ServicesNameType = "setLoadBalance"
	SetFactoryReset               Request2ServicesNameType = "setFactoryReset"
	NotifyEvent                   Request2ServicesNameType = "notifyEvent"
	NotifyReport                  Request2ServicesNameType = "notifyReport"
	CancelReservation             Request2ServicesNameType = "cancelReservation"
	ReserveNow                    Request2ServicesNameType = "reserveNow"
	QRCode                        Request2ServicesNameType = "qrcode"
	SendQRCode                    Request2ServicesNameType = "sendQRCode"
	SetIntellectCharge            Request2ServicesNameType = "setIntellectCharge"
	CancelIntellectCharge         Request2ServicesNameType = "cancelIntellectCharge"
	SetPriceScheme                Request2ServicesNameType = "setPriceScheme"
	TriggerMessage                Request2ServicesNameType = "triggerMessage"
	//TriggerMessage             Request2ServicesNameType = "callStatusNotification"
	// TODO: the name of this variable has not been defined.
	ChargeEncryInfoNotification Request2ServicesNameType = "chargeEncryInfoNotification"
	SetChargingProfile          Request2ServicesNameType = ""
	ClearChargingProfile        Request2ServicesNameType = ""
	BMSInfo                     Request2ServicesNameType = "bmsInfo"
	BMSLimit                    Request2ServicesNameType = "bmsLimit"
	Login                       Request2ServicesNameType = "equipLogin"
	GetIntellectCharge          Request2ServicesNameType = "getIntellectCharge"
	OffPeakCharge               Request2ServicesNameType = "offPeakCharging"
	// defined for Wukong
	UpdatedParaReport                  Request2ServicesNameType = "updatedParaReport"
	ConnectorPositionReport            Request2ServicesNameType = "connectorPositionReport"
	RequestStartDischargingTransaction Request2ServicesNameType = "remoteStartDischarging"
	RequestStopDischargingTransaction  Request2ServicesNameType = "remoteStopDischarging"
)

// FirstUpper is only for the interfaces having a regular category.
func (r Request2ServicesNameType) FirstUpper() string {
	s := r.String()
	return strings.ToUpper(s[:1]) + s[1:]
}

func (r Request2ServicesNameType) String() string {
	return string(r)
}

// func (r Request2ServicesNameType) Split() []string {
// 	switch r {
// 	case UpdateFirmware:
// 		return []string{"push", "firmware", "equipment"}
// 	case FirmwareStatusNotification:
// 		return []string{"push", "firmware", "notification"}
// 	case RemoteStartTransaction:
// 		return []string{"remote", "Start"}
// 	case RemoteStopTransaction:
// 		return []string{"remote", "stop"}
// 	case SendLocalList:
// 		return []string{"set", "local", "authorize"}
// 	case StartTransaction:
// 		return []string{"start", "transaction"}
// 	case StopTransaction:
// 		return []string{"stop", "transaction"}
// 	}
// 	for i := range len(r.String()) {
// 		str := r.String()[i : i+1]
// 		if str == strings.ToUpper(str) {
// 			switch r.String()[:i] {
// 			case "equip":
// 				return []string{strings.ToLower(r.String()[i:])}
// 			case "authorize":
// 				return []string{r.String()[:i]}
// 			default:
// 				return r.SplitName()
// 			}
// 		}
// 	}
// 	return []string{r.String()}
// }

// // SplitName will be used by the function above to parser all the regular attributes.
// func (r Request2ServicesNameType) SplitName() []string {
// 	var head, tail int
// 	var result []string
// 	for tail = 0; tail < len(r.String()); tail++ {
// 		str := r.String()[tail : tail+1]
// 		if str == strings.ToUpper(str) && tail != 0 {
// 			result = append(result, strings.ToLower(r.String()[head:tail]))
// 			head = tail
// 		}
// 	}
// 	if head < tail {
// 		result = append(result, strings.ToLower(r.String()[head:tail]))
// 	}
// 	return result
// }

func (r Request2ServicesNameType) GetCallbackCategory() string {
	return r.FirstUpper() + CallbackSuffix
}

type Base struct {
	EquipmentSn string    `json:"equipmentSn"`
	EquipmentId string    `json:"equipmentId"`
	Protocol    *Protocol `json:"protocol"`
	Category    string    `json:"category"`
	AccessPod   string    `json:"accessPod"`
	MsgID       string    `json:"msgId"`
}

type BaseConfig struct {
	equipmentSn string
	equipmentId string
	protocol    *Protocol
	category    string
	accessPod   string
	msgID       string
}

func NewBaseConfig() *BaseConfig {
	return &BaseConfig{}
}

func (b *BaseConfig) EquipmentSn(sn string) *BaseConfig {
	b.equipmentSn = sn
	return b
}

func (b *BaseConfig) EquipmentId(id string) *BaseConfig {
	b.equipmentId = id
	return b
}

func (b *BaseConfig) Protocol(p *Protocol) *BaseConfig {
	b.protocol = p
	return b
}

func (b *BaseConfig) Category(cate string) *BaseConfig {
	b.category = cate
	return b
}

// TODO: use a common string rather than a string with type Request2ServicesNameType.
func (b *BaseConfig) Categories(kind Request2ServicesNameType, isCallback bool) *BaseConfig {
	if !isCallback {
		b.category = kind.FirstUpper()
	} else {
		b.category = kind.FirstUpper() + CallbackSuffix
	}
	return b
}

func (b *BaseConfig) Hostname(hostname string) *BaseConfig {
	b.accessPod = hostname
	return b
}

func (b *BaseConfig) MsgID(msgID string) *BaseConfig {
	b.msgID = msgID
	return b
}

func (b *BaseConfig) Build() Base {
	return Base{
		EquipmentSn: b.equipmentSn,
		EquipmentId: b.equipmentId,
		Protocol:    b.protocol,
		Category:    b.category,
		AccessPod:   b.accessPod,
		MsgID:       b.msgID,
	}
}

type Protocol struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewIEC104Protocol(version string) *Protocol {
	return &Protocol{
		Name:    "IEC104",
		Version: version,
	}
}

func (p *Protocol) String() string {
	return p.Name + "" + p.Version
}

func (p *Protocol) Equal(p2 *Protocol) bool {
	return p.Name == p2.Name && p.Version == p2.Version
}

func (p *Protocol) UnmarshalJSON(data []byte) error {
	var v struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.Name {
	case "OCPP":
		if v.Version != "1.6" && v.Version != "2.0.1" {
			return errors.New("invalid OCPP version: " + v.Version)
		}
	case "IEC104":
		if v.Version != "0.1" && v.Version != "0.2" && v.Version != "0.3" && v.Version != "0.4" && v.Version != "0.5" && v.Version != "0.6" {
			return errors.New("invalid IEC104 version: " + v.Version)
		}
	default:
		return errors.New("invalid protocol name: " + v.Name)
	}
	p.Name = v.Name
	p.Version = v.Version

	return nil
}

var ocpp16p = &Protocol{Name: "OCPP", Version: "1.6"}
var ocpp201p = &Protocol{Name: "OCPP", Version: "2.0.1"}
var iec001 = &Protocol{Name: "IEC104", Version: "0.1"}
var iec002 = &Protocol{Name: "IEC104", Version: "0.2"}
var iec003 = &Protocol{Name: "IEC104", Version: "0.3"}
var iec004 = &Protocol{Name: "IEC104", Version: "0.4"}
var iec005 = &Protocol{Name: "IEC104", Version: "0.5"}

func OCPP16() *Protocol {
	return ocpp16p
}

func OCPP201() *Protocol {
	return ocpp201p
}

func IEC001() *Protocol {
	return iec001
}

func IEC002() *Protocol {
	return iec002
}

func IEC003() *Protocol {
	return iec003
}

func IEC004() *Protocol {
	return iec004
}

func IEC005() *Protocol {
	return iec005
}

// CB includes callback information
type CB struct {
	Status int                              `json:"status"`
	Code   *apierrors.CallbackErrorCodeType `json:"code,omitempty"`
	Msg    *string                          `json:"msg,omitempty"`
}

func NewCB(status int) CB {
	return CB{Status: status}
}

func NewCBError(err *apierrors.CallbackError) CB {
	code := err.Code()
	msg := err.Error()
	return CB{Status: CallbackError, Code: &code, Msg: &msg}
}

// boxing a callback request, the function can be invoked indirectly
// whichCallbackErr recognizes the type of the error, returning a corresponding callback error
func getCallbackError(clientId string, command string, err *apierrors.Error) *apierrors.CallbackError {
	switch err.Code {
	case apierrors.NotSupported:
		return apierrors.NewCallbackErrorNotSupported(clientId, command)
	case apierrors.NotImplemented:
		return apierrors.NewCallbackErrorNotImplemented(clientId, command)
	case apierrors.InternalError:
		return apierrors.NewCallbackErrorInternalError(clientId, command, err.Description)
	case apierrors.SecurityError:
		return apierrors.NewCallbackErrorSecurityError(clientId, command, err.Description)
	case apierrors.ProtocolError,
		apierrors.FormationViolation,
		apierrors.PropertyConstraintViolation,
		apierrors.TypeConstraintViolation:
		return apierrors.NewCallbackErrorPayloadError(clientId, command, err.Description)
	default:
		return apierrors.NewCallbackErrorGenericError(clientId, command, err.Description)
	}
}

// GetProperCallbackError turns an entering error into callback error
func GetProperCallbackError(clientId string, command string, err error) *apierrors.CallbackError {
	return typeCheckOr(err, func(e error) *apierrors.CallbackError {
		if err, ok := err.(*apierrors.Error); ok {
			return getCallbackError(clientId, command, err)
		}
		return apierrors.NewCallbackErrorOffline(clientId, command)
	}, err)
}

func typeCheckOr[T, B any](input any, f func(B) T, arg B) T {
	if t, ok := input.(T); ok {
		return t
	}
	return f(arg)
}

func GetSimpleURL(req Request) string {
	return api.ServicesUrl + Equip + "/" + req.GetName().String()
}

func GetCallbackURL(req Request) string {
	return api.ServicesUrl + Equip + "/" + Callback + "/" + req.GetName().String() + CallbackSuffix
}

func getURI(req Request) string {
	if req.IsCallback() {
		return fmt.Sprintf("%s/%s/%s%s", Equip, Callback, req.GetName(), CallbackSuffix)
	}
	return fmt.Sprintf("%s/%s", Equip, req.GetName())
}

func Transport(ctx context.Context, req Request) error {
	uri := getURI(req)
	result := api.ServiceClient.
		Post().
		RequestURI(uri).
		Body(req).
		SetHeader(api.TraceId, req.TraceId()).
		SetParam().
		Do(ctx)
	if result.Error() != nil {
		request, _ := json.Marshal(req)
		rBytes, err := result.Raw()
		return apierrors.GetFailedResponseUnmarshalError(uri, request, rBytes, err)
	}
	resp := &api.Response{}
	err := result.Into(resp)
	if err != nil {
		return err
	}
	if resp.Status == 1 {
		return errors.New(resp.Msg)
	}
	return err
}

func RequestWithoutResponse[T Response](ctx context.Context, req Request, url string, t T) (err error) {
	message, err := api.SendRequest(ctx, url, req, api.WithHeader(api.TraceId, req.TraceId()))
	if err != nil {
		return
	}
	err = json.Unmarshal(message, t)
	if err != nil {
		request, _ := json.Marshal(req)
		return apierrors.GetFailedResponseUnmarshalError(url, request, message, err)
	}

	// check whether it get an error from the services
	if t.GetStatus() == 1 {
		return errors.New(t.GetMsg())
	}
	return err
}

func RequestWithResponse[T Response](ctx context.Context, req Request, url string, t T) (resp T, err error) {
	message, err := api.SendRequest(ctx, url, req, api.WithHeader(api.TraceId, req.TraceId()))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(message, t)
	if err != nil {
		// The marshaling function has been verified before.
		request, _ := json.Marshal(req)
		return resp, apierrors.GetFailedResponseUnmarshalError(url, request, message, err)
	}
	resp = t
	// check whether it get an error from the services
	if resp.GetStatus() == 1 {
		return resp, errors.New(resp.GetMsg())
	}
	return resp, err
}
