package ui

import "github.com/gizak/termui"

const (
	finishedRedditDownload = "/lazarus/reddit/download/done"
	songListUpdated        = "/lazarus/update/songlist"
)

var (
	termuiSendCustomEvt = termui.SendCustomEvt
	termuiHandle        = termui.Handle
)

// UpdatePlayer Fires the event to redraw the player
func UpdatePlayer(player Player) {
	termuiSendCustomEvt(songListUpdated, player)
}

// FireFinishedRedditDownload Fires the event to begin the playlist display/download/play process
func FireFinishedRedditDownload(player Player) {
	termuiSendCustomEvt(finishedRedditDownload, player)
}

// EventHandler Registers all the event handlers
func EventHandler() {
	termuiHandle("/sys/kbd/q", func(termui.Event) { termui.StopLoop() })

	termuiHandle(finishedRedditDownload, paintSongList)
	termuiHandle(songListUpdated, paintSongList)
}
