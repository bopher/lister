package lister

func New() Lister {
	l := new(lDriver)
	l.init()
	return l
}
