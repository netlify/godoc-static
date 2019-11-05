package static

var GodocTemplate string = "<!DOCTYPE\x20html>\x0a<html>\x0a<head>\x0a<meta\x20http-equiv=\"Content-Type\"\x20content=\"text/html;\x20charset=utf-8\">\x0a<meta\x20name=\"viewport\"\x20content=\"width=device-width,\x20initial-scale=1\">\x0a<meta\x20name=\"theme-color\"\x20content=\"#375EAB\">\x0a{{with\x20.Tabtitle}}\x0a\x20\x20<title>{{html\x20.}}\x20-\x20Godocs</title>\x0a{{else}}\x0a\x20\x20<title>Godocs</title>\x0a{{end}}\x0a<link\x20type=\"text/css\"\x20rel=\"stylesheet\"\x20href=\"/lib/godoc/style.css\">\x0a{{if\x20.TreeView}}\x0a<link\x20rel=\"stylesheet\"\x20href=\"/lib/godoc/jquery.treeview.css\">\x0a{{end}}\x0a<script>window.initFuncs\x20=\x20[];</script>\x0a<script\x20src=\"/lib/godoc/jquery.js\"\x20defer></script>\x0a{{if\x20.TreeView}}\x0a<script\x20src=\"/lib/godoc/jquery.treeview.js\"\x20defer></script>\x0a<script\x20src=\"/lib/godoc/jquery.treeview.edit.js\"\x20defer></script>\x0a{{end}}\x0a\x0a{{if\x20.Playground}}\x0a<script\x20src=\"/lib/godoc/playground.js\"\x20defer></script>\x0a{{end}}\x0a{{with\x20.Version}}<script>var\x20goVersion\x20=\x20{{printf\x20\"%q\"\x20.}};</script>{{end}}\x0a<script\x20src=\"/lib/godoc/godocs.js\"\x20defer></script>\x0a</head>\x0a<body>\x0a\x0a<div\x20id='lowframe'\x20style=\"position:\x20fixed;\x20bottom:\x200;\x20left:\x200;\x20height:\x200;\x20width:\x20100%;\x20border-top:\x20thin\x20solid\x20grey;\x20background-color:\x20white;\x20overflow:\x20auto;\">\x0a...\x0a</div><!--\x20#lowframe\x20-->\x0a\x0a<div\x20id=\"topbar\"{{if\x20.Title}}\x20class=\"wide\"{{end}}><div\x20class=\"container\">\x0a<div\x20class=\"top-heading\"\x20id=\"heading-wide\"><a\x20href=\"/\">Netlify\x20Godocs</a></div>\x0a<div\x20class=\"top-heading\"\x20id=\"heading-narrow\"><a\x20href=\"/\">Go</a></div>\x0a<a\x20href=\"#\"\x20id=\"menu-button\"><span\x20id=\"menu-button-arrow\">&#9661;</span></a>\x0a<div\x20id=\"menu\">\x0a{{if\x20(and\x20.Playground\x20.Title)}}\x0a<a\x20id=\"playgroundButton\"\x20href=\"http://play.golang.org/\"\x20title=\"Show\x20Go\x20Playground\">Play</a>\x0a{{end}}\x0a</div>\x0a\x0a</div></div>\x0a\x0a{{if\x20.Playground}}\x0a<div\x20id=\"playground\"\x20class=\"play\">\x0a\x09<div\x20class=\"input\"><textarea\x20class=\"code\"\x20spellcheck=\"false\">package\x20main\x0a\x0aimport\x20\"fmt\"\x0a\x0afunc\x20main()\x20{\x0a\x09fmt.Println(\"Hello,\x20\xe4\xb8\x96\xe7\x95\x8c\")\x0a}</textarea></div>\x0a\x09<div\x20class=\"output\"></div>\x0a\x09<div\x20class=\"buttons\">\x0a\x09\x09<a\x20class=\"run\"\x20title=\"Run\x20this\x20code\x20[shift-enter]\">Run</a>\x0a\x09\x09<a\x20class=\"fmt\"\x20title=\"Format\x20this\x20code\">Format</a>\x0a\x09\x09{{if\x20not\x20$.GoogleCN}}\x0a\x09\x09<a\x20class=\"share\"\x20title=\"Share\x20this\x20code\">Share</a>\x0a\x09\x09{{end}}\x0a\x09</div>\x0a</div>\x0a{{end}}\x0a\x0a<div\x20id=\"page\"{{if\x20.Title}}\x20class=\"wide\"{{end}}>\x0a<div\x20class=\"container\">\x0a\x0a{{if\x20or\x20.Title\x20.SrcPath}}\x0a\x20\x20<h1>\x0a\x20\x20\x20\x20{{html\x20.Title}}\x0a\x20\x20\x20\x20{{html\x20.SrcPath\x20|\x20srcBreadcrumb}}\x0a\x20\x20</h1>\x0a{{end}}\x0a\x0a{{with\x20.Subtitle}}\x0a\x20\x20<h2>{{html\x20.}}</h2>\x0a{{end}}\x0a\x0a{{with\x20.SrcPath}}\x0a\x20\x20<h2>\x0a\x20\x20\x20\x20Documentation:\x20{{html\x20.\x20|\x20srcToPkgLink}}\x0a\x20\x20</h2>\x0a{{end}}\x0a\x0a{{/*\x20The\x20Table\x20of\x20Contents\x20is\x20automatically\x20inserted\x20in\x20this\x20<div>.\x0a\x20\x20\x20\x20\x20Do\x20not\x20delete\x20this\x20<div>.\x20*/}}\x0a<div\x20id=\"nav\"></div>\x0a\x0a{{/*\x20Body\x20is\x20HTML-escaped\x20elsewhere\x20*/}}\x0a{{printf\x20\"%s\"\x20.Body}}\x0a\x0a</div><!--\x20.container\x20-->\x0a</div><!--\x20#page\x20-->\x0a</body>\x0a</html>\x0a"