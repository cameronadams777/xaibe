package invite_status

type InviteStatus int

const (
	PENDING  InviteStatus = 0
	ACCEPTED InviteStatus = 1
	REJECTED InviteStatus = 2
)

func (i InviteStatus) String() string {
	return [...]string{"PENDING", "ACCEPTED", "REJECTED"}[i-1]
}

func (i InviteStatus) EnumIndex() int {
	return int(i)
}
