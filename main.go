package main

import (
	"addressGo/address"
	"os"
)

func main() {
	pwd, _ := os.Getwd() //현재 디렉토리

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
}
