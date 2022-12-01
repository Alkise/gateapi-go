package gateapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// SubAccountsAPIService SubAccountsAPI service
type SubAccountsAPIService service

func (a *SubAccountsAPIService) CreateSubAccount(ctx context.Context, subAccount SubAccount) (SubAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SubAccount
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &subAccount
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) ListSubAccounts(ctx context.Context) ([]SubAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SubAccount
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) GetSubAccount(ctx context.Context, userId int64) (SubAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SubAccount
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) CreateAPIKeyOfSubAccount(ctx context.Context, key APIV4Key) (APIV4Key, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  APIV4Key
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(key.UserId, "")), -1)
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &key
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) ListAllAPIKeysOfSubAccount(ctx context.Context, userId int64) ([]APIV4Key, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []APIV4Key
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) UpdateAPIKeyOfSubAccount(ctx context.Context, userId, key string, apiV4Key APIV4Key) (APIV4Key, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  APIV4Key
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/keys/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.QueryEscape(parameterToString(key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &apiV4Key
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) DeleteAPIKeyOfSubAccount(ctx context.Context, userId int64, key string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/keys/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.QueryEscape(parameterToString(key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) GetAPIKeyOfSubAccount(ctx context.Context, userId int64, key string) (APIV4Key, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  APIV4Key
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/keys/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.QueryEscape(parameterToString(key, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) LockSubAccount(ctx context.Context, userId int64) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

func (a *SubAccountsAPIService) UnlockSubAccount(ctx context.Context, userId int64) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/sub_accounts/{user_id}/unlock"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.QueryEscape(parameterToString(userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
