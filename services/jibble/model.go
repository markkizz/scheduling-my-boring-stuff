package jibble

import "time"

type UserAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type OrganizationIdResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		Name   string `json:"name"`
		Status string `json:"status"`
		ID     string `json:"id"`
	} `json:"value"`
}

type PersonIdResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		UserID         string `json:"userId"`
		OrganizationID string `json:"organizationId"`
		Title          string `json:"title"`
		Status         string `json:"status"`
		ID             string `json:"id"`
	} `json:"value"`
}

type PersonAccessToken struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	TokenType      string `json:"token_type"`
	RefreshToken   string `json:"refresh_token"`
	Scope          string `json:"scope"`
	PersonID       string `json:"personId"`
	OrganizationID string `json:"organizationId"`
}

type ClockingResponse struct {
	OdataContext           string      `json:"@odata.context"`
	BelongsToDate          string      `json:"belongsToDate"`
	LocalTime              time.Time   `json:"localTime"`
	PersonID               string      `json:"personId"`
	OrganizationID         string      `json:"organizationId"`
	ProjectID              interface{} `json:"projectId"`
	ActivityID             interface{} `json:"activityId"`
	LocationID             interface{} `json:"locationId"`
	KioskID                interface{} `json:"kioskId"`
	BreakID                interface{} `json:"breakId"`
	ClientType             string      `json:"clientType"`
	Type                   string      `json:"type"`
	Time                   time.Time   `json:"time"`
	Offset                 string      `json:"offset"`
	AutoClockOutTime       interface{} `json:"autoClockOutTime"`
	ClockInOutReminderTime interface{} `json:"clockInOutReminderTime"`
	IsOffline              bool        `json:"isOffline"`
	IsFaceRecognized       interface{} `json:"isFaceRecognized"`
	FaceSimilarity         interface{} `json:"faceSimilarity"`
	IsAutomatic            bool        `json:"isAutomatic"`
	IsManual               bool        `json:"isManual"`
	IsUnusual              bool        `json:"isUnusual"`
	IsEndOfDay             bool        `json:"isEndOfDay"`
	Note                   interface{} `json:"note"`
	Status                 string      `json:"status"`
	CreatedAt              time.Time   `json:"createdAt"`
	UpdatedAt              time.Time   `json:"updatedAt"`
	IsLocked               bool        `json:"isLocked"`
	ID                     string      `json:"id"`
	Coordinates            interface{} `json:"coordinates"`
	Picture                interface{} `json:"picture"`
	Platform               struct {
		ClientVersion interface{} `json:"clientVersion"`
		Os            interface{} `json:"os"`
		DeviceModel   interface{} `json:"deviceModel"`
		DeviceName    interface{} `json:"deviceName"`
	} `json:"platform"`
}

type UserAccessTokenRequest struct {
	username, password string
}

type PersonIdRequest struct {
	organizationId string
}

type PersonAccessTokenRequest struct {
	username, password, personId, refreshToken string
}
