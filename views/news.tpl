<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
    <meta name="viewport" content="width=device-width,initial-scale=1.0,maximum-scale=1.0,user-scalable=0">
    <style type="text/css">
        html, body, body div, span, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, abbr, address, cite, code, del, dfn, em, img, ins, kbd, q, samp, small, strong, sub, sup, var, b, i, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, figure, footer, header, menu, nav, section, time, mark, audio, video, details, summary {
            margin: 0;
            padding: 0;
            border: 0;
            font-size: 100%;
            font-weight: normal;
            vertical-align: baseline;
            background: transparent;
        }
        article, aside, figure, footer, header, nav, section, details, summary {display: block;}
        html {
            -ms-text-size-adjust: 100%;
            -webkit-text-size-adjust: 100%;
            box-sizing: border-box;
        }
        *,
        *:before,
        *:after {
            box-sizing: inherit;
        }
        img,
        object,
        embed {max-width: 100%;}
        html {overflow-y: scroll;}
        ul {list-style: none;}
        blockquote, q {quotes: none;}
        blockquote:before,
        blockquote:after,
        q:before,
        q:after {content: ''; content: none;}
        a {margin: 0; padding: 0; font-size: 100%; vertical-align: baseline; background: transparent;}
        del {text-decoration: line-through;}
        abbr[title], dfn[title] {border-bottom: 1px dotted #000; cursor: help;}
        table {border-collapse: separate; border-spacing: 0;}
        th {font-weight: bold; vertical-align: bottom;}
        td {font-weight: normal; vertical-align: top;}
        hr {display: block; height: 1px; border: 0; border-top: 1px solid #ccc; margin: 1em 0; padding: 0;}
        input, select {vertical-align: middle;}
        pre {
            white-space: pre; /* CSS2 */
            white-space: pre-wrap; /* CSS 2.1 */
            white-space: pre-line; /* CSS 3 (and 2.1 as well, actually) */
            word-wrap: break-word; /* IE */
        }
        input[type="radio"] {vertical-align: text-bottom;}
        input[type="checkbox"] {vertical-align: bottom;}
        select, input, textarea {font: 99% sans-serif;}
        table {font-size: inherit; font: 100%;}
        small {font-size: 85%;}
        strong {font-weight: bold;}
        td, td img {vertical-align: top;}/
        sub, sup {font-size: 75%; line-height: 0; position: relative;}
        sup {top: -0.5em;}
        sub {bottom: -0.25em;}
        pre, code, kbd, samp {font-family: monospace, sans-serif;}
        .clickable,
        label,
        input[type=button],
        input[type=submit],
        input[type=file],
        button {cursor: pointer;}
        button, input, select, textarea {margin: 0;}
        button,
        input[type=button] {width: auto; overflow: visible;}
        .clearfix:after { content: " "; display: block; clear: both; }

        body{
            line-height: 1.6;
            font-family: "Helvetica Neue",Helvetica,"Hiragino Sans GB","Microsoft YaHei",Arial,sans-serif;
            color: rgb(62, 62, 62);
            background-color: #f3f3f3;
        }
        #wraper{
            padding: 20px 15px 15px;
            font-size: 20px;
            background-color: #fff;
        }
        @media screen and (min-width: 1024px) {
            body{
                background-color: #fff;
            }
            #wraper {
                width: 740px;
                margin-left: auto;
                margin-right: auto;
            }
        }
        #header{
            margin-bottom: 15px;
            padding-bottom: 10px;
            border-bottom: 1px solid #e7e7eb;
        }
        #header h1{
            color: #000;
            line-height: 1.4;
            font-weight: 400;
            font-size: 24px;
        }
        #header .meta{
            padding-top: 15px;
            font-size: 16px;
            color: #8c8c8c;
        }
        @media (-webkit-min-device-pixel-ratio: 3) and (max-device-width: 736px) and (min-device-width: 414px) {
            #header h1 {
                font-size: 25px;
            }
        }
        #main img{
            display: block;
        }
        #main p,#main img{
            margin-bottom: 20px;
        }
        #main p>img{
            margin: 10px 0;
        }
    </style>
</head>
<body>
<div id="wraper">
    <div id="header">
        <h1>{{{.Mess.Title}}}</h1>
        <div class="meta">
            <span class="time">{{{.Mess.PublishTime}}}</span>&nbsp;&nbsp;<span class="category">{{{.Mess.MessFlagName}}}</span>
        </div>
    </div>
    <div id="main">
        {{{.Mess.Content}}}
    </div>
</div>
</body>
</html>