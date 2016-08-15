var Config = {
    magicTypeShow: false, // 是否可切换图表
    xLength: 10, // 图表横轴显示最大个数
    imgReg: /^(.)+\.((jpg)|(gif)|(png)|(jpeg))$/i, // 图片文件类型
    xlsReg: /^(.)+\.((xlsx)|(xls))$/i, // EXCEL文件类型
    reportReg: /^(.)+\.(.)+$/i, // 报告文件类型
    pwdStrength3: /^(?=.*[!(),-.:*=?_{}])(?=.*\d)(?=.*[a-zA-Z]).{6,20}$/, // 密码强度 强
    pwdStrength2: /^((?=.*[^a-zA-Z])(?=.*[a-zA-Z]).{6,20}|(?=.*[^0-9])(?=.*[0-9]).{6,20})$/, // 密码强度 中
    pwdStrength1: /^(.{1,5}|\d+|[a-zA-Z]+|[!(),-.:*=?_{}]+|(.)\2+)$/,  // 密码强度 弱
    homeworkReg: /^.{1,250}$/,  // 老师作业内容
    noEmptyReg: /^.{1,}$/,  // 非空
    passwordReg: /^[a-zA-Z0-9!(),-.:*=?_{}]{6,20}$/,  // 密码
    scoreReg: /^(\d+(\.\d{1,2})?$)|(^.{0}$)/,  // 成绩录入的成绩
    minSecScoreReg: /^([0-9]{1,2}$)|(^.{0}$)/,  // 成绩录入的成绩
    usernameReg: /^[a-zA-Z0-9_-]{3,16}$/,  // 用户名
    nicknameReg: /^.{0,20}$/,  // 昵称
    nationReg: /^[\u4e00-\u9fa5]{0,20}$/,  // 民族
    idnumberReg: /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)|(^.{0}$)/,  // 身份证号
    positionReg: /(^.{2,20}$)|(^.{0}$)/,  // 职位
    realnameReg: /^[\u4e00-\u9fa5]{2,20}$/,  // 真实姓名
    ageReg: /^\d{0,2}$/,  // 年龄
    addressReg: /(^.{2,50}$)|(^.{0}$)/,  // 地址
    phoneReg: /(^1\d{10}$)/,  // 手机号码
    emailReg: /^([a-zA-Z0-9_\.\-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/,  // 手机号码
    verifycodeReg: /^.+$/,  // 验证码
    titleReg: /^.{1,50}$/,  // 标题
    descriptionReg: /^.{1,250}$/,  // 描述
    feedbackReg: /^.{1,500}$/,  // 用户反馈内容
    numberReg: /^[a-zA-Z0-9]{1,20}$/,  // 学生证号
    teacherNumberReg: /(^[a-zA-Z0-9]{1,20}$)|(^.{0}$)/,  // 教师证号
    schoolReg: /^.{1,500}$/,  // 学校名称
    classReg: /^.{1,500}$/,  // 班级名称
};
//验证失败提示
var FailInfo = {
    homeworkReg: '输入的作业内容有误，作业内容为1-250个字符！',
    scoreReg: '请输入有效的成绩！',
    usernameReg: '请输入由3-16位，由字母、数字、_-组成的用户名！',
    nicknameReg: '用户昵称不得超过20个字符！',
    nationReg: '请输入一个有效的民族！',
    idnumberReg: '请输入一个有效的身份证号！',
    positionReg: '请输入一个有效的职位！',
    realnameReg: '请输入一个有效的真实姓名！',
    ageReg: '请输入一个有效的年龄！',
    addressReg: '请输入一个有效的地址！',
    passwordReg: '请输入由6-20位，由字母、数字、!(),-.:*=?_{}！组成的密码！',
    phoneReg: '请输入一个有效的手机号码！',
    emailReg: '请输入一个有效的邮箱！',
    verifycodeReg: '验证码不能为空！',
    titleReg: '输入的标题有误，标题为1-50个字符！',
    descriptionReg: '输入的描述有误，描述为1-250个字符！',
    feedbackReg: '输入的反馈内容有误，反馈内容为1-500个字符！',
    numberReg: '请输入一个有效的学生证号！',
    teacherNumberReg: '请输入一个有效的教师证号！',
    schoolReg: '学校名称不能为空！',
    classReg: '班级名称不能为空！',
};
renderType();
var ue = UE.getEditor('para2');

function renderType(){
    var _type = $('.btn-blue').attr('data-type');
    $.ajax({
        url: '/getFlags',
        type: 'POST',
        dataType: 'json',
        data: {type: _type}
    })
    .done(function(data) {
        if(GetAttr(data,'meta.code') == 200 || GetAttr(data,'meta.code') == 201){
            var _html = '';
            var _data = data.data || [];
            for(var i = 0, l = _data.length; i < l; i++){
                _html += '<option value="'+_data[i]['Id']+'">'+_data[i]['Name']+'</option>';
            }
            $('.news-type').html(_html);
        }else{
            var content = GetAttr(data,'meta.msg') || '系统处理异常！';
            ErrorCallback(content);
        }
    })
    .fail(function() {
        ErrorCallback();
    });
}

$('body').on('change','.file',function(){
    var _this = $(this);
    var fileName = _this.val();
    if(_this.hasClass('xls-file')){   
        if(!Config.xlsReg.test(fileName)){
            _this.val('');
            ErrorCallback('请选择xls、xlsx格式的文件上传！');
            return false;
        }
    }else if(_this.hasClass('report-file')){
        if(!Config.reportReg.test(fileName)){
            ErrorCallback('请选择报告文件上传！');
            _this.val('');
            return false;
        }
    }else{
        if(!Config.imgReg.test(fileName)){
            _this.val('');
            ErrorCallback('请选择jpg、gif、png、jpeg格式的文件上传！');
            return false;
        }
    }
    var fileNameLen = fileName.split('\\');
    _this.closest('.file-wrap').find('span').eq(0).html(fileNameLen[fileNameLen.length-1]);
});

$('.btn-top-box').on('click', '.btn', function() {
    if($(this).hasClass('btn-blue')){
        return;
    }
    $('.btn-top-box').find('.btn').attr('class', 'btn btn-green');
    $(this).attr('class', 'btn btn-blue');
    renderType();
});

$('.btn-submit').on('click', function() {
    var _this = $(this);
    var $item = _this.closest('.item');
    var _data = {};
    var _existfileids = [];
    _data.title = $.trim($item.find('.msg-name').val()); //标题
    _data.description = $.trim($item.find('.user-text').val()); //描述
    _data.flag = $.trim($item.find('.news-type').val()); //类型
    _data.content = ue.getContent(); //内容
    _data.type = $('.btn-blue').attr('data-type'); //消息类型 

    if(!ValidateText({dom: $item.find('.msg-name'),reg: Config.titleReg,text: FailInfo.titleReg})){
        return false;
    }

    if(!ValidateText({dom: $item.find('.user-text'),reg: Config.descriptionReg,text: FailInfo.descriptionReg})){
        return false;
    }

    if(!$item.find('.file').val()){
        ErrorCallback('请上传图片！',3);
        return false;
    }

    if(!_data.content){
        ErrorCallback('请编辑内容！',3);
        return false;
    }
    OpenLoad();
    _this.closest('form').ajaxSubmit({
        type:'post',
        timeout : 100000,
        data: _data,
        url: '/save',
        success:function(data){
            CloseLoad();
            if(GetAttr(data,'meta.code') == 200 || GetAttr(data,'meta.code') == 201){
                Smallpop('消息新增成功！');
            }else{
                var content = GetAttr(data,'meta.msg') || '系统处理异常！';
                ErrorCallback(content);
            }
        },
        error:function(){
            CloseLoad();
            ErrorCallback();
        }
    });
});

/**
* 解决不支持placeholder属性
* */
function JPlaceHolder($container){
    this.container = $container || $('body');
    this.init();
}
JPlaceHolder.prototype = {
    _check : function(){
        return 'placeholder' in document.createElement('input');
    },
    init : function(){
        if(!this._check()){
            this.fix();
        }
    },
    fix : function(){
        this.container.find('input[placeholder],textarea[placeholder]').each(function(index, element) {
            var self = $(this), txt = self.attr('placeholder');
            self.val(txt);
            self.focusin(function(e) {
                var _this = $(this);
                if(_this.val() == txt){
                    _this.val('');  
                }
            }).focusout(function(e) {
                var _this = $(this);
                if(_this.val() == ''){
                    _this.val(txt)  
                }
            });
        });
    }
};

/**  
* 弹框 
* type: 0为自定义内容，1为成功提示，2为失败提示，3为信息提示
* */
function Popupbox(options){
    var options = $.extend({
            'title': '系统提示',
            'type' : 3,
            'yesFn': function(){},
            'yesClose': true, //点确定的时候是否关闭弹框
            'closeFn': null,
            'noFn': null,
            'content': '操作有误！'
        },options);
    this.init(options);
    new JPlaceHolder($('#pop'));
    this.eventHandler(options);
}
Popupbox.prototype = {
    init: function(options){
        if($('#pop').length > 0 || $('#mask').length > 0){
            $('#mask').remove();
            $('#pop').remove();
        }
        var _this = this;
        var popHtml = ['<div id="pop">'];
        popHtml.push('<div class="pop-header">');
        popHtml.push(options.title+'<a class="fr" href="javascript:void(0)"></a></div>');
        popHtml.push('<div class="pop-body">');
        if(options.type == 0){
            popHtml.push(options.content);
        }else{
            var iconClass = 'pop-icon-info';
            if(options.type == 1){
                iconClass = 'pop-icon-suc';
            }else if((options.type == 2)){
                iconClass = 'pop-icon-err';
            }
            popHtml.push('<div class="pop-content">');
            popHtml.push('<em class="'+iconClass+' fl"></em>');
            popHtml.push('<p>'+options.content+'</p></div>');
        }
        popHtml.push('</div>');
        popHtml.push('<div class="pop-footer">');
        popHtml.push('<div class="btn-box-right">');
        if(options.yesFn){
            popHtml.push('<button class="btn btn-blue btn-yes">确定</button>');
        }
        if(options.noFn){
            popHtml.push('<button class="btn btn-gray btn-no">取消</button>');
        }
        popHtml.push('</div></div></div>');
        $('body').append('<div id="mask"></div>');
        $('body').append(popHtml.join(''));
        _this.resizeDialog();
        _this.noScroll();
    },
    eventHandler: function(options){
        var _this = this;
        $('body').off('click.popYes').on('click.popYes','#pop .btn-yes',function(){
            options.yesClose && _this.hidePop();
            options.yesFn && options.yesFn();
        });
        $('body').off('click.popNo').on('click.popNo','#pop .btn-no',function(){
            _this.hidePop();
            options.noFn && options.noFn();

        });
        $('body').off('click.popClose').on('click.popClose','#pop .pop-header a',function(){
            _this.hidePop();
            options.closeFn && options.closeFn();
        });
    },
    noScroll: function(){
        //禁止弹出后 滚动屏幕
        var vendor = (/webkit/i).test(navigator.appVersion) ? 'webkit' :
            (/firefox/i).test(navigator.userAgent) ? 'Moz' :
                'opera' in window ? 'O' : '';
        var WHEEL_EV = ("ontouchstart" in window)?"touchmove" :(vendor == 'Moz' ? 'DOMMouseScroll' : 'mousewheel');
        $('#pop').on(WHEEL_EV,function(e){                 
            e.preventDefault();
            e.stopPropagation();                    
        })
        $('#mask').on(WHEEL_EV,function(e){                 
            e.preventDefault();
            e.stopPropagation();                    
        })
    },
    hidePop: function(){
        $('#pop').remove();
        $('#mask').remove();
    },
    resizeDialog: function(){
        var $pop = $('#pop');
        if($(window).width() < 430){
            $pop.css({'width':'96%'});
        }
        if($pop){
            var _width = $pop.width()/2;
            var _height = $pop.height()/2;
            $pop.css({'marginTop':-_height,'marginLeft':-_width});
        }
    }
}
function CloseLoad(){
    if($('#loading').length > 0 || $('#loadMask').length > 0){
        $('#loadMask').remove();
        $('#loading').remove();
    }
}
function Smallpop(text,fn,time){
    if($('#smallpop').length > 0){
        $('#smallpop').remove();
    }
    var $smallpop = $('<div style="display: none;" id="smallpop">'+text+'</div>');
    var time = time || 800;
    $smallpop.appendTo($('body'));
    $smallpop.fadeIn(200);
    window.setTimeout(function(){
        fn && fn();
        $smallpop.fadeOut(200,function(){
            $smallpop.remove();
        });
    },time);
}
function GetURLValueByKeyName(name,params) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    params = params || window.location.href.split('?')[1] || '';
    var r = params.match(reg);
    if (r != null) return unescape(r[2]); return "";
}
function GetAttr(obj,attrStr) {
    if(obj == undefined || attrStr == undefined){
        return "";
    }
    try{
        var result = eval("obj." + attrStr);
        if( result == undefined){
            return "";
        }
        return result;
    } catch(ex) {
        return "";
    }
}
function OpenLoad(){
    if($('#loading').length > 0 || $('#loadMask').length > 0){
        $('#loadMask').remove();
        $('#loading').remove();
    }
    $('body').append('<div id="loadMask"></div><div id="loading"><img class="load-icon" src="/static/images/loading.gif"/>正在努力加载中...</div>');
}
function CloseLoad(){
    if($('#loading').length > 0 || $('#loadMask').length > 0){
        $('#loadMask').remove();
        $('#loading').remove();
    }
}
function ValidateText(obj){
    var obj = $.extend({
        dom: '', //文本DOM对象
        reg: '', //验证规则
        text: '', //提示文字
        isPop: false //dom是否在弹框内
    },obj);
    if(obj.dom.attr('type') == 'password'){
        var txt = obj.dom.val();
    }else{
        var txt = $.trim(obj.dom.val());
    }
    if(obj.reg.test(txt)){
        return true;
    }
    if(obj.isPop){
        obj.dom.focus();
        obj.dom.closest('#pop').find('.pop-tip').html(obj.text).fadeIn();
    }else{
        ErrorCallback(obj.text,3,function(){
            obj.dom.focus();
        });
    }
    return false;
}
function ErrorCallback(prompt,type,fn){
    var prompt = prompt || '系统处理异常！';
    var type = type || 2;
    var fn = fn || function(){};
    new Popupbox({
        'title': '系统提示',
        'type': type,
        'content':prompt,
        'yesFn': fn
    });
}