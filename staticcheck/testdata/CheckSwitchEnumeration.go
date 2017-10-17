package pkg

// StaticCheck: ENUM = MyEnumType
type MyEnumType int32

const (
	EnumType_BOOL           MyEnumType = iota
	EnumType_INT            MyEnumType = iota
	EnumType_FLOAT          MyEnumType = iota
	EnumType_DECIMAL        MyEnumType = iota
)


func fn() {
	// Test enum exemption
	var value MyEnumType
	switch value {
	case EnumType_BOOL:
		break
	case EnumType_INT:
		break
	case EnumType_DECIMAL:
		break
	default:
		// StaticCheck: ENUM_EXEMPT = [EnumType_FLOAT]
		break
	}

	// Test failure on an enum case
	switch value { // MATCH "missing case EnumType_DECIMAL"
	case EnumType_BOOL:
		break
	case EnumType_INT:
		break
	default:
		// StaticCheck: ENUM_EXEMPT = [EnumType_FLOAT]
		break
	}

	// TODO: Add case for Interface type switch... it's much more complicated :(
}
