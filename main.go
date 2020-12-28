package main
import (
	"fmt"
	"github.com/markbates/pkger"
)
func main() {
	run()
}
func run() error {
	fi, err := pkger.Open("/data/hello.txt")
	if err != nil {
		return err
	}
	defer fi.Close()
	info, err = fi.Stat()
	if err != nil {
		return err
	}
	fmt.Println("Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("ModTime: ", info.ModTime())
	return nil
}
