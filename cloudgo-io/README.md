# 关于项目
一个简单的cloudgo应用，主要包括：<br />
1. 静态文件服务
2. 简单的js访问
3. 提交表单，并输出一个表格
4. 自定义网络状态`501 Not Implemented`，使用`/unknwon`可验证

# 使用示例
1. 启动服务器
```
$ go run main.go

[negroni] listening on :9090
```

2. 访问静态文件示例
```
[negroni] 2017-11-21T16:10:05+08:00 | 200 |      499.3µs | localhost:9090 | GET /static/js/
[negroni] 2017-11-21T16:11:50+08:00 | 200 |      1.0032ms | localhost:9090 | GET /static/css/
```
![2](readme_assets/2.jpg)
![3](readme_assets/3.jpg)

3. js访问示例
hello.js
```
$(document).ready(function() {
    $.ajax({
        url: "/home/date"
    }).then(function(data) {
       $('.time').append(data);
    });
});
```
index.html(部分)
```
<html>
<head>
  <link rel="stylesheet" href="../static/css/main.css"/>
  <link rel="icon" href="../static/images/favicon.ico"/>
  <script src="http://code.jquery.com/jquery-latest.js"></script>
  <script src="../static/js/hello.js"></script>
</head>
<body>
  <div id="image"><img src="../static/images/image.jpg"/></div>
  <div id="hello">
      <h1>Hello,Guest!</h1>
      <p class="time">Today is </p>
      <form method="post">
          用户名:<input type="text" name="username"><br/>
          密　码:<input type="password" name="password"><br/>
          <input type="submit" value="登录">
      </form>
  </div>
</body>
</html>
```
访问index.html时，如图：
```
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      957.7µs | localhost:9090 | GET /
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      55.6485ms | localhost:9090 | GET /static/js/hello.js
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      62.667ms | localhost:9090 | GET /static/css/main.css
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      501.4µs | localhost:9090 | GET /static/images/image.jpg
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      540.2µs | localhost:9090 | GET /home/date
[negroni] 2017-11-21T16:19:26+08:00 | 200 |      424.9µs | localhost:9090 | GET /static/images/favicon.ico
```
![4](readme_assets/4.jpg)

4. 提交表单，返回表格
```
[negroni] 2017-11-21T16:34:18+08:00 | 200 |      500.4µs | localhost:9090 | POST /
```
![5](readme_assets/5.gif)

5.自定义网络状态
简单定义一个Nouimplented函数即可。
```
func NotImplemented(w http.ResponseWriter, r *http.Request) { 
     http.Error(w, "501 Not Implemented", 501) 
}
```
访问`/unknown`时，显示结果如图：
```
[negroni] 2017-11-21T16:36:38+08:00 | 501 |      1.0035ms | localhost:9090 | GET /unknown
```
![6](readme_assets/5.jpg)



