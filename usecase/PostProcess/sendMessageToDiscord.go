package postprocess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func createChangesMessage(changeContent ChangeContents) *string {
	var message string
	if !changeContent.isChanged() {
		message = "### 更新はありませんでした"
		return &message
	}
	del := changeContent.del
	add := changeContent.add
	update := changeContent.update
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
	return &message
}
