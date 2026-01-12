package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/filemanager"
	"os"
	"strconv"
	"strings"
)

func updateVersion(isVersionUp bool) error {
	if isVersionUp {
		return nil
	}
	bytes, err := filemanager.LoadFile(constants.VERSION_PATH)
	if err != nil {
		fmt.Println("error in usecase/postprocess/versionPath.go updateVersion/filemanager.LoadFile")
		return err
	}
	newVersion, err := bumpPatch(string(bytes))
	if err != nil {
		fmt.Println("error in usecase/postprocess/versionPath.go updateVersion/bumpPatch")
		return err
	}
	err = os.WriteFile(constants.VERSION_PATH, []byte(newVersion), 0644)
	if err != nil {
		fmt.Println("error in usecase/postprocess/versionPath.go updateVersion/os.WriteFile")
		return err
	}
	return nil
}

func bumpPatch(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}

	patch++
	parts[2] = strconv.Itoa(patch)

	return strings.Join(parts, "."), nil
}
