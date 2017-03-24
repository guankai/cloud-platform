## 云中台服务接口

### 服务注册

`url`  

/cp/v1/service/add

`method`  

post

`header`

UserName (用户添加自己的服务时,需要在header中添加username)

`formdata`

- serviceName 服务名称, 必须
- servicePic 服务Logo, 必须
- requestPath 访问路径, 必须 
- provider 提供者，必须
- version 版本，必须
- serviceDesc 服务描述
- upstreamUrl 服务backend地址，必须

`response`

```
"服务注册成功"
```

### 用户服务查询

该接口会查询出来所有系统已有的服务，如果用户已经开启的服务，则返回用户的对应该服务的access_key

`url`  

/cp/v1/relation/query

`method`  

post

`Header`

UserName

`formdata`

- page 查询当前页, 必须
- pageSize 每页行数, 必须


`response`

```
[
  {
    "typeId": 1,
    "typeName": "bigdata",
    "userServices": [
      {
        "service_id": "aa220f62-2a50-4627-8663-a4b86ec6e91c",
        "service_name": "baidu",
        "service_pic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
        "request_path": "/baidu",
        "api_id": "ea6c8f6d-3de2-4ef9-8216-9119ae62f02b",
        "provider": "baidu",
        "create_time": "2017-03-21T08:57:37+08:00",
        "version": "1.0",
        "service_desc": "baidu",
        "upstream_url": "http://www.baidu.com",
        "status": "1",
        "relation_id": "3f346c23-d808-4d07-9092-037c691f6810",
        "consumer_id": "9f6ad05b-4dcc-401d-94c9-f6131512e743",
        "api_key_id": "5f428f38-b037-4a21-bfd3-9371853306c4",
        "api_key": "f0ece73f4ff24a2daf2d7f6fa16fdacb",
        "type_name": "bigdata"
      },
      {
        "service_id": "b928b2e5-cd33-4066-9a1c-42f472b3637b",
        "service_name": "google",
        "service_pic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
        "request_path": "/google",
        "api_id": "975f60de-a8d5-4736-98d4-442dae8e1ba7",
        "provider": "google",
        "create_time": "2017-03-21T09:45:23+08:00",
        "version": "1.0",
        "service_desc": "google",
        "upstream_url": "http://www.google.com",
        "status": "0",
        "relation_id": "",
        "consumer_id": "",
        "api_key_id": "",
        "api_key": "",
        "type_name": "bigdata"
      }
    ],
    "pageInfo": {
      "Rows": 10,
      "Total": 1,
      "Page": 1,
      "Nums": 2
    }
  },
  {
    "typeId": 2,
    "typeName": "iot",
    "userServices": [
      {
        "service_id": "d3a286ae-ce02-4ab6-b003-74debf4997d4",
        "service_name": "bimworks",
        "service_pic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
        "request_path": "/bimworks",
        "api_id": "ff0988e3-7936-4237-be53-8997365c1b48",
        "provider": "bim",
        "create_time": "2017-03-21T08:51:09+08:00",
        "version": "1.0",
        "service_desc": "bimworks",
        "upstream_url": "http://www.bim.works:7001",
        "status": "0",
        "relation_id": "",
        "consumer_id": "",
        "api_key_id": "",
        "api_key": "",
        "type_name": "iot"
      }
    ],
    "pageInfo": {
      "Rows": 10,
      "Total": 1,
      "Page": 1,
      "Nums": 1
    }
  },
  {
    "typeId": 3,
    "typeName": "ai",
    "userServices": [
      {
        "service_id": "e81fec81-2b95-4fea-a755-f0d01d906a98",
        "service_name": "北京天气预报",
        "service_pic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
        "request_path": "/beijingweather",
        "api_id": "01e008fb-e620-4795-9e8a-a646d4ff60d9",
        "provider": "demo",
        "create_time": "2017-03-21T08:39:21+08:00",
        "version": "1.0",
        "service_desc": "北京天气预报",
        "upstream_url": "http://v.juhe.cn/weather/index?dtype=json&format=1&key=cd368ae09fe1a64e7da30f75af387e6c&cityname=北京",
        "status": "0",
        "relation_id": "",
        "consumer_id": "",
        "api_key_id": "",
        "api_key": "",
        "type_name": "ai"
      }
    ],
    "pageInfo": {
      "Rows": 10,
      "Total": 1,
      "Page": 1,
      "Nums": 1
    }
  }
]
```

