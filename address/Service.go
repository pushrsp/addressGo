package address

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Service struct {
	Source       string
	Dest         string
	FilePhase    FilePhase
	ShardService ShardService
}

func NewService(source string, dest string, filePhase FilePhase, numOfShards int) Service {
	service := Service{
		Source:    source,
		Dest:      dest,
		FilePhase: filePhase,
		ShardService: ShardService{
			ShardInfo: ShardInfo{
				Dest:        dest,
				NumOfShards: numOfShards,
			},
			FileCache: map[string]*os.File{},
		},
	}

	return service
}

func (s Service) Sort() {
	os.Mkdir(s.Dest, DefaultFileMode)

	rFile, _ := os.OpenFile(s.Source+"\\"+s.FilePhase.PreFix+".txt", DefaultFileFlag, DefaultFileMode)
	defer rFile.Close()

	csvReader := csv.NewReader(rFile)
	csvReader.Comma = '|'

	for {
		record, err := csvReader.Read()
		if record == nil || err == io.EOF {
			break
		}

		s.bucketlize(record)
	}

	s.doSort()
}

func (s Service) bucketlize(record []string) {
	id := record[s.FilePhase.FieldIdx]
	shardFile := s.ShardService.OpenFile(id, s.FilePhase.PreFix)
	defer shardFile.Close()
	if shardFile == nil {
		log.Fatalln("Error: OpenFile")
		os.Exit(-1)
	}

	wCsv := csv.NewWriter(shardFile)
	wCsv.Comma = '|'
	wCsv.Write(record)
	wCsv.Flush()
}

func (s Service) doSort() {
	files, isExist := groupFilesByPrefix(s.Dest, s.FilePhase)[s.FilePhase.PreFix]
	if !isExist {
		return
	}

	for _, fileInfo := range files {
		func() {
			file, _ := os.OpenFile(s.Dest+"/"+fileInfo.file.Name(), DefaultFileFlag, DefaultFileMode)
			defer file.Close()

			reader := csvReader(file)
			records, _ := reader.ReadAll()

			sorter := sorter{records: &records, idFieldIdx: s.FilePhase.FieldIdx}
			sorter.Sort()

			writer := csvWriter(file)
			writer.WriteAll(records)

			//for k, v := range records {
			//	fmt.Println(file.Name(), k, v)
			//}
		}()
	}
}

func csvReader(file *os.File) csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = '|'
	return *reader
}

func csvWriter(file *os.File) *csv.Writer {
	writer := csv.NewWriter(file)
	writer.Comma = '|'
	return writer
}

func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}
