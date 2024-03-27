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

// EnterpriseCommentReplyListV10ApiService EnterpriseCommentReplyListV10Api service
type EnterpriseCommentReplyListV10ApiService service

type ApiOpenApiV10EnterpriseCommentReplyListGetRequest struct {
	ctx         context.Context
	ApiService  *EnterpriseCommentReplyListV10ApiService
	ccAccountId *int64
	commentId   *int64
	eDouyinId   *string
	page        *int64
	pageSize    *int64
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) CcAccountId(ccAccountId int64) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.ccAccountId = &ccAccountId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) CommentId(commentId int64) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.commentId = &commentId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) EDouyinId(eDouyinId string) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.eDouyinId = &eDouyinId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) Page(page int64) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) PageSize(pageSize int64) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) Execute() (*EnterpriseCommentReplyListV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseCommentReplyListGet Method for OpenApiV10EnterpriseCommentReplyListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseCommentReplyListGetRequest
*/
func (a *EnterpriseCommentReplyListV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseCommentReplyListGetRequest {
	return &ApiOpenApiV10EnterpriseCommentReplyListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseCommentReplyListV10Response
func (a *EnterpriseCommentReplyListV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseCommentReplyListGetRequest) (*EnterpriseCommentReplyListV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseCommentReplyListV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/comment/reply/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ccAccountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cc_account_id", r.ccAccountId)
	}
	if r.commentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comment_id", r.commentId)
	}
	if r.eDouyinId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "e_douyin_id", r.eDouyinId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
