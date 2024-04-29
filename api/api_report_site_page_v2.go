/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// ReportSitePageV2ApiService ReportSitePageV2Api service
type ReportSitePageV2ApiService service

type ApiOpenApi2ReportSitePageGetRequest struct {
	ctx           context.Context
	ApiService    *ReportSitePageV2ApiService
	advertiserId  *int64
	inventoryType *ReportSitePageV2InventoryType
	siteId        *string
	timeDuration  *ReportSitePageV2TimeDuration
}

func (r *ApiOpenApi2ReportSitePageGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportSitePageGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportSitePageGetRequest) InventoryType(inventoryType ReportSitePageV2InventoryType) *ApiOpenApi2ReportSitePageGetRequest {
	r.inventoryType = &inventoryType
	return r
}

func (r *ApiOpenApi2ReportSitePageGetRequest) SiteId(siteId string) *ApiOpenApi2ReportSitePageGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApi2ReportSitePageGetRequest) TimeDuration(timeDuration ReportSitePageV2TimeDuration) *ApiOpenApi2ReportSitePageGetRequest {
	r.timeDuration = &timeDuration
	return r
}

func (r *ApiOpenApi2ReportSitePageGetRequest) Execute() (*ReportSitePageV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportSitePageGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportSitePageGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportSitePageGetRequest) WithLog(enable bool) *ApiOpenApi2ReportSitePageGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportSitePageGet Method for OpenApi2ReportSitePageGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportSitePageGetRequest
*/
func (a *ReportSitePageV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportSitePageGetRequest {
	return &ApiOpenApi2ReportSitePageGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportSitePageV2Response
func (a *ReportSitePageV2ApiService) getExecute(r *ApiOpenApi2ReportSitePageGetRequest) (*ReportSitePageV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportSitePageV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/site/page/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.inventoryType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inventory_type", r.inventoryType)
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
	}
	if r.timeDuration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_duration", r.timeDuration)
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
