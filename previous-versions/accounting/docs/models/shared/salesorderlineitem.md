# SalesOrderLineItem


## Fields

| Field                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                | Required                                                                                                                                                                                                                                                                                            | Description                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccountRef`                                                                                                                                                                                                                                                                                        | [*AccountRef](../../models/shared/accountref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.                                                                                                                                               |
| `Description`                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Description of the goods or services that have been ordered.                                                                                                                                                                                                                                        |
| `DiscountAmount`                                                                                                                                                                                                                                                                                    | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Value of any discounts applied.                                                                                                                                                                                                                                                                     |
| `DiscountPercentage`                                                                                                                                                                                                                                                                                | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Percentage rate (from 0 to 100) of any discounts applied to the unit amount.                                                                                                                                                                                                                        |
| `ItemRef`                                                                                                                                                                                                                                                                                           | [*ItemRef](../../models/shared/itemref.md)                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                                                                                                 |
| `Quantity`                                                                                                                                                                                                                                                                                          | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Number of units that have been ordered.                                                                                                                                                                                                                                                             |
| `SubTotal`                                                                                                                                                                                                                                                                                          | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Amount of the line, inclusive of discounts but exclusive of tax.                                                                                                                                                                                                                                    |
| `TaxAmount`                                                                                                                                                                                                                                                                                         | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Amount of tax for the line.                                                                                                                                                                                                                                                                         |
| `TaxRateRef`                                                                                                                                                                                                                                                                                        | [*TaxRateRef](../../models/shared/taxrateref.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Data types that reference a tax rate, for example invoice and bill line items, use a taxRateRef that includes the ID and name of the linked tax rate.<br/><br/>Found on:<br/><br/>- Bill line items<br/>- Bill Credit Note line items<br/>- Credit Note line items<br/>- Direct incomes line items<br/>- Invoice line items<br/>- Items |
| `TotalAmount`                                                                                                                                                                                                                                                                                       | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Total amount of the line, inclusive of discounts and tax.                                                                                                                                                                                                                                           |
| `Tracking`                                                                                                                                                                                                                                                                                          | [*SalesOrderLineItemTracking](../../models/shared/salesorderlineitemtracking.md)                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                                                                                                 |
| `UnitAmount`                                                                                                                                                                                                                                                                                        | **float64*                                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                  | Price of each unit.                                                                                                                                                                                                                                                                                 |