package themes

import (
	"strconv"

	"github.com/k0kubun/pp"
)

func Convert() {
	num := 256
	str := strconv.Itoa(num)
	pp.Println(num, str)

	format, err := strconv.Atoi("142")
	if err == nil {
		pp.Println(format, err)
	}
}
