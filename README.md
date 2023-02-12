# bookget

bookget 数字图书馆下载工具   

鸣谢：
[书格shuge.org](https://new.shuge.org/) 有品格的数字古籍图书馆。    
![](https://new.shuge.org/wp-content/themes/artview/images/layout/logo.png)

### 支持的数字图书馆：
1. [中国国家图书馆](http://read.nlc.cn/thematDataSearch/toGujiIndex)
2. [哈佛大学图书馆](https://hollis.harvard.edu/) [或哈佛燕京图书馆藏](https://gj.library.sh.cn/org/harvard)
3. [中国台北图书馆](http://rbook.ncl.edu.tw/)
4. [hathitrust数字图书馆](https://www.hathitrust.org/)
5. [普林斯顿大学图书馆](https://library.princeton.edu/)
6. [日本京都大学图书馆](https://rmda.kulib.kyoto-u.ac.jp/)
7. [美国国会图书馆](https://www.loc.gov/collections/chinese-rare-books/)
8. [日本国立国会图书馆](http://dl.ndl.go.jp/)
9. [日本E国宝e-Museum]( https://emuseum.nich.go.jp/)
10. [日本宫内厅书陵部](https://db2.sido.keio.ac.jp/kanseki/T_bib_search.php)
11. [日本东京大学东洋文化研究所](http://shanben.ioc.u-tokyo.ac.jp/list.php)
12. [中国香港中文大学图书馆](https://repository.lib.cuhk.edu.hk/sc/collection)
13. [牛津大学博德利图书馆](https://digital.bodleian.ox.ac.uk/collections/chinese-digitization-project/)
14. [日本国立公文书馆（内库文库）](https://www.digital.archives.go.jp/)
15. [日本东洋文库]( http://dsr.nii.ac.jp/toyobunko/index.html.ja)
16. [日本早稻田大学图书馆](https://www.wul.waseda.ac.jp/kotenseki/search.php)
17. [韩国国家图书馆](https://www.dlibrary.go.kr/) [或开放数据](https://lod.nl.go.kr/) 
    (注：请使用v0.2.6版。新版不再支持。)
18. [新日本古典籍综合数据库](https://kotenseki.nijl.ac.jp/)
19. [德国柏林国立图书馆](https://digital.staatsbibliothek-berlin.de)
20. [日本京都大学人文科学研究所 - 东方学数字图书博物馆](http://kanji.zinbun.kyoto-u.ac.jp/db-machine/toho/html/top.html)
21. [英国图书馆（藏有手稿本）](http://www.bl.uk/manuscripts/)
22. [中国香港科技大学图书馆](https://lbezone.ust.hk/)
23. [中国台北故宫博物院 – 善本古籍 ](https://rbk-doc.npm.edu.tw/)
24. [日本国立历史民俗博物馆](https://khirin-a.rekihaku.ac.jp/)
25. [日本本市立米泽图书馆](https://www.library.yonezawa.yamagata.jp/dg/zen.html)
26. [日本庆应义塾大学图书馆](https://dcollections.lib.keio.ac.jp/ja/kanseki)
27. [日本关西大学图书馆](https://www.iiif.ku-orcas.kansai-u.ac.jp/books)
28. [中国河南省洛阳市图书馆](http://221.13.137.120:8090/index.php)
29. [中国浙江省温州市图书馆-瓯越记忆](https://oyjy.wzlib.cn/pdf/)
30. [巴伐利亚州立图书馆](https://ostasien.digitale-sammlungen.de/)
31. [斯坦福大学图书馆](https://searchworks.stanford.edu/?f%5Baccess_facet%5D%5B%5D=Online&f%5Bbuilding_facet%5D%5B%5D=East+Asia&f%5Bformat_main_ssim%5D%5B%5D=Book&f%5Blanguage%5D%5B%5D=Chinese&utf8=%E2%9C%93)
32. [中国广东省深圳市图书馆-古籍](https://yun.szlib.org.cn/stgj2021/)
33. [familysearch.org 中國族譜收藏 1239-2014年](https://www.familysearch.org/search/collection/1787988)   
    [familysearch.org 家譜圖像](https://www.familysearch.org/records/images/)
34. [中国广东省广州大典](http://gzdd.gzlib.gov.cn/Hrcanton/)
35. [國際敦煌項目](http://idp.nlc.cn/)

## 用户手册
请参见以下文档：
1. [支持的URL格式](/doc/urls.md)
2. [IIIF自动检测下载](/doc/iiif.md)
3. [批量http下载](/doc/http.md)
4. [高级：自定义用户cookie ](/doc/cookie.md)
5. [旧版：PDF手册](/doc/pdf/) 适用于v0.2.6及更低版本。

### 下载 *bookget*
第一次使用，请按以下步骤操作。

1. 打开 [最新正式版网页](https://github.com/deweizhu/bookget/releases/latest), 下载匹配你操作系统的版本 (Windows, MacOS, 或 Linux),
2. 解压缩到电脑中任意文件夹下。
3. 双击运行，并按提示输入URL。（例如：欽定古今圖書集成 - 中国国家图书馆）。
```
$ bookget
Enter an URL:
-> http://read.nlc.cn/allSearch/searchDetail?searchType=1002&showType=1&indexName=data_892&fid=411999021002
```
4. 【可选】把 bookget 放到 C:\windows 目录下（Linux用户是 /usr/local/bin 或 /usr/bin/目录）。   
   在终端下输入命令：`bookget "URL"` （推荐用双引号包含网址）,按回车键开始下载。   
   Windows 终端：cmd / PowerShell   
   Linux / MacOS终端：bash / sh / zsh
```
$ bookget "http://read.nlc.cn/allSearch/searchDetail?searchType=1002&showType=1&indexName=data_892&fid=411999021002"
```
5. 【可选】批量下载多个URL。在终端内输入以下命令：
```
$ bookget -i urls.txt
```
提示：urls.txt可以是任意文件名，内容是要下载的图书URL，一行一个URL，回车换行。

6. 【可选】带上cookie下载：
```
$ bookget -c cookie.txt [URL]
```
### 支持的更多参数

```
$ bookget -h
Usage: bookget [OPTION]... [URL]...
  -a int
        自动检测下载URL。可选值[0|1|2]，;0=默认;
        1=通用批量下载（类似IDM、迅雷）;
        2= IIIF manifest.json 自动检测下载图片
  -c string
        指定cookie.txt文件路径
  -cdn int
        使用CDN加速，可选值[0|1]。0=否，1=是。仅对 www.loc.gov 有效。
  -ext string
        指定文件扩展名[.jpg|.tif|.png]等
  -fn int
        保存文件名规则。可选值[0|1]。0=中文名，1=数字名。仅对 read.nlc.cn 有效。 (default 1)
  -h    显示帮助
  -i string
        下载的URLs，指定任意本地文件，例如：urls.txt
  -mp int
        合并PDF文件下载，可选值[0|1]。0=否，1=是。仅对 rbk-doc.npm.edu.tw 有效。
  -n uint
        最大并发连接数 (default 16)
  -o string
        下载保存到目录 (default "D:/bookget/bookget")
  -rs string
        自定义dezoomify-rs路径，例如：D:/bookget/dezoomify-rs (default "dezoomify-rs")
  -seq int
        图书起始页面数字
  -ua string
        user-agent (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0")
  -v    显示版本
  -w int
        指定图片宽度像素。推荐2400，若>6400为最大图 (default 7000)
```


### 批量下载
`bookget -i urls.txt -c .\cookie.txt -a 1 -n 1 -ext ".jpg"`    
更多参数，请使用 bookget -h 查看。

urls.txt内容如下：   
在urls.txt文件中，毎行一个URL，回车换行，可以有多个URL。
```
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=1&page=(1-155)&twoThreeYn=N
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=2&page=(1-163)&twoThreeYn=N
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=3&page=(1-161)&twoThreeYn=N
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=4&page=(1-163)&twoThreeYn=N
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=5&page=(1-167)&twoThreeYn=N
http://viewer.nl.go.kr:8080/nlmivs/view_image.jsp?cno=CNTS-00047981911&vol=6&page=(1-135)&twoThreeYn=N
```

cookie.txt 格式如下：
```
Cookie: WMONID=soB981Rm1Zd; _ga=GA1.3.87528781.1649496227; PCID=f3195068-16ea-8747-eedd-b37cf8523975-1649496227656; _INSIGHT_CK_1101=a658ca0653f5817a32a1b3a6942409e8_96227|1cbbd600ff48120ce10fed8a58ea4686_80164:1650282843000; JSESSIONID="0cfPybFlA0z2qRiy8Fr7sJCtdJooLnY8oACN62iv.VWWAS1:tv-1"; _gid=GA1.3.1049050692.1659041876
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Safari/537.36 Edg/103.0.1264.71
```