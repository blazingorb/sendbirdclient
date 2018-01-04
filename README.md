Sendbird Client Package for Golang
==================================
[![Drone Build Status](https://droneci.blazingorb.com/api/badges/blazingorb/sendbirdclient/status.svg)](https://droneci.blazingorb.com/blazingorb/sendbirdclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/blazingorb/sendbirdclient)](https://goreportcard.com/report/github.com/blazingorb/sendbirdclient)

## Features

This library provides wrapper functions for Sendbird Platform API written in Golang and support following Sendbird APIs:

- [Users API]
- [User Metadata API]
- [Open Channels API]
- [Group Channels API]



## Requirements
- Go 1.5 or later.
- Sendbird API Key.

## Installation

To install SendbirdClient for Golang, please execute the following `go get` command.

```bash
    go get github.com/blazingorb/sendbirdclient
``` 

## Usage

Sample Usage of creating a new Sendbird User:

```go
package main

import (
	"flag"
	"fmt"
	sendbirdclient "sendbirdclient"
)

const (
	IssueAccessToken  = false
	IsDistinct        = true
)

var (
	apiKey   = flag.String("key", "", "API Key for using Sendbird Platform API")
	userID   = flag.String("id", "", "UserID for creating a new user")
	nickName = flag.String("name", "", "Nickname for creating a new user")
)

func main() {
	flag.Parse()

	testClient, err := sendbirdclient.NewClient(sendbirdclient.WithAPIKey(*apiKey))
	check(err)

	user, err := testClient.CreateAUserWithURL(&sendbirdclient.CreateAUserWithURLRequest{
		UserID:   *userID,
		NickName: *nickName,
		IssueAccessToken: IssueAccessToken,
	})
	check(err)
	fmt.Printf("User: %+v \n", user)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
```

For more detail sample cases, please refer to files under examples folder.

[Users API]: https://docs.sendbird.com/platform#user
[User Metadata API]: https://docs.sendbird.com/platform#user_metadata
[Open Channels API]: https://docs.sendbird.com/platform#open_channel
[Group Channels API]: https://docs.sendbird.com/platform#group_channel
