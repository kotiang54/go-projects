package handlers

import (
	"errors"
	"reflect"
	"school_management_api/pkg/utils"
	"strings"
)

// GetFieldNames extracts valid JSON field names from a struct's json tags using reflection.
func GetFieldNames(model interface{}) map[string]struct{} {
	val := reflect.TypeOf(model)
	validFields := make(map[string]struct{})
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			validFields[jsonTag] = struct{}{}
		}
	}
	return validFields
}

func CheckBlankFields(value interface{}) error {
	val := reflect.ValueOf(value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.Len() == 0 {
			return utils.ErrorHandler(errors.New("all fields are required"), "all fields are required")
		}
	}
	return nil
}
