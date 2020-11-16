## fmt包的学习

### 概述
fmt包使用与C语言的printf和scanf类似的功能实现格式化的I/O。格式"占位符(verbs)"派生自C语言，但是更加的简单

以下例子中用到的类型或变量定义:
```
type Website struct {
    Name string
}

// 定义结构体变量
var site = Website{Name:"studygolang"}
```

### Printing
占位符种类如下:

常用类型：
- %v    在打印结构体使用默认的格式来打印出对应的值，加号(%+v)将会额外打印出字段名称
- %#v   用Go语法来表示该值
- %T    用Go语言来表示该值的类型
- %%    就是字面意思上的百分号，不能代表任何其他意思
```
Printf("%v", site)  // {studygolang}
Printf("%+v", site)  // {Name:studygolang}
Printf("#v", site)  // main.Website{Name:"studygolang"}
Printf("%T", site)  // main.Website
Printf("%%")  // %
```

布尔类型：
- %t    输出单词true或者false
```
Printf("%t", true)  // true
```

整数类型
- %b    二进制表示 
- %c    相应的Unicode字符集中的码点表示的字符(如65代表字符'A')
- %d    十进制表示
- %o    八进制表示
- %O    带0o前缀的八进制表示
- %q    使用Go语法安全地转义的单引号字符文字
- %x    使用小写字符(a-f)的十六进制表示
- %X    使用大写字符(A-F)的十六进制表示
- %U    Unicode格式：U+1234
```
Printf("%b", 5)  // 101
Printf("%c", 0x4E2D)  // 中
Printf("%d", 0x12)  // 18
Printf("%o", 10)  // 12
Printf("%q", 0x4E2D)  // '中'
Printf("%x", 13)  // d
Printf("%X", 13)  // D
Printf("%U", 0x4E2D)  // U+4E2D  
```

浮点数和复数类型
- %b    无小数部分，指数为二的幂的科学计数法，与strconv.FormatFloat的'b'转换格式一致。例如：5742089524897382p-49
- %e    科学计数法(使用小写的e)。e.g. -1.234456e+78
- %E    科学计数法(使用大写的E)。e.g. -1.234456E+78
- %f    有小数点而无指数，例如：123.456
- %F    %f的同义词
- %g    大的指数用%e，反之用%f
- %G    大的指数用%E，反之用%F
    
