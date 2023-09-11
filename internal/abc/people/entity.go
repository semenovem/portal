package people

type UserPosition struct {
	ID       uint16
	Title    string
	ParentID uint16
}

type UserDept struct {
	ID       uint16
	Title    string
	ParentID uint16
	BossID   uint16
}

type UserBoss struct {
	ID            uint32
	Firstname     string
	Surname       string
	PositionTitle string
}
