package presence

import (
	"errors"
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"github.com/sech1p/AppleMusic_Discord_RPC/src/parser"
	"strconv"
	"time"
)

const RpcId string = "820367561975529483"

func Login() error {
	err := client.Login(RpcId)
	if err != nil {
		return errors.New("login error")
	}

	return nil
}

func Update(song parser.SongStruct) error {
	var state []string
	var stateString string
	var isRadio bool
	var timestamps client.Timestamps

	startTime := time.Now()

	duration, _ := strconv.ParseFloat(song.Duration, 64)
	position, _ := strconv.ParseFloat(song.Position, 64)

	elapsedTime := startTime.Add(-time.Second * time.Duration(position))
	endTime := startTime.Add(time.Second * time.Duration((duration)-(position)))

	if song.State == "playing" {
		state = []string{"play", "Playing"}
	} else if song.State == "paused" {
		state = []string{"pause", "Paused"}
	}

	// we probably caught one of the apple music radiostations
	if song.Duration == "missing value" {
		isRadio = true
	}

	if isRadio {
		stateString = "Radio"
		timestamps = client.Timestamps{
			Start: &elapsedTime,
		}
	} else {
		stateString = fmt.Sprintf("%s - %s (%s)", song.Artist, song.Album, song.Year)
		timestamps = client.Timestamps{
			Start: &startTime,
			End:   &endTime,
		}
	}

	if len(state) == 2 {
		err := client.SetActivity(client.Activity{
			State:      stateString,
			Details:    song.Name,
			LargeImage: "logo",
			LargeText:  "🎧 Apple Music Presence by sech1p",
			SmallImage: state[0],
			SmallText:  state[1],
			Timestamps: &timestamps,
		})

		if err != nil {
			return errors.New("rpc update fail")
		}
	}

	return nil
}
