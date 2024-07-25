# \TopicsApi

All URIs are relative to *https://kafka.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersTopicsDelete**](TopicsApi.md#ClustersTopicsDelete) | **Delete** /clusters/{clusterId}/topics/{topicId} | Delete Topic|
|[**ClustersTopicsFindById**](TopicsApi.md#ClustersTopicsFindById) | **Get** /clusters/{clusterId}/topics/{topicId} | Retrieve Topic|
|[**ClustersTopicsGet**](TopicsApi.md#ClustersTopicsGet) | **Get** /clusters/{clusterId}/topics | Retrieve all Topics|
|[**ClustersTopicsPost**](TopicsApi.md#ClustersTopicsPost) | **Post** /clusters/{clusterId}/topics | Create Topic|



## ClustersTopicsDelete

```go
var result  = ClustersTopicsDelete(ctx, clusterId, topicId)
                      .Execute()
```

Delete Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-kafka"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    topicId := "ae085c4c-3626-5f1d-b4bc-cc53ae8267ce" // string | The ID (UUID) of the Topic.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.TopicsApi.ClustersTopicsDelete(context.Background(), clusterId, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ClustersTopicsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |
|**topicId** | **string** | The ID (UUID) of the Topic. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersTopicsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TopicsApiService.ClustersTopicsDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TopicsApiService.ClustersTopicsDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TopicsApiService.ClustersTopicsDelete": {
    "port": "8443",
},
})
```


## ClustersTopicsFindById

```go
var result TopicRead = ClustersTopicsFindById(ctx, clusterId, topicId)
                      .Execute()
```

Retrieve Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-kafka"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    topicId := "ae085c4c-3626-5f1d-b4bc-cc53ae8267ce" // string | The ID (UUID) of the Topic.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TopicsApi.ClustersTopicsFindById(context.Background(), clusterId, topicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ClustersTopicsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersTopicsFindById`: TopicRead
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.ClustersTopicsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |
|**topicId** | **string** | The ID (UUID) of the Topic. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersTopicsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**TopicRead**](../models/TopicRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TopicsApiService.ClustersTopicsFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TopicsApiService.ClustersTopicsFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TopicsApiService.ClustersTopicsFindById": {
    "port": "8443",
},
})
```


## ClustersTopicsGet

```go
var result TopicReadList = ClustersTopicsGet(ctx, clusterId)
                      .Execute()
```

Retrieve all Topics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-kafka"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TopicsApi.ClustersTopicsGet(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ClustersTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersTopicsGet`: TopicReadList
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.ClustersTopicsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersTopicsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**TopicReadList**](../models/TopicReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TopicsApiService.ClustersTopicsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TopicsApiService.ClustersTopicsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TopicsApiService.ClustersTopicsGet": {
    "port": "8443",
},
})
```


## ClustersTopicsPost

```go
var result TopicRead = ClustersTopicsPost(ctx, clusterId)
                      .TopicCreate(topicCreate)
                      .Execute()
```

Create Topic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-kafka"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    topicCreate := *openapiclient.NewTopicCreate(*openapiclient.NewTopic("my-kafka-cluster-topic")) // TopicCreate | Topic to create.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TopicsApi.ClustersTopicsPost(context.Background(), clusterId).TopicCreate(topicCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsApi.ClustersTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersTopicsPost`: TopicRead
    fmt.Fprintf(os.Stdout, "Response from `TopicsApi.ClustersTopicsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersTopicsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **topicCreate** | [**TopicCreate**](../models/TopicCreate.md) | Topic to create. | |

### Return type

[**TopicRead**](../models/TopicRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"TopicsApiService.ClustersTopicsPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "TopicsApiService.ClustersTopicsPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "TopicsApiService.ClustersTopicsPost": {
    "port": "8443",
},
})
```

