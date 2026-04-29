package main

import (
	"fmt"
	"notion2atlas/presentation"
	postprocess "notion2atlas/usecase/PostProcess"
)

// func test() {
// 	initapp.InitApp()
// 	a, e := filemanager.LoadAndDecodeJson[[]domain.AtlBlockEntity](constants.PAGE_DATA_DIR + "/2f5a501ef33780069c27c8c9c3d5c599.dat")
// 	if e != nil {
// 		panic(e)
// 	}
// 	filemanager.WriteJson(a, "notion_data/test.json")
// }

// func decode[T any](path string) {
// 	godotenv.Load()
// 	ex := filemanager.FileExists(path)
// 	fmt.Println(ex)
// 	a, e := filemanager.LoadAndDecodeJson[T](path)
// 	if e != nil {
// 		panic(e)
// 	}
// 	filemanager.WriteJson(a, "notion_data/test.json")
// }

// func encode[T any](path string) {
// 	b, _ := filemanager.ReadJson[T](constants.NT_DATA_DIR + "/test.json")
// 	filemanager.EncodeAndSave(b, path)
// }

func main() {
	// test()
	// e := presentation.HandleUpdateData()
	// if e != nil {
	// 	panic(e)
	// }
	err := presentation.HandleUpdateData()
	if err != nil {
		err2 := postprocess.SendDiscordMessage("**❌ エラーが発生したため、ページの更新に失敗しました**")
		if err2 != nil {
			fmt.Printf("discord error: %s\n", err2)
		}
		panic(err)
	}
	// initapp.InitApp()
	// a, _ := notionUC.Test("331a501ef33780508cd9f67021f715ab")
	// filemanager.WriteJson(a, constants.TEST_JSON)
	// test()
	// decode[[]domain.InfoEntity]("notion_data/infos/info.dat")
	// encode[[]domain.BlockEntity]("notion_data/synced.dat")
	// p, _ := usecase.Test("24ba501ef33780edacc4d54914fb20d2")
	// filemanager.WriteJson(p, "notion_data/test.json")
}

//https://www.notion.so/331a501ef33780508cd9f67021f715ab?v=1ada501ef33781c291ee000c04c13f15&source=copy_link
