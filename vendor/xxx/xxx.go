package xxx

type a struct{ value string }

// A is package level variable
var A = a{"This is A from xxx"}

func init() {
	println("xxx.init() # vendored with main")
}
