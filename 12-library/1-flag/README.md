# Flag

`Flag`是`Go`语言中的一个标准库，它提供了方便的接口来解析命令行参数。定义格式如下：

```go
// 格式一
variableName := flag.[T]("cmd name", "default value", "some descriptive information")

// 格式二
var variableName T
flag.[T]Var(&variableName, "cmd name", "default value", "some descriptive information")
```



```go
// go run main.go -name golang -debug -age 20 -score 80 -timeout=10m
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [args]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// 定义一个字符串类型的命令行标志参数
	var name = flag.String("name", "defaultName", "The name to be used.")
	// 定义一个布尔类型的命令行标志参数
	var debug = flag.Bool("debug", false, "Enable debug mode.")
	var age = flag.Int("age", 18, "age")
	var score = flag.Float64("score", 99, "score")
	var timeout = flag.Duration("timeout", 5*time.Second, "a duration")

	flag.Usage = usage

	// 解析命令行标志参数
	flag.Parse()
	fmt.Println("name:", *name, "debug:", *debug, "age:", *age, "score:", *score, "timeout:", *timeout)
}

```

`flag.String()` - 用于定义一个字符串类型的命令行标志参数。例如：

```go
var name = flag.String("name", "defaultName", "The name to be used.")
```

上面的代码将定义一个名为`name`的命令行标志参数，其默认值为`defaultName`，如果用户未提供值，则使用默认值。第三个参数是该标志参数的描述。

`flag.Bool()` - 用于定义一个布尔类型的命令行标志参数。例如：

```go
var debug = flag.Bool("debug", false, "Enable debug mode.")
```

上面的代码将定义一个名为`debug`的命令行标志参数，其默认值为`false`，如果用户在命令行上使用`-debug`选项，则其值为`true`。

`flag.Int()`、`flag.Float64()`和`flag.Duration()` - 用于定义整型、浮点型和时间间隔类型的命令行标志参数。例如：

```go
var age = flag.Int("age", 18, "age")
var score = flag.Float64("score", 99, "score")
var timeout = flag.Duration("timeout", 5*time.Second, "a duration")
```

上面的代码分别定义了一个名为`num`的整型命令行标志参数、一个名为`pi`的浮点型命令行标志参数和一个名为`timeout`的时间间隔类型命令行标志参数。

`flag.Parse()` - 解析命令行标志参数并将其存储在相应的变量中。

```go
flag.Parse()
```

`flag.Usage` - 定义如何显示命令行标志参数的使用说明。例如：

```go
package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [args]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	// 解析命令行标志参数
	flag.Parse()
}
```

上面的代码定义了一个`usage()`函数，用于显示命令行标志参数的使用说明。通过将其赋值给`flag.Usage`变量，当用户在命令行上使用`-h`或`--help`选项时，将会显示此使用说明。

通过xxxVar解析：

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	name    string
	debug   bool
	age     int
	score   float64
	timeout time.Duration
)

// go run main.go -name zcw -debug -age 20 -score 80 -timeout 10m
func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&name, "name", "defaultName", "the name to be used")
	flag.IntVar(&age, "age", 18, "the age of the person")
	flag.Float64Var(&score, "score", 99.99, "the value of score")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "the timeout duration")
}

func main() {
	// 解析命令行标志参数
	flag.Parse()
	fmt.Println("name:", name, "debug:", debug, "age:", age, "score:", score, "timeout:", timeout)
}
```

# Pflag

`pflag`库使用方式基本和flag一样，只是提供更强大的功能，`pflag`库的命令行参数可以设置别名。如果设置了别名，则可以使用别名代替参数名称。

```go
package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
)

var (
	name    string
	debug   bool
	age     int
	score   float64
	timeout time.Duration
)

// go run main.go --name zcw --debug -a 20 -s 80 -t 10m
// 短命令和长命令需--和-区分
func init() {
	pflag.BoolVar(&debug, "debug", false, "enable debug mode")
	pflag.StringVar(&name, "name", "defaultName", "the name to be used")

	pflag.IntVarP(&age, "age", "a", 18, "the age of the person")
	pflag.Float64VarP(&score, "score", "s", 99.99, "the value of score")
	pflag.DurationVarP(&timeout, "timeout", "t", 5*time.Second, "the timeout duration")
}

func main() {
	// 解析命令行标志参数
	pflag.Parse()
	fmt.Println("name:", name, "debug:", debug, "age:", age, "score:", score, "timeout:", timeout)
}
```

以上代码使用了`pflag.IntVarP()`函数，其中第一个参数是参数指针，第二个参数是参数名称，第三个参数是参数别名，第四个参数是默认值，第五个参数是参数说明。

# Viper

`viper`是一个用于Go语言的配置解析库，它支持多种配置格式（如JSON、YAML、TOML等）的解析，并提供了一些方便的函数来获取配置项的值。下面是`viper`库的基础使用方法。

conf.yaml配置文件：

```yaml
StringKey: "golang tutorial"
IntKey: 101
Float64Key: 101.101
BoolKey: false
IntSliceKey:
  - 1
  - 2
  - 3
MapKey:
  host: 127.0.0.1
  port: 8080
