/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
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

// QianchuanAwemeEstimateProfitV10ApiService QianchuanAwemeEstimateProfitV10Api service
type QianchuanAwemeEstimateProfitV10ApiService service

type ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest struct {
	ctx             context.Context
	ApiService      *QianchuanAwemeEstimateProfitV10ApiService
	advertiserId    *int64
	marketingGoal   *QianchuanAwemeEstimateProfitV10MarketingGoal
	deliverySetting *QianchuanAwemeEstimateProfitV10DeliverySetting
	audience        *QianchuanAwemeEstimateProfitV10Audience
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) MarketingGoal(marketingGoal QianchuanAwemeEstimateProfitV10MarketingGoal) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

// 投放设置
func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) DeliverySetting(deliverySetting QianchuanAwemeEstimateProfitV10DeliverySetting) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	r.deliverySetting = &deliverySetting
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) Audience(audience QianchuanAwemeEstimateProfitV10Audience) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	r.audience = &audience
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) Execute() (*QianchuanAwemeEstimateProfitV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeEstimateProfitGet Method for OpenApiV10QianchuanAwemeEstimateProfitGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest
*/
func (a *QianchuanAwemeEstimateProfitV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest {
	return &ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeEstimateProfitV10Response
func (a *QianchuanAwemeEstimateProfitV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeEstimateProfitGetRequest) (*QianchuanAwemeEstimateProfitV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanAwemeEstimateProfitV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/estimate_profit/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.deliverySetting == nil {
		return localVarReturnValue, nil, ReportError("deliverySetting is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_setting", r.deliverySetting)
	if r.audience != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "audience", r.audience)
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
