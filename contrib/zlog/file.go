package zlog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

const (
	bufferSize    = 256 * 1024
	flushDuration = time.Second * 5
)

var _ io.WriteCloser = &FileBackend{}

type FileBackend struct {
	mu            sync.Mutex
	file          *os.File
	buffer        *bufio.Writer
	dir           string //directory for log files
	name          string
	filePath      string
	lastCheck     uint64
	flushDuration time.Duration
	closeCh       chan struct{}
}

func (p *FileBackend) Write(b []byte) (n int, err error) {
	p.mustFileExist()
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.buffer.Write(b)
}

func (p *FileBackend) Flush() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.buffer.Flush()
}
func (p *FileBackend) Close() error {
	close(p.closeCh)
	p.mu.Lock()
	defer p.mu.Unlock()
	p.buffer.Flush()
	p.file.Sync()
	return p.file.Close()
}

func (p *FileBackend) monitorFiles() {
	p.lastCheck = getLastCheck(time.Now())
	for range time.NewTicker(time.Second * 5).C {
		fileName := path.Join(p.dir, p.name)
		check := getLastCheck(time.Now())
		if p.lastCheck >= check {
			continue
		}
		p.mu.Lock()
		os.Rename(fileName, fileName+fmt.Sprintf(".%d", p.lastCheck))
		p.lastCheck = check
		newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		p.buffer.Flush()
		p.file.Close()
		p.file = newFile
		p.buffer.Reset(p.file)
		p.mu.Unlock()
	}
}
func (p *FileBackend) flushFile() {
	ticker := time.NewTicker(p.flushDuration)
	for {
		select {
		case <-ticker.C:
			p.Flush()
		case <-p.closeCh:
			return
		}
	}
}
func (p *FileBackend) mustFileExist() {
	timeStr := time.Now().Format(".20060102")
	filePath := path.Join(p.dir, p.name+timeStr)
	if filePath == p.filePath {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	newFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	p.buffer.Flush()
	p.file.Close()
	p.file = newFile
	p.buffer.Reset(p.file)
	p.filePath = filePath

}

func NewFileBackend(dir, name string) (*FileBackend, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	fb := new(FileBackend)
	fb.dir = dir
	fb.name = name
	fb.buffer = bufio.NewWriterSize(fb.file, bufferSize)
	fb.flushDuration = flushDuration
	fb.closeCh = make(chan struct{})
	fb.mustFileExist()

	// default
	//go fb.monitorFiles()
	go fb.flushFile()
	return fb, nil
}

func getLastCheck(now time.Time) uint64 {
	return uint64(now.Year())*1000000 + uint64(now.Month())*10000 + uint64(now.Day())*100 + uint64(now.Hour())
}
