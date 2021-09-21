package fssync
import (
	"fmt"
	"os"
)
func DirectoryExists(path string){
	fi, err := os.Stat(path)
	fmt.Println(fi, err)
	if fi, err := os.Stat(path); err == nil {
		fmt.Println(fi)
	}
}