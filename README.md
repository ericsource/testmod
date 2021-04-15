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

##### 修复版本
```
// Hi returns a friendly greeting
func Hi(name string) string {
-   return fmt.Sprintf("Hi, %s", name)
+   return fmt.Sprintf("Hi, %s!", name)
}
```
```
git commit -m "Emphasize our friendliness" testmod.go
git tag v1.0.1
git push --tags origin v1
```

##### 选择不同的版本
```
require github.com/robteix/testmod v1.0.1
```
