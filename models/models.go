package models

// 用户请求的结构体
type UserRequest struct {
	Prompt string `json:"prompt"` // 用户输入的文本
}

// DeepSeek API 请求的结构体
type DeepSeekRequest struct {
	Model    string    `json:"model"`    // 模型名称
	Messages []Message `json:"messages"` // 消息列表
	Stream   bool      `json:"stream"`   // 是否流式响应
}

// 消息结构体
type Message struct {
	Role    string `json:"role"`    // 角色（如 "user"）
	Content string `json:"content"` // 消息内容
}

// DeepSeek API 响应的结构体
type DeepSeekResponse struct {
	ID                string   `json:"id"`                 // 请求 ID
	Object            string   `json:"object"`             // 对象类型
	Created           int64    `json:"created"`            // 创建时间戳
	Model             string   `json:"model"`              // 模型名称
	Choices           []Choice `json:"choices"`            // 响应选择列表
	Usage             Usage    `json:"usage"`              // Token 使用情况
	SystemFingerprint string   `json:"system_fingerprint"` // 系统指纹
}

// 选择结构体
type Choice struct {
	Index        int     `json:"index"`         // 选择索引
	Message      Message `json:"message"`       // 消息内容
	Logprobs     *string `json:"logprobs"`      // Logprobs（可能为 null）
	FinishReason string  `json:"finish_reason"` // 结束原因
}

// Token 使用情况结构体
type Usage struct {
	PromptTokens          int `json:"prompt_tokens"`            // 输入的 Token 数量
	CompletionTokens      int `json:"completion_tokens"`        // 输出的 Token 数量
	TotalTokens           int `json:"total_tokens"`             // 总 Token 数量
	PromptCacheHitTokens  int `json:"prompt_cache_hit_tokens"`  // 缓存命中的 Token 数量
	PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"` // 缓存未命中的 Token 数量
}

// API 错误响应的结构体
type APIErrorResponse struct {
	Error string `json:"error"` // 错误信息
}
