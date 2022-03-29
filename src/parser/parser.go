package parser

import (
	"encoding/xml"
	"os"
	"strings"
)

// yup, everything below is a big one string

type SongStruct struct {
	XMLName  []xml.Name `xml:"song"`
	Type     string     `xml:"type,attr"`
	Name     string     `xml:"name"`
	Artist   string     `xml:"artist"`
	Album    string     `xml:"album"`
	Year     string     `xml:"year"`
	Duration string     `xml:"duration"`
	Position string     `xml:"position"`
	State    string     `xml:"state"`
}

func Parse(songXml string) (songParsed SongStruct) {
	var song SongStruct

	// xml doesn't like the & character :-(
	songXml = strings.ReplaceAll(songXml, "&", "&#x26;")

	err := xml.Unmarshal([]byte(songXml), &song)
	if err != nil {
		switch {
		// for example, on exiting Music,
		// the parser returns EOF because applescript has stopped
		// sending the song data list,
		// so we leave as if nothing had happened
		case strings.Contains(err.Error(), "EOF"):
			os.Exit(0)
		}
	}

	// umm, yeah, region-related stuff like different separators
	song.Duration = strings.ReplaceAll(song.Duration, ",", ".")
	song.Position = strings.ReplaceAll(song.Position, ",", ".")

	return song
}
