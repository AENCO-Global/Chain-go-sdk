# \NamespaceRoutesApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNamespace**](NamespaceRoutesApi.md#GetNamespace) | **Get** /namespace/{namespaceId} | Get namespace information
[**GetNamespacesFromAccount**](NamespaceRoutesApi.md#GetNamespacesFromAccount) | **Get** /account/{accountId}/namespaces | Get namespaces an account owns
[**GetNamespacesFromAccounts**](NamespaceRoutesApi.md#GetNamespacesFromAccounts) | **Post** /account/namespaces | Get namespaces information
[**GetNamespacesNames**](NamespaceRoutesApi.md#GetNamespacesNames) | **Post** /namespace/names | Get readable names for a set of namespaces


# **GetNamespace**
> NamespaceInfoDto GetNamespace(ctx, namespaceId)
Get namespace information

Returns NamespaceInfo for a given namespaceId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **namespaceId** | **string**| Namespace identifier. | 

### Return type

[**NamespaceInfoDto**](NamespaceInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacesFromAccount**
> []NamespaceInfoDto GetNamespacesFromAccount(ctx, accountId, optional)
Get namespaces an account owns

Returns an array of NamespaceInfo for an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accountId** | **string**| Account address or public key. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string**| Account address or public key. | 
 **pageSize** | **int32**| The number of namespaces to return. | 
 **id** | **string**| Identifier of the namespace after which we want the transactions to be returned. | 

### Return type

[**[]NamespaceInfoDto**](NamespaceInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacesFromAccounts**
> []NamespaceInfoDto GetNamespacesFromAccounts(ctx, addresses, optional)
Get namespaces information

Returns an array of NamespaceInfo for a given set of addresses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **addresses** | [**Addresses**](Addresses.md)| Accounts address array. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addresses** | [**Addresses**](Addresses.md)| Accounts address array. | 
 **pageSize** | **int32**| The number of namespaces to return. | 
 **id** | **string**| Identifier of the namespace after which we want the transactions to be returned. | 

### Return type

[**[]NamespaceInfoDto**](NamespaceInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacesNames**
> []NamespaceNameDto GetNamespacesNames(ctx, namespaceIds)
Get readable names for a set of namespaces

Returns names for namespaces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **namespaceIds** | [**NamespaceIds**](NamespaceIds.md)| Array of namespaceIds. | 

### Return type

[**[]NamespaceNameDto**](NamespaceNameDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

