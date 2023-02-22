package regmap

// Register describes a registry for a device.
type Register struct {
	Address     uint16 `json:"address" db:"address"`
	Symbol      string `json:"symbol" db:"symbol"`
	Display     string `json:"display" db:"display"`
	Description string `json:"description" db:"description"`
	Access      Mode   `json:"access" db:"access"`
}
