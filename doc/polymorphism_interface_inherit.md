> 问题： golang如何实现多态

通过**影子类**来实现多态, 父类接口不方便写的放到影子类中
解决复用痛点
（1）继承类不想重复实现接口中通用的函数
（2）继承类中的相同field希望写到父类中

## 1. 父类写成接口

```
type Parent interface{ 
    f1()
    f2()
    f3()
}
```
## 2. 影子类实现通用接口
（1）父类中成员字段可以写在这里，因为接口中不能有字段
（2）多态时要调用的共用函数写在这里，也就是Parent中的中定义的某些函数，只许要实现一些公共的需要的类
（3）父类中函数的默认实现写在这里

```
type ParentBack struct{
    Field1 string
    Field2 int
}
func (b*ParentBack)f1(){
    // 通用实现，不想在子类中实现的，又需要在多态时调用的函数
}

func (b*ParentBack) f4(){
    // 通用函数
}
func (b*ParentBack) f5(){
    // 通用函数
}
```

## 3.子类实现Parent，继承Parent_Back


```
type ChildA struct{
    ParentBack
	field3
}
// 只需要实现f2、f3 这些特有的，不需要实现f1, 因为可以通过 ParentBack 继承过来f1

func (b*ChildA) f2(){
	
}
func (b*ChildA) f3(){
	
}

type ChildB struct{
ParentBack
field3
}
// 只需要实现f2、f3 这些特有的，不需要实现f1, 因为可以通过 ParentBack 继承过来f1

func (b*ChildB) f2(){

}
func (b*ChildB) f3(){

}
```

## 4. 使用

```
// 加入集合
list []*Parent
a:= new(ChildA)
list = append(list, a )
b:= new(ChildB)
list = append(list, b)

// 多态使用
for _,item:=range list{
	item.f1()
	item.f2()
	item.f3()
}
```

## 5. struct多层继承问题
> 由于go中的struct继承其实是组合实现，因此不能实现类型转化，无法实现struct类型继承的多态

```
type GrandsonA struct{
    ChildA
	field4
}

```

```
var list []*ChildA
grandson:= new(GrandsonA)
// 会出错
list=append(list,grandson)
```

解决办法 使用接口来完成多态，可以定义一个IChildA来实现这一分支的多态