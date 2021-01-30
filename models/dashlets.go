package models

// User is a struct containing lastfm user data
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RealName   string `json:"realName"`
	URL        string `json:"url"`
	Country    string `json:"country"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	Subscriber string `json:"subscriber"`
	PlayCount  string `json:"playCount"`
	Playlists  string `json:"playlists"`
	Registered string `json:"registered"`
	Image      string `json:"image"`
}

// UserArtists default
type UserArtists struct {
	User   string
	Type   string
	Total  int
	Artist []ArtistInfo
}

// ArtistInfo nfo
type ArtistInfo struct {
	Rank      string
	Name      string
	PlayCount string
	URL       string
	Image     string
}
