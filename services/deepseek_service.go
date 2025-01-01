package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ai-s/config"
	"ai-s/models"
)

func CallDeepSeekAPI(prompt string, cfg *config.Config) (*models.DeepSeekResponse, error) {
	// 构建请求体
	requestBody, err := json.Marshal(models.DeepSeekRequest{
		Model: "deepseek-chat", // 模型名称
		Messages: []models.Message{
			{
				Role:    "user", // 用户角色
				Content: prompt, // 用户输入
			},
		},
		Stream: false, // 非流式响应
	})
	if err != nil {
		return nil, fmt.Errorf("error encoding request: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", cfg.DeepSeekAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// 打印响应内容
	fmt.Println("API Response:", string(body))

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		var apiError models.APIErrorResponse
		if err := json.Unmarshal(body, &apiError); err != nil {
			return nil, fmt.Errorf("API returned error: %s", string(body))
		}
		return nil, fmt.Errorf("API error: %s", apiError.Error)
	}

	// 解析响应
	var deepseekResp models.DeepSeekResponse
	if err := json.Unmarshal(body, &deepseekResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &deepseekResp, nil
}
