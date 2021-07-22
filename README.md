# LearnGolang
## Getting Start
### hello 
#### Don't communicate by sharing memory, share memory by communicating.
#### 在内部，接口值可以看做包含值和具体类型的元组：

(value, type)
接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

#### 底层值为 nil 的接口值
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil。

nil 接口值既不保存值也不保存具体类型。(其本质就是nil,go的哲学就是一切一视同仁)

为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
