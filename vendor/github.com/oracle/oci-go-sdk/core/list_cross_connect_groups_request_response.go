// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCrossConnectGroupsRequest wrapper for the ListCrossConnectGroups operation
type ListCrossConnectGroupsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListCrossConnectGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListCrossConnectGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState CrossConnectGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
}

func (request ListCrossConnectGroupsRequest) String() string {
	return common.PointerString(request)
}

// ListCrossConnectGroupsResponse wrapper for the ListCrossConnectGroups operation
type ListCrossConnectGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []CrossConnectGroup instance
	Items []CrossConnectGroup `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCrossConnectGroupsResponse) String() string {
	return common.PointerString(response)
}

// ListCrossConnectGroupsSortByEnum Enum with underlying type: string
type ListCrossConnectGroupsSortByEnum string

// Set of constants representing the allowable values for ListCrossConnectGroupsSortBy
const (
	ListCrossConnectGroupsSortByTimecreated ListCrossConnectGroupsSortByEnum = "TIMECREATED"
	ListCrossConnectGroupsSortByDisplayname ListCrossConnectGroupsSortByEnum = "DISPLAYNAME"
	ListCrossConnectGroupsSortByUnknown     ListCrossConnectGroupsSortByEnum = "UNKNOWN"
)

var mappingListCrossConnectGroupsSortBy = map[string]ListCrossConnectGroupsSortByEnum{
	"TIMECREATED": ListCrossConnectGroupsSortByTimecreated,
	"DISPLAYNAME": ListCrossConnectGroupsSortByDisplayname,
	"UNKNOWN":     ListCrossConnectGroupsSortByUnknown,
}

// GetListCrossConnectGroupsSortByEnumValues Enumerates the set of values for ListCrossConnectGroupsSortBy
func GetListCrossConnectGroupsSortByEnumValues() []ListCrossConnectGroupsSortByEnum {
	values := make([]ListCrossConnectGroupsSortByEnum, 0)
	for _, v := range mappingListCrossConnectGroupsSortBy {
		if v != ListCrossConnectGroupsSortByUnknown {
			values = append(values, v)
		}
	}
	return values
}

// ListCrossConnectGroupsSortOrderEnum Enum with underlying type: string
type ListCrossConnectGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListCrossConnectGroupsSortOrder
const (
	ListCrossConnectGroupsSortOrderAsc     ListCrossConnectGroupsSortOrderEnum = "ASC"
	ListCrossConnectGroupsSortOrderDesc    ListCrossConnectGroupsSortOrderEnum = "DESC"
	ListCrossConnectGroupsSortOrderUnknown ListCrossConnectGroupsSortOrderEnum = "UNKNOWN"
)

var mappingListCrossConnectGroupsSortOrder = map[string]ListCrossConnectGroupsSortOrderEnum{
	"ASC":     ListCrossConnectGroupsSortOrderAsc,
	"DESC":    ListCrossConnectGroupsSortOrderDesc,
	"UNKNOWN": ListCrossConnectGroupsSortOrderUnknown,
}

// GetListCrossConnectGroupsSortOrderEnumValues Enumerates the set of values for ListCrossConnectGroupsSortOrder
func GetListCrossConnectGroupsSortOrderEnumValues() []ListCrossConnectGroupsSortOrderEnum {
	values := make([]ListCrossConnectGroupsSortOrderEnum, 0)
	for _, v := range mappingListCrossConnectGroupsSortOrder {
		if v != ListCrossConnectGroupsSortOrderUnknown {
			values = append(values, v)
		}
	}
	return values
}