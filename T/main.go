package main

type F[T any, P any] struct {
	Name string
	*T
	P
}

type MyInterface interface {

}

type GenericsInterface[I MyInterface] interface {
	M1()
	I
}

func main() {
	var f F[string, string]
	_ = f
}

