package main

type C interface {
	T | T1
}

func GenericsFoo[P C] (p P) {
	p.M1()
}

type T struct {}

func (T) M1()  {
}

type T1 struct {
}
func (T1) M1()  {
}

func main()  {
	GenericsFoo(T{})
}
