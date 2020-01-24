package main

import (
	"fmt"
	"os"
)

func conv(s string) int {
	cur := 0
	for _, c := range s {
		w := 0
		for i := '1'; i <= c; i++ {
			w++
		}
		cur = cur*10 + w
	}
	return cur

}
func main() {
	arg := os.Args[1:]
	ln := 0
	for i := range arg {
		ln = i + 1
	}
	for i := 2; i < ln; i++ {
		file, err := os.Open(arg[i])
		if err != nil {
			fmt.Println(err)
		} else {
			stat, _ := file.Stat()
			sz := stat.Size()
			a := []byte{}
			for j := 1; j <= int(sz); j++ {
				a = append(a, '0')
			}
			file.Read(a)
			sl := int(sz) - conv(arg[1])
			if sl < 0 {
				sl = 0
			}
			fmt.Printf("%v\n", string(a[sl:]))
			file.Close()
		}
	}
}
tai