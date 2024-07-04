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

// BrandQueryStockV30ApiService BrandQueryStockV30Api service
type BrandQueryStockV30ApiService service

type ApiOpenApiV30BrandQueryStockGetRequest struct {
	ctx                 context.Context
	ApiService          *BrandQueryStockV30ApiService
	advertiserId        *int64
	classify            *BrandQueryStockV30Classify
	landingType         *BrandQueryStockV30LandingType
	appOrigin           *BrandQueryStockV30AppOrigin
	adForm              *BrandQueryStockV30AdForm
	gdSendType          *BrandQueryStockV30GdSendType
	policyNo            *string
	scheduleInfo        *BrandQueryStockV30ScheduleInfo
	audienceInfo        *BrandQueryStockV30AudienceInfo
	merchantIntentionNo *string
}

// 广告主ID
func (r *ApiOpenApiV30BrandQueryStockGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告类别
func (r *ApiOpenApiV30BrandQueryStockGetRequest) Classify(classify BrandQueryStockV30Classify) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.classify = &classify
	return r
}

// 推广目的
func (r *ApiOpenApiV30BrandQueryStockGetRequest) LandingType(landingType BrandQueryStockV30LandingType) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.landingType = &landingType
	return r
}

// 投放位置（媒体端）
func (r *ApiOpenApiV30BrandQueryStockGetRequest) AppOrigin(appOrigin BrandQueryStockV30AppOrigin) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.appOrigin = &appOrigin
	return r
}

// 版位
func (r *ApiOpenApiV30BrandQueryStockGetRequest) AdForm(adForm BrandQueryStockV30AdForm) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.adForm = &adForm
	return r
}

// 优化目标
func (r *ApiOpenApiV30BrandQueryStockGetRequest) GdSendType(gdSendType BrandQueryStockV30GdSendType) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.gdSendType = &gdSendType
	return r
}

// 政策编号
func (r *ApiOpenApiV30BrandQueryStockGetRequest) PolicyNo(policyNo string) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.policyNo = &policyNo
	return r
}

// 日期
func (r *ApiOpenApiV30BrandQueryStockGetRequest) ScheduleInfo(scheduleInfo BrandQueryStockV30ScheduleInfo) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.scheduleInfo = &scheduleInfo
	return r
}

// 定向
func (r *ApiOpenApiV30BrandQueryStockGetRequest) AudienceInfo(audienceInfo BrandQueryStockV30AudienceInfo) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.audienceInfo = &audienceInfo
	return r
}

// 招商意向单ID
func (r *ApiOpenApiV30BrandQueryStockGetRequest) MerchantIntentionNo(merchantIntentionNo string) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.merchantIntentionNo = &merchantIntentionNo
	return r
}

func (r *ApiOpenApiV30BrandQueryStockGetRequest) Execute() (*BrandQueryStockV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BrandQueryStockGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandQueryStockGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandQueryStockGetRequest) WithLog(enable bool) *ApiOpenApiV30BrandQueryStockGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandQueryStockGet Method for OpenApiV30BrandQueryStockGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandQueryStockGetRequest
*/
func (a *BrandQueryStockV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BrandQueryStockGetRequest {
	return &ApiOpenApiV30BrandQueryStockGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandQueryStockV30Response
func (a *BrandQueryStockV30ApiService) getExecute(r *ApiOpenApiV30BrandQueryStockGetRequest) (*BrandQueryStockV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandQueryStockV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/query_stock/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.classify == nil {
		return localVarReturnValue, nil, ReportError("classify is required and must be specified")
	}
	if r.landingType == nil {
		return localVarReturnValue, nil, ReportError("landingType is required and must be specified")
	}
	if r.appOrigin == nil {
		return localVarReturnValue, nil, ReportError("appOrigin is required and must be specified")
	}
	if r.adForm == nil {
		return localVarReturnValue, nil, ReportError("adForm is required and must be specified")
	}
	if r.gdSendType == nil {
		return localVarReturnValue, nil, ReportError("gdSendType is required and must be specified")
	}
	if r.policyNo == nil {
		return localVarReturnValue, nil, ReportError("policyNo is required and must be specified")
	}
	if r.scheduleInfo == nil {
		return localVarReturnValue, nil, ReportError("scheduleInfo is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "classify", r.classify)
	parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "app_origin", r.appOrigin)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_form", r.adForm)
	parameterAddToHeaderOrQuery(localVarQueryParams, "gd_send_type", r.gdSendType)
	if r.audienceInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience_info", r.audienceInfo)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "policy_no", r.policyNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "schedule_info", r.scheduleInfo)
	if r.merchantIntentionNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merchant_intention_no", r.merchantIntentionNo)
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
