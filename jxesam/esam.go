package jxesam

import (
	"strings"

	api "github.com/ForbiddenR/jxapi"
)

const (
	Equip     = "device"
	TicketKey = "ServiceInternalTickets"
)

type RequestNameEsamType string

const (
	Access RequestNameEsamType = "accessVerify"
)

var reusedPerm *string

func (r RequestNameEsamType) String() string {
	return string(r)
}

func (r RequestNameEsamType) Perm() string {
	if reusedPerm != nil {
		return *reusedPerm
	}
	permSlice := make([]string, 0, 4)
	permSlice = append(permSlice, api.Esam, Equip)
	permSlice = append(permSlice, r.Split()...)
	perm := strings.Join(permSlice, ":")
	reusedPerm = &perm
	return perm
}

func (r RequestNameEsamType) Split() []string {
	for i, v := range r {
		if string(v) == strings.ToUpper(string(v)) {
			return []string{string(r)[:i], string(r)[i:]}
		}
	}
	return nil
}
