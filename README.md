# web_template

### 某站项目框架模板


#### 运行参数：

-env 可选，根据配置文件而定
* prod
* test

```bash
-env prod -conf cmd/web_template.json
```

#### 数据库 Reverse

```bash
xorm reverse mysql dbUserName:password@tcp(hostname:3306)/databaseName?charset=utf8 pkg/dbmodel/goxorm pkg/dbmodel
xorm reverse mysql jerry:IloveJerry2019!@tcp(101.132.174.14:3306)/bilibili_comic?charset=utf8,utf8mb4 pkg/dbmodel/goxorm pkg/dbmodel
```
