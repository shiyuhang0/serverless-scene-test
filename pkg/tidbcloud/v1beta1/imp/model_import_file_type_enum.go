/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imp

import (
	"encoding/json"
	"fmt"
)

// ImportFileTypeEnum  - CSV: CSV type.  - SQL: SQL type.  - AURORA_SNAPSHOT: Aurora snapshot type.  - PARQUET: Parquet type.
type ImportFileTypeEnum string

// List of ImportFileType.Enum
const (
	IMPORTFILETYPEENUM_CSV             ImportFileTypeEnum = "CSV"
	IMPORTFILETYPEENUM_SQL             ImportFileTypeEnum = "SQL"
	IMPORTFILETYPEENUM_AURORA_SNAPSHOT ImportFileTypeEnum = "AURORA_SNAPSHOT"
	IMPORTFILETYPEENUM_PARQUET         ImportFileTypeEnum = "PARQUET"
)

// All allowed values of ImportFileTypeEnum enum
var AllowedImportFileTypeEnumEnumValues = []ImportFileTypeEnum{
	"CSV",
	"SQL",
	"AURORA_SNAPSHOT",
	"PARQUET",
}

func (v *ImportFileTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportFileTypeEnum(value)
	for _, existing := range AllowedImportFileTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportFileTypeEnum", value)
}

// NewImportFileTypeEnumFromValue returns a pointer to a valid ImportFileTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportFileTypeEnumFromValue(v string) (*ImportFileTypeEnum, error) {
	ev := ImportFileTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportFileTypeEnum: valid values are %v", v, AllowedImportFileTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportFileTypeEnum) IsValid() bool {
	for _, existing := range AllowedImportFileTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportFileType.Enum value
func (v ImportFileTypeEnum) Ptr() *ImportFileTypeEnum {
	return &v
}

type NullableImportFileTypeEnum struct {
	value *ImportFileTypeEnum
	isSet bool
}

func (v NullableImportFileTypeEnum) Get() *ImportFileTypeEnum {
	return v.value
}

func (v *NullableImportFileTypeEnum) Set(val *ImportFileTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableImportFileTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableImportFileTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportFileTypeEnum(val *ImportFileTypeEnum) *NullableImportFileTypeEnum {
	return &NullableImportFileTypeEnum{value: val, isSet: true}
}

func (v NullableImportFileTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportFileTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}