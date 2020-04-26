package files

type stringsByInt map[int64][]string

func newStringsByInt(size int) stringsByInt {
	return make(map[int64][]string, size)
}

func (l stringsByInt) register(rec record) {
	l[rec.size] = append(l[rec.size], rec.path)
}
