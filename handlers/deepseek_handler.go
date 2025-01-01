package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ai-s/config"
	"ai-s/models"
	"ai-s/services"
)

func DeepSeekHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	// 解析用户请求
	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding user request: %v", err), http.StatusBadRequest)
		return
	}

	// 调用 DeepSeek 服务
	deepseekResp, err := services.CallDeepSeekAPI(userReq.Prompt, cfg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error calling DeepSeek API: %v", err), http.StatusInternalServerError)
		return
	}

	// 返回响应给用户
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deepseekResp.Choices[0].Message.Content)
}
