package main

import (
	"flag"
	"fmt"
	sendbirdclient "sendbirdclient"
)

const (
	IssueAccessToken  = false
	IsDistinct        = true
	UserRoleMetaKey   = "role"
	UserRoleMetaValue = "admin"
	InitCustomType    = "idle"

	GroupChannelType       = "group_channels"
	ChannelStatusMetaKey   = "status"
	ChannelStatusMetaValue = "idle"
)

var (
	apiKey   = flag.String("key", "", "API Key for using Sendbird Platform API")
	userID   = flag.String("id", "", "UserID for admin user registeration")
	nickName = flag.String("name", "", "Nick name for admin user registeration")
	//profileURL = flag.String("profile", "", "The URL of the user’s profile image.")
)

func main() {
	flag.Parse()

	testClient, err := sendbirdclient.NewClient(sendbirdclient.WithAPIKey(*apiKey))
	check(err)

	user, err := testClient.CreateAUserWithURL(&sendbirdclient.CreateAUserWithURLRequest{
		UserID:   *userID,
		NickName: *nickName,
		//ProfileURL:       *profileURL,
		IssueAccessToken: IssueAccessToken,
	})
	check(err)
	fmt.Printf("User: %+v \n", user)

	meta := make(map[string]string)
	meta[UserRoleMetaKey] = UserRoleMetaValue

	returnMeta, err := testClient.CreateAnUserMetaData(user.UserID, &sendbirdclient.CreateAnUserMetaDataRequest{
		MetaData: meta,
	})
	check(err)
	fmt.Printf("Usermeta: %+v \n", returnMeta)

	chName := fmt.Sprintf("Channel_%s", user.UserID)

	groupCh, err := testClient.CreateAGroupChannelWithURL(&sendbirdclient.CreateAGroupChannelWithURLRequest{
		Name:       chName,
		UserIDs:    []string{user.UserID},
		CustomType: InitCustomType,
		IsDistinct: IsDistinct,
	})
	check(err)
	fmt.Printf("GroupChannel: %+v \n", groupCh)

	chMeta := make(map[string]string)
	chMeta[ChannelStatusMetaKey] = ChannelStatusMetaValue

	returnChMeta, err := testClient.CreateAChannelMetadata(GroupChannelType, groupCh.ChannelURL, &sendbirdclient.CreateAChannelMetadataRequest{
		Metadata: chMeta,
	})
	check(err)
	fmt.Printf("ChannelMetadata: %+v \n", returnChMeta)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
