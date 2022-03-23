package mypages

import "fmt"

type MyPackage struct {

}
func New()(obj *MyPackage)  {
	obj =&MyPackage{}
	return
}

func (m *MyPackage)Run()  {
	fmt.Println("MyPackage包  run 方法调用")
}