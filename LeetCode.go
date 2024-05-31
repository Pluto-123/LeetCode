package main

type User struct {
	name string `json:name-field`
	age  int
}

func main() {
	//user := &User{
	//	name: "fak er",
	//	age:  27,
	//}
	//field, ok := reflect.TypeOf(user).Elem().FieldByName("name")
	//if !ok {
	//	panic("")
	//}
}
