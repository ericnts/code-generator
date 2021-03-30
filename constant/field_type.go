package constant

type FieldType string

const (
	FTString FieldType = "string"
	FTInt    FieldType = "int"
	FTInt64  FieldType = "int64"
	FTFloat  FieldType = "float64"
	FTTime   FieldType = "time.Time"
	FTBool   FieldType = "bool"
)
