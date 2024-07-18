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

// StarReportCustomDataTopicReportV2ApiService StarReportCustomDataTopicReportV2Api service
type StarReportCustomDataTopicReportV2ApiService service

type ApiOpenApi2StarReportCustomDataTopicReportGetRequest struct {
	ctx        context.Context
	ApiService *StarReportCustomDataTopicReportV2ApiService
	starId     *int64
	workId     *int64
	demandId   *int64
	topics     *[]*StarReportCustomDataTopicReportV2Topics
}

func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) StarId(starId int64) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	r.starId = &starId
	return r
}

// 交付作品Id：如视频Id、直播间Id
func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) WorkId(workId int64) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	r.workId = &workId
	return r
}

// 任务id
func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) DemandId(demandId int64) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	r.demandId = &demandId
	return r
}

// 数据主题: BASIC_DATA：基础信息、 FLOW_DATA：流量表现、 CONVERT_DATA：转化表现、 SEARCH_DATA：搜索表现、 RECOMMEND_DATA： 种草表现、 DY_SHOP_DATA：抖音进店、 USER_DISTRIBUTION_DATA：用户画像、 INDEX_SCORE_DATA： 指数得分、 COMMENT_DATA：评论数据 直播用户画像仅保留近90天且直播时长 &gt;&#x3D; 25 分钟直播数据
func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) Topics(topics []*StarReportCustomDataTopicReportV2Topics) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	r.topics = &topics
	return r
}

func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) Execute() (*StarReportCustomDataTopicReportV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportCustomDataTopicReportGet Method for OpenApi2StarReportCustomDataTopicReportGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportCustomDataTopicReportGetRequest
*/
func (a *StarReportCustomDataTopicReportV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportCustomDataTopicReportGetRequest {
	return &ApiOpenApi2StarReportCustomDataTopicReportGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportCustomDataTopicReportV2Response
func (a *StarReportCustomDataTopicReportV2ApiService) getExecute(r *ApiOpenApi2StarReportCustomDataTopicReportGetRequest) (*StarReportCustomDataTopicReportV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarReportCustomDataTopicReportV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/custom_data_topic_report/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.workId == nil {
		return localVarReturnValue, nil, ReportError("workId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}
	if r.topics == nil {
		return localVarReturnValue, nil, ReportError("topics is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "work_id", r.workId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "topics", r.topics)
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
