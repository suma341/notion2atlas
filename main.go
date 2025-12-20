package main

import (
	"notion2atlas/presentation"

	"github.com/joho/godotenv"
)

func main() {
	var err error = nil
	godotenv.Load()
	err = presentation.HandleUpdateData()
	// aaa, err := usecase.Test("1ada501ef3378189ac56cfb744d1fd1b")
	// filemanager.WriteJson(aaa, "public/test.json")
	if err != nil {
		panic(err)
	}
}

//https://www.notion.so/Notion-1ada501ef33781428531e12a83e0195b?source=copy_link#1ada501ef3378191a728ecd1d680880c
//https://www.notion.so/1ada501ef337817c8daec3bd8ad9a5c3?source=copy_link
//https://www.notion.so/1ada501ef33780e7afabc3917c00a689?v=1ada501ef33781c291ee000c04c13f15&source=copy_link
//https://www.notion.so/Python-1ada501ef337814bac7dcbd864901138?v=1ada501ef33781c291ee000c04c13f15&source=copy_link
//https://www.notion.so/Web-1ada501ef3378162a825c0f0afcaf785?source=copy_link#2cfa501ef3378031a75dea58310fde01
//https://www.notion.so/1ada501ef3378182a8add5435b8cf059?source=copy_link#1ada501ef3378189ac56cfb744d1fd1b
