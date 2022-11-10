package scan

import "fmt"

func Maps(v interface{}, excluded ...string) (map[string]interface{}, error) {
	return maps(v, false, excluded...)
}

func MapsStrict(v interface{}, excluded ...string) (map[string]interface{}, error) {
	return maps(v, true, excluded...)
}

func maps(v interface{}, strict bool, excluded ...string) (map[string]interface{}, error) {
	if cols, err := columns(v, strict, excluded...); err != nil {
		return nil, err
	} else if vals, err := Values(cols, v); err != nil {
		return nil, err
	} else if len(cols) != len(vals) {
		return nil, fmt.Errorf("column and value length mismatch")
	} else {
		m := map[string]interface{}{}
		for i, col := range cols {
			m[col] = vals[i]
		}
		return m, nil
	}
}
