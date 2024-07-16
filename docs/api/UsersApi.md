# \UsersApi

All URIs are relative to *https://kafka.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersUsersAccessGet**](UsersApi.md#ClustersUsersAccessGet) | **Get** /clusters/{clusterId}/users/{userId}/access | Retrieve Kafka User with credentials.|
|[**ClustersUsersGet**](UsersApi.md#ClustersUsersGet) | **Get** /clusters/{clusterId}/users | Retrieve all Users|



## ClustersUsersAccessGet

```go
var result UserReadAccess = ClustersUsersAccessGet(ctx, clusterId, userId)
                      .Execute()
```

Retrieve Kafka User with credentials.



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
    userId := "d11db12c-2625-5664-afd4-a3599731b5af" // string | The ID (UUID) of the User.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersAccessGet(context.Background(), clusterId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersAccessGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersAccessGet`: UserReadAccess
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersAccessGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |
|**userId** | **string** | The ID (UUID) of the User. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersAccessGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**UserReadAccess**](../models/UserReadAccess.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"UsersApiService.ClustersUsersAccessGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "UsersApiService.ClustersUsersAccessGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "UsersApiService.ClustersUsersAccessGet": {
    "port": "8443",
},
})
```


## ClustersUsersGet

```go
var result UserReadList = ClustersUsersGet(ctx, clusterId)
                      .Execute()
```

Retrieve all Users



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
    resource, resp, err := apiClient.UsersApi.ClustersUsersGet(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersGet`: UserReadList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**UserReadList**](../models/UserReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"UsersApiService.ClustersUsersGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "UsersApiService.ClustersUsersGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "UsersApiService.ClustersUsersGet": {
    "port": "8443",
},
})
```

