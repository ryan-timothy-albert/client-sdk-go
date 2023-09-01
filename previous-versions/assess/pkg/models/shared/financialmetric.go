// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/assess/pkg/types"
)

// FinancialMetricErrorsType - Metric level error.
type FinancialMetricErrorsType string

const (
	FinancialMetricErrorsTypeUncategorizedAccounts FinancialMetricErrorsType = "UncategorizedAccounts"
	FinancialMetricErrorsTypeMissingInputData      FinancialMetricErrorsType = "MissingInputData"
	FinancialMetricErrorsTypeInputDataError        FinancialMetricErrorsType = "InputDataError"
)

func (e FinancialMetricErrorsType) ToPointer() *FinancialMetricErrorsType {
	return &e
}

func (e *FinancialMetricErrorsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UncategorizedAccounts":
		fallthrough
	case "MissingInputData":
		fallthrough
	case "InputDataError":
		*e = FinancialMetricErrorsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FinancialMetricErrorsType: %v", v)
	}
}

type FinancialMetricErrors struct {
	// Dictionary list outlining the missing properties or allowed values.
	Details map[string][]string `json:"details,omitempty"`
	// Description of the error.
	Message *string `json:"message,omitempty"`
	// Metric level error.
	Type *FinancialMetricErrorsType `json:"type,omitempty"`
}

type FinancialMetricKey string

const (
	FinancialMetricKeyUnknown                     FinancialMetricKey = "Unknown"
	FinancialMetricKeyEbitda                      FinancialMetricKey = "EBITDA"
	FinancialMetricKeyDebtServiceCoverageRatio    FinancialMetricKey = "DebtServiceCoverageRatio"
	FinancialMetricKeyCurrentRatioQuickRatio      FinancialMetricKey = "CurrentRatio QuickRatio"
	FinancialMetricKeyGrossProfitMargin           FinancialMetricKey = "GrossProfitMargin"
	FinancialMetricKeyFixedChargeCoverageRatio    FinancialMetricKey = "FixedChargeCoverageRatio"
	FinancialMetricKeyWorkingCapital              FinancialMetricKey = "WorkingCapital"
	FinancialMetricKeyFreeCashFlow                FinancialMetricKey = "FreeCashFlow"
	FinancialMetricKeyNetProfitMargin             FinancialMetricKey = "NetProfitMargin"
	FinancialMetricKeyReturnOnAssetsRatio         FinancialMetricKey = "ReturnOnAssetsRatio"
	FinancialMetricKeyReturnOnEquityRatio         FinancialMetricKey = "ReturnOnEquityRatio"
	FinancialMetricKeyOperatingProfitMargin       FinancialMetricKey = "OperatingProfitMargin"
	FinancialMetricKeyDeptToEquity                FinancialMetricKey = "DeptToEquity"
	FinancialMetricKeyDebtToAssets                FinancialMetricKey = "DebtToAssets"
	FinancialMetricKeyInterestCoverageRatio       FinancialMetricKey = "InterestCoverageRatio"
	FinancialMetricKeyCashRatio                   FinancialMetricKey = "CashRatio"
	FinancialMetricKeyInventoryTurnoverRatio      FinancialMetricKey = "InventoryTurnoverRatio"
	FinancialMetricKeyAssetTurnoverRatio          FinancialMetricKey = "AssetTurnoverRatio"
	FinancialMetricKeyWorkingCapitalTurnoverRatio FinancialMetricKey = "WorkingCapitalTurnoverRatio"
	FinancialMetricKeyDaysSalesOutstanding        FinancialMetricKey = "DaysSalesOutstanding"
	FinancialMetricKeyDaysPayablesOutstanding     FinancialMetricKey = "DaysPayablesOutstanding"
)

func (e FinancialMetricKey) ToPointer() *FinancialMetricKey {
	return &e
}

