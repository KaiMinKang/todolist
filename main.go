package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		slog.Error("加载配置文件失败", "err", err)
		os.Exit(1)
	}
	fmt.Printf("解析结果：Addr=%q, DSN=%q ", cfg.Server.Addr, cfg.MySQL.DSN)
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	slog.Info("服务靠配置文件启动，不是写死的", "cfg.Server.Addr", cfg.Server.Addr)
	r.Run(cfg.Server.Addr)
}
