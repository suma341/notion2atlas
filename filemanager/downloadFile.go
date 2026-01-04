package filemanager

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func DownloadFile(fileURL string, dir string, filename string, spareExt string) (string, error) {
	// URL パース
	parsed, err := url.Parse(fileURL)
	if err != nil {
		fmt.Println("error in filemanager/downloadFile.go: DownloadFile/url.Parse\nfileURL:" + fileURL)
		return "", err
	}

	// path.Ext はクエリを含まない URL.Path で取得
	ext := path.Ext(parsed.Path)
	if ext == "" {
		ext = spareExt
	}

	// 保存先パス
	fullPath := filepath.Join(dir, filename+ext)

	// ディレクトリ作成
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("error in filenamager/downloadFile.go: DownloadFile/os.MkdirAll\nfileURL:" + fileURL)
			return "", err
		}
	}

	// ファイル作成
	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Println("error in filemanager/downloadFile.go: DownloadFile/http.Get\nfileURL:" + fileURL)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download: status code %d\nfileURL:%s", resp.StatusCode, fileURL)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("error in filemanager/downloadFile.go: DownloadFile/os.Create\nfileURL:" + fileURL)
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("error in filemanager/downloadFile.go: DownloadFile/io.Copy\nfileURL:" + fileURL)
		return "", err
	}

	return filename + ext, nil
}

func SavePNG(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
