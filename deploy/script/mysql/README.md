## 生成mysql的model代码

`goctl model`为`goctl`提供的数据库模型代码生成指令，目前支持`MySQL`、`PostgreSQL`、`Mongo`的代码生成，`MySQL`支持从`sql`文件和数据库连接两种方式生成，`PostgreSQL`仅支持从数据库连接生成。

脚本使用`goctl model mysql datasource`指令从数据库连接生成`model`代码（需要保证数据库已开启，并已创建对应数据库和表）。

```shell
goctl model mysql datasource --help
Generate model from datasource

Usage:
  goctl model mysql datasource [flags]

Flags:
      --branch string   The branch of the remote repo, it does work with --remote
  -c, --cache           Generate code with cache [optional]
  -d, --dir string      The target dir
  -h, --help            help for datasource
      --home string     The goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
      --idea            For idea plugin [optional]
      --remote string   The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
                        The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
      --style string    The file naming format
  -t, --table strings   The table or table globbing patterns in the database
      --url string      The data source of database,like "root:password@tcp(127.0.0.1:3306)/database"


Global Flags:
  -i, --ignore-columns strings   Ignore columns while creating or updating rows (default [create_at,created_at,create_time,update_at,updated_at,update_time])
      --strict                   Generate model in strict mode
```

脚本参数：脚本 创建的数据库名后缀 表名

```shell
./genModel.sh usercenter user
./genModel.sh usercenter user_auth
```
再将`./genModel`下的文件剪切到对应服务的`model`目录里面，并修改`package`为`model`。