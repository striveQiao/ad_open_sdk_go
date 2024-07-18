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

// CgTransferWalletTransferListV30ApiService CgTransferWalletTransferListV30Api service
type CgTransferWalletTransferListV30ApiService service

type ApiOpenApiV30CgTransferWalletTransferListGetRequest struct {
	ctx               context.Context
	ApiService        *CgTransferWalletTransferListV30ApiService
	accountId         *int64
	accountType       *CgTransferWalletTransferListV30AccountType
	bizRequestNo      *string
	queryBeginTime    *string
	queryEndTime      *string
	queryWalletIdList *[]int64
	pageInfo          *CgTransferWalletTransferListV30PageInfo
	payeeId           *int64
	remitterId        *int64
}

// 鉴权账户id
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) AccountId(accountId int64) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.accountId = &accountId
	return r
}

// 鉴权账户类型
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) AccountType(accountType CgTransferWalletTransferListV30AccountType) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.accountType = &accountType
	return r
}

// 请求id，推荐uuid，方便请求链路对齐
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) BizRequestNo(bizRequestNo string) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.bizRequestNo = &bizRequestNo
	return r
}

// 查询开始时间(转账创建时间) yyyy-MM-dd HH:mm:ss
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) QueryBeginTime(queryBeginTime string) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.queryBeginTime = &queryBeginTime
	return r
}

// 查询结束时间(转账创建时间) yyyy-MM-dd HH:mm:ss
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) QueryEndTime(queryEndTime string) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.queryEndTime = &queryEndTime
	return r
}

// 需要查询的大小钱包id，加款方与减款方都需要包含在内
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) QueryWalletIdList(queryWalletIdList []int64) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.queryWalletIdList = &queryWalletIdList
	return r
}

// 分页信息
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) PageInfo(pageInfo CgTransferWalletTransferListV30PageInfo) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.pageInfo = &pageInfo
	return r
}

// 加款方id，可选，需要包含在query_wallet_id_list
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) PayeeId(payeeId int64) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.payeeId = &payeeId
	return r
}

// 减款方id，可选，需要包含在query_wallet_id_list
func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) RemitterId(remitterId int64) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.remitterId = &remitterId
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) Execute() (*CgTransferWalletTransferListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferWalletTransferListGet Method for OpenApiV30CgTransferWalletTransferListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferWalletTransferListGetRequest
*/
func (a *CgTransferWalletTransferListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30CgTransferWalletTransferListGetRequest {
	return &ApiOpenApiV30CgTransferWalletTransferListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferWalletTransferListV30Response
func (a *CgTransferWalletTransferListV30ApiService) getExecute(r *ApiOpenApiV30CgTransferWalletTransferListGetRequest) (*CgTransferWalletTransferListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferWalletTransferListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/wallet/transfer/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.bizRequestNo == nil {
		return localVarReturnValue, nil, ReportError("bizRequestNo is required and must be specified")
	}
	if r.queryBeginTime == nil {
		return localVarReturnValue, nil, ReportError("queryBeginTime is required and must be specified")
	}
	if r.queryEndTime == nil {
		return localVarReturnValue, nil, ReportError("queryEndTime is required and must be specified")
	}
	if r.queryWalletIdList == nil {
		return localVarReturnValue, nil, ReportError("queryWalletIdList is required and must be specified")
	}
	if r.pageInfo == nil {
		return localVarReturnValue, nil, ReportError("pageInfo is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "biz_request_no", r.bizRequestNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_begin_time", r.queryBeginTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_end_time", r.queryEndTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "query_wallet_id_list", r.queryWalletIdList)
	if r.payeeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payee_id", r.payeeId)
	}
	if r.remitterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "remitter_id", r.remitterId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_info", r.pageInfo)
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
