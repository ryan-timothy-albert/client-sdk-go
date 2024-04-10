// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Locale string

const (
	LocaleEnUs Locale = "en-us"
	LocaleFrFr Locale = "fr-fr"
)

func (e Locale) ToPointer() *Locale {
	return &e
}

func (e *Locale) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "en-us":
		fallthrough
	case "fr-fr":
		*e = Locale(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Locale: %v", v)
	}
}