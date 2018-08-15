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

type NamespaceRoutesApiService service

// NamespaceRoutesApiService Get namespace information
// param namespaceId Namespace identifier.
// Returns NamespaceInfo for a given namespaceId.
// return NamespaceInfoDto
func (a *NamespaceRoutesApiService) GetNamespace(namespaceId string) (NamespaceInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     NamespaceInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/namespace/{namespaceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespaceId"+"}", fmt.Sprintf("%v", namespaceId), -1)

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

// NamespaceRoutesApiService Get namespaces an account owns
// Returns an array of NamespaceInfo for an account.
// param accountId Account address or public key.
// param "pageSize" (int32) The number of namespaces to return.
// param "id" (string) Identifier of the namespace after which we want the transactions to be returned.
// return []NamespaceInfoDto
func (a *NamespaceRoutesApiService) GetNamespacesFromAccount(accountId string, pageSize int,
	id string) ([]NamespaceInfoDto,
	error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []NamespaceInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{accountId}/namespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

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

// NamespaceRoutesApiService Get namespaces information
// param addresses Accounts address array.
// param "pageSize" (int32) The number of namespaces to return.
// param "id" (string) Identifier of the namespace after which we want the transactions to be returned.
// Returns an array of NamespaceInfo for a given set of addresses.
// return []NamespaceInfoDto
func (a *NamespaceRoutesApiService) GetNamespacesFromAccounts(addresses Addresses, pageSize int,
	id string) ([]NamespaceInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []NamespaceInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/namespaces"

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
	// body params
	localVarPostBody = &addresses
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

// NamespaceRoutesApiService Get readable names for a set of namespaces
// Returns names for namespaces.
// param namespaceIds Array of namespaceIds.
// return []NamespaceNameDto
func (a *NamespaceRoutesApiService) GetNamespacesNames(namespaceIds NamespaceIds) ([]NamespaceNameDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []NamespaceNameDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/namespace/names"

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
	localVarPostBody = &namespaceIds
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
