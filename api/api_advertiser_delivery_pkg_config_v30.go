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

// AdvertiserDeliveryPkgConfigV30ApiService AdvertiserDeliveryPkgConfigV30Api service
type AdvertiserDeliveryPkgConfigV30ApiService service

type ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest struct {
	ctx              context.Context
	ApiService       *AdvertiserDeliveryPkgConfigV30ApiService
	firstIndustryId  *int64
	secondIndustryId *int64
	thirdIndustryId  *int64
	advertiserId     *int64
}

// 一级行业id
func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) FirstIndustryId(firstIndustryId int64) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	r.firstIndustryId = &firstIndustryId
	return r
}

// 二级行业id
func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) SecondIndustryId(secondIndustryId int64) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	r.secondIndustryId = &secondIndustryId
	return r
}

// 三级行业id
func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) ThirdIndustryId(thirdIndustryId int64) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	r.thirdIndustryId = &thirdIndustryId
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) Execute() (*AdvertiserDeliveryPkgConfigV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserDeliveryPkgConfigGet Method for OpenApiV30AdvertiserDeliveryPkgConfigGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest
*/
func (a *AdvertiserDeliveryPkgConfigV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest {
	return &ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserDeliveryPkgConfigV30Response
func (a *AdvertiserDeliveryPkgConfigV30ApiService) getExecute(r *ApiOpenApiV30AdvertiserDeliveryPkgConfigGetRequest) (*AdvertiserDeliveryPkgConfigV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserDeliveryPkgConfigV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/delivery_pkg_config/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.firstIndustryId == nil {
		return localVarReturnValue, nil, ReportError("firstIndustryId is required and must be specified")
	}
	if r.secondIndustryId == nil {
		return localVarReturnValue, nil, ReportError("secondIndustryId is required and must be specified")
	}
	if r.thirdIndustryId == nil {
		return localVarReturnValue, nil, ReportError("thirdIndustryId is required and must be specified")
	}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "first_industry_id", r.firstIndustryId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "second_industry_id", r.secondIndustryId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "third_industry_id", r.thirdIndustryId)
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
