/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsAdvertiserStoreSearchV2ApiService ToolsAdvertiserStoreSearchV2Api service
type ToolsAdvertiserStoreSearchV2ApiService service

type ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdvertiserStoreSearchV2ApiService
	advertiserId *int64
	page         *int64
	pageSize     *int64
	type_        *ToolsAdvertiserStoreSearchV2Type
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) Page(page int64) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) Type_(type_ ToolsAdvertiserStoreSearchV2Type) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	r.type_ = &type_
	return r
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) Execute() (*ToolsAdvertiserStoreSearchV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdvertiserStoreSearchGet Method for OpenApi2ToolsAdvertiserStoreSearchGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest
*/
func (a *ToolsAdvertiserStoreSearchV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest {
	return &ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdvertiserStoreSearchV2Response
func (a *ToolsAdvertiserStoreSearchV2ApiService) getExecute(r *ApiOpenApi2ToolsAdvertiserStoreSearchGetRequest) (*ToolsAdvertiserStoreSearchV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdvertiserStoreSearchV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/advertiser_store/search/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_)
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
