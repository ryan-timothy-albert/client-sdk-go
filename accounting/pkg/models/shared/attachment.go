// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Attachment - The Codat API supports pulling and pushing of file attachments for invoices, bills, direct costs, and direct incomes.
//
// > **Retrieving attachments**
// >
// > If a company is authorized, you can query the Codat API to read, download, and upload attachments without requiring a fresh sync of data.
//
// Unlike other data types, Codat doesn't support sync settings for attachments.
//
// Note that different integrations have different requirements to file size and extension of attachments.
//
// | Integration       | File size | File extension                                                                                                                                               |
// |-------------------|-----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
// | Xero              | 4 MB      | 7Z, BMP, CSV, DOC, DOCX, EML, GIF, JPEG, JPG, KEYNOTE, MSG, NUMBERS, ODF,   ODS, ODT, PAGES, PDF, PNG, PPT, PPTX, RAR, RTF, TIF, TIFF, TXT, XLS, XLSX,   ZIP |
// | QuickBooks Online | 100 MB    | AI, CSV, DOC, DOCX, EPS, GIF, JPEG, JPG, ODS, PAGES, PDF, PNG, RTF, TIF,   TXT, XLS, XLSX, XML                                                               |
// | NetSuite          | 100 MB    | BMP, CSV, XLS, XLSX, JSON, PDF, PJPG, PJPEG, PNG, TXT, SVG, TIF, TIFF,   DOC, DOCX, ZIP                                                                      |
//
// View the coverage for accounts in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=chartOfAccounts" target="_blank">Data coverage explorer</a>.
type Attachment struct {
	// File type of the attachment. This is represented by appending the file type to the [IETF standard file naming requirements](https://tools.ietf.org/html/rfc6838). For example, for a jpeg file the output is **image/jpeg**.
	//
	// Supported file types vary per platform.
	ContentType *string `json:"contentType,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	DateCreated *string `json:"dateCreated,omitempty"`
	// File size in bytes. For example, if this reads **46153**, then the file size is 46kb.
	FileSize *int `json:"fileSize,omitempty"`
	// Identifier for the attachment, unique for the company in the accounting platform.
	ID *string `json:"id,omitempty"`
	// If `true`, then the attachment is included with the associated invoice, bill or direct costs when it is printed, emailed, or sent to a customer, if the underlying accounting platform allows this.
	IncludeWhenSent *bool   `json:"includeWhenSent,omitempty"`
	ModifiedDate    *string `json:"modifiedDate,omitempty"`
	// Name of the attachment file.
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}
