package lister

func New() Lister {
	l := new(lister)
	l.init()
	return l
}
