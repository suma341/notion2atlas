package main

import (
	"fmt"
	"notion2atlas/presentation"
	postprocess "notion2atlas/usecase/PostProcess"
)

// func test() {
// 	a, e := filemanager.LoadAndDecodeJson[[]domain.AtlPageEntity](constants.PAGE_DAT_PATH)
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
	err := presentation.HandleUpdateData()
	if err != nil {
		err2 := postprocess.SendDiscordMessage("**❌ エラーが発生したため、ページの更新に失敗しました**")
		if err2 != nil {
			fmt.Printf("discord error: %s\n", err2)
		}
		panic(err)
	}
	// test()
	// decode[[]domain.InfoEntity]("notion_data/infos/info.dat")
	// encode[[]domain.BlockEntity]("notion_data/synced.dat")
	// p, _ := usecase.Test("24ba501ef33780edacc4d54914fb20d2")
	// filemanager.WriteJson(p, "notion_data/test.json")
}

//https://www.notion.so/1-24ba501ef33780edacc4d54914fb20d2?source=copy_link
//https://www.notion.so/Python-256a501ef337802e8fcaf378c366fb03?v=1ada501ef33781c291ee000c04c13f15&source=copy_link
