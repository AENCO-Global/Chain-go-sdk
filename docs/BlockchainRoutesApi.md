# \BlockchainRoutesApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockByHeight**](BlockchainRoutesApi.md#GetBlockByHeight) | **Get** /block/{height} | Get block information
[**GetBlockTransactions**](BlockchainRoutesApi.md#GetBlockTransactions) | **Get** /block/{height}/transactions | Get transactions from a block
[**GetBlockchainHeight**](BlockchainRoutesApi.md#GetBlockchainHeight) | **Get** /chain/height | Get the current height of the chain
[**GetBlockchainScore**](BlockchainRoutesApi.md#GetBlockchainScore) | **Get** /chain/score | Get the current score of the chain
[**GetBlocksByHeightWithLimit**](BlockchainRoutesApi.md#GetBlocksByHeightWithLimit) | **Get** /blocks/{height}/limit/{limit} | Get blocks information
[**GetDiagnosticStorage**](BlockchainRoutesApi.md#GetDiagnosticStorage) | **Get** /diagnostic/storage | Get the storage information


# **GetBlockByHeight**
> BlockInfoDto GetBlockByHeight(ctx, height)
Get block information

Returns BlockInfo for a given block height.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **height** | **int64**| Block height. | 

### Return type

[**BlockInfoDto**](BlockInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockTransactions**
> []interface{} GetBlockTransactions(ctx, height, optional)
Get transactions from a block

Returns array of transactions included in a block for a block height.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **height** | **int64**| Block height | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **int64**| Block height | 
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

# **GetBlockchainHeight**
> HeightDto GetBlockchainHeight(ctx, )
Get the current height of the chain

Returns the current blockchain height.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HeightDto**](HeightDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockchainScore**
> BlockchainScoreDto GetBlockchainScore(ctx, )
Get the current score of the chain

Returns the current chain score.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BlockchainScoreDto**](BlockchainScoreDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlocksByHeightWithLimit**
> []BlockInfoDto GetBlocksByHeightWithLimit(ctx, height, limit)
Get blocks information

Returns an array of BlockInfo for a given block height and limit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **height** | **int64**| Block height. | 
  **limit** | **int32**| Number of following blocks to be returned. | 

### Return type

[**[]BlockInfoDto**](BlockInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticStorage**
> BlockchainStorageInfoDto GetDiagnosticStorage(ctx, )
Get the storage information

Returns statistical information about the blockchain.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BlockchainStorageInfoDto**](BlockchainStorageInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

