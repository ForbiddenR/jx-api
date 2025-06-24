package equip

const maxUint8 = ^uint8(0)

type IdTokenStatusTypeEnum int

const (
	Accepted IdTokenStatusTypeEnum = iota
	Blocked
	Expired
	Invalid
	ConcurrentTx
)

type CustomVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Component struct {
	Name string `json:"name"`
	Evse *EVSE  `json:"evse,omitempty"`
}

type EVSE struct {
	Id          string `json:"serial"`
	ConnectorId string `json:"connectorSerial"`
}

type Intellect struct {
	ID        int64  `json:"id"`
	LimitSoC  *int   `json:"limitSoC,omitempty"`
	LimitElec *int   `json:"limitElectricity,omitempty"`
	StopTime  *int64 `json:"stopTime,omitempty"`
}

type IdTokenTypeEnums int

const (
	IdTokenTypeEnumsLocalPlug          IdTokenTypeEnums = 0
	IdTokenTypeEnumsLocalAdmin         IdTokenTypeEnums = 1
	IdTokenTypeEnumsLocalIdentityCard  IdTokenTypeEnums = 2
	IdTokenTypeEnumsOnlineIdentityCard IdTokenTypeEnums = 3
	IdTokenTypeEnumsLocalWalletCard    IdTokenTypeEnums = 4
	IdTokenTypeEnumsLocalVIN           IdTokenTypeEnums = 5
	IdTokenTypeEnumsOnlineVIN          IdTokenTypeEnums = 6
	IdTokenTypeEnumsBluetooth          IdTokenTypeEnums = 7
	IdTokenTypeEnumsMAC                IdTokenTypeEnums = 8
	IdTokenTypeEnumsRemoteAdmin        IdTokenTypeEnums = 9
	IdTokenTypeEnumsRemoteUser         IdTokenTypeEnums = 10
	IdTokenTypeEnumsIntellect          IdTokenTypeEnums = 11
	IdTokenTypeEnumsUnknown            IdTokenTypeEnums = 98
	IdTokenTypeEnumsCentral            IdTokenTypeEnums = 99
	IdTokenTypeEnumsLocal              IdTokenTypeEnums = 12
	IdTokenTypeEnumsEMAID              IdTokenTypeEnums = 13
	IdTokenTypeEnumsOffPeak            IdTokenTypeEnums = 14
)

type AuthorizationMode int32

type IdTokenType struct {
	IdToken string            `json:"idToken"`
	Type    *IdTokenTypeEnums `json:"type,omitempty"`
}

// type IdTokenTypeEnum int

type IdTokenInfo struct {
	ExpiryDate    *int64                `json:"expiryDate,omitempty"`
	ParentIdToken *IdTokenType          `json:"parentIdToken,omitempty"`
	GroupIdToken  *IdTokenType          `json:"groupIdToken,omitempty"`
	Status        IdTokenStatusTypeEnum `json:"status"`
}

type VariableAttribute struct {
	Value    string     `json:"value"`
	Readonly Mutability `json:"mutability"`
}

type Mutability int

const (
	MutabilityReadOnly  Mutability = 0
	MutabilityReadWrite Mutability = 1
)

type MeterValue struct {
	Timestamp   int64                            `json:"timestamp"`
	SampleValue []MeterValueElemSampledValueElem `json:"sampleValue"`
}

type MeterValueElemSampledValueElem struct {
	// Context corresponds to the JSON schema field "context".
	Context MeterValueElemSampledValueElemContext `json:"context,omitempty" yaml:"context,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format *MeterValueElemSampledValueElemFormat `json:"format,omitempty" yaml:"format,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *MeterValueElemSampledValueElemLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand MeterValueElemSampledValueElemMeasurand `json:"measurand,omitempty" yaml:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *MeterValueElemSampledValueElemPhase `json:"phase,omitempty" yaml:"phase,omitempty"`

	// Unit corresponds to the JSON schema field "unit".
	Unit MeterValueElemSampledValueElemUnit `json:"unit,omitempty" yaml:"unit,omitempty"`

	// Value corresponds to the JSON schema field "value".
	Value string `json:"value" yaml:"value"`
}

type MeterValueElemSampledValueElemContext string

const MeterValueElemSampledValueElemContextInterruptionBegin MeterValueElemSampledValueElemContext = "Interruption.Begin"
const MeterValueElemSampledValueElemContextInterruptionEnd MeterValueElemSampledValueElemContext = "Interruption.End"
const MeterValueElemSampledValueElemContextOther MeterValueElemSampledValueElemContext = "Other"
const MeterValueElemSampledValueElemContextSampleClock MeterValueElemSampledValueElemContext = "Sample.Clock"
const MeterValueElemSampledValueElemContextSamplePeriodic MeterValueElemSampledValueElemContext = "Sample.Periodic"
const MeterValueElemSampledValueElemContextTransactionBegin MeterValueElemSampledValueElemContext = "Transaction.Begin"
const MeterValueElemSampledValueElemContextTransactionEnd MeterValueElemSampledValueElemContext = "Transaction.End"
const MeterValueElemSampledValueElemContextTrigger MeterValueElemSampledValueElemContext = "Trigger"

