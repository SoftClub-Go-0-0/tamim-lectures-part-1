package lecture8

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var file_name string
	var a, d int
	fmt.Scan(&file_name)
	fmt.Scan(&a, &d)
	file, _ := os.Create(file_name)
	for i := 0; i < 10; i++ {
		s32 := strconv.FormatFloat()
	}
}
