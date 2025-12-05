package utils

import (
	"fmt"
	"notion2atlas/constants"
)

func GetDownloadPath(pageId string, fileName string) string {
	return fmt.Sprintf("%s/assets/%s/%s", constants.DEPLOY_URL, pageId, fileName)
}
