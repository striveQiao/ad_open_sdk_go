/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// FileMaterialAuditResultGetV2ApiService FileMaterialAuditResultGetV2Api service
type FileMaterialAuditResultGetV2ApiService service

type ApiOpenApi2FileMaterialAuditResultGetGetRequest struct {
	ctx           context.Context
	ApiService    *FileMaterialAuditResultGetV2ApiService
	advertiserId  *int64
	materialInfos *[]*FileMaterialAuditResultGetV2MaterialInfosInner
}

func (r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileMaterialAuditResultGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) MaterialInfos(materialInfos []*FileMaterialAuditResultGetV2MaterialInfosInner) *ApiOpenApi2FileMaterialAuditResultGetGetRequest {
	r.materialInfos = &materialInfos
	return r
}

func (r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) Execute() (*FileMaterialAuditResultGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileMaterialAuditResultGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileMaterialAuditResultGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileMaterialAuditResultGetGet Method for OpenApi2FileMaterialAuditResultGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileMaterialAuditResultGetGetRequest
*/
func (a *FileMaterialAuditResultGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileMaterialAuditResultGetGetRequest {
	return &ApiOpenApi2FileMaterialAuditResultGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileMaterialAuditResultGetV2Response
func (a *FileMaterialAuditResultGetV2ApiService) getExecute(r *ApiOpenApi2FileMaterialAuditResultGetGetRequest) (*FileMaterialAuditResultGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileMaterialAuditResultGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/material/audit_result/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.materialInfos != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "material_infos", r.materialInfos)
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
