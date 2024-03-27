/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

// StarOrderListByCampaignV2ApiService StarOrderListByCampaignV2Api service
type StarOrderListByCampaignV2ApiService service

type ApiOpenApi2StarOrderListByCampaignGetRequest struct {
	ctx         context.Context
	ApiService  *StarOrderListByCampaignV2ApiService
	starId      *int64
	campaignIds *[]int64
	page        *int32
	limit       *int32
}

// 星图客户ID
func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) StarId(starId int64) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	r.starId = &starId
	return r
}

// 需求ID
func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) CampaignIds(campaignIds []int64) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	r.campaignIds = &campaignIds
	return r
}

// 分页页码
func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) Page(page int32) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	r.page = &page
	return r
}

// 页大小
func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) Limit(limit int32) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) Execute() (*StarOrderListByCampaignV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarOrderListByCampaignGetRequest) WithLog(enable bool) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarOrderListByCampaignGet Method for OpenApi2StarOrderListByCampaignGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarOrderListByCampaignGetRequest
*/
func (a *StarOrderListByCampaignV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarOrderListByCampaignGetRequest {
	return &ApiOpenApi2StarOrderListByCampaignGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarOrderListByCampaignV2Response
func (a *StarOrderListByCampaignV2ApiService) getExecute(r *ApiOpenApi2StarOrderListByCampaignGetRequest) (*StarOrderListByCampaignV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarOrderListByCampaignV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/order/list_by_campaign/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.campaignIds == nil {
		return localVarReturnValue, nil, ReportError("campaignIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_ids", r.campaignIds)
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
