package enums

var PriorityOptions = []string{"Low", "Medium", "High", "Urgent"}

type Priority int

const (
	Low Priority = iota
	Medium
	High
	Urgent
)

func (p Priority) String() string {

	if p < 0 || int(p) >= len(PriorityOptions) {
		return "Unknown"
	}
	return PriorityOptions[p]
}