type MeterValueElemSampledValueElemFormat string

const MeterValueElemSampledValueElemFormatRaw MeterValueElemSampledValueElemFormat = "Raw"
const MeterValueElemSampledValueElemFormatSignedData MeterValueElemSampledValueElemFormat = "SignedData"

type MeterValueElemSampledValueElemLocation string

const MeterValueElemSampledValueElemLocationBody MeterValueElemSampledValueElemLocation = "Body"
const MeterValueElemSampledValueElemLocationCable MeterValueElemSampledValueElemLocation = "Cable"
const MeterValueElemSampledValueElemLocationEV MeterValueElemSampledValueElemLocation = "EV"
const MeterValueElemSampledValueElemLocationInlet MeterValueElemSampledValueElemLocation = "Inlet"
const MeterValueElemSampledValueElemLocationOutlet MeterValueElemSampledValueElemLocation = "Outlet"

type MeterValueElemSampledValueElemMeasurand string

const MeterValueElemSampledValueElemMeasurandCurrentExport MeterValueElemSampledValueElemMeasurand = "Current.Export"
const MeterValueElemSampledValueElemMeasurandCurrentImport MeterValueElemSampledValueElemMeasurand = "Current.Import"
const MeterValueElemSampledValueElemMeasurandCurrentOffered MeterValueElemSampledValueElemMeasurand = "Current.Offered"
const MeterValueElemSampledValueElemMeasurandEnergyActiveExportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyActiveExportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Active.Export.Register"
const MeterValueElemSampledValueElemMeasurandEnergyActiveImportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Register"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveExportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyActiveImportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Active.Import.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveExportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Export.Register"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveImportInterval MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Interval"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveImportRegister MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Import.Register"
const MeterValueElemSampledValueElemMeasurandEnergyActiveNet MeterValueElemSampledValueElemMeasurand = "Energy.Active.Net"
const MeterValueElemSampledValueElemMeasurandEnergyReactiveNet MeterValueElemSampledValueElemMeasurand = "Energy.Reactive.Net"
const MeterValueElemSampledValueElemMeasurandEnergyApparentNet MeterValueElemSampledValueElemMeasurand = "Energy.Apparent.Net"
const MeterValueElemSampledValueElemMeasurandEnergyApparentImport MeterValueElemSampledValueElemMeasurand = "Energy.Apparent.Import"
const MeterValueElemSampledValueElemMeasurandEnergyApparentExport MeterValueElemSampledValueElemMeasurand = "Energy.Apparent.Export"
const MeterValueElemSampledValueElemMeasurandFrequency MeterValueElemSampledValueElemMeasurand = "Frequency"
const MeterValueElemSampledValueElemMeasurandPowerActiveExport MeterValueElemSampledValueElemMeasurand = "Power.Active.Export"
const MeterValueElemSampledValueElemMeasurandPowerActiveImport MeterValueElemSampledValueElemMeasurand = "Power.Active.Import"
const MeterValueElemSampledValueElemMeasurandPowerFactor MeterValueElemSampledValueElemMeasurand = "Power.Factor"
const MeterValueElemSampledValueElemMeasurandPowerOffered MeterValueElemSampledValueElemMeasurand = "Power.Offered"
const MeterValueElemSampledValueElemMeasurandPowerReactiveExport MeterValueElemSampledValueElemMeasurand = "Power.Reactive.Export"
const MeterValueElemSampledValueElemMeasurandPowerReactiveImport MeterValueElemSampledValueElemMeasurand = "Power.Reactive.Import"
const MeterValueElemSampledValueElemMeasurandRPM MeterValueElemSampledValueElemMeasurand = "RPM"
const MeterValueElemSampledValueElemMeasurandSoC MeterValueElemSampledValueElemMeasurand = "SoC"
const MeterValueElemSampledValueElemMeasurandTemperature MeterValueElemSampledValueElemMeasurand = "Temperature"
const MeterValueElemSampledValueElemMeasurandVoltage MeterValueElemSampledValueElemMeasurand = "Voltage"

type MeterValueElemSampledValueElemPhase string

