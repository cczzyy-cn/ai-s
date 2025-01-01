package main

import (
	"fmt"
	"net/http"

	"ai-s/config"
	"ai-s/handlers"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// 注册路由
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeepSeekHandler(w, r, cfg)
	})

	// 启动 HTTP 服务
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
