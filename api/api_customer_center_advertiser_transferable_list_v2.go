/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

// CustomerCenterAdvertiserTransferableListV2ApiService CustomerCenterAdvertiserTransferableListV2Api service
type CustomerCenterAdvertiserTransferableListV2ApiService service

type ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest struct {
	ctx                 context.Context
	ApiService          *CustomerCenterAdvertiserTransferableListV2ApiService
	advertiserId        *int64
	transferAccountType *CustomerCenterAdvertiserTransferableListV2TransferAccountType
	page                *int32
	pageSize            *int32
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) TransferAccountType(transferAccountType CustomerCenterAdvertiserTransferableListV2TransferAccountType) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	r.transferAccountType = &transferAccountType
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) Page(page int32) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) PageSize(pageSize int32) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) Execute() (*CustomerCenterAdvertiserTransferableListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) AccessToken(accessToken string) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) WithLog(enable bool) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CustomerCenterAdvertiserTransferableListGet Method for OpenApi2CustomerCenterAdvertiserTransferableListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest
*/
func (a *CustomerCenterAdvertiserTransferableListV2ApiService) Get(ctx context.Context) *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest {
	return &ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CustomerCenterAdvertiserTransferableListV2Response
func (a *CustomerCenterAdvertiserTransferableListV2ApiService) getExecute(r *ApiOpenApi2CustomerCenterAdvertiserTransferableListGetRequest) (*CustomerCenterAdvertiserTransferableListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CustomerCenterAdvertiserTransferableListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/customer_center/advertiser/transferable/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.transferAccountType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transfer_account_type", r.transferAccountType)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
