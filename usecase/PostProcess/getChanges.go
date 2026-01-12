package postprocess

import (
	"fmt"
	"notion2atlas/constants"
	"notion2atlas/filemanager"
	"notion2atlas/usecase/fileUC"
)

type ChangeContents struct {
	del    []string
	add    []string
	update []string
}

func (c ChangeContents) isChanged() bool {
	isChanged := len(c.add) != 0 || len(c.del) != 0 || len(c.update) != 0
	return isChanged
}

func getChanges() (*ChangeContents, error) {
	changes, err := filemanager.ReadJson[[]fileUC.ChangeItem](constants.TMP_CHANGE_PATH)
	if err != nil {
		fmt.Println("error in usecase/PostProcess.sendMessageToDiscord.go: createChangesMessage/filemanager.ReadJson")
		return nil, err
	}
	var del []string
	var add []string
	var update []string
	for _, c := range changes {
		switch c.Type {
		case "add":
			t := "- " + c.Title
			add = append(add, t)
		case "delete":
			t := "- " + c.Title
			del = append(del, t)
		case "update":
			t := "- " + c.Title
			update = append(update, t)
		}
	}
	content := ChangeContents{
		del:    del,
		add:    add,
		update: update,
	}
	return &content, nil
}
