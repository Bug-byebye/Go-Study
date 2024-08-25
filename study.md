# Golang Study

## 基础语法

### 1.变量与常量

1. **变量：**
    * > var a int = 1
    * > var b = 2
    * > var a, b = 666, "123"
    * > c := "abc"
    * >var{
        a int
        b bool
            > }

    函数外的每个语句都必须以关键字开始（var、func 等），因此 := 结构不能在函数外使用。

2. **常量：**

    具体方式与变量类似
    > const a = 6

    当一次声明多个时，若省略了值则默认与其哪一个相等

    ```go.golang
    const {
        a = 1
        b
    }
    ```

    `iota`为一种特殊常量,可理解为`const`的索引块，`const`中每新增一行常量声明`iota`将会加一

### 2.数据类型

除了基本数据类型，Go语言中还有

* 指针类型（Pointer）
* 数组类型
* 结构化类型(struct)
* Channel 类型
* 函数类型
* 切片类型
* 接口类型（interface）
* Map 类型

`printf`中  
> `%v` 直接输出 *(对于字符，`%v`输出的为其对应的编码，`%c`才能输出其值)*
> `%d` 十进制输出
> `%b` 二进制输出
> `%o` 八进制输出
> `%x` 十六进制输出
> `%t` bool型输出
> `%T` 输出类型
> `%f` 输出`float`类型，`%.nf`保留`n`位小数输出

可通过`unsafe.Sizeof()`得知变量所占字节数

类型间的转换除了常规方式（`eg: a:=11 b=float(a)`）和`Sprintf`，还可以通过`strconv`等

#### 1. 整型

与其他语言基本一致
> Go语言中不允许将整形转化为布尔型，布尔型也无法参与数值运算和与其他类型进行转换

#### 2. 浮点型

> Go默认`float`类型为64位

`float`的最大值用`math.MaxFloat32`、`math.MaxFloat64`常量定义

#### 3. 字符串

> `\r` 回车符（返回行首）
> `\n` 换行符
> `\f` 换页符
> `\t` 制表符
> `\v` 垂直制表符
> `\'`、`\"` 单、双引号
> `\\` 反斜杠

>多行字符串定义：
> a := &#96;
> line1
> line2
> ...
> &#96;

**常用操作：**
> `len(str)` 求长度
> `+`/`fmt.Sprintf` 拼接字符串
> `strings.Split` 字符串分割
> `strings.contains` 判断是否包含
> `strings.HasPrefix`/`strings.HasSuffix` 前/后缀判断
> `strings.Index`/`strings.LastIndex` 子串出现位置
> `string.Join(a[]string,sep string)` join操作

