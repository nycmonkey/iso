//go:generate protoc -I ./proto/iso/protobuf country.proto --go_out=../../../

package iso

// Scan implements the sql.Scanner interface
func (c *Country) Scan(v interface{}) error {
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
	case "00":
		*c = Country_SUPRANATIONAL
	case "99":
		*c = Country_NOT_APPLICABLE
	default:
		i, ok := Country_value[s]
		if !ok {
			return nil
		}
		*c = Country(i)
	}
	return nil
}
