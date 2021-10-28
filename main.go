package main

import (
	"addressGo/address"
	"log"
	"os"
)

/*
진행 순서: utf-8 변환 -> 주소, 부가정보, 도로명, 지번 머지 ->
*/
func main() {
	pwd, _ := os.Getwd() //현재 디렉토리
	dest := pwd + "\\result"

	err := os.RemoveAll(dest)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
		os.Exit(-1)
	}

	juso := address.FilePhase{
		PreFix:   "주소",
		Column:   address.Columns["주소"],
		FieldIdx: 0,
		Encoding: "utf-8",
	}
	jibun := address.FilePhase{
		PreFix:   "지번",
		Column:   address.Columns["지번"],
		FieldIdx: 0,
		Encoding: "CP949",
	}
	buga := address.FilePhase{
		PreFix:   "부가정보",
		Column:   address.Columns["부가정보"],
		FieldIdx: 0,
		Encoding: "CP949",
	}
	doro := address.FilePhase{
		PreFix:   "개선",
		Column:   address.Columns["개선"],
		FieldIdx: 0,
		Encoding: "CP949",
	}

	collector := &address.Collector{
		Src:        pwd + "\\example",
		Dest:       pwd + "\\result",
		FilePhases: []address.FilePhase{juso, jibun, buga, doro},
	}

	collector.Merge()

	numOfShards := 5

	source := dest
	jusoService := address.NewService(source, source+"\\chunk", juso, numOfShards)
	//jibunService := address.NewService(source, source+"\\chunk", jibun, numOfShards)
	//bugaService := address.NewService(source, source+"\\chunk", buga, numOfShards)
	//doroService := address.NewService(source, source+"\\chunk", doro, numOfShards)

	jusoService.Sort()
	//jibunService.Sort()
	//bugaService.Sort()
	//doroService.Sort()
}
