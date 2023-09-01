// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// LoanSummary - OK
type LoanSummary struct {
	ReportInfo *LoanSummaryReportInfo `json:"reportInfo,omitempty"`
	// Returns a summary of all loan activity for that integration type
	ReportItems []LoanSummaryReportItem `json:"reportItems,omitempty"`
}

func (o *LoanSummary) GetReportInfo() *LoanSummaryReportInfo {
	if o == nil {
		return nil
	}
	return o.ReportInfo
}

func (o *LoanSummary) GetReportItems() []LoanSummaryReportItem {
	if o == nil {
		return nil
	}
	return o.ReportItems
}
