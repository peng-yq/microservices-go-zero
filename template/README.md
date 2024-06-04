## 使用模板（Template）定制项目

[goctl template](https://go-zero.dev/docs/tutorials/cli/template)

项目对`template`部分的修改：

1. `api`：
   1. `handler.tpl`
   2. `main.tpl`
2. `model`
   1. `err.tpl`
   2. `import.tpl`
   3. `import-no-cache.tpl`
   4. `interface-update.tpl`
   5. `update.tpl`
3. `rpc`
   1. `main/tpl`

**初始化模板**

```shell
goctl template init
```

**将模板放置在与安装`goctl`版本一致的文件夹如`1.6.5`，并在使用`goctl`命令时加上`--home=`；或覆盖`~/.goctl/`下的默认模板（不推荐）**