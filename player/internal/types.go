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
	RoundStatusUnknown  RoundStatus = 0
	RoundStatusJoin     RoundStatus = 1
	RoundStatusIncluded RoundStatus = 2
	RoundStatusExcluded RoundStatus = 3
	roundStatusSentinel RoundStatus = 9
)

