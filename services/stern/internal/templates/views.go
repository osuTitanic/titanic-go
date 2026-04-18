package templates

import "github.com/osuTitanic/titanic-go/internal/config"

type Statistics struct {
	TotalUsers  int
	OnlineUsers int
	TotalScores int
}

type DefaultView struct {
	Stats  Statistics
	Config *config.Config
}

type HomeView struct {
	DefaultView
	// TODO: Add news, chat, most played maps, ...
}
