package main

import (
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/presentation"
)

func test() {
	a, e := filemanager.LoadAndDecodeJson[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH)
	if e != nil {
		panic(e)
	}
	filemanager.WriteJson(a, "notion_data/test.json")
}

func decode[T any](path string) {
	a, _ := filemanager.LoadAndDecodeJson[T](path)
	filemanager.WriteJson(a, "notion_data/test.json")
}

func encode[T any](path string) {
	b, _ := filemanager.ReadJson[T](constants.ANSWER_DIR + "/data.json")
	filemanager.EncodeAndSave(b, path)
}

func main() {
	err := presentation.HandleUpdateData()
	if err != nil {
		panic(err)
	}
	// test()
	// decode[[]domain.AnswerEntity]("notion_data/answer.dat")
	// encode[[]domain.AnswerEntity]("notion_data/answer.dat")
	// p, _ := usecase.Test("24ba501ef33780edacc4d54914fb20d2")
	// filemanager.WriteJson(p, "notion_data/test.json")
}

//https://www.notion.so/1-24ba501ef33780edacc4d54914fb20d2?source=copy_link
