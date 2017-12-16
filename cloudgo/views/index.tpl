<!DOCTYPE html>

<html>
<head>
  <title>Hello</title>

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }
    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 30px;
      line-height: 40px;
      background-color: #fff;
    }
    .authour{
      width: 800px;
      height: 300px;
      margin:80px auto;
    }
  </style>
</head>

<body>
    <div class="author">
      Hello,{{.Sex}}{{.Name}}</br>
      Nice to meet you!
    </div>
</body>
</html>
