<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1603680752190" ID="ID_631656097" MODIFIED="1603681191904" TEXT="golang">
<node CREATED="1603681926583" ID="ID_1019289924" MODIFIED="1603681936322" POSITION="right" TEXT="learn">
<node CREATED="1603681936325" ID="ID_305119485" MODIFIED="1603681944903" TEXT="basic">
<node CREATED="1603680771357" ID="ID_1488548989" MODIFIED="1603680775484" TEXT="introduction">
<node CREATED="1603766287982" ID="ID_1282373928" MODIFIED="1603793142801" TEXT="https://www.liwenzhou.com/posts/Go/go_menu/">
<node CREATED="1603681727781" ID="ID_1900463217" MODIFIED="1603681730189" TEXT="https://www.bilibili.com/video/BV19k4y1k7U9"/>
</node>
<node CREATED="1603680776367" ID="ID_857625971" MODIFIED="1603680782622" TEXT="what is go">
<node CREATED="1603680787191" ID="ID_471060761" MODIFIED="1603680825671" TEXT="from google in 2007"/>
<node CREATED="1603680818474" ID="ID_1044701241" MODIFIED="1603680859936" TEXT="compile"/>
<node CREATED="1603680947181" ID="ID_673042342" MODIFIED="1603680970882" TEXT="first dual-core processor from 2005"/>
</node>
<node CREATED="1603681709663" ID="ID_314488805" MODIFIED="1603681714030" TEXT="reference">
<node CREATED="1603681714037" ID="ID_1547301602" MODIFIED="1603681715414" TEXT="https://www.liwenzhou.com/"/>
</node>
</node>
<node CREATED="1603681944904" ID="ID_716923861" MODIFIED="1603681948744" TEXT="installation">
<node CREATED="1603683169380" ID="ID_1046665188" MODIFIED="1603683172774" TEXT="https://www.liwenzhou.com/posts/Go/install_go_dev/"/>
<node CREATED="1603681948745" ID="ID_440484750" MODIFIED="1603681950156" TEXT="https://gomirrors.org/">
<node CREATED="1603683024295" ID="ID_1593088755" MODIFIED="1603683025303" TEXT="https://dl.google.com/go/go1.15.3.windows-amd64.zip "/>
</node>
<node CREATED="1603683027338" ID="ID_1788043532" MODIFIED="1603683036507" TEXT="environment variables">
<node CREATED="1603683795959" ID="ID_46908109" MODIFIED="1603683810959" TEXT="system variables">
<node CREATED="1603683036515" ID="ID_544727202" MODIFIED="1603683149155">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      GOBOOT=D:\program\go
    </p>
    <p>
      Path=;%GOROOT%\bin
    </p>
    <p>
      # open new cmd
    </p>
    <p>
      go env
    </p>
  </body>
</html></richcontent>
</node>
</node>
<node CREATED="1603683816614" ID="ID_1549554467" MODIFIED="1603683825977" TEXT="user variables">
<node CREATED="1603683333130" ID="ID_939969610" MODIFIED="1603683984772">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      GOPATH=D:\project\go
    </p>
    <p>
      Path=;%GOPATH%\bin
    </p>
  </body>
</html></richcontent>
</node>
</node>
<node CREATED="1603683753222" ID="ID_869135888" MODIFIED="1603683830537">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      $ go env
    </p>
    <p>
      set GOPATH=D:\project\go
    </p>
    <p>
      set GOROOT=D:\program\go
    </p>
  </body>
</html></richcontent>
</node>
</node>
<node CREATED="1603697239262" ID="ID_959713400" MODIFIED="1603697244530" TEXT="command">
<node CREATED="1603697244532" ID="ID_383587170" MODIFIED="1603697248430" TEXT="go run"/>
<node CREATED="1603697249465" ID="ID_1117416702" MODIFIED="1603697253694" TEXT="go build"/>
<node CREATED="1603697254689" ID="ID_1304369973" MODIFIED="1603697257166" TEXT="go install"/>
<node CREATED="1603697258465" ID="ID_1801364179" MODIFIED="1603697260727" TEXT="go fmt"/>
<node CREATED="1603703539603" ID="ID_888086229" MODIFIED="1603703540975" TEXT="go fmt main.go &amp;&amp; go run main.go"/>
</node>
<node CREATED="1603852981118" ID="ID_257629358" MODIFIED="1603856548914" TEXT="syntax">
<node CREATED="1603852987749" ID="ID_1211807221" MODIFIED="1603852991024" TEXT="closure">
<node CREATED="1603852991025" ID="ID_881248922" MODIFIED="1603855492592" TEXT="internal function closure external function&apos;s variable reference">
<node CREATED="1603853804640" ID="ID_51544694" MODIFIED="1603853845318">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <div style="color: #d4d4d4; background-color: #1e1e1e; font-family: Consolas, Courier New, monospace; font-weight: normal; font-size: 14px; line-height: 19px; white-space: pre">
      <div>
        <font color="#569cd6">func</font><font color="#d4d4d4">&#160;</font><font color="#dcdcaa">cloo</font><font color="#d4d4d4">(name&#160;</font><font color="#4ec9b0">string</font><font color="#d4d4d4">)&#160;</font><font color="#569cd6">func</font><font color="#d4d4d4">()&#160;{</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;</font><font color="#c586c0">return</font><font color="#d4d4d4">&#160;</font><font color="#569cd6">func</font><font color="#d4d4d4">()&#160;{</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;fmt.</font><font color="#dcdcaa">Println</font><font color="#d4d4d4">(</font><font color="#ce9178">&quot;</font><font color="#d7ba7d">\n</font><font color="#ce9178">closure&quot;</font><font color="#d4d4d4">,&#160;name)</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;}</font>
      </div>
      <div>
        <font color="#d4d4d4">}</font>
      </div>
    </div>
  </body>
