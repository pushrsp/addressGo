package address

import (
	"github.com/djimenez/iconv-go"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	DefaultFileFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	DefaultFileMode = os.ModePerm
)

type Collector struct {
	Src        string
	Dest       string
	FilePhases []FilePhase
}

type workableFile struct {
	file      os.DirEntry
	filePhase FilePhase
}

type workableFileGroups map[string][]workableFile

func (c *Collector) Merge() {
	groupFiles := groupFilesByPrefix(c.Src, c.FilePhases...)

	//for _, workableFiles := range groupFiles {
	//	for _, workableFile := range workableFiles {
	//		c.copyFiles(workableFile)
	//	}
	//} 10초

	var wg sync.WaitGroup
	wg.Add(len(groupFiles))

	for _, workableFiles := range groupFiles {
		workableFiles := workableFiles
		go func() {
			for _, workableFile := range workableFiles {
				c.copyFiles(workableFile)
			}
			wg.Done()
		}()
	}

	wg.Wait() //6초
}

func (c Collector) copyFiles(wf workableFile) {
	osFile := wf.file
	filePhase := wf.filePhase

	rFile, _ := os.OpenFile(c.Src+"\\"+osFile.Name(), DefaultFileFlag, DefaultFileMode)
	_ = os.Mkdir(c.Dest, DefaultFileMode)
	wFile, _ := os.OpenFile(c.Dest+"\\"+filePhase.PreFix+".txt", DefaultFileFlag, DefaultFileMode)

	rStat, _ := rFile.Stat()
	bytes := make([]byte, rStat.Size())

	for {
		n, err := rFile.Read(bytes)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("while reading, error occured cause: %s\n", err.Error())
				os.Exit(-1)
			}
		}

		line := bytes[:n]
		byteCode, err := encodingConverter(string(line), filePhase.Encoding)
		if err != nil {
			log.Fatalf("while reading, error occured cause: %s\n", err.Error())
			os.Exit(-1)
		}

		wFile.Write(byteCode)
	}
}

func groupFilesByPrefix(src string, filePhases ...FilePhase) workableFileGroups {
	files, _ := os.ReadDir(src)
	ret := workableFileGroups{}

	for _, file := range files {
		filePhase := hasPrefix(file.Name(), filePhases...)
		if filePhase == nil {
			continue
		}

		workableFiles, isExist := ret[filePhase.PreFix]
		if !isExist {
			workableFiles = []workableFile{}
			ret[filePhase.PreFix] = workableFiles
		}

		ret[filePhase.PreFix] = append(workableFiles, workableFile{file: file, filePhase: *filePhase})
	}

	return ret
}

func hasPrefix(fileName string, files ...FilePhase) *FilePhase {
	for _, file := range files {
		if strings.HasPrefix(fileName, file.PreFix) {
			return &file
		}
	}

	return nil
}

func encodingConverter(line string, encoding string) ([]byte, error) {
	str, err := iconv.ConvertString(line, encoding, "utf-8")

	return []byte(str), err
}
