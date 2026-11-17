package apierrors

import "slices"

func IsOneOf(err error, expects ...error) bool {
	return slices.Contains(expects, err)
}
