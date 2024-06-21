# \JabaliAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckHealth**](JabaliAPI.md#CheckHealth) | **Get** /healthz | 
[**CreateGame**](JabaliAPI.md#CreateGame) | **Post** /game | Create a new Game
[**DeleteGame**](JabaliAPI.md#DeleteGame) | **Delete** /game/{gameId} | Delete a game
[**GetGame**](JabaliAPI.md#GetGame) | **Get** /game/{gameId} | Get a Game
[**ListGames**](JabaliAPI.md#ListGames) | **Get** /games | List Games
[**UpdateGame**](JabaliAPI.md#UpdateGame) | **Put** /game/{gameId} | Update a game



## CheckHealth

> string CheckHealth(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JabaliAPI.CheckHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.CheckHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckHealth`: string
	fmt.Fprintf(os.Stdout, "Response from `JabaliAPI.CheckHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckHealthRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGame

> Game CreateGame(ctx).Game(game).Execute()

Create a new Game



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {
	game := *openapiclient.NewCreateGameRequest("Name_example", "Displayname_example", "Url_example") // CreateGameRequest | The Game to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JabaliAPI.CreateGame(context.Background()).Game(game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.CreateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGame`: Game
	fmt.Fprintf(os.Stdout, "Response from `JabaliAPI.CreateGame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | [**CreateGameRequest**](CreateGameRequest.md) | The Game to create. | 

### Return type

[**Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGame

> DeleteGame(ctx, gameId).Execute()

Delete a game



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {
	gameId := "gameId_example" // string | ID of the game to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JabaliAPI.DeleteGame(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.DeleteGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | ID of the game to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGame

> Game GetGame(ctx, gameId).Execute()

Get a Game



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {
	gameId := "gameId_example" // string | ID of the game to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JabaliAPI.GetGame(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.GetGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGame`: Game
	fmt.Fprintf(os.Stdout, "Response from `JabaliAPI.GetGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | ID of the game to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGames

> []Game ListGames(ctx).Execute()

List Games



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JabaliAPI.ListGames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.ListGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGames`: []Game
	fmt.Fprintf(os.Stdout, "Response from `JabaliAPI.ListGames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGamesRequest struct via the builder pattern


### Return type

[**[]Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGame

> Game UpdateGame(ctx, gameId).Game(game).Execute()

Update a game



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/barrygear/jabali"
)

func main() {
	gameId := "gameId_example" // string | ID of the game to get
	game := *openapiclient.NewCreateGameRequest("Name_example", "Displayname_example", "Url_example") // CreateGameRequest | The Game to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JabaliAPI.UpdateGame(context.Background(), gameId).Game(game).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JabaliAPI.UpdateGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGame`: Game
	fmt.Fprintf(os.Stdout, "Response from `JabaliAPI.UpdateGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | ID of the game to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | [**CreateGameRequest**](CreateGameRequest.md) | The Game to update. | 

### Return type

[**Game**](Game.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

