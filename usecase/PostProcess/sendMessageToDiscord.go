package postprocess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notion2atlas/constants"
	"notion2atlas/filemanager"
	"notion2atlas/usecase/fileUC"
	"os"
)

type DiscordWebhookPayload struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

func SendDiscordMessage(message string) error {
	baseWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	threadID := os.Getenv("DISCORD_THREAD_ID")
	webhookURL := fmt.Sprintf("%s?thread_id=%s", baseWebhookURL, threadID)

	payload := DiscordWebhookPayload{
		Content:  message,
		Username: "HorizonAtlas",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error in usecase/PostProcess.sendMessageToDiscord.go: sendDiscordMessage/json.Marshal")
		return err
	}

	resp, err := http.Post(
		webhookURL,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook failed: status %d", resp.StatusCode)
	}

	return nil
}

func createChangesMessage() (*string, error) {
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
			t := "- " + c.Title + "\n"
			add = append(add, t)
		case "delete":
			t := "- " + c.Title + "\n"
			del = append(del, t)
		case "update":
			t := "- " + c.Title + "\n"
			update = append(update, t)
		}
	}
	var message string
	if len(del) == 0 && len(add) == 0 && len(update) == 0 {
		message = "### 更新はありませんでした"
		return &message, nil
	}
	message += "# 更新内容\n"
	if len(del) != 0 {
		message += "### 削除：\n"
		for _, d := range del {
			message += d + "\n"
		}
	}
	if len(add) != 0 {
		message += "### 追加：\n"
		for _, a := range add {
			message += a + "\n"
		}
	}
	if len(update) != 0 {
		message += "### 更新：\n"
		for _, a := range update {
			message += a + "\n"
		}
	}
	return &message, nil
}
