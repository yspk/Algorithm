package main

type MyInterface interface {
	M1()
}

type GenericsInterface interface {
	~int | MyInterface | float64
}

func main() {

}
