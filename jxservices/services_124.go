//go:build go1.24

package jxservices

import (
	"strings"

	"github.com/ForbiddenR/jxapi/v2/jxutils/cache"
)

var categoryCache *cache.Cache[Request2ServicesNameType, string]

func init() {
	categoryCache = cache.NewCache[Request2ServicesNameType, string]()
}

func (r Request2ServicesNameType) firstUpper() string {
	s := r.String()
	return strings.ToUpper(s[:1]) + s[1:]
}

func (r Request2ServicesNameType) FirstUpper2() string {
	if v, ok := categoryCache.Get(r); ok {
		return v
	}
	category := r.firstUpper()
	categoryCache.Set(r, category)
	return category
}

func (r Request2ServicesNameType) GetCallbackCategory2() string {
	if v, ok := categoryCache.Get(r); ok {
		return v
	}
	category := r.firstUpper() + CallbackSuffix
	categoryCache.Set(r, category)
	return category
}

// TODO: use a common string rather than a string with type Request2ServicesNameType.
func (b *BaseConfig) Categories2(kind Request2ServicesNameType, isCallback bool) *BaseConfig {
	if !isCallback {
		b.category = kind.FirstUpper()
	} else {
		b.category = kind.GetCallbackCategory()
	}
	return b
}
