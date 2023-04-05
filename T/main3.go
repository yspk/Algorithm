package main

type C interface {
	T2 | T3
}

func GenericsFoo[P C] (p P) {
	_ = p.Name
}

type T2 struct {
	Name string
}

type T3 struct {
	Name string
}

func main()  {
	GenericsFoo(T{})
}
