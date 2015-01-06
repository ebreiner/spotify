package spotify

// SimpleTrack contains basic info about a track.
type SimpleTrack struct {
	// The artists who performed the track.
	Artists []SimpleArtist `json:"artists"`
	// A list of the countries in which the track
	// can be played, identified by their ISO 3166-1
	// alpha-2 codes.
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually 1 unless the album
	// consists of more than one disc).
	DiscNumber int `json:"disc_number"`
	// The length of the track, in milliseconds. TODO: time package?
	Duration int `json:"duration_ms"`
	// Whether or not the track has explicit lyrics.
	// true => yes, it does; false => no, it does not.
	Explicit bool `json:"explicit"`
	// External URLs for this track.
	ExternalURLs ExternalURL `json:"external_urls"`
	// A link to the Web API endpoint providing full
	// details for this track.
	Endpoint string `json:"href"`
	// The Spotify ID for the track.
	ID ID `json:"id"`
	// The name of the track
	Name string `json:"name"`
	// A URL to a 30 second preview (MP3) of the track.
	PreviewURL string `json:"preview_url"`
	// The number of the track.  If an album has several
	// discs, the track number is the number on the specified
	// DiscNumber.
	TrackNumber int `json:"track_number"`
	// The Spotify URI for the track.
	URI URI `json:"uri"`
}

// FullTrack provides extra track data in addition
// to what is provided by SimpleTrack.
type FullTrack struct {
	SimpleTrack
	// Known external IDs for the track.
	ExternalIDs ExternalID `json:"exernal_ids"`
	// Popularity of the trakc.  The value will be
	// between 0 and 100, with 100 being the most
	// popular.  The popularity is calculated from
	// both total plays and most recent plays.
	Popularity int `json:"popularity"`
}

// PlaylistTrack contains info about a track
// in a playlist.
type PlaylistTrack struct {
	// The date and time the track was added to the playlist.
	// TODO: very old playlists may return null here.
	AddedAt Timestamp `json:"added_at"`
	// The Spotify user who added the track to the playlist.
	// TODO: very old playlists may return null here
	AddedBy User `json:"added_by"`
	// Information about the track.
	Track FullTrack `json:"track"`
}

// SavedTrack provides info about a track saved
// to a user's account.
type SavedTrack struct {
}

func (t *SimpleTrack) String() string {
	return "SimpleTrack" + t.Name
}

func (t *FullTrack) String() string {
	return "FullTrack" + t.Name
}