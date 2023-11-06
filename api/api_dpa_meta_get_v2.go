/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
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

// DpaMetaGetV2ApiService DpaMetaGetV2Api service
type DpaMetaGetV2ApiService service

type ApiOpenApi2DpaMetaGetGetRequest struct {
	ctx          context.Context
	ApiService   *DpaMetaGetV2ApiService
	advertiserId *int64
	platformId   *int64
	indexable    *int64
	extractable  *int64
	industry     *int64
	status       *int64
	mediaType    *int64
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) PlatformId(platformId int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.platformId = &platformId
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) Indexable(indexable int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.indexable = &indexable
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) Extractable(extractable int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.extractable = &extractable
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) Industry(industry int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.industry = &industry
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) Status(status int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) MediaType(mediaType int64) *ApiOpenApi2DpaMetaGetGetRequest {
	r.mediaType = &mediaType
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) Execute() (*DpaMetaGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaMetaGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaMetaGetGetRequest) WithLog(enable bool) *ApiOpenApi2DpaMetaGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaMetaGetGet Method for OpenApi2DpaMetaGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaMetaGetGetRequest
*/
func (a *DpaMetaGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaMetaGetGetRequest {
	return &ApiOpenApi2DpaMetaGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaMetaGetV2Response
func (a *DpaMetaGetV2ApiService) getExecute(r *ApiOpenApi2DpaMetaGetGetRequest) (*DpaMetaGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DpaMetaGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/meta/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.platformId == nil {
		return localVarReturnValue, nil, ReportError("platformId is required and must be specified")
	}
	if *r.platformId < 1 {
		return localVarReturnValue, nil, ReportError("platformId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "platform_id", r.platformId)
	if r.indexable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "indexable", r.indexable)
	}
	if r.extractable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "extractable", r.extractable)
	}
	if r.industry != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "industry", r.industry)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.mediaType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mediaType", r.mediaType)
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