```

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("12-library/1-flag/ch_3/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件失败的情况
		log.Panicf("read conf error %s", err.Error())
	}

	// 获取字符串类型的配置项的值
	fmt.Printf("String: %+v\n", viper.GetString("StringKey"))
	// 获取整数类型的配置项的值
	fmt.Printf("Int: %+v\n", viper.GetInt("IntKey"))
	// 获取浮点数类型的配置项的值
	fmt.Printf("Float64: %+v\n", viper.GetFloat64("Float64Key"))
	// 获取布尔类型的配置项的值
	fmt.Printf("Bool: %+v\n", viper.GetBool("BoolKey"))
	// 获取切片类型的配置项的值
	fmt.Printf("IntSlice: %+v\n", viper.GetIntSlice("IntSliceKey"))
	// 获取Interface类型的配置项的值
	fmt.Printf("Map: %+v\n", viper.Get("MapKey"))
	// 获取映射类型的配置项的值
	fmt.Printf("Map: %+v\n", viper.GetStringMap("MapKey"))

	// 获取映射类型的配置项的值到结构体
	type MapKey struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
	var mk MapKey
	viper.UnmarshalKey("MapKey", &mk)
	fmt.Printf("Map: %+v\n", mk)

	// 获取所有配置项
	settings := viper.AllSettings()
	for key, value := range settings {
		fmt.Printf("%s=%v\n", key, value)
	}

	// 绑定环境变量 GOPATH
	viper.BindEnv("GOPATH")
	fmt.Printf("GOPATH: %+v\n", viper.Get("GOPATH"))

	// 绑定环境变量 GOROOT 到 root 这个Key上
	viper.BindEnv("root", "GOROOT")
	fmt.Printf("GOROOT: %+v\n", viper.Get("root"))
}
```

首先，导入`viper`库：

```go
import "github.com/spf13/viper"
```

然后，我们需要设置`viper`的配置项。`viper`支持多种配置格式的解析，包括JSON、YAML、TOML等，我们可以使用`SetConfigFile()`函数来指定配置文件的路径，并使用`ReadInConfig()`函数读取配置文件。例如，如果我们有一个名为`conf.yaml`的yaml格式的配置文件，我们可以这样设置`viper`的配置项：

```go
viper.SetConfigFile("12-library/1-flag/ch_3/conf.yaml")
if err := viper.ReadInConfig(); err != nil {
  // 处理读取配置文件失败的情况
	log.Panicf("read conf error %s", err.Error())
}
```

然后，我们就可以使用`viper`提供的函数来获取配置项的值了。`viper`提供了多种函数来获取不同类型的配置项的值，例如：

- `viper.GetString(key string) string`：获取字符串类型的配置项的值
- `viper.GetInt(key string) int`：获取整数类型的配置项的值
- `viper.GetFloat64(key string) float64`：获取浮点数类型的配置项的值
- `viper.GetBool(key string) bool`：获取布尔类型的配置项的值
- `viper.GetIntSlice(key string) []int`：获取Int切片类型的配置项的值
- `viper.Get(key string) interface{}`：获取Interface类型的配置项的值
- `viper.GetStringMap(key string) map[string]interface{}`：获取映射类型的配置项的值

除了获取配置项的值，`viper`库还提供了一些其他的基础用法，例如：将配置直接映射到结构体。

```go
// 获取映射类型的配置项的值到结构体
type MapKey struct {
  Host string `json:"host"`
  Port int    `json:"port"`
}
var mk MapKey
viper.UnmarshalKey("MapKey", &mk)
fmt.Printf("Map: %+v\n", mk)
```

`viper`还允许我们为配置项设置默认值，这样在配置文件中没有设置该配置项时，就会使用默认值。可以使用`SetDefault(key string, value interface{})`函数来设置默认值，例如：

```go
viper.SetDefault("key", "default value")
```

`viper`还允许我们获取所有配置值，使用`AllSettings()`函数获取所有配置值，它返回一个`map[string]interface{}`类型的值，其中键为配置项的名称，值为配置项的值。例如：

```go
// 获取所有配置项
settings := viper.AllSettings()
for key, value := range settings {
  fmt.Printf("%s=%v\n", key, value)
}
```

`viper`库还支持使用环境变量来设置配置项的值。可以使用`BindEnv()`函数将一个配置项绑定到一个环境变量上，例如：

```go
// 绑定环境变量 GOPATH
viper.BindEnv("GOPATH")
fmt.Printf("GOPATH: %+v\n", viper.Get("GOPATH"))

// 绑定环境变量 GOROOT 到 root 这个Key上
viper.BindEnv("root", "GOROOT")
fmt.Printf("GOROOT: %+v\n", viper.Get("root"))
```

上面的代码将一个名为`root`的配置项绑定到一个名为`GOROOT`的环境变量上。当环境变量被设置时，`viper`会自动使用环境变量的值覆盖配置文件中的值。

`viper`库还提供了一个`OnConfigChange()`函数，可以用来监听配置文件的变化。当配置文件发生变化时，`viper`会自动重新读取配置文件，并执行我们指定的回调函数。例如：

```go
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

func main() {
	viper.SetConfigFile("12-library/1-flag/ch_4/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件失败的情况
		log.Panicf("read conf error %s", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		fmt.Printf("String: %+v\n", viper.GetString("StringKey"))
	})

	// 获取字符串类型的配置项的值
	fmt.Printf("String: %+v\n", viper.GetString("StringKey"))

	stopper := make(chan os.Signal, 1)
	signal.Notify(stopper, os.Interrupt)

	<-stopper
}
```

需要注意的是，要使用`OnConfigChange()`函数监听配置文件变化，需要先调用`viper.WatchConfig()`函数开启配置文件的监听。

