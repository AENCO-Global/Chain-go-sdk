# \MosaicRoutesApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMosaic**](MosaicRoutesApi.md#GetMosaic) | **Get** /mosaic/{mosaicId} | Get mosaic information
[**GetMosaics**](MosaicRoutesApi.md#GetMosaics) | **Post** /mosaic | Get information for a set of mosaics
[**GetMosaicsFromNamespace**](MosaicRoutesApi.md#GetMosaicsFromNamespace) | **Get** /namespace/{namespaceId}/mosaics | Get mosaics information.
[**GetMosaicsName**](MosaicRoutesApi.md#GetMosaicsName) | **Post** /mosaic/names | Get readable names for a set of mosaics


# **GetMosaic**
> MosaicInfoDto GetMosaic(ctx, mosaicId)
Get mosaic information

Returns MosaicInfo for a given mosaicId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mosaicId** | **string**| Mosaic identifier. | 

### Return type

[**MosaicInfoDto**](MosaicInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMosaics**
> []MosaicInfoDto GetMosaics(ctx, mosaicIds)
Get information for a set of mosaics

Returns MosaicInfo for a given set of mosaicIds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mosaicIds** | [**MosaicIds**](MosaicIds.md)| Array of mosaicIds. | 

### Return type

[**[]MosaicInfoDto**](MosaicInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMosaicsFromNamespace**
> []MosaicInfoDto GetMosaicsFromNamespace(ctx, namespaceId, optional)
Get mosaics information.

Returns an array of MosaicInfo from mosaics created under provided namespace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **namespaceId** | **string**| Namespace identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceId** | **string**| Namespace identifier. | 
 **pageSize** | **int32**| The number of mosaics to return. | 
 **id** | **string**| Identifier of the mosaic after which we want the transactions to be returned. | 

### Return type

[**[]MosaicInfoDto**](MosaicInfoDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMosaicsName**
> []MosaicNameDto GetMosaicsName(ctx, mosaicIds)
Get readable names for a set of mosaics

Returns names for mosaics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mosaicIds** | [**MosaicIds**](MosaicIds.md)| Array of mosaicIds. | 

### Return type

[**[]MosaicNameDto**](MosaicNameDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

