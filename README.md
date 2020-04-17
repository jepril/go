# golang的context

## 一个简单的并发
[举个例子](https://play.golang.org/p/yBiFEXzpu5b)

## 什么是context

我理解的context类似于一个保存了状态的object，他被翻译为上下文，但其实它更像环境状态，记载了当前情况下程序执行状态，也就是上文的状态，会影响下文的发展。
每个Goroutine在执行之前，都要先知道“程序当前的执行状态”。通常，将这些执行状态封装在一个Context变量中，传递到要执行的Goroutine中。当状态发生了变化，和这个状态进行过交互的一个或者多个Goroutine也会发生相应的变化。

或者，再简单一点，代码c = a + b
a,b的value就是context。

## 结束一个goroutine

我们都知道一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的。

这种情况化，一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台goroutine不停的检查这个变量，如果发现被通知关闭了，就自我结束。

这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 。

[这是一个chan+select来结束一个goroutine的方式](https://play.golang.org/p/ZZ6EjeZjfj7)

虽然可以用channel+select来从外部杀死某个线程，但是某些情况下会比较麻烦，例如由一个请求衍生出多个线程并且之间需要满足一定的约束关系，以实现一些诸如：有效期，中止线程树，传递请求全局变量之类的功能。这就是context的优势。

[context重写](https://play.golang.org/p/cuCURu2Rh48)

每一个都使用了Context进行跟踪，当我们使用cancel函数通知取消时，这3个goroutine都会被结束。这就是Context的控制能力，它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，这时就可以进行清理操作了，最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题。

[刚刚是控制一个goroutine，现在控制三个](https://play.golang.org/p/128CPyFdi-P)

## golang内置的context包
context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。

向服务器的传入请求应创建一个上下文，而对服务器的传出调用应接受一个上下文。它们之间的函数调用链必须传播Context，可以选择将其替换为使用WithCancel，WithDeadline，WithTimeout或WithValue创建的派生Context。取消上下文后，从该上下文派生的所有上下文也会被取消。

WithCancel，WithDeadline和WithTimeout函数采用Context（父级）并返回派生的Context（子级）和CancelFunc。调用CancelFunc会取消该子代及其子代，删除父代对该子代的引用，并停止所有关联的计时器。未能调用CancelFunc会使子代及其子代泄漏，直到父代被取消或计时器触发。

Goroutine的创建和调用关系是分层级的。更靠顶部的Goroutine应有办法主动关闭其下属的Goroutine的执行（否则，程序就可能失控）。为了实现这种关系，Context结构像一棵树，叶子节点须总是由根节点衍生出来的。

所有的context的父对象，也叫根对象，根节点，是一个空的context，它不能被取消，它没有值，从不会被取消，也没有超时时间，它常常作为处理request的顶层context存在，然后通过WithCancel、WithTimeout函数来创建子对象来获得cancel、timeout的能力.这也就是后面提到的emptyCtx。



### context.Context
```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

|字段|含义|
|-----|-----|
Deadline|返回一个time.Time，表示当前Context应该结束的时间，ok则表示有结束时间
Done|当Context被取消或者超时时候返回的一个close的channel，告诉给context相关的函数要停止当前工作然后返回了。(这个有点像全局广播)
Err|context被取消的原因
Value|context实现共享数据存储的地方，是协程安全的
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

### 基本数据结构
context的创建者称为root节点，其一般是一个处理上下文的独立goroutine。root节点负责创建Context的具体对象，并将其传递到其下游调用的goroutine. 下游的goroutine可以继续封装改Context对象，再传递更下游的goroutine.这些下游goroutine的Context 对象实例都要逐层向上注册。这样通过root节点的Context对象就可以遍历整个Context对象树，所以通知也能通知到下游的goroutine.


#### 4种context
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

### 继承衍生

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```
这四个With函数，接收的都有一个partent参数，就是父Context，我们要基于这个父Context创建出子Context的意思，这种方式可以理解为子Context对父Context的继承，也可以理解为基于父Context的衍生。

通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。而这四个，就是用来创建它的子节点、孙节点。

WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。 
WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context

#### context.Withcancel
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

#### context.Withvalue
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

![image.png](https://upload-images.jianshu.io/upload_images/22969962-6acf85fd36b2b0ca.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 原理
Context 的调用应该是链式的，通过WithCancel，WithDeadline，WithTimeout或WithValue派生出新的 Context。当父 Context 被取消时，其派生的所有 Context 都将取消。

通过4个context.WithXXX都将返回新的 Context 和 CancelFunc。调用 CancelFunc 将取消子代，移除父代对子代的引用，并且停止所有定时器。未能调用 CancelFunc 将泄漏子代，直到父代被取消或定时器触发。go vet工具检查所有流程控制路径上使用 CancelFuncs。

## 小结
1.不要把Context放在结构体中，要以参数的方式传递
2.以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
3.给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
4.Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
5.Context是线程安全的，可以放心的在多个goroutine中传递
6.为了保证父Context对象的创建环境获得对子Context将要被传递到的Goroutine的撤销权，当通过父Context对象创建子Context对象时，可同时获得子Context的一个撤销函数。

- 为什么不应该放在结构体？

Context 最基本的作用，是对一些 不那么全局的全局变量 的打包，把它放到结构体，其生存周期和作用域是无法控制的，相当于把它变成了它所在包的一个全局变量，那和最开始的目的不久矛盾了吗。

- 为什么 HTTP 包的 Request 结构体持有 context？

Request 本身就是一堆参数的集合，只不过参数太多单独写成结构体了而已，这堆参数在请求结束时或者读写超时时，就应该释放，需要一个可超时的 Context 来协助。那为什么不把请求参数都放在 Context 呢，因为可读性是非常重要的。

- 为什么是并发安全的？

Context 本身的实现是不可变的，既然不可变，那当然是线程安全的。并且通过 Context.Done() 返回的通道可以协调 goroutine 的行为。

