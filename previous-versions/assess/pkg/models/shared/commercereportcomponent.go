// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CommerceReportComponent struct {
	Components           []CommerceReportComponent `json:"components,omitempty"`
	Dimension            *int64                    `json:"dimension,omitempty"`
	DimensionDisplayName *string                   `json:"dimensionDisplayName,omitempty"`
	Item                 *int64                    `json:"item,omitempty"`
	ItemDisplayName      *string                   `json:"itemDisplayName,omitempty"`
	Measures             []ReportComponentMeasure  `json:"measures,omitempty"`
}

func (o *CommerceReportComponent) GetComponents() []CommerceReportComponent {
	if o == nil {
		return nil
	}
	return o.Components
}

func (o *CommerceReportComponent) GetDimension() *int64 {
	if o == nil {
		return nil
	}
	return o.Dimension
}

func (o *CommerceReportComponent) GetDimensionDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DimensionDisplayName
}

func (o *CommerceReportComponent) GetItem() *int64 {
	if o == nil {
		return nil
	}
	return o.Item
}

func (o *CommerceReportComponent) GetItemDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.ItemDisplayName
}

func (o *CommerceReportComponent) GetMeasures() []ReportComponentMeasure {
	if o == nil {
		return nil
	}
	return o.Measures
}
