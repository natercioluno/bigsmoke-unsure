package internal

type RoundStatus int

func (rs RoundStatus) Enum() int {
	return int(rs)
}

func (rs RoundStatus) ShiftStatus() {}

func (rs RoundStatus) ReflexType() int {
	return int(rs)
}

const (
	RoundStatusUnknown  RoundStatus = iota
	RoundStatusJoin
	RoundStatusIncluded
	RoundStatusExcluded
	RoundStatusCollect
	roundStatusSentinel
)

