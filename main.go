package main

import (
	"notion2atlas/constants"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"notion2atlas/usecase/initapp"
)

func test() {
	initapp.InitApp()
	a, e := filemanager.LoadAndDecodeJson[[]domain.AtlBlockEntity](constants.PAGE_DATA_DIR + "/2f5a501ef33780069c27c8c9c3d5c599.dat")
	if e != nil {
		panic(e)
	}
	filemanager.WriteJson(a, "notion_data/test.json")
}

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
	test()
	// e := presentation.HandleUpdateData()
	// if e != nil {
	// 	panic(e)
	// }
	// err := presentation.HandleUpdateData()
	// if err != nil {
	// 	err2 := postprocess.SendDiscordMessage("**❌ エラーが発生したため、ページの更新に失敗しました**")
	// 	if err2 != nil {
	// 		fmt.Printf("discord error: %s\n", err2)
	// 	}
	// 	panic(err)
	// }
	// test()
	// decode[[]domain.InfoEntity]("notion_data/infos/info.dat")
	// encode[[]domain.BlockEntity]("notion_data/synced.dat")
	// p, _ := usecase.Test("24ba501ef33780edacc4d54914fb20d2")
	// filemanager.WriteJson(p, "notion_data/test.json")
}

//https://www.notion.so/2f5a501ef33780069c27c8c9c3d5c599?v=1ada501ef33781c291ee000c04c13f15&source=copy_link#2f5a501ef33780609d4ae06e1a207595
//https://www.notion.so/1-24ba501ef33780edacc4d54914fb20d2?source=copy_link
//https://www.notion.so/Python-256a501ef337802e8fcaf378c366fb03?v=1ada501ef33781c291ee000c04c13f15&source=copy_link
