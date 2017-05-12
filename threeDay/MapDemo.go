package main

func main() {
	//m:=map[int]struct{
	//	name string
	//	age int
	//}{
	//	1:{"user1",10},
	//	2:{"user2",20},
	//}
	//println(m[1].age)
	test()

}
func  test()  {
	m:=map[string]int{"a":1,}
	if v,ok:=m["a"];ok{
		println(v,ok)
	}
	println(m["c"])

	m["b"]=2

	delete(m,"c")

	println(len(m))

	for k,v:=range m{
		println(k,v)
	}
}

