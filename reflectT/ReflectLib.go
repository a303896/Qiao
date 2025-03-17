package reflectT

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Permission `json:"permission"`
}

type Permission struct {
	Read  bool `json:"readable"`
	Write bool `json:"writeable"`
}

type Person interface {
	Call()
}

func (u *User) Call() {
	fmt.Printf("user name is %s, age is %d\n", u.Name, u.Age)
}

func (u *User) Eat(food string) {
	fmt.Printf("%s eat %s\n", u.Name, food)
}

func test() {
	user := User{"Lihua", 10, Permission{true, false}}
	//user.Call()
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	fmt.Printf("v:%v  t:%v NumField:%d NumMethod:%d \n", v, t, t.NumField(), t.NumMethod())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("f:%v name:%s  type:%v kind:%v tag:%v\n", f, f.Name, f.Type, f.Type.Kind(), f.Tag.Get("json"))
	}
	t2 := reflect.TypeOf(&user)
	for i := 0; i < t2.NumMethod(); i++ {
		m := t2.Method(i)
		fmt.Printf("m:%v name:%s  type:%v\n", m, m.Name, m.Type)
		if m.Name == "Call" {
			m.Func.Call([]reflect.Value{reflect.ValueOf(&user)})
		}
		if m.Name == "Eat" {
			m.Func.Call([]reflect.Value{reflect.ValueOf(&user), reflect.ValueOf("Apple")})
		}
	}
}

func test2(p Person) error {
	res, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("json marshal fail")
	}
	fmt.Printf("json result:%s\n", res)
	var u User
	err = json.Unmarshal(res, &u)
	if err != nil {
		return fmt.Errorf("json unmarshal fail")
	}
	fmt.Printf("json unmarshal result:%v\n", u)
	return nil
}
