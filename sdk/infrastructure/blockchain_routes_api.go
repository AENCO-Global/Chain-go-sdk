package infrastructure

import (
	"io/ioutil"
	"net/url"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

// Linger please
var (
	_ context.Context
)

type BlockchainRoutesApiService service

// BlockchainRoutesApiService Get block information
// Returns BlockInfo for a given block height.
// param height Block height.
// return BlockInfoDto
func (a *BlockchainRoutesApiService) GetBlockByHeight(height int64) (BlockInfoDto, error) {

	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  BlockInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/block/{height}"
	localVarPath = strings.Replace(localVarPath, "{"+"height"+"}", fmt.Sprintf("%v", height), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

// BlockchainRoutesApiService Get transactions from a block
// param height Block height
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// param "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// Returns array of transactions included in a block for a block height.
// return []TransactionInfoDto
func (a *BlockchainRoutesApiService) GetBlockTransactions(height int64, pageSize int, id string) ([]TransactionInfoDto,
error) {
	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/block/{height}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"height"+"}", fmt.Sprintf("%v", height), -1)

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
	localVarHttpContentTypes := []string{  }

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
		return nil, err
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

// BlockchainRoutesApiService Get the current height of the chain
// Returns the current blockchain height.
// return HeightDto
func (a *BlockchainRoutesApiService) GetBlockchainHeight() (HeightDto, error) {
	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  HeightDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chain/height"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

// BlockchainRoutesApiService Get the current score of the chain
// Returns the current chain score.
// return BlockchainScoreDto
func (a *BlockchainRoutesApiService) GetBlockchainScore() (BlockchainScoreDto, error) {
	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  BlockchainScoreDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/chain/score"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

// BlockchainRoutesApiService Get blocks information
// Returns an array of BlockInfo for a given block height and limit.
// param height Block height.
// param limit Number of following blocks to be returned.
// return []BlockInfoDto
func (a *BlockchainRoutesApiService) GetBlocksByHeightWithLimit(height int64, limit int32) ([]BlockInfoDto, error) {
	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []BlockInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/blocks/{height}/limit/{limit}"
	localVarPath = strings.Replace(localVarPath, "{"+"height"+"}", fmt.Sprintf("%v", height), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if height < 0 {
		return successPayload, reportError("height must be greater than 0")
	}
	if limit < 0 {
		return successPayload, reportError("limit must be greater than 0")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

// BlockchainRoutesApiService Get the storage information
// return BlockchainStorageInfoDto
// Returns statistical information about the blockchain.
func (a *BlockchainRoutesApiService) GetDiagnosticStorage() (BlockchainStorageInfoDto, error) {
	var (
		ctx context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  BlockchainStorageInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/diagnostic/storage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

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

