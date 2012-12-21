package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"
	"unsafe"
)

type FileEvent struct {
	mask uint32
	File string
}

func (e *FileEvent) isCreate() {
	return e.mask&syscall.IN_CREATE == syscall.IN_CREATE
}

func add_watch_r(inotify_fd int, f string, paths map[int]string) {
	fi, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	n, err := syscall.InotifyAddWatch(inotify_fd, f, syscall.IN_ALL_EVENTS)
	paths[n] = f
	if err != nil {
		log.Fatal(err)
	}

	s, _ := fi.Stat()
	if s.IsDir() {
		files, _ := ioutil.ReadDir(f)
		for _, file := range files {
			if file.IsDir() {
				f := path.Join(f, file.Name())
				add_watch_r(inotify_fd, f, paths)
				// log.Print(path.Join(f, file.Name()))
			}
		}
	}
}

func printEvent(file string, mask uint32) {
	if mask&syscall.IN_ACCESS != 0 {
		log.Printf("Access %s", file)
	}

	if mask&syscall.IN_ATTRIB != 0 {
		log.Printf("ATTRIB %s", file)
	}

	if mask&syscall.IN_OPEN != 0 {
		log.Printf("OPEN %s", file)
	}

	// log.Print(file, mask)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	inotify_fd, err := syscall.InotifyInit()
	if err != nil {
		log.Fatal(err)
	}
	paths := make(map[int]string)
	add_watch_r(inotify_fd, "/home/feng/workspace/rssminer", paths)
	log.Print("watcher added")

	for {
		var (
			buf [syscall.SizeofInotifyEvent * 4096]byte
		)

		n, _ := syscall.Read(inotify_fd, buf[0:])

		offset := 0
		for offset <= n-syscall.SizeofInotifyEvent {
			raw := (*syscall.InotifyEvent)(unsafe.Pointer(&buf[offset]))
			wd := int(raw.Wd)
			file := paths[wd]
			mask := raw.Mask
			printEvent(file, mask)
			offset += syscall.SizeofInotifyEvent + int(raw.Len)
		}
	}
}
