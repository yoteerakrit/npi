package domain

type TableName string

const (
	TableNameSchedule TableName = "schedule"
)

func (t TableName) String() string {
	return string(t)
}
