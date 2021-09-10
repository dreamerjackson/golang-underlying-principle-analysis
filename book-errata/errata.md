# 书籍勘误表
## 前言添加版本号
本书由21章组成 ->  本书基于Go1.14版本编写，由21章组成...
## 第1章
* (1.1) 第1页第3段第1行： 通过了解Go语言编辑器 -> 通过了解Go语言编译器
* (1.4) 第5页图1-4： ConstDeal -> ConstDecl 
## 第2章
* (2.3.2) 灰色第一个代码： \\n\\n -> \n
* (2.10) 40页正文第二行：56.78可以表示为5678×10^2 ->56.78可以表示为5678×10^-2,并且^应该为上标。

## 第7章
* (7.1.1) 第74页第2段： date ->data

## 第8章 哈希表与Go语言实现机制
* 86页最后一行:  
Go语言中的哈希表采用的是开放寻址法中的线性探测（Linear Probing）策略，线性探测策略是顺序（每次探测间隔为1）的。由于良好的CPU高速缓存利用率和高性能，该算法是现代计算机体系中使用最广泛的结构。在第12章还会看到，接口使用的全局itab哈希表采用了开放寻址法中的二次方探测策略。

修改为：

Go语言中的哈希表采用的是优化的拉链法，每一个桶中存储了8个元素用于加速访问。在第12章还会看到，接口使用的全局itab哈希表采用了开放寻址法中的二次方探测策略。

* 104页总结： 
Go语言为了解决哈希冲突使用了开放式寻址中的线性地址探测策略。
修改为：
Go语言为了解决哈希冲突使用了优化后的拉链法


## 第12章
* (12.4.4) 161页灰色第三个代码:  三角形周长 -> 长方形周长
* (12.4.5) 正文第3行: Writer实现了一系列方法 -> os.File实现了一系列方法
* (12.5.2) 正文第2段: data字段存储了接口中动态类型的函数指针 -> data字段存储了接口中动态类型的数据指针
## 第15章
* (15.5.2) 236页正文第2行: 那么调度器会将本地运行队列的一半放入局部运行队列 -> 那么调度器会将本地运行队列的一半放入全局运行队列
## 第17章
* (17.5)  第2段： trace工具 -> race工具
## 第20章
图20-23  p.k=&k -> p.x=&k

图20-25  p.k=&k -> p.x=&k
## 代码缩进问题

第5页第2段代码
```
AssignStmt struct {
	Op       Operator 
	Lhs, Rhs Expr
	simpleStmt
}
```
第10页第3段
```
case OCLOSURE,
     OCALLPART, 
     ORANGE,
     OFOR,
     OFORUNTIL,
     OSELECT,
     OTYPESW,
     OGO,
     ODEFER,
     ODCLTYPE,
     OBREAK,
     ORETJMP:
     v.reason = "unhandled op " + n.Op.String()
     return true
```

第27页第1段：
```
var f1 float64 = 0.3
var f2 float64 = 0.6
fmt.Println(f1 + f2)
```


第46页第2段代码
```
AssignStmt struct {
	Op       Operator 
	Lhs, Rhs Expr
	simpleStmt
}
```

第46页第三段：
```
type Node struct {
	Left  *Node
	Right *Node
	Ninit Nodes
	Nbody Nodes
	List  Nodes
	Rlist Nodes
	Type  *types.Type
	E     interface{}
	...
}
```

第50页第1段：
```
const k = 5
address := &k
```

第54页第2段：
```
a := "123" + 12
b := true + 12
c := nil + 12
```
第77页第3段
```
numbers:= []int{1,2,3,4}
numbers = append(numbers,5)
fmt.Println(len(numbers),cap(numbers))
```

82页第3段
```
old := make([]int64,3,3)
new := old[1:3]
fmt.Printf("%p %p",arr,slice) // 0xc000018140 0xc000018148
```

102页 第2段：
```
if h.extra != nil && h.extra.overflow != nil {
	h.extra.oldoverflow = h.extra.overflow
	h.extra.overflow = nil
}
```

103页 第1段：
```
for {
	b.tophash[i] = emptyRest
	if i == 0 {
		if b == bOrig {
			break // beginning of initial bucket, we're done.
		}
		// Find previous bucket, continue at its last entry.
		c := b
		for b = bOrig; b.overflow(t) != c; b = b.overflow(t) {
		}
		i = bucketCnt - 1
	} else {
		i--
	}
	if b.tophash[i] != emptyOne {
		break
	}
}
```

126页第2段：
```
	before:= time.Now()
	defer func(){
		after := time.Now()
		fmt.Println(after.After(before))
	}()
```

127页第3段：

```
	m.Lock()
	defer m.Lock()
	n.Lock()
	defer n.Lock()
```

130页第3段
```
for i:=0;i<3;i++{
	CALL    runtime.deferproc(SB)
}
CALL    runtime.deferreturn(SB)
RET
```


137页第3段：
```
	0x002f 00047 (stack.go:10)  MOVL    $16, ""..autotmp_0+104(SP)
	0x0037 00055 (stack.go:10)  LEAQ    "".add·f(SB), AX
	0x003e 00062 (stack.go:10)  MOVQ    AX, ""..autotmp_0+128(SP)
	0x0046 00070 (stack.go:10)  MOVQ    $3, ""..autotmp_0+176(SP)
	0x0052 00082 (stack.go:10)  MOVQ    $4, ""..autotmp_0+184(SP)
	0x005e 00094 (stack.go:10)  LEAQ    ""..autotmp_0+104(SP), AX
	0x0063 00099 (stack.go:10)  MOVQ    AX, (SP)
	0x0067 00103 (stack.go:10)  CALL    runtime.deferprocStack(SB)
```

第161页第3段
```
var s Shape
s = Rectangle{3, 4}
fmt.Printf("长方形周长:%v, 面积:%v \n",s.perimeter(),s.area())
s = Triangle{3, 4, 5}
fmt.Printf("三角形周长:%v, 面积:%v",s.perimeter(),s.area())
```

第161页第5段：
```
func (r Rectangle) getHeight() float64 {
	return  r.a
}
var s Shape
s = Rectangle{3, 4}
```

第164页最后一段
```
var a1 Empty = Cat{"Mimi"}
var a3 Empty = "Learn golang with me!"
var a4 Empty = 100
var a5 Empty = 3.14
```

195最后一段
```
type Type interface {
	Align() int
	FieldAlign() int
	Method(int) Method
	MethodByName(string) (Method, bool)
	String() string
	...
```

256页第3段
```
func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/‎",
		"https://www.taobao.com/",
	}
	c:=make(chan string)
	for _, link := range links {
		go checkLink(link,c)
	}
	for{
		go checkLink(<-c,c)
	}
}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c<-link
		return
	}
	fmt.Println(link, "is up!")
	c<-link
}

```

304页 第1段
```
off := c.tinyoffset
if off+size <= maxTinySize && c.tiny != 0 {
	x = unsafe.Pointer(c.tiny + off)
	c.tinyoffset = off + size
	return x
}

```

329页第2段：
删除空格

330页第1段
```
for sp := s.specials; sp != nil; sp = sp.next {
	if sp.kind != _KindSpecialFinalizer {
		continue
	}
	...
	scanobject(p, gcw)
	scanblock(uintptr(unsafe.Pointer(&spf.fn)), sys.PtrSize, &oneptrmask[0], gcw, nil)
}

```

356页第1段：
```
func main(){
	f, err := os.Create(*cpuProfile)
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	busyLoop()
}
```
