#### Golang Review



##### array part

**如何创建一个三维数组（二维数组也是类似的）**

```go

	tiles := make([][][]int, n)

	for i := range tiles {
		tiles[i] = make([][]int, n)
		for j := range tiles[i] {
			tiles[i][j] = make([]int, n)
			for k := range tiles[i][j] {
				tiles[i][j][k] = math.MinInt8
			}
		}
	}
```



##### SLICE PART

**SLICE 置空**

```go
// my perfer
a := [1,2,3]
a = nil
fmt.Println(a, len(a), cap(a) // [] 0 0

// second way
a := [1,2,3]
a = a[:0]
fmt.Println(a, len(a), cap(a) // [] 0 3
fmt.Println(a[:1]) // [1]
```

**SCLICE 在函数中修改**

需要在函数中处理后返回，否则会如下； 原因golang值拷贝，"[]" 实际上是个slice header的数据结构，那么实际的地址发生改变后，原有引用的值不会发生变化

```go
package main

import (
	"fmt"
)

func main() {
	// var s []string
	s := make([]string, 10000)
	stringHelper(s)
	fmt.Printf("%s\n", s)
	fmt.Println("Hello, playground")
}

func stringHelper(s []string) {
	s = append(s, "hello World")
	s = append(s, []string{"a", "b"}...)
}

/*
[         ]
Hello, playground
*/
```



##### 字符串相关



>Go语言的字符有以下两种：
>
>- 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
>- 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

** 类似于C++ 的IsNumber**

```go
import "unicode"
import "fmt"
fmt.Println(unicode.IsDigit(r))
fmt.Println(unicode.IsNumber(r))
```

