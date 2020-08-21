# LedgerRecord

## Properties
<<<<<<< HEAD
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] 
**Txid** | **string** | Hash record of the withdrawal | [optional] 
**Timestamp** | **string** | Record time | [optional] 
=======

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Record ID | [optional] [readonly] 
**Txid** | **string** | Hash record of the withdrawal | [optional] [readonly] 
**Timestamp** | **string** | Record time | [optional] [readonly] 
>>>>>>> upstream/master
**Amount** | **string** | Trade amount | 
**Currency** | **string** | Record currency | 
**Address** | **string** | Withdrawal address. Required for withdrawals | [optional] 
**Memo** | **string** | Extra withdrawal memo | [optional] 
<<<<<<< HEAD
**Status** | **string** | Record status.  - DONE: done - CANCEL: cancelled - REQUEST: requesting - MANUAL: waiting for manual approval - BCODE: GateCode operation | [optional] 
=======
**Status** | **string** | Record status.  - DONE: done - CANCEL: cancelled - REQUEST: requesting - MANUAL: waiting for manual approval - BCODE: GateCode operation | [optional] [readonly] 
>>>>>>> upstream/master

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


