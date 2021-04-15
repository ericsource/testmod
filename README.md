本文来源：[https://roberto.selbach.dev/intro-to-go-modules/]

##### 创建一个可以供其他项目使用的项目：testmod


```
package testmod

import (
    "fmt"
)

func Hi(name string) string {
    return fmt.Sprintf("Hi, %s", name)
}
```
```
git clone https://github.com/ericsource/testmod.git
git add testmod.go
git commit -m "testmod"
git push
git tag v1.0.0
git push --tags
```

##### 创建一个使用testmod包的项目 mymod.go
```
package main

import (
    "fmt"
    "github.com/robteix/testmod"
)

func main() {
    fmt.Printf(testmod.Hi("roberto"))
}
```

```
go mod init mymod
go mod tidy
go run mymod.go
```

#####  更新版本
```
//添加新方法
func New(name string) string {
    return fmt.Sprintf("New, %s", name)
}
```
```
git commit -m "add New func" testmod.go
git tag v1.0.1
git push --tags origin
```

##### 选择不同的版本
```
//更新到最新版本
go get -u github.com/ericsource/testmod
//go.mod文件结果
require github.com/robteix/testmod v1.0.1

//更新到指定版本
go get github.com/ericsource/testmod@v1.0.0
```
