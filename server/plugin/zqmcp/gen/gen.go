package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/zqmcp/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "zqmcp", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.McpServer), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.McpTool),
	)
	g.Execute()
}