const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1 MeterValueElemSampledValueElemPhase = "L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1L2 MeterValueElemSampledValueElemPhase = "L1-L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL1N MeterValueElemSampledValueElemPhase = "L1-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2 MeterValueElemSampledValueElemPhase = "L2"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2L3 MeterValueElemSampledValueElemPhase = "L2-L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL2N MeterValueElemSampledValueElemPhase = "L2-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3 MeterValueElemSampledValueElemPhase = "L3"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3L1 MeterValueElemSampledValueElemPhase = "L3-L1"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseL3N MeterValueElemSampledValueElemPhase = "L3-N"
const MeterValuesJsonMeterValueElemSampledValueElemPhaseN MeterValueElemSampledValueElemPhase = "N"

type MeterValueElemSampledValueElemUnit string

const MeterValueElemSampledValueElemUnitA MeterValueElemSampledValueElemUnit = "A"
const MeterValueElemSampledValueElemUnitCelcius MeterValueElemSampledValueElemUnit = "Celcius"
const MeterValueElemSampledValueElemUnitCelsius MeterValueElemSampledValueElemUnit = "Celsius"
const MeterValueElemSampledValueElemUnitFahrenheit MeterValueElemSampledValueElemUnit = "Fahrenheit"
const MeterValueElemSampledValueElemUnitK MeterValueElemSampledValueElemUnit = "K"
const MeterValueElemSampledValueElemUnitKVA MeterValueElemSampledValueElemUnit = "kVA"
const MeterValueElemSampledValueElemUnitKW MeterValueElemSampledValueElemUnit = "kW"
const MeterValueElemSampledValueElemUnitKWh MeterValueElemSampledValueElemUnit = "kWh"
const MeterValueElemSampledValueElemUnitKvar MeterValueElemSampledValueElemUnit = "kvar"
const MeterValueElemSampledValueElemUnitKvarh MeterValueElemSampledValueElemUnit = "kvarh"
const MeterValueElemSampledValueElemUnitPercent MeterValueElemSampledValueElemUnit = "Percent"
const MeterValueElemSampledValueElemUnitV MeterValueElemSampledValueElemUnit = "V"
const MeterValueElemSampledValueElemUnitVA MeterValueElemSampledValueElemUnit = "VA"
const MeterValueElemSampledValueElemUnitVar MeterValueElemSampledValueElemUnit = "var"
const MeterValueElemSampledValueElemUnitVarh MeterValueElemSampledValueElemUnit = "varh"
const MeterValueElemSampledValueElemUnitW MeterValueElemSampledValueElemUnit = "W"
const MeterValueElemSampledValueElemUnitWh MeterValueElemSampledValueElemUnit = "Wh"

type Tariff struct {
	Id            int64   `json:"id"`
	TotalMoney    float64 `json:"totalMoney"`
	ServiceMoney  float64 `json:"serviceMoney"`
	ElectricMoney float64 `json:"electricMoney"`
	SharpMoney    float64 `json:"sharpMoney"`
	PeakMoney     float64 `json:"peakMoney"`
	FlatMoney     float64 `json:"flatMoney"`
	ValleyMoney   float64 `json:"valleyMoney"`
	Sharp         float64 `json:"sharp"`
	Peak          float64 `json:"peak"`
	Flat          float64 `json:"flat"`
	Valley        float64 `json:"valley"`
}

type Variable struct {
	Key      string  `json:"key"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value,omitempty"`
}

type OffPeak struct {
	SharpCurrent  int `json:"sharpCurrent"`
	PeakCurrent   int `json:"peakCurrent"`
	FlatCurrent   int `json:"flatCurrent"`
	ValleyCurrent int `json:"valleyCurrent"`
}

type TriggerMessageEnumType uint8

const (
	TriggerMessageEnumTypeCCCP TriggerMessageEnumType = iota
	TriggerMessageEnumTypeDiagnosticsNotification
	TriggerMessageEnumTypeFirmwareNotification
	TriggerMessageEnumTypeMeterValues
	TriggerMessageEnumTypeTransactionEvent
	TriggerMessageEnumTypeHeartbeat
)

type Temperatures struct {
	Temperature1 float64 `json:"temperature1"`
	Temperature2 float64 `json:"temperature2"`
	Temperature3 float64 `json:"temperature3"`
	Temperature4 float64 `json:"temperature4"`
	Temperature5 float64 `json:"temperature5"`
}

type IntellectCharge struct {
	EVSE           EVSE    `json:"evse"`
	IntellectType  uint8   `json:"strategy"`
	IntellectId    string  `json:"intellectId"`
	StartTime      string  `json:"startTime"`
	EndTime        *string `json:"endTime"`
	EndElectricity *int    `json:"endElectricity"`
	EndSoc         *int    `json:"endSOC"`
	Status         uint8   `json:"status"`
}
