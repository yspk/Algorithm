package main

type F struct {

}

func (F) M1[T any](t T){}

func main() {
	var f F[string]
	f.M1("hello")
}
