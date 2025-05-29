package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetPriceSchemeRequest struct {
	services.Base
	Data *EquipSetPriceSchemeRequestDetail `json:"data"`
}

type EquipSetPriceSchemeRequestDetail struct {
	VendorId      string         `json:"vendorId"`
	TariffId      uint64         `json:"tariffId"`
	BaseTime      int64          `json:"baseTime"`
	ChargeTariffs []ChargeTariff `json:"chargeTariffs"`
	TimePrices    []TimePrice    `json:"timePrices"`
}

// TODO: the format of TImeStart.
type ChargeTariff struct {
	TimeStart int `json:"timeStart"`
	Tag       int `json:"tag"`
}

type TimePrice struct {
	ElecPrice    float64 `json:"elecPrice"`
	ServicePrice float64 `json:"servicePrice"`
	Tag          int     `json:"tag"`
}

func (s *EquipSetPriceSchemeRequest) Unmarshal(data []byte) error {
	type Alias EquipSetPriceSchemeRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}