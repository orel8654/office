package types

import "time"

type OfficeData struct {
	UUID       string
	Name       string
	Address    string
	Created_at time.Time
}

type OfficeList struct {
	Result []OfficeData
}

type OfficeMake struct {
	Name    string
	Address string
}