func (e *FinancialMetricKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "EBITDA":
		fallthrough
	case "DebtServiceCoverageRatio":
		fallthrough
	case "CurrentRatio QuickRatio":
		fallthrough
	case "GrossProfitMargin":
		fallthrough
	case "FixedChargeCoverageRatio":
		fallthrough
	case "WorkingCapital":
		fallthrough
	case "FreeCashFlow":
		fallthrough
	case "NetProfitMargin":
		fallthrough
	case "ReturnOnAssetsRatio":
		fallthrough
	case "ReturnOnEquityRatio":
		fallthrough
	case "OperatingProfitMargin":
		fallthrough
	case "DeptToEquity":
		fallthrough
	case "DebtToAssets":
		fallthrough
	case "InterestCoverageRatio":
		fallthrough
	case "CashRatio":
		fallthrough
	case "InventoryTurnoverRatio":
		fallthrough
	case "AssetTurnoverRatio":
		fallthrough
	case "WorkingCapitalTurnoverRatio":
		fallthrough
	case "DaysSalesOutstanding":
		fallthrough
	case "DaysPayablesOutstanding":
		*e = FinancialMetricKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FinancialMetricKey: %v", v)
	}
}

type FinancialMetricMetricUnit string

const (
	FinancialMetricMetricUnitRatio FinancialMetricMetricUnit = "Ratio"
	FinancialMetricMetricUnitMoney FinancialMetricMetricUnit = "Money"
)

func (e FinancialMetricMetricUnit) ToPointer() *FinancialMetricMetricUnit {
	return &e
}

func (e *FinancialMetricMetricUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Ratio":
		fallthrough
	case "Money":
		*e = FinancialMetricMetricUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FinancialMetricMetricUnit: %v", v)
	}
}

// FinancialMetricPeriodsErrorsType - Period error type.
type FinancialMetricPeriodsErrorsType string

const (
	FinancialMetricPeriodsErrorsTypeMissingAccountData FinancialMetricPeriodsErrorsType = "MissingAccountData"
	FinancialMetricPeriodsErrorsTypeDatesOutOfRange    FinancialMetricPeriodsErrorsType = "DatesOutOfRange"
)

func (e FinancialMetricPeriodsErrorsType) ToPointer() *FinancialMetricPeriodsErrorsType {
	return &e
}

func (e *FinancialMetricPeriodsErrorsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MissingAccountData":
		fallthrough
	case "DatesOutOfRange":
		*e = FinancialMetricPeriodsErrorsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FinancialMetricPeriodsErrorsType: %v", v)
	}
}

type FinancialMetricPeriodsErrors struct {
	// Dictionary list outlining the missing properties or allowed values.
	Details map[string][]string `json:"details,omitempty"`
	// Description of the error.
	Massage *string `json:"massage,omitempty"`
	// Period error type.
	Type *FinancialMetricPeriodsErrorsType `json:"type,omitempty"`
}

type FinancialMetricPeriodsInputs struct {
	// The name of the metric input e.g. “Current Assets”, “Capital Expenditure”.
	Name *string `json:"name,omitempty"`
	// The positive or negative number of the input value.
	Value *float64 `json:"value,omitempty"`
}

type FinancialMetricPeriods struct {
	Errors []FinancialMetricPeriodsErrors `json:"errors,omitempty"`
	// The date from which the report starts.
	FromDate *types.Date                    `json:"fromDate,omitempty"`
	Inputs   []FinancialMetricPeriodsInputs `json:"inputs,omitempty"`
	// The date on which the report ends (inclusive of day).
	ToDate *types.Date `json:"toDate,omitempty"`
	// The top level metric value that is calculated for the specified period.
	//
	// If the system cannot calculate for that period, the value will be null. The system will still show the metric inputs.
	Value *float64 `json:"value,omitempty"`
}

type FinancialMetric struct {
	Errors     []FinancialMetricErrors    `json:"errors,omitempty"`
	Key        *FinancialMetricKey        `json:"key,omitempty"`
	MetricUnit *FinancialMetricMetricUnit `json:"metricUnit,omitempty"`
	// Metric name.
	Name    *string                  `json:"name,omitempty"`
	Periods []FinancialMetricPeriods `json:"periods,omitempty"`
}