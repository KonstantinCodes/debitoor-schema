# debitoor-schema
Collection of Go Structs for Debitoor API

This Project is in no way endorsed or supported by Debitoor.
Also there is no guarantee that any of the code works like it is intended to.

## Usage
Example how to make a draft invoice in version 3:
```go
package main
import (
    "encoding/json"
    "github.com/KonstantinCodes/debitoor-schema"
)

func main() {
    draftInvoice := debitoor.SalesDraftinvoiceV3DraftInvoice{
            PaymentTermsId: 4,
            Lines: []debitoor.SalesDraftinvoiceV3Lines{
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
        
    draftInvoiceJson, err = json.Marshal(draftInvoice)
    ...
}
```

## Contribution
As of writing this document, schemas provided through the api denote required fields with the `required` boolean in each property.
This leads to JSON-Schema Draft 0.4 parsing problems. See [Section 6.17 of draft-wright-json-schema-validation-01](http://json-schema.org/latest/json-schema-validation.html#rfc.section.6.17). To solve this issue, transform these markers into a `required` array containing the names of the properties under the main object, as specified.

Since Go requires variables to have a type, we can not support multiple possible types. In the case that debitoor allows either *"null"* or *"string"* for example, we must choose *"string"*. Should *"string"* and *"number"* be allowed, this project chooses to take the *"string"* type, as it is able to contain both values.

Also, the *paymentaccount/v1* schema includes the keyword "oneOf", which is not part of the Spec. To work around that, just spilt listed objects into separate files.

## Transforming schemas automatically
Usage of [converter/main.go](converter/main.go):
 * Get JSON Schema: `curl https://api.debitoor.com/api/schemas/customer/v1 > schema/customer_v1.json`
 * `go run converter/main.go`
 * There should now be a new compliant JSON Schema File named `schema-processed/customer_v1.json`