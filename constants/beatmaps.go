package constants

type BeatmapStatus int

const (
	BeatmapStatusInactive  BeatmapStatus = -3
	BeatmapStatusGraveyard BeatmapStatus = -2
	BeatmapStatusWIP       BeatmapStatus = -1
	BeatmapStatusPending   BeatmapStatus = 0
	BeatmapStatusRanked    BeatmapStatus = 1
	BeatmapStatusApproved  BeatmapStatus = 2
	BeatmapStatusQualified BeatmapStatus = 3
	BeatmapStatusLoved     BeatmapStatus = 4
)

type BeatmapServer int

const (
	BeatmapServerBancho  BeatmapServer = 0
	BeatmapServerTitanic BeatmapServer = 1
)
