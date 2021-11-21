package events

import "reflect"

func filterOptionalFields(data map[string]interface{}, fields ...string) {
	for _, field := range fields {
		if val, ok := data[field]; ok {
			r := reflect.ValueOf(val)
			if r.IsZero() {
				delete(data, field)
			}
		}
	}
}
