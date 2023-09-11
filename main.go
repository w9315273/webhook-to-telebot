package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxRequestBodySize = 10 << 10
	defaultTextCount   = 20
)

var (
	botToken    = os.Getenv("BOT_TOKEN")
	telegramURL = "https://api.telegram.org/bot" + botToken + "/sendMessage"
	port        = os.Getenv("PORT")
	webhookPath = os.Getenv("WEBHOOK_PATH")
	chatID      = os.Getenv("CHAT_ID")
	authToken   = os.Getenv("AUTH_TOKEN")
	textCount   = getTextCount()
)

func getTextCount() int {
	countStr := os.Getenv("TEXT_COUNT")
	if count, err := strconv.Atoi(countStr); err == nil {
		return count
	}
	return defaultTextCount
}

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

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)

	token := r.Header.Get("Authorization")
	if token != authToken {
		http.Error(w, "·¢ËÍÊ§°Ü", http.StatusUnauthorized)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "·¢ËÍÊ§°Ü", http.StatusUnsupportedMediaType)
		return
	}

	data := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		if err.Error() == "http: request body too large" {
			http.Error(w, "ÇëÇóÌåÌ«´ó", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(w, "·¢ËÍÊ§°Ü", http.StatusInternalServerError)
		}
		return
	}

	textParts := []string{}
	for i := 1; i <= textCount; i++ {
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
		http.Error(w, "·¢ËÍÊ§°Ü", http.StatusInternalServerError)
		return
	}

	const maxRetries = 5
	var backoffFactor time.Duration = 2
	var resp *http.Response

	for i := 0; i < maxRetries; i++ {
		resp, err = http.Post(telegramURL, "application/json", bytes.NewBuffer(jsonData))
		if resp != nil && (resp.StatusCode == 429 || resp.StatusCode >= 500) {
			time.Sleep(backoffFactor * time.Second)
			backoffFactor *= 2
		} else {
			break
		}
	}

	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "·¢ËÍÊ§°Ü", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "·¢ËÍ³É¹¦")
}

func main() {
	if port == "" {
		port = "5000"
	}
	http.HandleFunc(webhookPath, webhookHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}