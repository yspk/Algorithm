package main

import (
	"container/list"
	"fmt"
)

type ArgList struct {
	*list.List
	Sum int64
}

func (ar *ArgList)ListInit() {
	ar.Init()
	ar.Sum = 0
}

func (ar *ArgList)ListPushBack(val int64) {
	ar.PushBack(val)
	ar.Sum += val
}

func (ar *ArgList)ListRemove()  {
	if trash := ar.Front(); trash != nil {
		ar.Sum -= trash.Value.(int64)
		ar.Remove(trash)
	}
}

func main()  {
	var i int64
	argList := &ArgList{list.New(),0}
	for ;i<100;i++ {
		argList.PushBack(i)
	}
	fmt.Println(argList.Sum/int64(argList.Len()-1))
}
