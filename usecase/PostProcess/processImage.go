package postprocess

import (
	"errors"
	"fmt"
	"notion2atlas/domain"
	"notion2atlas/filemanager"
	"os"
	"strings"
)

func processImage(item domain.BlockEntity, pageEntities []domain.PageEntity) (*domain.AtlBlockEntity, error) {
	atlEntity := processNormalParent(item, pageEntities)
	url := atlEntity.Data.Image.Url
	if strings.HasPrefix(url, "https://raw.githubusercontent.com/Ryukoku-Horizon/notion2atlas/main/notion_data/assets") {
		image_path := strings.Replace(url, "https://raw.githubusercontent.com/Ryukoku-Horizon/notion2atlas/main/", "", 1)
		image_size, err := filemanager.MeasureLocalImageSize(image_path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("file is not Exists: %s\n", image_path)
				return nil, nil
			}
			fmt.Println("❌ error in postprocess/ProcessParent/filemanager.MeasureImageSize")
			return nil, err
		}
		atlEntity.Data.Image.Width = image_size.Width
		atlEntity.Data.Image.Height = image_size.Height
	} else {
		if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
			image_size, err := filemanager.MeasureImageSizeFromURL(url)
			if err != nil {
				switch {
				case errors.Is(err, filemanager.ErrRequestFailed):
					return nil, nil
				case errors.Is(err, filemanager.ErrHTTPStatus):
					return nil, nil
				case errors.Is(err, filemanager.ErrNotImage):
					return nil, nil
				case errors.Is(err, filemanager.ErrDecodeImage):
					return nil, nil
				default:
					// 想定外
					fmt.Println("❌ error in postprocess/ProcessParent/filemanager.MeasureImageSizeFromURL")
					return nil, err
				}
			}
			atlEntity.Data.Image.Width = image_size.Width
			atlEntity.Data.Image.Height = image_size.Height
		}
	}
	if atlEntity.Data.Image.Width == 0 || atlEntity.Data.Image.Height == 0 {
		return nil, nil
	}
	return &atlEntity, nil
}
