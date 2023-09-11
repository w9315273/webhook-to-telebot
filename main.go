package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	botToken    = os.Getenv("BOT_TOKEN")
	telegramURL = "https://api.telegram.org/bot" + botToken + "/sendMessage"
	webhookPath = os.Getenv("WEBHOOK_PATH")
	chatID      = os.Getenv("CHAT_ID")
	authToken   = os.Getenv("AUTH_TOKEN")
	port        = os.Getenv("PORT")
)

func getRealIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	ipAddress := getRealIP(r)
	log.Println("Received request from IP address:", ipAddress)

	token := r.Header.Get("Authorization")
	if token != authToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	data := map[string]string{}
	json.NewDecoder(r.Body).Decode(&data)
	textParts := []string{}
	for i := 1; i <= 20; i++ {
		if val, exists := data[fmt.Sprintf("text%d", i)]; exists {
			textParts = append(textParts, val)
		}
	}
	message := strings.Join(textParts, "\n")

	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(telegramURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending request to Telegram:", err)
		http.Error(w, "Failed to send message to Telegram", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("Telegram API responded with status: %d and message: %s\n", resp.StatusCode, bodyString)
		http.Error(w, "Telegram API rejected the message.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "OK")
}

func main() {
	if port == "" {
		port = "5000"
	}
	http.HandleFunc(webhookPath, webhookHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}