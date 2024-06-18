/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
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

// CampaignGetV2ApiService CampaignGetV2Api service
type CampaignGetV2ApiService service

type ApiOpenApi2CampaignGetGetRequest struct {
	ctx          context.Context
	ApiService   *CampaignGetV2ApiService
	advertiserId *int64
	fields       *[]string
	filtering    *CampaignGetV2Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2CampaignGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CampaignGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) Fields(fields []string) *ApiOpenApi2CampaignGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) Filtering(filtering CampaignGetV2Filtering) *ApiOpenApi2CampaignGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) Page(page int64) *ApiOpenApi2CampaignGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2CampaignGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) Execute() (*CampaignGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CampaignGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2CampaignGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CampaignGetGetRequest) WithLog(enable bool) *ApiOpenApi2CampaignGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CampaignGetGet Method for OpenApi2CampaignGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CampaignGetGetRequest
*/
func (a *CampaignGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2CampaignGetGetRequest {
	return &ApiOpenApi2CampaignGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CampaignGetV2Response
func (a *CampaignGetV2ApiService) getExecute(r *ApiOpenApi2CampaignGetGetRequest) (*CampaignGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CampaignGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/campaign/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
