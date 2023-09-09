package people_provider

type DeptModel struct {
	id          uint16
	title       string
	description string
	parentID    *uint16
	deleted     bool
}
