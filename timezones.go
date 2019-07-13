package worldtime

import "time"

type Timezone string

type TimezoneInfo struct {
	WeekNumber   int       `json:"week_number"`
	UTCOffset    string    `json:"utc_offset"`
	UTCDateTime  time.Time `json:"utc_datetime"`
	UnixTime     int       `json:"unixtime"`
	Timezone     Timezone  `json:"timezone"`
	RawOffset    int       `json:"raw_offset"`
	DSTUntil     time.Time `json:"dst_until"`
	DSTOffset    int       `json:"dst_offset"`
	DSTFrom      time.Time `json:"dst_from"`
	DST          bool      `json:"dst"`
	DayOfYear    int       `json:"day_of_year"`
	DayOfWeek    int       `json:"day_of_week"`
	Datetime     time.Time `json:"datetime"`
	ClientIP     string    `json:"client_ip"`
	Abbreviation string    `json:"abbreviation"`
}
