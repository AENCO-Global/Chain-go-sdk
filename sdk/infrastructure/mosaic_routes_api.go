package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type MosaicRoutesApiService service

// MosaicRoutesApiService Get mosaic information
// Returns MosaicInfo for a given mosaicId.
// param mosaicId Mosaic identifier.
// return MosaicInfoDto
func (a *MosaicRoutesApiService) GetMosaic(mosaicId string) (MosaicInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     MosaicInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mosaic/{mosaicId}"
	localVarPath = strings.Replace(localVarPath, "{"+"mosaicId"+"}", fmt.Sprintf("%v", mosaicId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, reportError("{\"Status\": %v, \"Body\": %s}", localVarHttpResponse.StatusCode, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, err
	}

	return successPayload, err
}

// MosaicRoutesApiService Get information for a set of mosaics
// Returns MosaicInfo for a given set of mosaicIds.
// param mosaicIds Array of mosaicIds.
// return []MosaicInfoDto
func (a *MosaicRoutesApiService) GetMosaics(mosaicIds MosaicIds) ([]MosaicInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []MosaicInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mosaic"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &mosaicIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, reportError("{\"Status\": %v, \"Body\": %s}", localVarHttpResponse.StatusCode, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, err
	}

	return successPayload, err
}

// MosaicRoutesApiService Get mosaics information.
// Returns an array of MosaicInfo from mosaics created under provided namespace.
// param namespaceId Namespace identifier.
// param "pageSize" (int32) The number of mosaics to return.
// param "id" (string) Identifier of the mosaic after which we want the transactions to be returned.
// return []MosaicInfoDto
func (a *MosaicRoutesApiService) GetMosaicsFromNamespace(namespaceId string, pageSize int, id string) ([]MosaicInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []MosaicInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/namespace/{namespaceId}/mosaics"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceId"+"}", fmt.Sprintf("%v", namespaceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if !utils.IsEmpty(pageSize) {
		localVarQueryParams.Add("pageSize", parameterToString(pageSize, ""))
	}

	if !utils.IsEmpty(id) {
		localVarQueryParams.Add("id", parameterToString(id, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, reportError("{\"Status\": %v, \"Body\": %s}", localVarHttpResponse.StatusCode, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, err
	}

	return successPayload, err
}

// MosaicRoutesApiService Get readable names for a set of mosaics
// Returns names for mosaics.
// param mosaicIds Array of mosaicIds.
// return []MosaicNameDto
func (a *MosaicRoutesApiService) GetMosaicsName(mosaicIds MosaicIds) ([]MosaicNameDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []MosaicNameDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/mosaic/names"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &mosaicIds
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, reportError("{\"Status\": %v, \"Body\": %s}", localVarHttpResponse.StatusCode, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, err
	}

	return successPayload, err
}
