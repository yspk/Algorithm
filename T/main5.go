package main

func GenericsFoo[M any](s M) M {
	type bar int
	var a bar
	println(a)
	return s
}

func main()  {
	GenericsFoo("string")
}
