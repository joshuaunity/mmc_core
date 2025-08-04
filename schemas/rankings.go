package s

import (
	"encoding/json"
	"time"
)

var WebAppUrl = "https://www.mymusicchart.com"

type WeightRecord struct {
	Ref    string  `json:"ref"`
	Weight float64 `json:"weight"`
}

// Custom Query Ranking Durations
type QueryRankingDurationType string

const (
	LastWeek    QueryRankingDurationType = "lastWeek"
	LastMon     QueryRankingDurationType = "lastMon"
	CurrentYear QueryRankingDurationType = "currentYear"
	CustomYear  QueryRankingDurationType = "customYear"
)

// Yearly Quaters
type YearlyQuater string

const (
	Q1 YearlyQuater = "Q1"
	Q2 YearlyQuater = "Q2"
	Q3 YearlyQuater = "Q3"
	Q4 YearlyQuater = "Q4"
)

// year quaters list
var YearlyQuaterList = []YearlyQuater{Q1, Q2, Q3, Q4}

// Months of the year
type Month string

const (
	Jan Month = "January"
	Feb Month = "February"
	Mar Month = "March"
	Apr Month = "April"
	May Month = "May"
	Jun Month = "June"
	Jul Month = "July"
	Aug Month = "August"
	Sep Month = "September"
	Oct Month = "October"
	Nov Month = "November"
	Dec Month = "December"
)

// Months of the year list
var MonthList = []Month{Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec}

// RankingDurationType represents the enum for duration types
type RankingDurationType string

const (
	DAILY   RankingDurationType = "daily"
	WEEKLY  RankingDurationType = "weekly"
	MONTHLY RankingDurationType = "monthly"
	YEARLY  RankingDurationType = "yearly"
)

// RankingType represents the enum for ranking types
type RankingType string

const (
	SONGS       RankingType = "songs"
	ARTISTS     RankingType = "artists"
	ALBUMS      RankingType = "albums"
	ALBUM_SONGS RankingType = "album_songs"
)

// Ranking represents the ranking model in MongoDB
type Ranking struct {
	ID           string
	User         string
	DurationType RankingDurationType
	Entries      []Entry
	CreatedAt    time.Time
}

type NomineeList struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Nominees   []Nominee `json:"nominees"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Nominee struct {
	ID            string          `json:"id"` // ID is a string
	Position      int             `json:"position"`
	UserID        string          `json:"user_id"`
	NomineeListID string          `json:"nominee_list_id"`
	SongRef       string          `json:"song_ref"`
	ArtistRefs    json.RawMessage `gorm:"type:jsonb" json:"artist_refs"`
	AlbumRef      string          `json:"album_ref"`
	Song          json.RawMessage `gorm:"type:jsonb" json:"song"` // Storing raw JSON in PostgreSQL
	CreatedAt     time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

// EntryStatus represents the status of an entry
type EntryStatus string

const (
	UP       EntryStatus = "up"
	DOWN     EntryStatus = "down"
	SAME     EntryStatus = "same"
	NEW      EntryStatus = "new"
	RE_ENTRY EntryStatus = "reEntry"
)

// ExternalUrls represents the external URLs for a song
type ExternalUrls struct {
	Spotify string `bson:"spotify" json:"spotify"`
}

// Artist represents an artist for a song
type Artist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type ArtistV2 struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Image      `json:"images"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

// Image represents an image for an album
type Image struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

// Album represents an album for a song
type Album struct {
	AlbumType            string       `json:"album_type"`
	Artists              []Artist     `json:"artists"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	TotalTracks          int          `json:"total_tracks"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
}

// ExternalIds represents external IDs for a song
type ExternalIds struct {
	ISRC string `bson:"isrc" json:"isrc"`
}

