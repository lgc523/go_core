### 闭包

- 闭包 = 函数 + 引用环境
- 闭包对闭包外的环境引入是直接引用，编译器会检测到闭包，将闭包引用的外部变量好分配到堆上
- 如果函数返回的闭包引用了函数内的局部变量(参数或内部变量)
  - 多次调用函数，返回闭包引用的外部变量是多个副本，每次调用都会为局部变量分配内存
  - 一个闭包函数调用多次，闭包函数共享外部引用
