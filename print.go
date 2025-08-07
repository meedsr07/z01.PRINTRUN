package z01

import("github.com/01-edu/z01")

func Printrstr(str string) {

	for _,v:=range str{
		z01.PrintRune(v)
	}
	

}