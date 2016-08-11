<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
    <title>UEDITOR</title>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.js"></script>
    <script type="text/javascript" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" src="/static/ueditor/ueditor.all.js"></script>
    <script type="text/javascript" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
</head>
<body>
<form action="/user/add" method="post">
    <script type="text/plain" id="para2" name="para2"></script>
    <script type="text/javascript">
        var options = {"initialFrameWidth": "60%","initialFrameHeight":"400","enableAutoSave": true,"saveInterval": 500}
        var ue = UE.getEditor('para2', options);
    </script>
</form>
</body>
</html>