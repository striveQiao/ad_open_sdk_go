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

// QianchuanDmpAudiencesGetV10ApiService QianchuanDmpAudiencesGetV10Api service
type QianchuanDmpAudiencesGetV10ApiService service

type ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest struct {
	ctx                 context.Context
	ApiService          *QianchuanDmpAudiencesGetV10ApiService
	advertiserId        *int64
	retargetingTagsType *int32
	offset              *int64
	limit               *int64
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) RetargetingTagsType(retargetingTagsType int32) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	r.retargetingTagsType = &retargetingTagsType
	return r
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) Offset(offset int64) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	r.offset = &offset
	return r
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) Limit(limit int64) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) Execute() (*QianchuanDmpAudiencesGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanDmpAudiencesGetGet Method for OpenApiV10QianchuanDmpAudiencesGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest
*/
func (a *QianchuanDmpAudiencesGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest {
	return &ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanDmpAudiencesGetV10Response
func (a *QianchuanDmpAudiencesGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanDmpAudiencesGetGetRequest) (*QianchuanDmpAudiencesGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanDmpAudiencesGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/dmp/audiences/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.retargetingTagsType == nil {
		return localVarReturnValue, nil, ReportError("retargetingTagsType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "retargeting_tags_type", r.retargetingTagsType)
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset)
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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
