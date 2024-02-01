//go:build darwin

package platform

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/matishsiao/goInfo"
)

func GetSongData() (string, error) {
	osInfo, err := goInfo.GetInfo()
	kernelVersion, _ := strconv.Atoi(strings.Split(osInfo.Core, ".")[0])
	musicApplication := ""
	if kernelVersion <= 18 {
		musicApplication = "iTunes"
	} else {
		musicApplication = "Music"
	}

	cmd := fmt.Sprintf(`¬
		log "<song>"
		tell application "%s"
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
		log "</song>"`, musicApplication)

	out, err := mack.Tell(musicApplication, cmd)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Can’t get name of current track."):
			return "", errors.New("no song is playing")
		}
	}

	return out, nil
}
