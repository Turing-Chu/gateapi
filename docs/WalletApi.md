# WalletApi

All URIs are relative to *https://api.gateio.ws/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDepositAddress**](WalletApi.md#GetDepositAddress) | **Get** /wallet/deposit_address | Generate currency deposit address
<<<<<<< HEAD
[**ListDeposits**](WalletApi.md#ListDeposits) | **Get** /wallet/deposits | Retrieve deposit records. Time range cannot exceed 30 days
[**ListWithdrawals**](WalletApi.md#ListWithdrawals) | **Get** /wallet/withdrawals | Retrieve withdrawal records. Time range cannot exceed 30 days
[**Transfer**](WalletApi.md#Transfer) | **Post** /wallet/transfers | Transfer between accounts
[**TransferWithSubAccount**](WalletApi.md#TransferWithSubAccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts


# **GetDepositAddress**
> DepositAddress GetDepositAddress(ctx, currency)
=======
[**ListWithdrawals**](WalletApi.md#ListWithdrawals) | **Get** /wallet/withdrawals | Retrieve withdrawal records
[**ListDeposits**](WalletApi.md#ListDeposits) | **Get** /wallet/deposits | Retrieve deposit records
[**Transfer**](WalletApi.md#Transfer) | **Post** /wallet/transfers | Transfer between trading accounts
[**ListSubAccountTransfers**](WalletApi.md#ListSubAccountTransfers) | **Get** /wallet/sub_account_transfers | Transfer records between main and sub accounts
[**TransferWithSubAccount**](WalletApi.md#TransferWithSubAccount) | **Post** /wallet/sub_account_transfers | Transfer between main and sub accounts


## GetDepositAddress

> DepositAddress GetDepositAddress(ctx, currency)

>>>>>>> upstream/master
Generate currency deposit address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currency** | **string**| Currency name | 
=======
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string**| Currency name | 
>>>>>>> upstream/master

### Example

```golang
<<<<<<< HEAD
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi
currency := "currency_example"; // string - Currency name

result, _, err := api.GetDepositAddress(nil, currency)
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
    currency := "currency_example" // string - Currency name
    
    result, _, err := client.WalletApi.GetDepositAddress(ctx, currency)
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

[**DepositAddress**](DepositAddress.md)

### Authorization

<<<<<<< HEAD
Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeposits**
> []LedgerRecord ListDeposits(ctx, optional)
Retrieve deposit records. Time range cannot exceed 30 days
=======
[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListWithdrawals

> []LedgerRecord ListWithdrawals(ctx, optional)

Retrieve withdrawal records

Record time range cannot exceed 30 days
>>>>>>> upstream/master

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
 **from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
 **to** | **optional.Int64**| Time range ending, default to current time | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
=======
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListWithdrawalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWithdrawalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
**from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
>>>>>>> upstream/master

### Example

```golang
<<<<<<< HEAD
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi

result, _, err := api.ListDeposits(nil, nil)
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
    
    result, _, err := client.WalletApi.ListWithdrawals(ctx, nil)
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

[**[]LedgerRecord**](LedgerRecord.md)

### Authorization

<<<<<<< HEAD
Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWithdrawals**
> []LedgerRecord ListWithdrawals(ctx, optional)
Retrieve withdrawal records. Time range cannot exceed 30 days
=======
[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListDeposits

> []LedgerRecord ListDeposits(ctx, optional)

Retrieve deposit records

Record time range cannot exceed 30 days
>>>>>>> upstream/master

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
<<<<<<< HEAD
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWithdrawalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListWithdrawalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
 **from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
 **to** | **optional.Int64**| Time range ending, default to current time | 
 **limit** | **optional.Int32**| Maximum number of record returned in one list | [default to 100]
 **offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
=======
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**currency** | **optional.String**| Filter by currency. Return all currency records if not specified | 
**from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]
>>>>>>> upstream/master

### Example

```golang
<<<<<<< HEAD
client := gateapi.NewAPIClient(gateapi.NewConfiguration())
// uncomment the next line if your are testing against other hosts
// client.ChangeBasePath("https://some-other-host")
client.SetKeySecret("YOUR API KEY", "YOUR API SECRET")
api := client.WalletApi

result, _, err := api.ListWithdrawals(nil, nil)
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
    
    result, _, err := client.WalletApi.ListDeposits(ctx, nil)
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

[**[]LedgerRecord**](LedgerRecord.md)

### Authorization

<<<<<<< HEAD
Authentication with API key and secret is required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Transfer**
=======
[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Transfer

>>>>>>> upstream/master
> Transfer(ctx, transfer)

<<<<<<< HEAD
Transfer between different accounts. Currently support transfers between the following:  1. spot - margin 2. spot - futures
=======
Transfer between trading accounts

Transfer between different accounts. Currently support transfers between the following:  1. spot - margin 2. spot - futures(perpetual) 3. spot - delivery
>>>>>>> upstream/master

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transfer** | [**Transfer**](Transfer.md)|  | 

### Example

```golang
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
    transfer := gateapi.Transfer{} // Transfer - 
    
    result, _, err := client.WalletApi.Transfer(ctx, transfer)
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


### Return type

 (empty response body)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSubAccountTransfers

> []SubAccountTransfer ListSubAccountTransfers(ctx, optional)

Transfer records between main and sub accounts

Record time range cannot exceed 30 days  > Note: only records after 2020-04-10 can be retrieved

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | **ListSubAccountTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubAccountTransfersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**subUid** | **optional.String**| Sub account user ID. Return records related to all sub accounts if not specified | 
**from** | **optional.Int64**| Time range beginning, default to 7 days before current time | 
**to** | **optional.Int64**| Time range ending, default to current time | 
**limit** | **optional.Int32**| Maximum number of records returned in one list | [default to 100]
**offset** | **optional.Int32**| List offset, starting from 0 | [default to 0]

### Example

```golang
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
    
    result, _, err := client.WalletApi.ListSubAccountTransfers(ctx, nil)
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


### Return type

[**[]SubAccountTransfer**](SubAccountTransfer.md)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TransferWithSubAccount

> TransferWithSubAccount(ctx, subAccountTransfer)

Transfer between main and sub accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subAccountTransfer** | [**SubAccountTransfer**](SubAccountTransfer.md)|  | 

### Example

```golang
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
    subAccountTransfer := gateapi.SubAccountTransfer{} // SubAccountTransfer - 
    
    result, _, err := client.WalletApi.TransferWithSubAccount(ctx, subAccountTransfer)
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


### Return type

 (empty response body)

### Authorization

[apiv4](../README.md#apiv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
