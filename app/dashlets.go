package app

import (
	"strconv"
	"time"

	"github.com/lcmps/hippyfm/models"
	"github.com/shkh/lastfm-go/lastfm"
)

// GetUserInfo retuns information about the user
func GetUserInfo(conn *lastfm.Api, u string) (models.User, error) {
	var nfo models.User
	opts := lastfm.P{"user": u}

	info, err := conn.User.GetInfo(opts)
	if err != nil {
		return nfo, err
	}

	t, _ := strconv.ParseInt(info.Registered.Unixtime, 10, 64)

	nfo = models.User{
		ID:         info.Id,
		Name:       info.Name,
		RealName:   info.RealName,
		URL:        info.Url,
		Country:    info.Country,
		Age:        info.Age,
		Gender:     info.Gender,
		Subscriber: info.Subscriber,
		PlayCount:  info.PlayCount,
		Playlists:  info.Playlists,
		Registered: time.Unix(t, 0).Format("2006-01-02"),
		Image:      getLargestImage(info.Images),
	}

	return nfo, nil
}

// GetUserTopArtists retrieves user most listened artists
func GetUserTopArtists(conn *lastfm.Api, u, p string, l int) (models.UserArtists, error) {
	var ua models.UserArtists
	var artist []models.ArtistInfo
	opts := lastfm.P{"user": u, "limit": l}

	art, err := conn.User.GetTopArtists(opts)
	if err != nil {
		return ua, err
	}

	for _, item := range art.Artists {
		a := models.ArtistInfo{
			Rank:      item.Rank,
			Name:      item.Name,
			PlayCount: item.PlayCount,
			URL:       item.Url,
			Image:     getLargestImage(item.Images),
		}
		artist = append(artist, a)
	}

	ua = models.UserArtists{
		User:   art.User,
		Type:   art.Type,
		Total:  art.Total,
		Artist: artist,
	}

	return ua, nil
}
