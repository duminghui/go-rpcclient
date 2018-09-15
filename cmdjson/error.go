package cmdjson

import "fmt"

type ErrorCode int

const (
	ErrDuplicateMethod ErrorCode = iota
	ErrInvalidUsageFlags
	ErrInvalidType
	ErrEmbeddedType
	ErrUnexportedField
	ErrUnsupportedFieldType
	ErrNonOptionalField
	ErrNonOptionalDefault
	ErrMismatchedDefault
	ErrUnregisteredMethod
	ErrMissingDescription
	ErrNumParams
)

var errorCodeStrings = map[ErrorCode]string{
	ErrDuplicateMethod:      "ErrDuplicateMethod",
	ErrInvalidUsageFlags:    "ErrInvalidUsageFlags",
	ErrInvalidType:          "ErrInvalidType",
	ErrEmbeddedType:         "ErrEmbeddedType",
	ErrUnexportedField:      "ErrUnexportedField",
	ErrUnsupportedFieldType: "ErrUnsupportedFieldType",
	ErrNonOptionalField:     "ErrNonOptionalField",
	ErrNonOptionalDefault:   "ErrNonOptionalDefault",
	ErrMismatchedDefault:    "ErrMismatchedDefault",
	ErrUnregisteredMethod:   "ErrUnregisteredMethod",
	ErrMissingDescription:   "ErrMissingDescription",
	ErrNumParams:            "ErrNumParams",
}

func (e ErrorCode) String() string {
	s, ok := errorCodeStrings[e]
	if !ok {
		return fmt.Sprintf("Unknown ErrorCode (%d)", int(e))
	}
	return s
}

type Error struct {
	ErrorCode   ErrorCode
	Description string
}

func (e Error) Error() string {
	return e.Description
}

func makeError(c ErrorCode, desc string) Error {
	return Error{ErrorCode: c, Description: desc}
}
