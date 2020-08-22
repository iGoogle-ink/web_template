# web_template

### 某站项目框架模板（修改版）

#### 运行参数：

```bash
-conf cmd/config.yaml
```

### 接口地址

**Ping：GET**
- http://www.igoogle.ink/ping

**获取学生信息列表：GET**
- http://www.igoogle.ink/student/list

**获取学生信息ByID：GET**
- 参数：id，
- http://www.igoogle.ink/student/id?id=1

**获取教师信息列表：GET**
- http://www.igoogle.ink/teacher/list

**获取教师信息ByID：GET**
- 参数：id，
- http://www.igoogle.ink/teacher/id?id=1

**添加学生信息：POST**
- 请求体：JSON
```bash
{
	"name":"学生G",
	"teacher_id":3
}
```
- http://www.igoogle.ink/student/add

**添加教师信息：POST**
- 请求体：JSON
- http://www.igoogle.ink/teacher/add
```bash
{
	"name":"英语老师",
	"subject":"语文"
}
```