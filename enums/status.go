package enums

var StatusOptions = []string{"New", "In Progress", "On Hold", "Pending", "Closed", "Done"}

type Status int

const (
	New Status = iota
	InProgress
	OnHold
	Pending
	Closed
	Done
)

func (s Status) String() string {

	if s < 0 || int(s) >= len(StatusOptions) {
		return ""
	}
	return StatusOptions[s]
}
