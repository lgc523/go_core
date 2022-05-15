# defer

- defer 后面必须是函数或方法带哦用，不能是语句
- defer 函数的实参在注册时通过值拷贝传递进去
- defer 位于 return 之后，因为没有注册不会执行
- 主动调用 os.Exit(int) defer 不再被执行
- defer 不要操作有名返回值参数