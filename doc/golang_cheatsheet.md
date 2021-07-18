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

**LowerBound**

一直以为golang里没有特别好的数据结构和算法库，比如sort，lowerbound，但是今天惊奇的发现，竟然有。但是由于没有模板的原因，目前没有看到非int的对应算法能力

```golang
func sliceLowerBound(arr []int, target int) int {
	rec := append(sort.IntSlice(nil), arr...)
	fmt.Printf("before sort %v\n", rec)
	rec.Sort()
	fmt.Printf("after sort %v\n", rec)
	res := rec.Search(target)  // 等同于lowerbound                                                                                                                                                                                                 
	return res
}
```



**排序相关**

在c++里有经典的排序算法，golang里也逐步熟悉了它的排序，也基本上是lambda表达式的对比.  下面先看一下C++, 然后看一下golang的

```c++
int main()
{
    std::array<int, 10> s = {5, 7, 4, 2, 8, 6, 1, 9, 0, 3};
 
    auto print = [&s](std::string_view const rem) {
        for (auto a : s) {
            std::cout << a << ' ';
        }
        std::cout << ": " << rem << '\n';
    };
 
    std::sort(s.begin(), s.end());
    print("sorted with the default operator<");
 
    std::sort(s.begin(), s.end(), std::greater<int>());
    print("sorted with the standard library compare function object");
 
    struct {
        bool operator()(int a, int b) const { return a < b; }
    } customLess;
    std::sort(s.begin(), s.end(), customLess);
    print("sorted with a custom function object");
 
    std::sort(s.begin(), s.end(), [](int a, int b) {
        return a > b;
    });
    print("sorted with a lambda expression");
}
```

```go
	type pair struct {
		atom string
		num  int
	}
	pairs := make([]pair, 0, len(c))
	for k, v := range c {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].atom < pairs[j].atom })
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

r := rune(s[i])
fmt.Println(unicode.IsDigit(r))
fmt.Println(unicode.IsNumber(r))
```

```c++
#include <cctype>

std::islower();
std::isupper();
std::isnumber();
std::isalpha();
std::isalnum();
```

