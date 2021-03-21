package napodate

import (
	"context"
	"encoding/json"
	"net/http"
)

// 在文件的第一部分中，我们将请求和响应映射到它们的JSON payload
type getRequest struct {
}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err, omitempty"`
}

type statusRequest struct {
}

type statusResponse struct {
	Status string `json:"status"`
}

// 在第二部分中，我们将为传入的请求编写「解码器」 decoders
func decodeGetRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

func decodeValidateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// 最后，我们有响应输出的编码器
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
