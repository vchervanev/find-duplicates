package files

type list struct {
	bySize map[int64][]string
}

func newList() (result *list) {
	result = &list{}
	result.bySize = make(map[int64][]string, 1000)
	return
}

func (l list) register(rec record) {
	l.bySize[rec.size] = append(l.bySize[rec.size], rec.path)
}