字符串中字符修改可通过`byte`或`rune`
> `rune` 表示一个`Unicode`码位, 对于小写英文字母，字符`'a'`的`Unicode`值是97, `'b'`是98，以此类推，直到`'z'`的`Unicode`值为 122。
当进行字符相减时，实际上是在做字符减法运算：'c' - 'a' 的结果是 99 - 97 = 2。
（相关使用可见[leetcode-3144的灵神题解](https://leetcode.cn/problems/minimum-substring-partition-of-equal-character-frequency/solutions/2775377/hua-fen-xing-dpji-yi-hua-sou-suo-di-tui-s1nq0/)）

#### 4. 切片

类型`[]T`表示一个元素类型为`T`的切片。
切片通过两个下标来界定，二者以冒号分隔：`a[low : high]`（切片下界的默认值为 0，上界则是该切片的长度。），它会选出一个半闭半开区间，包括第一个元素，排除最后一个元素。

切片类似数组的引用，本身并不存储任何数据，它只是描述了底层数组中的一段。更改切片的元素会修改其底层数组中对应的元素。和它共享底层数组的切片都会观测到这些修改。
>当切片容量大于底层数组容量时，会自动创建一个新的底层数组，取消对原数组的引用

切片拥有`长度` 和`容量`。其中长度就是它所包含的元素个数，容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
> 切片s的长度和容量可通过表达式`len(s)`和`cap(s)`来获取。
> 切片的零值是`nil`。`nil`切片的长度和容量为0且没有底层数组。
切片在创建时，如果不指定容量则会自动计算，
创建时计算容量的公式为：`创建时的长度 * 2`。在追加时，会先判断容量够不够，如果容量足够则容量不变；如果超出容量那么判断长度是否小于 `1024` ，小于则 `容量 * 2` ，大于则 `容量 * 1.25`

切片可以用内置函数`make`来创建，这也是创建动态数组的方式。`make`函数会分配一个元素为零值的数组并返回一个引用了它的切片：
> `a := make([]int, 5)`    // len(a)=5

要指定它的容量，需向`make`传入第三个参数：

> `b := make([]int, 0, 5)`    // len(b)=0, cap(b)=5
    `b = b[:cap(b)]`  // len(b)=5, cap(b)=5
    `b = b[1:]`        // len(b)=4, cap(b)=4`

为切片追加新的元素可通过内置的`append`函数。
`append`的第一个参数是元素类型为`T`的切片，其余类型为`T`的值将会追加到该切片的末尾，其返回的结果是一个包含原切片所有元素加上新添加元素的切片。当接受追加的切片对应的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。 返回的切片会指向这个新分配的数组。

#### 5. 其他结构

> Go 没有指针运算
> 结构体：`type 结构体名 struct {}`，其字段可通过`.`来访问。使用`字段名 ： 值`语法可以仅列出部分字段（字段名的顺序无关）。前缀 & 返回一个指向结构体的指针
>

### 3.运算符

与其他语言基本一致
> `^`为异或运算、按位取反
> Go 的`*`、`/`、`%`、`<<`、`>>`、`&`、`^`都是最高优先级
> `++`和`--`在go中不是运算符，且不能前置使用
> 科学计数法表示：`var a float = 2.2e3`

对于精度丢失问题，可通过引入第三方包解决，例如`github/shopspring/decimal`
> `a3 := decimal.NewFormFloat(a1).ADD(decimal.NewFormFloat(a2))`

### 4.语句

> go中的`if`和`for`语句条件判断无`()`,且`{}`不能省略;

> `if`语句中可在判断语句前添加初始语句，`for`语句中初始语句与结束语句均可去掉,做`while`语句使用（也可省略所有循环条件，实现无限循环）

**range 遍历**：`for`循环的`range`形式可遍历切片或映射。eg：`for i, v := range pow {...}`
`range`循环会在开始时获取切片的副本，副本包括切片的起始地址和长度，即使在循环过程中向切片中添加了元素, `range`仍然会根据最初的长度进行遍历
>当使用`for`循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。可以用`_`来忽略对应返回值。
`for i, _ := range pow {...}`
`for _, value := range pow {...}`
若只需要索引，忽略第二个变量即可。
`for i := range po`

**map映射**将键映射到值,其零值为`nil`(`nil`映射既没有键，也不能添加键。)
> `make`函数会返回给定类型的映射，并将其初始化备用。eg：`m = make(map[string]Vertex)`
在映射 m 中插入或修改元素：`m[key] = elem`
获取元素：`elem = m[key]`
删除元素：`delete(m, key)`
通过双赋值检测某个键是否存在：
`elem, ok = m[key]`
若 key 在 m 中，ok 为 true ；否则，ok 为 false。
若 key 不在映射中，则 elem 是该映射元素类型的零值。
（若 elem 或 ok 还未声明，可使用短变量声明`elem, ok := m[key]`）

> 可通过设置`label`控制`break`、`continue`、`goto`对应跳转位置

> Go 的 switch 语句只会运行选定的 case，而非之后所有的 case。在效果上相当于为每个 case 自动添加 break，除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。

**defer**语句会将函数推迟到外层函数返回之后执行。
> 推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。



### 5.函数

函数可以返回任意数量的返回值。
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
返回值的命名应当能反应其含义，它可以作为文档使用。
没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值。

**函数闭包**：Go 函数可以是一个`闭包`。`闭包`是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量值，换句话说，该函数被“绑定”到了这些变量。

```go.golang

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

### 6.方法与接口

Go 没有类，不过可以为类型定义`方法`。
`方法`为一类带特殊的`接收者`参数的函数。方法接收者在其参数列表内，位于`func`和方法名之间。
eg：`func (v Vertex) Abs() float64 {return math.Sqrt(v.X*v.X + v.Y*v.Y)}`
**其中`接收者`的类型（包括`int`、`float`等）定义和方法声明必须在同一包内。**
此外，同样可以为`指针类型`的接收者声明方法。
接收者为值或指针的的方法被调用时，接收者均既能是值又能是指针。

### 7.包

### 8.其他

Go中变量、属性等的首字母大写为公有，小写为私有

## Web编程

### `JSON`

`JSON`*（JavaScript Object Notation）*是一种轻量级的数据交换格式，易于人类阅读和编写，同时也易于机器解析和生成。它最初是基于 JavaScript 的一种数据格式，但现在几乎所有的编程语言都支持 JSON，并且广泛应用于各种场景中，如 API 数据交换、配置文件、数据存储等

`JSON`由两种结构组成：

对象（Object）：
由`{}`包围, 以键值对的形式表示数据，键和值之间用冒号`:`分隔，键必须是字符串类型，值可以是任意合法的`JSON`数据类型, 多个键值对之间用逗号`,`分隔。

```
{
    "name": "Alice",
    "age": 30,
    "isStudent": false
}
```

数组（Array）：由 [] 包围, 以逗号`,`分隔的一组有序值。

```
[
    "apple",
    "banana",
    "cherry"
]
```

JSON 支持以下几种数据类型：
>字符串：由双引号 "" 包围, 支持`Unicode`字符，可以包含转义字符（如 \n, \t, \" 等）。
>数值：可以是整数或浮点数，不带引号。
>布尔值（`true`和`false`）：
>空值`null`
>对象、数组


```
type Greeting struct {
    Message string `json:"message"`
}
// `json:"message"`为结构体标签，
//指定在结构体进行JSON编码/解码时“Message”字段对应的JSON键为“message”
```

## 实践

**strings.Fields 是 strings 包中的一个函数，用于将字符串按空白字符（如空格、制表符、换行符等）分割成一个子字符串（单词）的切片。**
*练习：实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。 函数 wc.Test 会为此函数执行一系列测试用例，并输出成功还是失败。*

```go.golang
    package main

    import (
        "strings"
        "golang.org/x/tour/wc"
    )

    func WordCount(s string) map[string]int {
        // 创建一个映射来存储单词计数
        wordCount := make(map[string]int)
    
        // 使用 strings.Fields 将字符串分割成单词
        words := strings.Fields(s)
    
        // 遍历每个单词并更新计数
        for _, word := range words {
            wordCount[word]++
        }
    
        return wordCount
    }
    
    func main() {
        // 使用 wc.Test 进行测试
        wc.Test(WordCount)
    }
```
