# Test

在 Go 语言中，可以使用内置的测试框架来进行单元测试。

## 目录

- 简单测试
- 子测试
- 表驱动测试

## 简单测试

如下是 Go 测试的基本用法：

1. 创建测试文件：测试代码应该与被测试的代码放在同一个包中，并以 `_test.go` 结尾。例如，要测试名为 `add` 文件中的函数，可以创建一个名为 `add_test.go` 的文件。
2. 编写测试函数：测试函数应该以 `Test` 开头，后跟被测试函数的名称。例如，要测试名为 `add` 的函数，可以创建一个名为 `TestAdd` 的测试函数。
3. 写测试代码：测试函数应该包含测试代码，以验证被测试函数的行为是否符合预期。测试代码应该使用 `t.Error` 或 `t.Fail` 等函数来报告测试失败。
4. 运行测试：在命令行中使用 `go test` 命令来运行测试。如果测试通过，将输出 `PASS`。如果测试失败，将输出失败的测试用例的详细信息。

示例代码：

```
ch_1
|-- add.go
|-- add_test.go
```

```go
// add.go
package ch_1

func add(x, y int) int {
	return x + y
}
```

```go
// add_test.go
package ch_1

import "testing"

// go test
func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) returned %d, expected %d", result, expected)
	}
}
```

这个示例测试了 `Add` 函数是否能正确地计算两个整数的和。

## 子测试

除了基本的测试用例，Go 语言的测试框架还提供了一些高级功能，以帮助更好地组织和管理测试代码。其中子测试：可以在一个测试函数中创建多个子测试，每个子测试都可以独立运行和失败。使用 `t.Run` 函数来创建子测试。

```go
func TestAdd(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		result := add(2, 3)
		expected := 5
		if result != expected {
			t.Errorf("add(2, 3) returned %d, expected %d", result, expected)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		result := add(3, 3)
		expected := 6
		if result != expected {
			t.Errorf("add(3, 3) returned %d, expected %d", result, expected)
		}
	})
}
```

## 表驱动测试

表驱动测试：可以使用表格驱动测试来测试多组输入和输出。创建一个包含输入和预期输出的表格，然后使用 `for` 循环遍历表格中的每一行并运行测试。

```go
func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1+2",
			args{
				a: 1,
				b: 2,
			},
			3,
		},
		{
			"10+10",
			args{
				a: 10,
				b: 10,
			},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

## 可选命令

 默认情况下, `go test` 命令会运行指定包中的所有测试, 并输出测试结果和覆盖率报告。如果测试成功，则输出 PASS；如果测试失败，则输出 FAIL。如果测试运行期间出现 panic，测试会被中止并输出相关信息。

`go test` 命令的一些常用标志如下：

- `-v`：输出测试详情；
- `-run`：指定要运行的测试函数；
- `-cover`：输出测试覆盖率报告；
- `-bench`：运行基准测试。

例如，以下命令将运行 `example_test.go` 文件中的所有测试并输出覆盖率报告：

```
go test -cover example_test.go
```

`go test` 命令是 Go 语言测试的核心命令，可以帮助开发者编写、运行和管理测试代码，确保代码的质量和稳定性。





## 思考题

## 参考
https://bingdoal.github.io/backend/2022/05/unit-test-on-golang/

