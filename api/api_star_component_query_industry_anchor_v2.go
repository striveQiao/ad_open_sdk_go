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

// StarComponentQueryIndustryAnchorV2ApiService StarComponentQueryIndustryAnchorV2Api service
type StarComponentQueryIndustryAnchorV2ApiService service

type ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest struct {
	ctx          context.Context
	ApiService   *StarComponentQueryIndustryAnchorV2ApiService
	starId       *int64
	anchorStatus *int32
	anchorType   *int32
	page         *int32
	limit        *int32
}

// 星图客户ID
func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) StarId(starId int64) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.starId = &starId
	return r
}

// 组件状态 0 &#x3D; 待送审 1 &#x3D; 有效 2 &#x3D;审核中 3 &#x3D; 审核失败 不传则为全部
func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) AnchorStatus(anchorStatus int32) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.anchorStatus = &anchorStatus
	return r
}

// 行业锚点类型       电商 &#x3D; 1    汽车 &#x3D; 2    网服 &#x3D; 3    保险 &#x3D; 4    教育 &#x3D; 5    家装 &#x3D; 6    剪映锚点 &#x3D; 7    招商加盟组件 &#x3D; 8    旅游 &#x3D; 9    游戏行业组件 &#x3D; 10    通信行业组件 &#x3D; 11    房产行业组件 &#x3D; 12    西瓜行业组件 &#x3D; 13    轻颜行业组件 &#x3D; 14    醒图行业组件  &#x3D; 15 不传则为全部
func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) AnchorType(anchorType int32) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.anchorType = &anchorType
	return r
}

// 分页页码 正整数，默认1
func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) Page(page int32) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.page = &page
	return r
}

// 页大小 正整数且≤50，默认10
func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) Limit(limit int32) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) Execute() (*StarComponentQueryIndustryAnchorV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) WithLog(enable bool) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarComponentQueryIndustryAnchorGet Method for OpenApi2StarComponentQueryIndustryAnchorGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest
*/
func (a *StarComponentQueryIndustryAnchorV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest {
	return &ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarComponentQueryIndustryAnchorV2Response
func (a *StarComponentQueryIndustryAnchorV2ApiService) getExecute(r *ApiOpenApi2StarComponentQueryIndustryAnchorGetRequest) (*StarComponentQueryIndustryAnchorV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarComponentQueryIndustryAnchorV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/component/query_industry_anchor/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.anchorStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "anchor_status", r.anchorStatus)
	}
	if r.anchorType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "anchor_type", r.anchorType)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