### 服务开启

`url`  

/cp/v1/relation/startup

`method`  

post

`Header`

UserName

`formdata`

- serviceId 服务ID, 必须


`response`

```
{
  "apiKey": "91818c5141c5439e9662f98755f5f45a",
  "msg": "服务开启成功",
  "requestPath": "http://223.202.32.56:8055/beijingweather"
}
```

### 服务关闭

`url`  

/cp/v1/relation/shutdown

`method`  

post

`Header`

UserName

`formdata`

- serviceId 服务ID, 必须


`response`

```
{
  "msg": "服务关闭成功",
  "requestPath": "http://223.202.32.56:8055/baidu"
}
```

### 获取用户已开启的服务

`url`  

/cp/v1/relation/getOpened

`method`  

post

`Header`

UserName

`formdata`

- serviceId 服务ID, 必须


`response`

```
[
  {
    "relationId": "3f346c23-d808-4d07-9092-037c691f6810",
    "userName": "patrick",
    "Service": {
      "serviceId": "aa220f62-2a50-4627-8663-a4b86ec6e91c",
      "serviceName": "baidu",
      "servicePic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
      "requestPath": "/baidu",
      "apiId": "ea6c8f6d-3de2-4ef9-8216-9119ae62f02b",
      "Plugin": {
        "pluginId": "1",
        "pluginName": "key-auth",
        "pluginDesc": "api验证插件"
      },
      "provider": "baidu",
      "createTime": "2017-03-21T08:57:37+08:00",
      "version": "1.0",
      "serviceDesc": "baidu",
      "upstreamUrl": "http://www.baidu.com",
      "callPath": "",
      "Type": {
        "typeId": 1,
        "typeName": "bigdata",
        "typeDesc": "大数据"
      }
    },
    "apiKeyId": "5f428f38-b037-4a21-bfd3-9371853306c4",
    "apiKey": "f0ece73f4ff24a2daf2d7f6fa16fdacb",
    "consumerId": "9f6ad05b-4dcc-401d-94c9-f6131512e743",
    "status": "0",
    "accessNum": 0
  },
  {
    "relationId": "39a6ea33-a18a-4b31-bf1a-af9c8da55b8c",
    "userName": "patrick",
    "Service": {
      "serviceId": "e81fec81-2b95-4fea-a755-f0d01d906a98",
      "serviceName": "北京天气预报",
      "servicePic": "https://s3.cn-north-1.amazonaws.com.cn/cloudos-file-server/guankai/other.png",
      "requestPath": "/beijingweather",
      "apiId": "01e008fb-e620-4795-9e8a-a646d4ff60d9",
      "Plugin": {
        "pluginId": "1",
        "pluginName": "key-auth",
        "pluginDesc": "api验证插件"
      },
      "provider": "demo",
      "createTime": "2017-03-21T08:39:21+08:00",
      "version": "1.0",
      "serviceDesc": "北京天气预报",
      "upstreamUrl": "http://v.juhe.cn/weather/index?dtype=json&format=1&key=cd368ae09fe1a64e7da30f75af387e6c&cityname=北京",
      "callPath": "",
      "Type": {
        "typeId": 3,
        "typeName": "ai",
        "typeDesc": "人工智能"
      }
    },
    "apiKeyId": "191d9be1-ff91-4c60-95f7-af0b0622e9aa",
    "apiKey": "91818c5141c5439e9662f98755f5f45a",
    "consumerId": "9f6ad05b-4dcc-401d-94c9-f6131512e743",
    "status": "1",
    "accessNum": 0
  }
]
```














