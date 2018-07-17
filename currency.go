//go:generate protoc -I ./proto/iso/protobuf currency.proto --go_out=../../../

package iso

// Scan implements the sql.Scanner interface
func (c *Currency) Scan(v interface{}) error {
	if v == nil {
		return nil // defaults to UNKNOWN
	}
	s, ok := v.(string)
	if !ok {
		return nil
	}
	switch s {
	// two of the values aren't alphanum strings and therefore invalid proto enums.
	// Those can't be resolved with the generated map.
	case "999":
		*c = Currency_UNKNOWN_CURRENCY
	default:
		i, ok := Currency_value[s]
		if !ok {
			return nil
		}
		*c = Currency(i)
	}
	return nil
}
