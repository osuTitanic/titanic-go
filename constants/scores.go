package constants

import "fmt"

type Mode int8

const (
	ModeOsu   Mode = 0
	ModeTaiko Mode = 1
	ModeCatch Mode = 2
	ModeMania Mode = 3
)

func (m Mode) String() string {
	switch m {
	case ModeOsu:
		return "osu!"
	case ModeTaiko:
		return "Taiko"
	case ModeCatch:
		return "CatchTheBeat"
	case ModeMania:
		return "osu!mania"
	default:
		return fmt.Sprintf("%d", m)
	}
}

func (m Mode) Alias() string {
	switch m {
	case ModeOsu:
		return "osu"
	case ModeTaiko:
		return "taiko"
	case ModeCatch:
		return "catch"
	case ModeMania:
		return "mania"
	default:
		return fmt.Sprintf("%d", m)
	}
}

type Grade int8

const (
	GradeXH Grade = 0
	GradeSH Grade = 1
	GradeX  Grade = 2
	GradeS  Grade = 3
	GradeA  Grade = 4
	GradeB  Grade = 5
	GradeC  Grade = 6
	GradeD  Grade = 7
	GradeF  Grade = 8
	GradeN  Grade = 9
)

func (g Grade) String() string {
	switch g {
	case GradeXH:
		return "XH"
	case GradeSH:
		return "SH"
	case GradeX:
		return "X"
	case GradeS:
		return "S"
	case GradeA:
		return "A"
	case GradeB:
		return "B"
	case GradeC:
		return "C"
	case GradeD:
		return "D"
	case GradeF:
		return "F"
	case GradeN:
		return "N"
	default:
		return fmt.Sprintf("Unknown(%d)", g)
	}
}

type ScoreStatus int

const (
	ScoreStatusHidden    ScoreStatus = -1
	ScoreStatusFailed    ScoreStatus = 0
	ScoreStatusExited    ScoreStatus = 1
	ScoreStatusSubmitted ScoreStatus = 2
	ScoreStatusBest      ScoreStatus = 3
	ScoreStatusMods      ScoreStatus = 4
)