// Song represents a song
type Song struct {
	Album        Album        `json:"album"`
	Artists      []Artist     `json:"artists"`
	DiskNumber   int          `json:"disc_number"`
	DurationMS   int          `json:"duration_ms"`
	Explicit     bool         `json:"explicit"`
	ExternalIds  ExternalIds  `json:"external_ids"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Position     int          `json:"position"`
	IsLocal      bool         `json:"is_local"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	PreviewURL   string       `json:"preview_url"`
	TrackNumber  int          `json:"track_number"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type ArtistRef struct {
	ID string `json:"id"`
}

type TopSongEntry struct {
	Count int
	Song  Song
}


type Ranking struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	DurationType string    `json:"duration_type"`
	RankingType  string    `json:"ranking_type"`
	Entries      []Entry   `json:"entries"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Entry struct {
	ID         string    `json:"id"`
	Position   int       `json:"position"`
	UserID     string    `json:"user_id"`
	RankingID  string    `json:"ranking_id"`
	SongRef    string    `json:"songRef"`
	ArtistRefs []string  `json:"artistRefs"`
	Status     string    `json:"status"`
	Song       Song      `gorm:"type:jsonb" json:"song"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


// TrendingItem represents an item that is trending.
type TrendingItem struct {
	ID     string  `json:"id"`     // Unique identifier for the trending item
	Name   string  `json:"name"`   // Name of the trending item
	Weight float64 `json:"weight"` // Weight of the trending item
	Images []Image `json:"images"` // List of images associated with the trending item
}

// EntryWeights defines the weights for different positions.
const (
	EntryWeightFirst   = 1.0
	EntryWeightSecond  = 0.9
	EntryWeightThird   = 0.8
	EntryWeightFourth  = 0.7
	EntryWeightFifth   = 0.6
	EntryWeightSixth   = 0.5
	EntryWeightSeventh = 0.4
	EntryWeightEighth  = 0.3
	EntryWeightNinth   = 0.2
	EntryWeightTenth   = 0.1
)

// EntryPositions defines the positions from first to tenth.
const (
	EntryPositionFirst   = 1
	EntryPositionSecond  = 2
	EntryPositionThird   = 3
	EntryPositionFourth  = 4
	EntryPositionFifth   = 5
	EntryPositionSixth   = 6
	EntryPositionSeventh = 7
	EntryPositionEighth  = 8
	EntryPositionNinth   = 9
	EntryPositionTenth   = 10
)

type Request struct {
	Action string `json:"action"`
	Data   Data   `json:"data"`
}

type Data struct {
	Query string `json:"query"`
}

// TrendingItem represents an item that is trending.
type TrendingItem struct {
	ID     string  `json:"id"`     // Unique identifier for the trending item
	Name   string  `json:"name"`   // Name of the trending item
	Weight float64 `json:"weight"` // Weight of the trending item
	Images []Image `json:"images"` // List of images associated with the trending item
}

type UserTrending struct {
	ID           string               `json:"id"`
	User         string               `json:"user"`
	DurationType string               `json:"duration_type"`
	Songs        []TrendingItem `json:"songs"`
	Artists      []TrendingItem `json:"artists"`
	CreatedAt    string               `json:"createdAt"`
	UpdatedAt    string               `json:"updatedAt"`
}

type RankingSearch struct {
	ID           string `json:"id"`
	User         string `json:"user"`
	DurationType string `json:"duration_type"`
	Entries      string `json:"entries"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type YearInMusic struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Year      int             `json:"year"`
	Data      YearInMusicData `json:"data"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type YearInMusicData struct {
	MostNo1Artist           ArtistV2       `json:"most_no1_artist"`
	MostNo1ArtistEntries    int            `json:"most_no1_artist_entries"`
	MostTop3Artist          ArtistV2       `json:"most_top3_artist"`
	MostTop3ArtistEntries   int            `json:"most_top3_artist_entries"`
	MostTop5Artist          ArtistV2       `json:"most_top5_artist"`
	MostTop5ArtistEntries   int            `json:"most_top5_artist_entries"`
	MostBelow5Artist        ArtistV2       `json:"most_below5_artist"`
	MostBelow5ArtistEntries int            `json:"most_below5_artist_entries"`
	LongestNo1Song          *Song          `json:"longest_no1_song"`
	LongestNo1SongEntries   int            `json:"longest_no1_song_entries"`
	MostOccurringSongs      []TopSongEntry `json:"most_occurring_songs"`
	TotalRankings           int64          `json:"total_rankings"`
	TrendingArtists         []TrendingItem `json:"trending_artists"`
	TrendingSongs           []TrendingItem `json:"trending_songs"`
}

// Streaming Serice Token
type ServiceToken struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Service      string    `json:"service"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}