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

type AccountRoutesApiService service

// AccountRoutesApiService Get account information
// Returns AccountInfo for an account.
// param accountId Account address or publicKey.
// return AccountInfoDto
func (a *AccountRoutesApiService) GetAccountInfo(accountId string) (AccountInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     AccountInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{accountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

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

// AccountRoutesApiService Get multisig account information
// param ctx context.Context for authentication, logging, tracing, etc.
// param accountId Account address or public key.
// Returns MultisigAccountInfo for an account.
// return MultisigAccountInfoDto
func (a *AccountRoutesApiService) GetAccountMultisig(accountId string) (MultisigAccountInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     MultisigAccountInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{accountId}/multisig"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

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

// AccountRoutesApiService Get multisig account graph information
// param accountId Account address or public key.
// Returns MultisigAccountGraphInfo for an account.
// return []MultisigAccountGraphInfoDto
func (a *AccountRoutesApiService) GetAccountMultisigGraph(accountId string) ([]MultisigAccountGraphInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []MultisigAccountGraphInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{accountId}/multisig/graph"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

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

// AccountRoutesApiService Get accounts information
// param addresses Slice of addresses.
// Returns AccountsInfo for different accounts.
// return []AccountInfoDto
func (a *AccountRoutesApiService) GetAccountsInfo(addresses Addresses) ([]AccountInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []AccountInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account"

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

// AccountRoutesApiService Get incoming transactions information
// Gets an array of transactions for which an account is the recipient.
// A transaction is said to be incoming regarding an account if the account is the recipient of a transaction.
// param publicKey Account publicKey.
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// param "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// return []TransactionInfoDto
func (a *AccountRoutesApiService) IncomingTransactions(publicKey string, pageSize int, id string) ([]TransactionInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{publicKey}/transactions/incoming"
	localVarPath = strings.Replace(localVarPath, "{"+"publicKey"+"}", fmt.Sprintf("%v", publicKey), -1)

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

// AccountRoutesApiService Get outgoing transactions information
// Gets an array of transactions for which an account is the sender.
// A transaction is said to be outgoing egarding an account if the account is the sender of a transaction.
// param publicKey Account publicKey.
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// param "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// return []TransactionInfoDto
func (a *AccountRoutesApiService) OutgoingTransactions(publicKey string, pageSize int, id string) ([]TransactionInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{publicKey}/transactions/outgoing"
	localVarPath = strings.Replace(localVarPath, "{"+"publicKey"+"}", fmt.Sprintf("%v", publicKey), -1)

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

// AccountRoutesApiService Get aggregate bonded transactions information
// Gets an array of aggregate bonded transactions for which an account is the sender or
// has signed the transaction. A transaction is said to be aggregate bonded regarding an
// account if announced but there are missing signatures.
// param publicKey Account publicKey.
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// aram "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// return []TransactionInfoDto
func (a *AccountRoutesApiService) PartialTransactions(publicKey string, pageSize int, id string) ([]TransactionInfoDto,
	error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{publicKey}/transactions/partial"
	localVarPath = strings.Replace(localVarPath, "{"+"publicKey"+"}", fmt.Sprintf("%v", publicKey), -1)

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

// AccountRoutesApiService Get confirmed transactions information
// Gets an array of confirmed transaction for which an account is signer or recipient.
// param publicKey Account publicKey.
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// param "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// return []TransactionInfoDto
func (a *AccountRoutesApiService) Transactions(publicKey string, pageSize int, id string) ([]TransactionInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{publicKey}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"publicKey"+"}", fmt.Sprintf("%v", publicKey), -1)

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

// AccountRoutesApiService Get unconfirmed transactions information
// Gets the array of transactions for which an account is the sender or
// receiver and which have not yet been included in a block.
// param publicKey Account publicKey.
// param "pageSize" (int32) The number of transactions to return. Should be between 10 and 100, otherwise 10.
// param "id" (string) Identifier of the transaction after which we want the transactions to be returned.
// return []TransactionInfoDto
func (a *AccountRoutesApiService) UnconfirmedTransactions(publicKey string, pageSize int, id string) ([]TransactionInfoDto, error) {
	var (
		ctx                context.Context
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []TransactionInfoDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/{publicKey}/transactions/unconfirmed"
	localVarPath = strings.Replace(localVarPath, "{"+"publicKey"+"}", fmt.Sprintf("%v", publicKey), -1)

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
