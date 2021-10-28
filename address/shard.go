package address

import (
	"os"
	"strconv"
)

type ShardInfo struct {
	NumOfShards int
	Dest        string
}

type ShardService struct {
	ShardInfo ShardInfo
	FileCache map[string]*os.File
}

func (ss ShardService) OpenFile(id string, preFix string) *os.File {
	shardId := preFix + "_" + strconv.Itoa(hash(id)%ss.ShardInfo.NumOfShards)

	fileCache, isExist := ss.FileCache[shardId]
	if !isExist {
		fileCache, _ = os.OpenFile(ss.ShardInfo.Dest+"\\"+shardId+".txt", DefaultFileFlag, DefaultFileMode)
		ss.FileCache[shardId] = fileCache
	}
	//fileCache.Seek(0, 0)

	return fileCache
}
