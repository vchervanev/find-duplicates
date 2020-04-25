package files

// import "sync"

// type manager struct{
// 	workers []worker
// 	dirsQueue chan string
// 	wg sync.WaitGroup
// }

// func (m manager) perform(path string, filesQueue chan os.FileInfo, size int) {
// 	m.dirsQueue = make(chan string, 1000)
// 	m.workers = make([]worker, size)
// 	for id := 0; i++; i < size {
// 		w := worker{id: id, mgr: m}
// 		m.workers[id] = w
// 		go w.start()
// 	}
// 	m.addDir(path)
// 	m.wg.Wait()
// 	close(m.dirsQueue)
// }

// func (m manager) addDir(path string) {
// 	m.dirsQueue <- path
// 	m.wg.Add(1)
// }

// type worker struct {
// 	id int
// 	mgr *manager
// }

// func (w worker) start() {
// 	for path := range w.mgr.dirsQueue {
// 		fmt.Println("before", len(w.mgr.dirsQueue))
// 		f.visit(path)
// 		fmt.Println("after", len(w.mgr.dirsQueue))
// 	}
// }

// func (w worker) visit(path string) {
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	for _, file := range files {
// 		if file.IsDir() {
// 			w.mgr.addDir(path + "/" + file.Name())
// 		} else {
// 			// fmt.Println(formatFileInfo(path, file))
// 		}
// 	}
// }
