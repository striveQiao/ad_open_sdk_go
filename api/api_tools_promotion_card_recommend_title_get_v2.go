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

// ToolsPromotionCardRecommendTitleGetV2ApiService ToolsPromotionCardRecommendTitleGetV2Api service
type ToolsPromotionCardRecommendTitleGetV2ApiService service

type ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsPromotionCardRecommendTitleGetV2ApiService
	adId         *int64
	advertiserId *int64
	contentType  *ToolsPromotionCardRecommendTitleGetV2ContentType
	externalUrl  *string
	industryId   *int64
	textType     *ToolsPromotionCardRecommendTitleGetV2TextType
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) ContentType(contentType ToolsPromotionCardRecommendTitleGetV2ContentType) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.contentType = &contentType
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) ExternalUrl(externalUrl string) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.externalUrl = &externalUrl
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) IndustryId(industryId int64) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.industryId = &industryId
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) TextType(textType ToolsPromotionCardRecommendTitleGetV2TextType) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.textType = &textType
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) Execute() (*ToolsPromotionCardRecommendTitleGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPromotionCardRecommendTitleGetGet Method for OpenApi2ToolsPromotionCardRecommendTitleGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest
*/
func (a *ToolsPromotionCardRecommendTitleGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest {
	return &ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPromotionCardRecommendTitleGetV2Response
func (a *ToolsPromotionCardRecommendTitleGetV2ApiService) getExecute(r *ApiOpenApi2ToolsPromotionCardRecommendTitleGetGetRequest) (*ToolsPromotionCardRecommendTitleGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPromotionCardRecommendTitleGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/promotion_card/recommend_title/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content_type", r.contentType)
	}
	if r.externalUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_url", r.externalUrl)
	}
	if r.industryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "industry_id", r.industryId)
	}
	if r.textType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "text_type", r.textType)
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
