<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<title>青少年体质健康中心</title>
<link rel="alternate icon" type="image/png" href="/static/images/i_icon.png">
<link rel="stylesheet" href="/static/css/css.css">
</head>
<body>
<div id="content">
  <div class="item" style="display: block;">
      <h2 class="top-title">消息中心/新增</h2>
      <p class="top-tip"></p>
      <hr/>
      <form method="post" action="#" enctype="multipart/form-data">
          <div class="top-operate btn-box btn-top-box">
            <a data-type="1" class="btn btn-blue" href="javascript:;">校园新闻</a>
            <a data-type="2" class="btn btn-green" href="javascript:;">公益频道</a>
          </div>
          <div class="box user-box cf">
              <ul>
                  <li class="cf">
                      <span>*标题：</span><div><input class="msg-name" type="text"></div>
                  </li>
                  <li class="cf regions-li">
                      <span>*描述：</span>
                      <div>
                          <textarea class="user-text" maxlength="250" placeholder=""></textarea>
                      </div>
                  </li>
                  <li class="cf">
                      <span>*类型：</span>
                      <div>
                          <select class="select-on news-type">
                          </select>
                      </div>
                  </li>
                  <li class="cf file-li">
                      <span>*图片：</span>
                      <div>
                          <div class="file-box first-file file-wrap">
                              <a href="javascript:;" class="btn btn-file">选择文件</a>
                              <input class="file" name="Image" type="file" accept="image/*">
                              <span></span>
                              <i>图片格式：212*146 jpg、png的图片</i>
                          </div>
                      </div>
                  </li>
                  <li class="cf" style="position: relative;">
                      <span style="position: absolute;">*内容：</span>
                      <div style="float: none;padding-left: 100px;">
                        <script type="text/plain" id="para2" style="width: 100%; height: 400px;"></script>
                      </div>
                  </li>
                  <li class="cf">
                      <div class="user-btn-box btn-box">
                          <a class="btn btn-skyblue btn-submit" href="javascript:void(0)">提交</a>
                      </div>
                  </li>
              </ul>
          </div>
      </form> 
  </div>
</div>
<script type="text/javascript" src="/static/js/jquery.js"></script>
<script type="text/javascript" src="/static/js/jquery-form.js"></script>
<script type="text/javascript" src="/static/ueditor/ueditor.config.js"></script>
<script type="text/javascript" src="/static/ueditor/ueditor.all.js"></script>
<script type="text/javascript" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
<script type="text/javascript" src="/static/js/index.js"></script>
</body>
</html>