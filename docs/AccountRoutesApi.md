# \AccountRoutesApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInfo**](AccountRoutesApi.md#GetAccountInfo) | **Get** /account/{accountId} | Get account information
[**GetAccountMultisig**](AccountRoutesApi.md#GetAccountMultisig) | **Get** /account/{accountId}/multisig | Get multisig account information
[**GetAccountMultisigGraph**](AccountRoutesApi.md#GetAccountMultisigGraph) | **Get** /account/{accountId}/multisig/graph | Get multisig account graph information
[**GetAccountsInfo**](AccountRoutesApi.md#GetAccountsInfo) | **Post** /account | Get accounts information
[**IncomingTransactions**](AccountRoutesApi.md#IncomingTransactions) | **Get** /account/{publicKey}/transactions/incoming | Get incoming transactions information
[**OutgoingTransactions**](AccountRoutesApi.md#OutgoingTransactions) | **Get** /account/{publicKey}/transactions/outgoing | Get outgoing transactions information
[**PartialTransactions**](AccountRoutesApi.md#PartialTransactions) | **Get** /account/{publicKey}/transactions/partial | Get aggregate bonded transactions information
[**Transactions**](AccountRoutesApi.md#Transactions) | **Get** /account/{publicKey}/transactions | Get confirmed transactions information
[**UnconfirmedTransactions**](AccountRoutesApi.md#UnconfirmedTransactions) | **Get** /account/{publicKey}/transactions/unconfirmed | Get unconfirmed transactions information


# **GetAccountInfo**
> AccountInfoDto GetAccountInfo(ctx, accountId)
Get account information

Returns AccountInfo for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Account address or publicKey. | 

### Return type

[**AccountInfoDto**](AccountInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountMultisig**
> MultisigAccountInfoDto GetAccountMultisig(ctx, accountId)
Get multisig account information

Returns MultisigAccountInfo for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Account address or public key. | 

### Return type

[**MultisigAccountInfoDto**](MultisigAccountInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountMultisigGraph**
> []MultisigAccountGraphInfoDto GetAccountMultisigGraph(ctx, accountId)
Get multisig account graph information

Returns MultisigAccountGraphInfo for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Account address or public key. | 

### Return type

[**[]MultisigAccountGraphInfoDto**](MultisigAccountGraphInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountsInfo**
> []AccountInfoDto GetAccountsInfo(ctx, addresses)
Get accounts information

Returns AccountsInfo for different accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **addresses** | [**Addresses**](Addresses.md)| Array of addresses. | 

### Return type

[**[]AccountInfoDto**](AccountInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncomingTransactions**
> []interface{} IncomingTransactions(ctx, publicKey, optional)
Get incoming transactions information

Gets an array of transactions for which an account is the recipient. A transaction is said to be incoming regarding an account if the account is the recipient of a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **publicKey** | **string**| Account publicKey. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string**| Account publicKey. | 
 **pageSize** | **int32**| The number of transactions to return. Should be between 10 and 100, otherwise 10. | 
 **id** | **string**| Identifier of the transaction after which we want the transactions to be returned. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OutgoingTransactions**
> []interface{} OutgoingTransactions(ctx, publicKey, optional)
Get outgoing transactions information

Gets an array of transactions for which an account is the sender. A transaction is said to be outgoing regarding an account if the account is the sender of a transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **publicKey** | **string**| Account publicKey. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string**| Account publicKey. | 
 **pageSize** | **int32**| The number of transactions to return. Should be between 10 and 100, otherwise 10. | 
 **id** | **string**| Identifier of the transaction after which we want the transactions to be returned. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartialTransactions**
> []interface{} PartialTransactions(ctx, publicKey, optional)
Get aggregate bonded transactions information

Gets an array of aggregate bonded transactions for which an account is the sender or has signed the transaction. A transaction is said to be aggregate bonded regarding an account if announced but there are missing signatures.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **publicKey** | **string**| Account publicKey. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string**| Account publicKey. | 
 **pageSize** | **int32**| The number of transactions to return. Should be between 10 and 100, otherwise 10. | 
 **id** | **string**| Identifier of the transaction after which we want the transactions to be returned. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Transactions**
> []interface{} Transactions(ctx, publicKey, optional)
Get confirmed transactions information

Gets an array of confirmed transaction for which an account is signer or recipient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **publicKey** | **string**| Account publicKey. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string**| Account publicKey. | 
 **pageSize** | **int32**| The number of transactions to return. Should be between 10 and 100, otherwise 10. | 
 **id** | **string**| Identifier of the transaction after which we want the transactions to be returned. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnconfirmedTransactions**
> []interface{} UnconfirmedTransactions(ctx, publicKey, optional)
Get unconfirmed transactions information

Gets the array of transactions for which an account is the sender or receiver and which have not yet been included in a block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **publicKey** | **string**| Account publicKey. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string**| Account publicKey. | 
 **pageSize** | **int32**| The number of transactions to return. Should be between 10 and 100, otherwise 10. | 
 **id** | **string**| Identifier of the transaction after which we want the transactions to be returned. | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

