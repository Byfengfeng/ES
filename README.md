
<h1>介绍</h1>
<p>ES是一基于elastic封装的api。
   内置简单查询、根据条件单条查询、根据时间范围查询、设置获取数据的数量查询所有数据</p>

<h1>配置参数列表</h1>

```
    Host string json:"host"

    UserName string json:"user_name"

    PassWord string json:"pass_word"

    IndexDBName string mapstructure:"index-db-name"
```

<h1>使用方法</h1>
     <p>去获取<a href="github.com/Byfengfeng/es">github.com/Byfengfeng/es</a></p>
<h1>使用</h1>

```
	data := Es.EsData{
		"http://127.0.0.1:9200/",
		"esuser",
		"123456",
		"华山",
	}
	esClient := Es.NewEsClient(&data)
	queryAll := EsService.QueryAll(esClient, Log{}, 15, data.IndexDBName)
	fmt.Println(queryAll)
```


<h1>查询结果示例</h1>


```
{1605593790 INFO main/sasd.go:18 测试1 李四 20 123 1}
{1605593797 INFO main/sasd.go:18 测试2 李四 20 123 2}
{1605593901 INFO main/sasd.go:18 测试2 李四 20 123 3}
{1605593964 INFO main/sasd.go:18 测试3 李四 20 456 4}
{1605594119 INFO main/sasd.go:18 测试4 李四 20 456 5}
{1605594163 INFO main/sasd.go:18 测试2 李四 20 789 6}
{1605593790 INFO main/sasd.go:18 测试1 李四 20 123 1}
{1605593797 INFO main/sasd.go:18 测试2 李四 20 123 2}
{1605593901 INFO main/sasd.go:18 测试2 李四 20 123 3}
{1605593964 INFO main/sasd.go:18 测试3 李四 20 456 4}
{1605594119 INFO main/sasd.go:18 测试4 李四 20 456 5}
{1605594163 INFO main/sasd.go:18 测试2 李四 20 789 6}
{1605593790 INFO main/sasd.go:18 测试1 李四 20 123 1}
{1605593797 INFO main/sasd.go:18 测试2 李四 20 123 2}
{1605593901 INFO main/sasd.go:18 测试2 李四 20 123 3}
```
