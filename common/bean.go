package common

import (
	"time"
)

type TelemetryUserAnalyticsDto struct {
	Id              int32     `json:"id"`
	UPID            string    `json:"upid"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	EventType       string    `json:"eventType,omitempty"` //startup,normal,frequency
	ServerVersion   string    `json:"serverVersion"`
	DevtronVersion  string    `json:"devtronVersion"`
	ActiveSince     time.Time `json:"activeSince,omitempty"`
	LastActive      time.Time `json:"lastActive,omitempty"`
	Clusters        int       `json:"clusters"`
	Environments    int       `json:"environments"`
	NoOfProdApps    int       `json:"prodApps"`
	NoOfNonProdApps int       `json:"nonProdApps"`
	Users           int       `json:"users"`
}
