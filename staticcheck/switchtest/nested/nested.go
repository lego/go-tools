package nested

// StaticCheck: ENUM = MyEnumType
type MyEnumType int32

const (
	EnumType_BOOL    MyEnumType = iota
	EnumType_INT     MyEnumType = iota
	EnumType_FLOAT   MyEnumType = iota
	EnumType_DECIMAL MyEnumType = iota
)
