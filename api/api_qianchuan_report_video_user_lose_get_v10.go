/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

// QianchuanReportVideoUserLoseGetV10ApiService QianchuanReportVideoUserLoseGetV10Api service
type QianchuanReportVideoUserLoseGetV10ApiService service

type ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanReportVideoUserLoseGetV10ApiService
	fields       *[]string
	startDate    *string
	endDate      *string
	filtering    *QianchuanReportVideoUserLoseGetV10Filtering
	advertiserId *int64
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) Filtering(filtering QianchuanReportVideoUserLoseGetV10Filtering) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) Execute() (*QianchuanReportVideoUserLoseGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportVideoUserLoseGetGet Method for OpenApiV10QianchuanReportVideoUserLoseGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest
*/
func (a *QianchuanReportVideoUserLoseGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest {
	return &ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportVideoUserLoseGetV10Response
func (a *QianchuanReportVideoUserLoseGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportVideoUserLoseGetGetRequest) (*QianchuanReportVideoUserLoseGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanReportVideoUserLoseGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/video_user_lose/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
