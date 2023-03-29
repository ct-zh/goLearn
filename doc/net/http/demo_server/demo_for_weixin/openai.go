package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenAiReq struct {
	Content string `json:"content"`
	Key     string `json:"key"`
}

func AskForOpenAI(ctx context.Context, user, text string) (string, error) {
	openaiReq := &OpenAiReq{
		Content: text,
		Key:     cfg.RequestKey,
	}
	jsonStr, err := json.Marshal(openaiReq)

	req, err := http.NewRequest("POST", cfg.RequestIP, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("http.NewRequest err=%+v", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err=%+v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
