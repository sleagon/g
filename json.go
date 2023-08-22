package g

import (
	"bytes"
	"encoding/json"
	"strings"
)

type JMap map[string]any

func (j JMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

func (j *JMap) UnmarshalJSON(bts []byte) error {

	tmpVal := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewBuffer(bts))
	d.UseNumber()

	if err := d.Decode(&tmpVal); err != nil {
		return err
	}

	for k, v := range tmpVal {

		parsedValue := v
		var err error

		switch tv := v.(type) {
		case json.Number:
			if strings.Contains(string(tv), ".") {
				if parsedValue, err = tv.Float64(); err != nil {
					return err
				}
			} else {
				if parsedValue, err = tv.Int64(); err != nil {
					return err
				}
			}
		default:
		}

		(*j)[k] = parsedValue
	}

	return nil
}
