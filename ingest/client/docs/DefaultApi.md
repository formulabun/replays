# \DefaultApi

All URIs are relative to *http://info.replay*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphQLGet**](DefaultApi.md#GraphQLGet) | **Get** /graphQL | Query replays
[**GraphQLPost**](DefaultApi.md#GraphQLPost) | **Post** /graphQL | Create new replay entry



## GraphQLGet

> Graphql GraphQLGet(ctx).GraphQLQuery(graphQLQuery).Execute()

Query replays



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    graphQLQuery := "graphQLQuery_example" // string | graphQL query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GraphQLGet(context.Background()).GraphQLQuery(graphQLQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GraphQLGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphQLGet`: Graphql
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GraphQLGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphQLGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLQuery** | **string** | graphQL query | 

### Return type

[**Graphql**](Graphql.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphQLPost

> GraphQLPost(ctx).Replay(replay).Execute()

Create new replay entry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    replay := *openapiclient.NewReplay() // Replay |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GraphQLPost(context.Background()).Replay(replay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GraphQLPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGraphQLPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replay** | [**Replay**](Replay.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

