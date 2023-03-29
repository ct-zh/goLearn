package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Req struct {
	Content string `json:"content"`
	Key     string `json:"key"`
}

var (
	cfg    *Config
	openAi *OpenAi
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("err= %+v \n", err)
				w.Write([]byte("error"))
				return
			}
			req := &Req{}
			log.Printf("body=%s", body)
			err = json.Unmarshal(body, &req)
			if err != nil {
				log.Printf("err= %+v \n", err)
				w.Write([]byte("error"))
				return
			}
			log.Printf("Req = %+v", req)
			if req.Key != cfg.RequestKey {
				w.Write([]byte("error"))
				return
			}
			if req.Content == "" {
				w.Write([]byte("error"))
				return
			}
			content, err := openAi.AskForOpenAI(context.Background(), "", req.Content)
			if err != nil {
				log.Printf("err= %+v \n", err)
				w.Write([]byte("error"))
				return
			}
			w.Write([]byte(content))
			return
		}
	})

	cfg = NewConfig()
	openAi = NewOpenAi()
	log.Fatal(http.ListenAndServe(":12333", nil))
}
