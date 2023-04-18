// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ReportBasisEnum - Accounting method used when aggregating the report data. In this case, `Cash`.
type ReportBasisEnum string

const (
	ReportBasisEnumUnknown ReportBasisEnum = "Unknown"
	ReportBasisEnumAccrual ReportBasisEnum = "Accrual"
	ReportBasisEnumCash    ReportBasisEnum = "Cash"
)

func (e *ReportBasisEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Accrual":
		fallthrough
	case "Cash":
		*e = ReportBasisEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ReportBasisEnum: %s", s)
	}
}