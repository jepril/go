# go
#golang的context

##一个简单的并发
[举个例子](https://play.golang.org/p/yBiFEXzpu5b)

##什么是context
[这是一个chan+select来结束一个goroutine的方式]
（https://play.golang.org/p/ZZ6EjeZjfj7）

上面说的这种场景是存在的，比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。所以我们需要一种可以跟踪goroutine的方案，才可以达到控制他们的目的，这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文。

[context重写]（https://play.golang.org/p/cuCURu2Rh48）
[刚刚是控制一个goroutine，现在控制三个]
（https://play.golang.org/p/128CPyFdi-P）

##golang内置的context包
context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。

向服务器的传入请求应创建一个上下文，而对服务器的传出调用应接受一个上下文。它们之间的函数调用链必须传播Context，可以选择将其替换为使用WithCancel，WithDeadline，WithTimeout或WithValue创建的派生Context。取消上下文后，从该上下文派生的所有上下文也会被取消。

WithCancel，WithDeadline和WithTimeout函数采用Context（父级）并返回派生的Context（子级）和CancelFunc。调用CancelFunc会取消该子代及其子代，删除父代对该子代的引用，并停止所有关联的计时器。未能调用CancelFunc会使子代及其子代泄漏，直到父代被取消或计时器触发。审核工具检查所有控制流路径上是否都使用了CancelFuncs。

使用上下文的程序应遵循以下规则，以使各个包之间的接口保持一致，并启用静态分析工具来检查上下文传播：

不要将上下文存储在结构类型中；而是将上下文明确传递给需要它的每个函数。Context应该是第一个参数，通常命名为ctx

###context.Context
>type Context interface {
>    Deadline() (deadline time.Time, ok bool)
>    Done() <-chan struct{}
>    Err() error
>    Value(key interface{}) interface{}
>}

|字段|含义|
|-----|-----|
Deadline|返回一个time.Time，表示当前Context应该结束的时间，ok则表示有结束时间
Done|当Context被取消或者超时时候返回的一个close的channel，告诉给context相关的函数要停止当前工作然后返回了。(这个有点像全局广播)
Err|context被取消的原因
Value|context实现共享数据存储的地方，是协程安全的（还记得之前有说过map是不安全的？所以遇到map的结构,如果不是sync.Map,需要加锁来进行操作）
以上常用的就是Done，如果Context取消的时候，我们就可以得到一个关闭的chan，关闭的chan是可以读取的，所以只要可以读取的时候，就意味着收到Context取消的信号了，以下是这个方法的经典用法。
```
  func Stream(ctx context.Context, out chan<- Value) error {
  	for {
  		v, err := DoSomething(ctx)
  		if err != nil {
  			return err
  		}
  		select {
  		case <-ctx.Done():
  			return ctx.Err()
  		case out <- v:
  		}
  	}
  }
  ```

4种context
![image.png](https://upload-images.jianshu.io/upload_images/22969962-563664c09081e6b2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Context接口并不需要我们实现，Go内置已经帮我们实现了2个，我们代码中最开始都是以这两个内置的作为最顶层的partent context，衍生出更多的子Context。
```
func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```
这两个私有变量都是通过 `new(emptyCtx)` 语句初始化的，它们是指向私有结构体 `context.emptyCtx`的指针，这是最简单、最常用的上下文类型：
```
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```
从源代码来看，[`context.Background`] 和 [`context.TODO`] 函数其实也只是互为别名，没有太大的差别。它们只是在使用和语义上稍有不同：

*   [`context.Background`]是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
*   [`context.TODO`]应该只在不确定应该使用哪种上下文时使用；

在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 [`context.Background`]作为起始的上下文向下传递。

###继承衍生

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
这四个With函数，接收的都有一个partent参数，就是父Context，我们要基于这个父Context创建出子Context的意思，这种方式可以理解为子Context对父Context的继承，也可以理解为基于父Context的衍生。

通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。

WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。 
WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context

####context.Withcancel
`context.WithCancel`函数能够从 `context.Context` 中衍生出一个新的子上下文并返回用于取消该上下文的函数（CancelFunc）。一旦我们执行返回的取消函数，当前上下文以及它的子上下文都会被取消，所有的 Goroutine 都会同步收到这一取消信号。
![image.png](https://upload-images.jianshu.io/upload_images/22969962-3bc5b3b63fdc8960.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}
```
这个代码的实现是一个套娃！！！我放个链接，有空的可以看看，[戳这]([https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/)
)

####context.Withvalue
传值方法
先上代码
```
func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}
```
我们可以使用context.WithValue方法附加一对K-V的键值对，这里Key必须是等价性的，也就是具有可比性；Value值要是线程安全的。

这样我们就生成了一个新的Context，这个新的Context带有这个键值对，在使用的时候，可以通过Value方法读取ctx.Value(key)。

记住，使用WithValue传值，一般是必须的值，不要什么值都传递。


##小结
1.不要把Context放在结构体中，要以参数的方式传递
2.以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
3.给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
4.Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
5.Context是线程安全的，可以放心的在多个goroutine中传递

