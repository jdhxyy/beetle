# beetle

## 1. 介绍
命令行解析工具，支持长选项，短选项，参数三种类型。

## 2. 使用方法
假设要解析的字符串是
```
boardtype --set PCB1
```

则首先需要注册命令字的回调函数：
```go
bettle.Register("boardtype", dealBoardtype)

func dealBoardtype() string {
    if bettle.GetItemNum() == 3 {
        item := bettle.GetItem(1)
        if (item.Type == bettle.TypeLongOption && item.Value == "set") ||
            (item.Type == bettle.TypeShortOption && item.Value == "s") {
            return fmt.Sprintln("set boardtype success", bettle.GetItem(2).Value)
        }
    }
    return ""
}
```
