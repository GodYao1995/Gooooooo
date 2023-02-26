package main

import (
	_ "github.com/cd-home/Goooooo/api/admin"
	"github.com/cd-home/Goooooo/cmd/admin/cmd"
)

// InitRouter @title Goooooo-Admin
// @contact.name God Yao
// @contact.email liyaoo1995@163.com
// @version 1.0.0
// @description this is Goooooo-Admin Sys.
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	cmd.Exeute()
}
