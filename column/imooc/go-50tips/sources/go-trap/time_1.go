package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("%Y-%m-%d %H:%M:%S"))
	fmt.Println(time.Now().Format("2006年01月02日 15时04分05秒"))
}
