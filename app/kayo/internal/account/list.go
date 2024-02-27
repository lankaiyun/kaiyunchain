package account

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/lankaiyun/kaiyunchain/db"
)

func ListAccount() {
	files, _ := filepath.Glob(db.KeystoreDataPath + "/*")
	if len(files) == 0 {
		color.Red("You don't have an account yet!")
		fmt.Println()
		return
	}
	for i := 0; i < len(files); i++ {
		fmt.Printf("Address%d: %s\n", i+1, files[i][25:])
	}
	fmt.Println()
}
