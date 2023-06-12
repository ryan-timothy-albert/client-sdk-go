# CreditNoteLineItem


## Fields

| Field                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                | Required                                                                                                                                                                                                                                                                                            | Description                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccountRef`                                                                                                                                                                                                                                                                                        | [*AccountRef](../../models/shared/accountref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.                                                                                                                                               |
| `Description`                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Friendly name of each line item. For example, the goods or service for which credit has been issued.                                                                                                                                                                                                |
| `DiscountAmount`                                                                                                                                                                                                                                                                                    | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Value of any discounts applied.                                                                                                                                                                                                                                                                     |
| `DiscountPercentage`                                                                                                                                                                                                                                                                                | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Percentage rate of any discount applied to the line item.                                                                                                                                                                                                                                           |
| `IsDirectIncome`                                                                                                                                                                                                                                                                                    | **bool*                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                                                                                                 |
| `ItemRef`                                                                                                                                                                                                                                                                                           | [*ItemRef](../../models/shared/itemref.md)                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                                                                                                 |
| `Quantity`                                                                                                                                                                                                                                                                                          | *float64*                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                  | Number of units of the goods or service for which credit has been issued.                                                                                                                                                                                                                           |
| `SubTotal`                                                                                                                                                                                                                                                                                          | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Amount of credit associated with the line item, including discounts but excluding tax.                                                                                                                                                                                                              |
| `TaxAmount`                                                                                                                                                                                                                                                                                         | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Amount of tax associated with the line item.                                                                                                                                                                                                                                                        |
| `TaxRateRef`                                                                                                                                                                                                                                                                                        | [*TaxRateRef](../../models/shared/taxrateref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.<br/><br/>Found on:<br/><br/>- Bill line items<br/>- Bill Credit Note line items<br/>- Credit Note line items<br/>- Direct incomes line items<br/>- Invoice line items<br/>- Items |
| `TotalAmount`                                                                                                                                                                                                                                                                                       | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Total amount of the line item, including discounts and tax.                                                                                                                                                                                                                                         |
| `Tracking`                                                                                                                                                                                                                                                                                          | [*CreditNoteLineItemTracking](../../models/shared/creditnotelineitemtracking.md)                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Categories, and a project and customer, against which the item is tracked.                                                                                                                                                                                                                          |
| `TrackingCategoryRefs`                                                                                                                                                                                                                                                                              | [][TrackingCategoryRef](../../models/shared/trackingcategoryref.md)                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Reference to the tracking categories to which the line item is linked.                                                                                                                                                                                                                              |
| `UnitAmount`                                                                                                                                                                                                                                                                                        | *float64*                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                  | Unit price of the goods or service.                                                                                                                                                                                                                                                                 |