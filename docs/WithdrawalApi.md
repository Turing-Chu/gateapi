<<<<<<< HEAD
# \WithdrawalApi
=======
# WithdrawalApi
>>>>>>> upstream/master

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Withdraw**](WithdrawalApi.md#Withdraw) | **Post** /withdrawals | Withdraw


<<<<<<< HEAD
# **Withdraw**
> LedgerRecord Withdraw(ctx, ledgerRecord)
=======
## Withdraw

> LedgerRecord Withdraw(ctx, ledgerRecord)

>>>>>>> upstream/master
Withdraw

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ledgerRecord** | [**LedgerRecord**](LedgerRecord.md)|  | 
=======
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledgerRecord** | [**LedgerRecord**](LedgerRecord.md)|  | 
>>>>>>> upstream/master

### Example

```golang
<<<<<<< HEAD
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WithdrawalApi
ledgerRecord := new (gateapi.LedgerRecord); // LedgerRecord - 

result, _, err := api.Withdraw(nil, ledgerRecord)
if err != nil {
    fmt.Println(err.Error())
} else {
    fmt.Println(result)
}
```

=======
package main

import (
    "context"
    "fmt"
    "gateapi"
)

func main() {
    client := gateapi.NewAPIClient(gateapi.NewConfiguration())
    // uncomment the next line if your are testing against testnet
    // client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
    ctx := context.WithValue(context.Background(),
                             gateapi.ContextGateAPIV4,
                             gateapi.GateAPIV4{
                                 Key:    "YOUR_API_KEY",
                                 Secret: "YOUR_API_SECRET",
                             }
                            )
    ledgerRecord := gateapi.LedgerRecord{} // LedgerRecord - 
    
    result, _, err := client.WithdrawalApi.Withdraw(ctx, ledgerRecord)
    if err != nil {
        if e, ok := err.(gateapi.GateAPIError); ok {
            fmt.Printf("gate api error: %s\n", e.Error())
        } else {
            fmt.Printf("generic error: %s\n", err.Error())
        }
    } else {
        fmt.Println(result)
    }
}
```


>>>>>>> upstream/master
### Return type

[**LedgerRecord**](LedgerRecord.md)

### Authorization

<<<<<<< HEAD
Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

=======
[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
>>>>>>> upstream/master
