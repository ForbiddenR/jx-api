package apierrors

type JsonErrorType int

const (
	// SyntaxJsonError json语法解析错误
	// SyntaxJsonError happens when json messages fail to be parsed for grammar errors.
	SyntaxJsonError JsonErrorType = 1
	// RequiredJsonError json必须要的字段未解析到
	// RequiredJsonError happens because of a lack of necessary fields.
	RequiredJsonError JsonErrorType = 2
	// FormatValueJsonError json一些字段的值解析错误 e.g 时间类型必须是RFC3339, 还有类型错误
	// FormatValueJsonError happens because the values of some json fields fails to be parsed like type error or
	// time fields' violating the RFC3339 standard etc.
	FormatValueJsonError JsonErrorType = 3
	// EnumValueJsonError json一些字段是枚举值 解析出来的内容不符合枚举值
	// EnumValueJsonError occurs when some parsed values do not include in its enum domain.
	EnumValueJsonError JsonErrorType = 4
)

type JsonError struct {
	t     JsonErrorType
	inner error
}

func (e *JsonError) Type() JsonErrorType {
	return e.t
}

func (e *JsonError) Error() string {
	return e.inner.Error()
}

func WrapSyntaxJsonError(err error) *JsonError {
	return &JsonError{inner: err, t: SyntaxJsonError}
}
func WrapRequiredJsonError(err error) *JsonError {
	return &JsonError{inner: err, t: RequiredJsonError}
}
func WrapFormatValueJsonError(err error) *JsonError {
	return &JsonError{inner: err, t: FormatValueJsonError}
}

func WrapEnumValueJsonError(err error) *JsonError {
	return &JsonError{inner: err, t: EnumValueJsonError}
}

func IsSyntaxJsonError(err error) bool {
	if e, ok := IsJsonError(err); ok {
		return e.t == SyntaxJsonError
	}
	return false
}

func IsRequiredJsonError(err error) bool {
	if e, ok := IsJsonError(err); ok {
		return e.t == RequiredJsonError
	}
	return false
}

func IsFormatValueJsonError(err error) bool {
	if e, ok := IsJsonError(err); ok {
		return e.t == FormatValueJsonError
	}
	return false
}

func IsEnumValueJsonError(err error) bool {
	if e, ok := IsJsonError(err); ok {
		return e.t == EnumValueJsonError
	}
	return false
}

func IsJsonError(err error) (*JsonError, bool) {
	e, ok := err.(*JsonError)
	if ok {
		return e, ok
	} else {
		return nil, ok
	}
}
