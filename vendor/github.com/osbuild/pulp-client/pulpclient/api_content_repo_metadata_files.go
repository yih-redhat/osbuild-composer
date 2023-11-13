/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentRepoMetadataFilesAPIService ContentRepoMetadataFilesAPI service
type ContentRepoMetadataFilesAPIService service

type ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest struct {
	ctx context.Context
	ApiService *ContentRepoMetadataFilesAPIService
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) Limit(limit int32) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) Offset(offset int32) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) Ordering(ordering []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) PulpHrefIn(pulpHrefIn []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) PulpIdIn(pulpIdIn []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) RepositoryVersion(repositoryVersion string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) Fields(fields []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) ExcludeFields(excludeFields []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) Execute() (*PaginatedrpmRepoMetadataFileResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmRepoMetadataFilesListExecute(r)
}

/*
ContentRpmRepoMetadataFilesList List repo metadata files

RepoMetadataFile Viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest
*/
func (a *ContentRepoMetadataFilesAPIService) ContentRpmRepoMetadataFilesList(ctx context.Context) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest {
	return ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedrpmRepoMetadataFileResponseList
func (a *ContentRepoMetadataFilesAPIService) ContentRpmRepoMetadataFilesListExecute(r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesListRequest) (*PaginatedrpmRepoMetadataFileResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmRepoMetadataFileResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRepoMetadataFilesAPIService.ContentRpmRepoMetadataFilesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/repo_metadata_files/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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

	localVarHTTPResponse, err := a.client.callAPI(req)
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

type ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest struct {
	ctx context.Context
	ApiService *ContentRepoMetadataFilesAPIService
	rpmRepoMetadataFileHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest) Fields(fields []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest) ExcludeFields(excludeFields []string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest) Execute() (*RpmRepoMetadataFileResponse, *http.Response, error) {
	return r.ApiService.ContentRpmRepoMetadataFilesReadExecute(r)
}

/*
ContentRpmRepoMetadataFilesRead Inspect a repo metadata file

RepoMetadataFile Viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmRepoMetadataFileHref
 @return ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest
*/
func (a *ContentRepoMetadataFilesAPIService) ContentRpmRepoMetadataFilesRead(ctx context.Context, rpmRepoMetadataFileHref string) ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest {
	return ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmRepoMetadataFileHref: rpmRepoMetadataFileHref,
	}
}

// Execute executes the request
//  @return RpmRepoMetadataFileResponse
func (a *ContentRepoMetadataFilesAPIService) ContentRpmRepoMetadataFilesReadExecute(r ContentRepoMetadataFilesAPIContentRpmRepoMetadataFilesReadRequest) (*RpmRepoMetadataFileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmRepoMetadataFileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentRepoMetadataFilesAPIService.ContentRpmRepoMetadataFilesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{rpm_repo_metadata_file_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_repo_metadata_file_href"+"}", parameterValueToString(r.rpmRepoMetadataFileHref, "rpmRepoMetadataFileHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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

	localVarHTTPResponse, err := a.client.callAPI(req)
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