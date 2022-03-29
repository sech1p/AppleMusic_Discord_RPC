//go:build darwin

package platform

import (
	"errors"
	"github.com/andybrewer/mack"
	"strings"
)

func GetSongData() (string, error) {
	const cmd = `¬
		log "<song>"
		tell application "Music"
			set key_array to {"name", ¬
				"artist", ¬
				"album", ¬
				"year", ¬
				"duration", ¬
				"position", ¬
				"state"}
			set song_array to {name of current track, ¬
				artist of current track, ¬
				album of current track, ¬
				year of current track, ¬
				duration of current track, ¬
				player position, ¬
				player state}
			repeat with x from 1 to length of song_array
				set i to item x of song_array
				log "<" & item x of key_array & ">" ¬
					& i & ¬
				"</" & item x of key_array & ">"
			end repeat
		end tell
		log "</song>"`

	out, err := mack.Tell("Music", cmd)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Can’t get name of current track."):
			return "", errors.New("no song is playing")
		}
	}

	return out, nil
}
