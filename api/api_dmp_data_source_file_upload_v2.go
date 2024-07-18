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

// DmpDataSourceFileUploadV2ApiService DmpDataSourceFileUploadV2Api service
type DmpDataSourceFileUploadV2ApiService service

type ApiOpenApi2DmpDataSourceFileUploadPostRequest struct {
	ctx           context.Context
	ApiService    *DmpDataSourceFileUploadV2ApiService
	advertiserId  *int64
	file          *FormFileInfo
	fileSignature *string
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) File(file *FormFileInfo) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	r.file = file
	return r
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) FileSignature(fileSignature string) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	r.fileSignature = &fileSignature
	return r
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) Execute() (*DmpDataSourceFileUploadV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) AccessToken(accessToken string) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) WithLog(enable bool) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DmpDataSourceFileUploadPost Method for OpenApi2DmpDataSourceFileUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DmpDataSourceFileUploadPostRequest
*/
func (a *DmpDataSourceFileUploadV2ApiService) Post(ctx context.Context) *ApiOpenApi2DmpDataSourceFileUploadPostRequest {
	return &ApiOpenApi2DmpDataSourceFileUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DmpDataSourceFileUploadV2Response
func (a *DmpDataSourceFileUploadV2ApiService) postExecute(r *ApiOpenApi2DmpDataSourceFileUploadPostRequest) (*DmpDataSourceFileUploadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DmpDataSourceFileUploadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dmp/data_source/file/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, ReportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.file != nil {
		formFiles["file"] = r.file
	}
	if r.fileSignature != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "file_signature", r.fileSignature)
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
