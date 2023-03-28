// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListTrackingCategoriesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type ListTrackingCategories200ApplicationJSONLinksHypertextReference struct {
	Href *string `json:"href,omitempty"`
}

type ListTrackingCategories200ApplicationJSONLinks struct {
	Current  ListTrackingCategories200ApplicationJSONLinksHypertextReference  `json:"current"`
	Next     *ListTrackingCategories200ApplicationJSONLinksHypertextReference `json:"next,omitempty"`
	Previous *ListTrackingCategories200ApplicationJSONLinksHypertextReference `json:"previous,omitempty"`
	Self     ListTrackingCategories200ApplicationJSONLinksHypertextReference  `json:"self"`
}

// ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum - Current state of the tracking category.
type ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum string

const (
	ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnumUnknown  ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnumActive   ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum = "Active"
	ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnumArchived ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

func (e *ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Unknown":
		fallthrough
	case "Active":
		fallthrough
	case "Archived":
		*e = ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum: %s", s)
	}
}

// ListTrackingCategories200ApplicationJSONSourceModifiedDate - Details of a category used for tracking transactions.
//
// > Language tip
// >
// > Parameters used to track types of spend in various parts of an organization can be called  **dimensions**, **projects**, **classes**, or **locations** in different accounting platforms. In Codat, we refer to these as tracking categories.
//
// View the coverage for tracking categories in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=trackingCategories" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Tracking categories are used to monitor cost centres and control budgets that sit outside the standard chart of accounts. Customers may use tracking categories to group together and track the income and costs of specific departments, projects, locations or customers.
//
// From their accounting system, customers can:
//
// - Create and maintain tracking categories and tracking category types.
// - View all tracking categories that are available for use.
// - View the relationships between the categories.
// - Assign invoices, bills, credit notes, or bill credit notes to one or more categories.
// - View the categories that a transaction belongs to.
// - View all transactions in a tracking category.
type ListTrackingCategories200ApplicationJSONSourceModifiedDate struct {
	// Boolean value indicating whether this category has SubCategories
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The identifier for the item, unique per tracking category
	ID *string `json:"id,omitempty"`
	// The date on which this record was last modified in Codat.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the tracking category
	Name *string `json:"name,omitempty"`
	// The identifier for this item's immediate parent
	ParentID *string `json:"parentId,omitempty"`
	// The date on which this record was last modified in the originating system
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the tracking category.
	Status *ListTrackingCategories200ApplicationJSONSourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
}

// ListTrackingCategories200ApplicationJSON - Success
type ListTrackingCategories200ApplicationJSON struct {
	Links        ListTrackingCategories200ApplicationJSONLinks                `json:"_links"`
	PageNumber   int64                                                        `json:"pageNumber"`
	PageSize     int64                                                        `json:"pageSize"`
	Results      []ListTrackingCategories200ApplicationJSONSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                                        `json:"totalResults"`
}

type ListTrackingCategoriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	ListTrackingCategories200ApplicationJSONObject *ListTrackingCategories200ApplicationJSON
}
