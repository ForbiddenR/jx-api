package jxesam

import (
	"strings"
	"unique"

	api "github.com/ForbiddenR/jxapi/v2"
)

const (
	Equip     = "device"
	TicketKey = "ServiceInternalTickets"
)

type RequestNameEsamType string

const (
	Access RequestNameEsamType = "accessVerify"
)

var esamPerm *unique.Handle[string]

func (r RequestNameEsamType) String() string {
	return string(r)
}

func Perm() string {
	if esamPerm == nil {
		permSlice := make([]string, 0, 4)
		permSlice = append(permSlice, api.Esam, Equip)
		permSlice = append(permSlice, split(Access)...)
		maker := unique.Make(strings.Join(permSlice, ":"))
		esamPerm = &maker
	}
	return esamPerm.Value()
}

func split(r RequestNameEsamType) []string {
	for i := range len(r) {
		str := r.String()[i : i+1]
		if str == strings.ToUpper(str) {
			return []string{r.String()[:i], strings.ToLower(r.String()[i:])}
		}
	}
	return nil
}
