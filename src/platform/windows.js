var iTunesInstance = WScript.CreateObject("iTunes.Application");

var currentTrack = iTunesInstance.currentTrack;
var Track = {
	name: currentTrack.name,
	artist: currentTrack.artist,
	album: currentTrack.album,
	year: currentTrack.year,
	duration: currentTrack.duration,
	position: iTunesInstance.PlayerPosition,
	state: ""
};

switch (iTunesInstance.PlayerState) {
	case 0:
		if (iTunesInstance.currentTrack)
			Track.state = "paused";
		else
			Track.state = "stopped";
		break;
	case 1:
		Track.state = "playing";
		break;
}

WScript.Echo("<song>")
WScript.Echo("<name>" + Track.name + "</name>");
WScript.Echo("<artist>" + Track.artist + "</artist>");
WScript.Echo("<album>" + Track.album + "</album>");
WScript.Echo("<year>" + Track.year + "</year>");
WScript.Echo("<duration>" + Track.duration + "</duration>");
WScript.Echo("<position>" + Track.position + "</position>");
WScript.Echo("<state>" + Track.state + "</state>");
WScript.Echo("</song>");