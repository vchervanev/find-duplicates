package files

type filesByMD5 map[string][]string

func newFilesByMD5(size int) filesByMD5 {
	return make(map[string][]string, size)
}

func (l filesByMD5) register(path string) {
	//  TODO error handling
	md5 := md5sum(path)
	l[md5] = append(l[md5], path)
}
