package sort

type Direction string

const (
	Asc  Direction = "ASC"
	Desc Direction = "DESC"
)

func (d Direction) String() string {
	return string(d)
}
