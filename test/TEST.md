### Test

- 文件名以 _test.go 结尾
    - 功能测试函数
        - Test 前缀命名
    - 基准测试函数
        - Benchmark 前缀命名
    - 示例测试函数
        - Example 前缀命名

```text
go test 工具扫描 *_test文件来寻找特殊函数，并生成一个临时的 main 包来调用、编译运行，最后清空临时文件
只能在扫描当前文件夹下的 *_test 文件
```

### cli

- ``-v`` 显示测试用例名称和执行的时间
- ``-run`` 正则匹配，测试函数名匹配的函数
    - ``go test -v -run="TestNonPalindrome"
      ``
- ``go test -v -test_xxx.go -test.run Testxxx`` 测试指定方法
- ``go test -v xxx.go xxx_test.go`` 测试单个文件，将依赖文件放前面

### notes

- t.ErrorF 输出的信息不会导致程序终止，测试用例是彼此独立的，如果一个条目造成测试失败，其他条目仍然会继续测试
- 想要终止，可以使用 t.Fatal、t.Fatalf，这些函数的调用必须和 Test 函数在同一个 goroute 中
- 测试错误信息一般格式是 ``"f(x)=y,want z"``