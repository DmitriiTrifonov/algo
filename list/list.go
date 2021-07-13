package list

type List interface {
	Add()
	Reverse() List
	Length() int
	Slice() []List
}
