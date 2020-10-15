package sessions

import (
	"fmt"
	"testing"
)

func TestBegin(t *testing.T) {
	a := Begin()
	a.Create("xyz", "aghla", "1.1.1.1", 2)
	a.Create("xxx", "aghla", "1.1.1.1", 3)
	a.Create("yyy", "aghla", "1.1.1.1", 4)
	bt, _ := a.AllSessions()
	fmt.Println(string(bt))
	a.Delete("xyz")
	bt, _ = a.AllSessions()
	fmt.Println(string(bt))

}
