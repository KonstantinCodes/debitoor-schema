# debitoor-schema
Collection of Go Structs for Debitoor API

This Project is in no way endorsed or supported by Debitoor.
Also there is no guarantee that any of the code works like it is intended to.

## Usage
Example:
```go
package main
import "github.com/KonstantinCodes/debitoor-schema"

func main() {
    draft_invoice := debitoor.DraftInvoice{
            PaymentTermsId: 4,
            Lines: []debitoor.Lines{
                {
                    TaxEnabled: true,
                    TaxRate: 19,
                    Quantity: 1,
                    UnitId: "12",
                    UnitName: "Bit",
                    UnitNetPrice: 5,
                    UnitGrossPrice: 5.95,
                    ProductSku: "RB16",
                    ProductName: "Rando-Bits",
                    Description: "Randomly generated Bits.",
                },
            },
        }
    ...
}
```

## Contribution / Transforming schemas
As of writing this document, schemas provided through the api denote required fields with the `required` boolean in each property.
This leads to JSON-Schema Draft 0.4 parsing problems. See [Section 6.17 of draft-wright-json-schema-validation-01](http://json-schema.org/latest/json-schema-validation.html#rfc.section.6.17). To solve this issue, transform these markers into a `required` array containing the names of the properties under the main object, as specified.

Since Go required variables to have a type, we can not support multiple possible types. In the case that debitoor allows either *"null"* or *"string"* for example, we must choose *"string"*. Should *"string"* and *"number"* be allowed, this project chooses to take the *"string"* type, as it is able to contain both values.