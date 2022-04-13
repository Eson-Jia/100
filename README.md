# 100

## 1. 交替打印数字和字⺟

### 问题描述

使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果 如下：

```bash
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
```

### 错误方式，使用一个阻塞式 chan 同步

[参考源码](./main.go)
使用一个阻塞式 chan 的问题是，两个 goroutine 往同一个 os.Stdout 写入数据的时机是并发的。 这两个 goroutine 在同一个通道一读一写之后，会同时进行下一步。解决方法是使用两个通道，nums 写完后要等
alphabet 写完才能再接着写。

## 2. 判断字符串中字符是否全都不同

### 问题描述

请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都不同， false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。 字符串的⻓度⼩于等于【3000】。

## 4. 判断两个给定的字符串排序后是否⼀致

### 问题描述

给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string s1和⼀个string s2，请返回⼀个bool，
代表两串是否重新排列后可相同。 保证两串的⻓度都⼩于等于5000

## 7. 下⾯代码能运⾏吗？为什么

```go
type Param map[string]interface{}
type Show struct {
Param
}
func main1() {
s := new(Show)
s.Param["RMB"] = 10000
}
```

## 9.写出打印结果

```go

```

## 11. 请找出下⾯代码的问题所在

```go
func main() {
ch := make(chan int, 1000)
go func () {
for i := 0; i < 10; i++ {
ch <- i
}
}()
go func () {
for {
a, ok := <-ch
if !ok {
fmt.Println("close")
return
}
fmt.Println("a: ", a)
}
}()
close(ch)
fmt.Println("ok")
time.Sleep(time.Second * 100)
}
```

在 main.goroutine 关闭之后，还有可能写入引发 panic

## 12. 请说明下⾯代码书写是否正确

### 13.下⾯的程序运⾏后为什么会爆异常

### 15.请说出下⾯代码，执⾏时为什么会报错

```go
package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
}
```

#### 解析：

map 的 value 本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在 map 改变时会变得无效。 故如果需要修改map值，可以将map中的非指针类型 value,改为指针类型，比如使用 map[string]*
Student

16. 请说出下⾯的代码存在什么问题？

```go
package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}
func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
```

### 解析

依据4个goroutine的启动后执⾏效率，很可能打印111func4，但其他的111func*也可能先执⾏， exec只会返回⼀ 条信息。

## 19.以下代码有什么问题，说明原因

```go
package main

import "fmt"

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Print(&stu)
		m[stu.Name] = &stu
	}
}
```

### 解析

`for ... range`会复用变量 stu,每次循环迭代都会将 slice 中的值拷贝的 stu 中。但是需要注意，stu 变量的地址一直保持固定不变。 所以每次对 stu 取地址得到的都是同一个地址。而 stu
在循环结束后的值定格在 stus 最后一个元素 wang,所以 map 中所有的键都指向 wang.



















