package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJson(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("error in filemanager/WriteJson/os.Create")
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error in filemanager/WriteJson/encoder.Encode")
		return err
	}
	return nil
}

func ReadJson[T any](path string) (T, error) {
	var zero T

	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return zero, fmt.Errorf("read error: %w", err)
	}

	if len(jsonBytes) == 0 {
		return zero, fmt.Errorf("file is empty")
	}

	var obj T
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		return zero, fmt.Errorf("json unmarshal error: %w", err)
	}

	return obj, nil
}

func ReadJSONToMapArray(path string) ([]map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result []map[string]any

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("error in filemanager/CreateFile/os.Create path: " + path)
		return err
	}
	defer file.Close()
	return nil
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("error in filemanager/CreateDir/os.MkdirAll")
		return err
	}
	return nil
}

func ClearDir(path string) error {
	exist := DirExists(path)
	if !exist {
		fmt.Println("ClearDir: dir not exists")
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error in filemanager/ClearDir/os.ReadDir")
		return err
	}
	for _, e := range entries {
		err := os.RemoveAll(path + "/" + e.Name()) // ファイル・フォルダ両対応
		if err != nil {
			fmt.Println("error in filemanager/ClearDir/RemoveAll")
			return err
		}
	}
	return nil
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CreateFileIfNotExist(path string) (exist bool, err error) {
	var is_exists = FileExists(path)
	if !is_exists {
		err := CreateFile(path)
		if err != nil {
			fmt.Println("error in filemanager/CreateFileIfNotExist/CreateFile")
			return false, err
		}
	}
	return is_exists, nil
}

func CreateDirIfNotExist(path string) (exist bool, err error) {
	var is_exists = DirExists(path)
	if !is_exists {
		err := CreateDir(path)
		if err != nil {
			fmt.Println("error in filemanager/CreateDirIfNotExist/CreateDir")
			return false, err
		}
	}
	return is_exists, nil
}

func DelFile(path string) error {
	var is_exists = FileExists(path)
	if is_exists {
		err := os.Remove(path)
		if err != nil {
			fmt.Println("error in filemanager/DelFile/os.Remove")
			return err
		}
	} else {
		fmt.Println(path + " is not exists")
	}
	return nil
}

func DelDir(path string) error {
	var is_exists = DirExists(path)
	if is_exists {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("error in filemanager/DelDir/os.RemoveAll")
			return err
		}
	} else {
		fmt.Println(path + " is not exists")
	}
	return nil
}
