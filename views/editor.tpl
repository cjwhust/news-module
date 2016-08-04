<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <title>UEDITOR</title>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.js"></script>
    <script type="text/javascript" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" src="/static/ueditor/ueditor.all.js"></script>
    <script type="text/javascript" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
</head>
<body>
    <header>
        <h1>Demo</h1>
    </header>
    <div>
        <form action="/user/add" method="post">
            <input name="para1" value="test1"/>
            <script type="text/plain" id="para2" name="para2"></script>
            <script type="text/javascript">
//                var options = {"fileUrl":"/article/upload","filePath":"","imageUrl":"/articleImage/upload","imagePath":""
//                    ,"initialFrameWidth": "60%","initialFrameHeight":"400","enableAutoSave": true,"saveInterval": 500}
                var options = {"initialFrameWidth": "60%","initialFrameHeight":"400","enableAutoSave": true,"saveInterval": 500}
                var ue = UE.getEditor('para2', options)
            </script>
            <button type="submit">submit</button>
        </form>
    </div>
</body>
</html>