package beetle

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	Register("uname", dealUname)
	fmt.Println(Parse("uname -v"))
}

func dealUname() string {
	fmt.Println(GetItemNum())
	for i := 0; i < GetItemNum(); i++ {
		fmt.Println(GetItem(i))
	}
	return "v1.0"
}

func TestCase2(t *testing.T) {
	Register("boardtype", dealBoardtype)
	fmt.Println("resp1", Parse("boardtype --help"))
	fmt.Println("resp2", Parse("boardtype -s PCH1"))
	fmt.Println("resp3", Parse("boardtype --set PCH2"))
}

func dealBoardtype() string {
	fmt.Println(GetItemNum())
	for i := 0; i < GetItemNum(); i++ {
		fmt.Println(GetItem(i))
	}

	if GetItemNum() == 2 {
		item := GetItem(1)
		if (item.Type == TypeLongOption && item.Value == "help") ||
			(item.Type == TypeShortOption && item.Value == "h") {
			return `Usage:
boardtype            print board type
boardtype -h         print the help
boardtype -s <type>  set board type
\n
Option:
-h,--help
-s,--set

Parameter:
type  board type is a string,e.g.:PCB1H2
`
		}
	}

	if GetItemNum() == 3 {
		item := GetItem(1)
		if (item.Type == TypeLongOption && item.Value == "set") ||
			(item.Type == TypeShortOption && item.Value == "s") {
			return fmt.Sprintln("set boardtype success", GetItem(2).Value, GetItem(2).Type)
		}
	}
	return ""
}
