package switchtest

import (
	"nested"
)

func fn() {
	// Test enum exemption
	var value nested.MyEnumType
	switch value {
	case nested.EnumType_BOOL:
		break
	case nested.EnumType_INT:
		break
	case nested.EnumType_DECIMAL:
		break
	default:
		// StaticCheck: ENUM_EXEMPT = [nested.EnumType_FLOAT]
		break
	}

	// Test failure on an enum case
	switch value { // MATCH "missing case nested.EnumType_DECIMAL"
	case nested.EnumType_BOOL:
		break
	case nested.EnumType_INT:
		break
	default:
		// StaticCheck: ENUM_EXEMPT = [nested.EnumType_FLOAT]
		break
	}

	// TODO: Add case for Interface type switch... it's much more complicated :(
}
