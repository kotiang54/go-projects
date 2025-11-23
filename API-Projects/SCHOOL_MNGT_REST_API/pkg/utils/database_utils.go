package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func GetIDFromMap(m map[string]interface{}) (int, error) {
	idVal, exists := m["id"]
	if !exists {
		return 0, fmt.Errorf("id field is missing")
	}

	// Try multiple numeric type assertions
	switch v := idVal.(type) {
	case float64:
		return int(v), nil
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case int32:
		return int(v), nil
	case float32:
		return int(v), nil
	default:
		return 0, fmt.Errorf("id field is not a valid numeric type, got %T", v)
	}
}

// validateUpdateFields checks if the fields in the update map are valid and of correct type
func ValidateUpdateFields(model interface{}, validFields map[string]int, update map[string]interface{}) error {
	for key, value := range update {
		if key == "id" {
			continue
		}
		fieldIdx, ok := validFields[key]
		if !ok {
			return fmt.Errorf("invalid field: %s", key)
		}
		fieldType := reflect.TypeOf(model).Field(fieldIdx).Type
		val := reflect.ValueOf(value)
		if !val.Type().ConvertibleTo(fieldType) {
			return fmt.Errorf("type mismatch for field: %s", key)
		}
	}
	return nil
}

// Extracted function for building ORDER BY clause from sortby query parameters
func BuildOrderByClause(r *http.Request) string {
	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) == 0 {
		return ""
	}
	orderBy := " ORDER BY "
	for i, param := range sortParams {
		parts := strings.Split(param, ":")
		if len(parts) == 2 {
			field := parts[0]
			order := strings.ToUpper(parts[1])
			if order != "ASC" && order != "DESC" {
				order = "ASC"
			}
			if i > 0 {
				orderBy += ", "
			}
			orderBy += fmt.Sprintf("%s %s", field, order)
		}
	}
	return orderBy
}

// buildValidFieldsMap builds a map of valid JSON field names to struct field indices
func BuildValidFieldsMap(person interface{}) map[string]int {
	personType := reflect.TypeOf(person)
	validFields := make(map[string]int)
	for i := 0; i < personType.NumField(); i++ {
		jsonTag := strings.Split(personType.Field(i).Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			validFields[jsonTag] = i
		}
	}
	return validFields
}

func ApplyUpdateToStruct(person interface{}, validFields map[string]int, update map[string]interface{}) {
	for key, value := range update {
		if key == "id" {
			continue
		}
		fieldIdx := validFields[key]
		fieldVal := reflect.ValueOf(person).Elem().Field(fieldIdx)
		val := reflect.ValueOf(value)
		fieldVal.Set(val.Convert(fieldVal.Type()))
	}
}

// getStructValues returns a slice of values from a given model
func GetStructValues(model interface{}) []interface{} {
	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()

	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
		modelValue = modelValue.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return nil // or panic / log an error
	}

	var values []interface{}
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbTag := strings.Split(field.Tag.Get("db"), ",")[0]
		dbTag = strings.TrimSpace(dbTag)
		if dbTag != "" && dbTag != "id" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}
	return values
}

// generateInsertQuery generates an INSERT query for a given model
func GenerateInsertQuery(model interface{}, tableName string) string {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return "" // or panic / log an error
	}

	var columns, placeholders string

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbTag := strings.Split(field.Tag.Get("db"), ",")[0]
		dbTag = strings.TrimSpace(dbTag)
		if dbTag != "" && dbTag != "id" {
			if len(columns) > 0 {
				columns += ", "
				placeholders += ", "
			}
			columns += dbTag
			placeholders += "?"
		}

	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, placeholders)
	return query
}
