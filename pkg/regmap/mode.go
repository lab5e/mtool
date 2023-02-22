package regmap

// Mode defines the access modes available for an address.
type Mode byte

// Mode enum values
const (
	ModeUnspecified Mode = iota
	ModeRead
	ModeWrite
	ModeReadWrite
)
