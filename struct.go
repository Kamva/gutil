package gutil

import "encoding/json"

// StructToMap convert the struct to a map by marshall and un-marshall
func StructToMap(input interface{}) map[string]interface{} {
	var m = map[string]interface{}{}

	encodedJson, _ := json.Marshal(input)
	_ = json.Unmarshal(encodedJson, &m)

	return m
}

// MapToStruct convert map to struct by json marshal and Unmarshal
func MapToStruct(m map[string]interface{}, s interface{}) error {
	encodedJson, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(encodedJson, s)
}