</html></richcontent>
</node>
</node>
</node>
<node CREATED="1603856168283" ID="ID_914127082" MODIFIED="1603856186269" TEXT="defer recover() panic">
<node CREATED="1603856186271" ID="ID_716604064" MODIFIED="1603856193552" TEXT="try catch"/>
<node CREATED="1603856362724" ID="ID_180037255" MODIFIED="1603856366556">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <div style="color: #d4d4d4; background-color: #1e1e1e; font-family: Consolas, Courier New, monospace; font-weight: normal; font-size: 14px; line-height: 19px; white-space: pre">
      <div>
        <font color="#569cd6">func</font><font color="#d4d4d4">&#160;</font><font color="#dcdcaa">panicDeferRecover</font><font color="#d4d4d4">()&#160;{</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;</font><font color="#c586c0">defer</font><font color="#d4d4d4">&#160;</font><font color="#569cd6">func</font><font color="#d4d4d4">()&#160;{</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;</font><font color="#9cdcfe">err</font><font color="#d4d4d4">&#160;:=&#160;</font><font color="#dcdcaa">recover</font><font color="#d4d4d4">()</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;</font><font color="#c586c0">if</font><font color="#d4d4d4">&#160;err&#160;!=&#160;</font><font color="#569cd6">nil</font><font color="#d4d4d4">&#160;{</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;fmt.</font><font color="#dcdcaa">Println</font><font color="#d4d4d4">(</font><font color="#ce9178">&quot;err&quot;</font><font color="#d4d4d4">,&#160;err)</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;&#160;&#160;&#160;&#160;}</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;}()</font>
      </div>
      <div>
        <font color="#d4d4d4">&#160;&#160;&#160;&#160;</font><font color="#dcdcaa">panic</font><font color="#d4d4d4">(</font><font color="#ce9178">&quot;panicDefer&quot;</font><font color="#d4d4d4">)</font>
      </div>
      <div>
        <font color="#d4d4d4">}</font>
      </div>
    </div>
  </body>
</html></richcontent>
</node>
</node>
<node CREATED="1603856436007" ID="ID_1070098273" MODIFIED="1603856445842" TEXT="pointer">
<node CREATED="1603856445846" ID="ID_1760292647" MODIFIED="1603856478772" TEXT="safe pointer"/>
<node CREATED="1603856497296" ID="ID_39785255" MODIFIED="1603856511897" TEXT="operator">
<node CREATED="1603856511898" ID="ID_935758613" MODIFIED="1603856516150" TEXT="&amp;"/>
<node CREATED="1603856519839" ID="ID_535532433" MODIFIED="1603856524092" TEXT="*"/>
</node>
<node CREATED="1603856695085" ID="ID_76540252" MODIFIED="1603856721906" TEXT="value type -&gt; pointer">
<node CREATED="1603856697497" ID="ID_1057438112" MODIFIED="1603856744374" TEXT="int, float, bool, array, string, struct"/>
</node>
</node>
<node CREATED="1603856548916" ID="ID_1876256517" MODIFIED="1603856552623" TEXT="func">
<node CREATED="1603856552624" ID="ID_1843021524" MODIFIED="1603856591772" TEXT="all variables is value parameter (copy)"/>
</node>
<node CREATED="1603865601849" ID="ID_1170435144" MODIFIED="1603865610350" TEXT="memory">
<node CREATED="1603865612719" ID="ID_460956138" MODIFIED="1603865614893" TEXT="new">
<node CREATED="1603865653905" ID="ID_450959133" MODIFIED="1603865658904" TEXT="return pointer"/>
</node>
<node CREATED="1603865616849" ID="ID_151323477" MODIFIED="1603865619576" TEXT="make">
<node CREATED="1603865619577" ID="ID_626723739" MODIFIED="1603865884255" TEXT="slice, map, channel"/>
</node>
</node>
<node CREATED="1603874639196" ID="ID_684555717" MODIFIED="1603874642616" TEXT="struct">
<node CREATED="1603874642617" ID="ID_215677360" MODIFIED="1603874647467" TEXT="receiver">
<node CREATED="1603874647468" ID="ID_736294522" MODIFIED="1603874650438" TEXT="value"/>
<node CREATED="1603874651537" ID="ID_1958535311" MODIFIED="1603874659062" TEXT="pointer"/>
</node>
<node CREATED="1603942573654" ID="ID_1973676699" MODIFIED="1603942586672" TEXT="privilege">
<node CREATED="1603942586673" ID="ID_253454430" MODIFIED="1603942591063" TEXT="upper">
<node CREATED="1603942591985" ID="ID_1371514531" MODIFIED="1603942596118" TEXT="public"/>
</node>
<node CREATED="1603942597370" ID="ID_111046672" MODIFIED="1603942600146" TEXT="lower">
<node CREATED="1603942600146" ID="ID_1761229085" MODIFIED="1603942601839" TEXT="private"/>
</node>
</node>
</node>
</node>
</node>
</node>
</node>
</node>
</map>
