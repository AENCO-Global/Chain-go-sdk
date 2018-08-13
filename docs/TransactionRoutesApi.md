# \TransactionRoutesApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnounceCosignatureTransaction**](TransactionRoutesApi.md#AnnounceCosignatureTransaction) | **Put** /transaction/cosignature | Announce a cosignature transaction
[**AnnouncePartialTransaction**](TransactionRoutesApi.md#AnnouncePartialTransaction) | **Put** /transaction/partial | Announce an aggregate bonded transaction
[**AnnounceTransaction**](TransactionRoutesApi.md#AnnounceTransaction) | **Put** /transaction | Announce a  new transaction
[**GetTransaction**](TransactionRoutesApi.md#GetTransaction) | **Get** /transaction/{transactionId} | Get transaction information
[**GetTransactionStatus**](TransactionRoutesApi.md#GetTransactionStatus) | **Get** /transaction/{hash}/status | Get transaction status
[**GetTransactions**](TransactionRoutesApi.md#GetTransactions) | **Post** /transaction | Get transactions information
[**GetTransactionsStatuses**](TransactionRoutesApi.md#GetTransactionsStatuses) | **Post** /transaction/statuses | Get transactions status.


# **AnnounceCosignatureTransaction**
> interface{} AnnounceCosignatureTransaction(ctx, payload)
Announce a cosignature transaction

Announces a cosignature transaction to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **payload** | [**TransactionPayload**](TransactionPayload.md)| Transaction payload. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnnouncePartialTransaction**
> interface{} AnnouncePartialTransaction(ctx, payload)
Announce an aggregate bonded transaction

Announces an aggregate bonded transaction to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **payload** | [**TransactionPayload**](TransactionPayload.md)| Transaction payload. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnnounceTransaction**
> interface{} AnnounceTransaction(ctx, payload)
Announce a  new transaction

Announces a transaction to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **payload** | [**TransactionPayload**](TransactionPayload.md)| Transaction payload. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> interface{} GetTransaction(ctx, transactionId)
Get transaction information

Returns transaction given its transactionId or hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transactionId** | **string**| TransactionId or hash. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionStatus**
> TransactionStatusDto GetTransactionStatus(ctx, hash)
Get transaction status

Returns transaction status for a given transactionId or hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hash** | **string**| Transaction hash. | 

### Return type

[**TransactionStatusDto**](TransactionStatusDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactions**
> []interface{} GetTransactions(ctx, transactionIds)
Get transactions information

Returns transaction information for a given set of transactionId or hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transactionIds** | [**TransactionIds**](TransactionIds.md)| Array of transactionIds or hashes. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionsStatuses**
> []TransactionStatusDto GetTransactionsStatuses(ctx, transactionHashes)
Get transactions status.

Returns an array of transaction statuses for a given set of transactionId or hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **transactionHashes** | [**TransactionHashes**](TransactionHashes.md)| Array of transactionIds or hashes. | 

### Return type

[**[]TransactionStatusDto**](TransactionStatusDto.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

