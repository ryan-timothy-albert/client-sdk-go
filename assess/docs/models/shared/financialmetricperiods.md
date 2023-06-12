# FinancialMetricPeriods


## Fields

| Field                                                                                                                                                                                           | Type                                                                                                                                                                                            | Required                                                                                                                                                                                        | Description                                                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Errors`                                                                                                                                                                                        | [][FinancialMetricPeriodsErrors](../../models/shared/financialmetricperiodserrors.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                                              | N/A                                                                                                                                                                                             |
| `FromDate`                                                                                                                                                                                      | [*types.Date](../../types/date.md)                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                              | The date from which the report starts.                                                                                                                                                          |
| `Inputs`                                                                                                                                                                                        | [][FinancialMetricPeriodsInputs](../../models/shared/financialmetricperiodsinputs.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                                              | N/A                                                                                                                                                                                             |
| `ToDate`                                                                                                                                                                                        | [*types.Date](../../types/date.md)                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                              | The date on which the report ends (inclusive of day).                                                                                                                                           |
| `Value`                                                                                                                                                                                         | **float64*                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                              | The top level metric value that is calculated for the specified period. <br/><br/>If the system cannot calculate for that period, the value will be null. The system will still show the metric inputs. |