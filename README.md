# Go API client for jabali_sdk

API Proposal for Jabali.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 0.0.1
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import jabali_sdk "github.com/barrygear/jabali"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `jabali_sdk.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), jabali_sdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `jabali_sdk.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), jabali_sdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `jabali_sdk.ContextOperationServerIndices` and `jabali_sdk.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), jabali_sdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), jabali_sdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*JabaliAPI* | [**CheckHealth**](docs/JabaliAPI.md#checkhealth) | **Get** /healthz | 
*JabaliAPI* | [**CreateGame**](docs/JabaliAPI.md#creategame) | **Post** /game | Create a new Game
*JabaliAPI* | [**DeleteGame**](docs/JabaliAPI.md#deletegame) | **Delete** /game/{gameId} | Delete a game
*JabaliAPI* | [**GetGame**](docs/JabaliAPI.md#getgame) | **Get** /game/{gameId} | Get a Game
*JabaliAPI* | [**ListGames**](docs/JabaliAPI.md#listgames) | **Get** /games | List Games
*JabaliAPI* | [**UpdateGame**](docs/JabaliAPI.md#updategame) | **Put** /game/{gameId} | Update a game


## Documentation For Models

 - [CreateGameRequest](docs/CreateGameRequest.md)
 - [Game](docs/Game.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



