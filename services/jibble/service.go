package jibble

import (
	"context"
	"fmt"

	globalConfig "github.com/markkizz/time-tracker-automation/config"
)

func Client(options struct{ token string }) *JibbleClient {
	jibbleClient := JibbleClient{}
	jibbleConfig := globalConfig.Config().Jibble
	client := jibbleClient.NewJibbleClient(
		JibbleClientOptions{
			identityUrl:     jibbleConfig.ApiIdentityUrl,
			timetrackingUrl: jibbleConfig.ApiTimeTrackingUrl,
			token:           options.token,
		},
	)
	return client
}

type TimeTracker interface {
	Login(username string, password string)
	ClockIn()
	ClockOut()
	clock(clockType string)
}

type JibbleService struct {
	userId, personId, orgId, accessToken, refreshToken string
}

func (service *JibbleService) Login(username string, password string) {
	jibbleClient := Client(struct{ token string }{})

	ctx := context.Background()
	request := UserAccessTokenRequest{
		username,
		password,
	}
	userAccessToken, err := jibbleClient.GetUserAccessToken(ctx, request)
	if err != nil {
		fmt.Printf("[Error Login]: %v\n", err)
		return
	}

	service.accessToken = userAccessToken.AccessToken
	service.refreshToken = userAccessToken.RefreshToken

	jibbleClient = Client(struct{ token string }{service.accessToken})
	organizationId, err := jibbleClient.GetOrganizationId(ctx)
	if err != nil {
		fmt.Printf("[Error Login]: %v\n", err)
		return
	}

	service.orgId = organizationId.Value[0].ID

	personId, err := jibbleClient.GetPersonId(ctx, PersonIdRequest{service.orgId})
	if err != nil {
		fmt.Printf("[Error Login]: %v\n", err)
		return
	}

	service.personId = personId.Value[0].ID
	service.userId = personId.Value[0].UserID

	personAccessToken, err := jibbleClient.GetPersonAccessToken(
		ctx,
		PersonAccessTokenRequest{
			username:     username,
			password:     password,
			personId:     service.personId,
			refreshToken: service.refreshToken,
		},
	)
	if err != nil {
		fmt.Printf("[Error Login]: %v\n", err)
		return
	}

	service.accessToken = personAccessToken.AccessToken
	service.refreshToken = personAccessToken.RefreshToken
}

func (service *JibbleService) Clock(clockType string) {

	requestPayload := ClockingRequest{
		PersonID: service.personId,
		Type:     clockType, // In | Out
	}

	ctx := context.Background()
	jibbleClient := Client(struct{ token string }{service.accessToken})
	err := jibbleClient.Clocking(ctx, requestPayload)
	if err != nil {
		fmt.Printf("[Error ClockIn]: %v\n", err)
		return
	}
}
