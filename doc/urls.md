# bookget
bookget 数字图书馆下载工具

#### 支持的数字图书馆URL格式：

原则上，以你能在浏览器中【在线阅读】书籍正文的URL为下载地址。

```
http://read.nlc.cn/allSearch/searchDetail?searchType=1002&showType=1&indexName=data_892&fid=411999021002
http://read.nlc.cn/OutOpenBook/OpenObjectBook?aid=403&bid=70621.0
https://babel.hathitrust.org/cgi/pt?id=uc1.c087423515&view=1up&seq=1&skin=2021
https://iiif.lib.harvard.edu/manifests/view/drs:53262215
https://iiif.lib.harvard.edu/manifests/drs:53954875
https://rmda.kulib.kyoto-u.ac.jp/item/rb00024956
http://kanji.zinbun.kyoto-u.ac.jp/db-machine/toho/ShiSanJingZhuShu/html/A002menu.html
https://www.loc.gov/item/2014514163/
https://catalog.princeton.edu/catalog/9940468523506421
https://dpul.princeton.edu/catalog/99915a8b423b596e47540e3feeee19b8
https://dl.ndl.go.jp/info:ndljp/pid/8929985
https://rbook.ncl.edu.tw/NCLSearch/Search/SearchDetail?item=422a7598bd0046aebf2684ae0f945d25fDcyODIz0&image=1&page=&whereString=&sourceWhereString=&SourceID=
https://emuseum.nich.go.jp/detail?content_base_id=100168&content_part_id=009&langId=zh&webView=
https://db2.sido.keio.ac.jp/kanseki/T_bib_frame.php?id=006754
http://shanben.ioc.u-tokyo.ac.jp/main_p.php?nu=C5613401&order=rn_no&no=00870
https://repository.lib.cuhk.edu.hk/sc/item/cuhk-412225#page/1/mode/2up
https://digital.bodleian.ox.ac.uk/objects/310cb04e-6bce-44e3-85b5-03417c9644a8/
https://www.digital.archives.go.jp/DAS/meta/listPhoto?LANG=default&BID=F1000000000000095447&ID=&NO=&TYPE=
https://archive.wul.waseda.ac.jp/kosho/ri08/ri08_01899/
http://dsr.nii.ac.jp/toyobunko/XI-6-A-16/V-1/
http://lod.nl.go.kr/page/CNTS-00076977176
https://kotenseki.nijl.ac.jp/biblio/100270332/viewer/1
https://kotenseki.nijl.ac.jp/biblio/100270332
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770&PHYSID=PHYS_0001
http://www.bl.uk/manuscripts/Viewer.aspx?ref=or_6814!1_fs001r
https://lbezone.ust.hk/bib/b1129168
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-1
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-173-1
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-(1-90)
https://khirin-a.rekihaku.ac.jp/sohankanjo/h-173-(1-61)
https://www.library.yonezawa.yamagata.jp/dg/AA001_view.html
https://www.library.yonezawa.yamagata.jp/dg/AA002_view.html
https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1
https://www.iiif.ku-orcas.kansai-u.ac.jp/books/210185040#?page=1
http://221.13.137.120:8090/productshow.php?cid=4&id=112
https://oyjy.wzlib.cn/resource/?id=61e4c764505415b2e6921e5e
https://oyjy.wzlib.cn/resource/?id=62c56bb357de1ef36b1f5614
https://ostasien.digitale-sammlungen.de/view/bsb11129280/1
https://searchworks.stanford.edu/view/4182111   
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=4&book_id=18269  
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=1&book_id=18017
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-C8KN?wc=3X27-MNY%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022419701&cc=1787988
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-CRG9?wc=3X2Q-BZ7%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022421801&cc=1787988
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9S9-WS92?view=explore&groupId=M94X-6HR
https://www.familysearch.org/records/images/image-details?rmsId=M94F-78D&jiapuOnly=true&surname=%E6%9C%B1&place=2013&showUnknown=true&page=1&pageSize=100&imageIndex=0
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultDetail?BookId=GZDD022601004
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultSummary?bookid=GZDD022601004&filename=GZDD022601004#
http://idp.nlc.cn/database/oo_scroll_h.a4d?uid=47355195088;recnum=0;index=2
```


1.中国国家图书馆：
```
整书多册URL：http://read.nlc.cn/allSearch/searchDetail?searchType=1002&showType=1&indexName=data_892&fid=411999021002
或者单册URL：http://read.nlc.cn/OutOpenBook/OpenObjectBook?aid=403&bid=70621.0
```

2.  hathitrust 数字图书馆-图书单册URL 

```
https://babel.hathitrust.org/cgi/pt?id=uc1.c087423515&view=1up&seq=1&skin=2021
```


3.  哈佛大学图书馆-图书在线阅读（分享）URL
```
https://iiif.lib.harvard.edu/manifests/view/drs:53262215
```


4.  日本京东大学图书馆-图书在线阅读URL
```
https://rmda.kulib.kyoto-u.ac.jp/item/rb00024956
```


5.  日本京都大学人文科学研究所-图书在线阅读URL
```
http://kanji.zinbun.kyoto-u.ac.jp/db-machine/toho/ShiSanJingZhuShu/html/A002menu.html
```


6. 美国国会图书馆    
   注：中国大陆访问此网站需自备海外VPN，免VPN方法需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
https://www.loc.gov/item/2014514163/
```


7. 普林斯顿大学图书馆 – 图书在线阅读URL
```
https://catalog.princeton.edu/catalog/9940468523506421
https://dpul.princeton.edu/catalog/99915a8b423b596e47540e3feeee19b8
```


8.  日本国立国会图书馆 – 部分图书在线阅读URL（其它的可以手动打印下载）
```
https://dl.ndl.go.jp/info:ndljp/pid/8929985
```


9.  中国台北图书馆古典与特藏文献 –（白天很慢，可夜间或清晨下载）
```
https://rbook.ncl.edu.tw/NCLSearch/Search/SearchDetail?item=422a7598bd0046aebf2684ae0f945d25fDcyODIz0&image=1&page=&whereString=&sourceWhereString=&SourceID=
```


10.  日本E国宝 – 画册在线阅读URL（部分单图有误，暂未修复）
```
https://emuseum.nich.go.jp/detail?content_base_id=100168&content_part_id=009&langId=zh&webView=
```


11.  日本宫内厅书陵部 – 图书在线阅读URL
```
https://db2.sido.keio.ac.jp/kanseki/T_bib_frame.php?id=006754
```


12.  日本东京大学东洋文化研究所 汉籍善本 – 图书在线阅读URL   
```
http://shanben.ioc.u-tokyo.ac.jp/main_p.php?nu=C5613401&order=rn_no&no=00870
```


13.  中国香港中文大学图书馆 – 图书在线阅读URL（需自备VPN，从海外访问）
```
https://repository.lib.cuhk.edu.hk/sc/item/cuhk-412225#page/1/mode/2up
```


14.  牛津大学博德利图书馆 – 图书在线阅读URL
```
https://digital.bodleian.ox.ac.uk/objects/310cb04e-6bce-44e3-85b5-03417c9644a8/
```


15.  日本国立公文书馆（内库文库） - 图书在线阅读URL
```
https://www.digital.archives.go.jp/DAS/meta/listPhoto?LANG=default&BID=F1000000000000095447&ID=&NO=&TYPE=
```


16.  日本早稻田大学图书馆 – 图书在线阅读URL
```
https://archive.wul.waseda.ac.jp/kosho/ri08/ri08_01899/
```


17.  日本东洋文库（丝绸之路项目） - 图书在线阅读URL
```
http://dsr.nii.ac.jp/toyobunko/XI-6-A-16/V-1/
```


18.  韩国国家图书馆 （必须[参考pdf文档](/doc/pdf/03.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E9%9F%A9%E5%9B%BD%E5%9B%BE%E4%B9%A6%E9%A6%86%E5%9B%BE%E4%B9%A6.pdf)）
```
http://lod.nl.go.kr/page/CNTS-00076977176
```
注：请使用v0.2.6版。新版不再支持。

19.  新日本古典籍综合数据库（[参考pdf文档](/doc/pdf/04.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E6%96%B0%E6%97%A5%E6%9C%AC%E5%8F%A4%E5%85%B8%E5%9B%BE%E4%B9%A6.pdf)）
```
https://kotenseki.nijl.ac.jp/biblio/100270332/viewer/1
https://kotenseki.nijl.ac.jp/biblio/100270332
```


20.  德国柏林图书馆URL
```
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770
https://digital.staatsbibliothek-berlin.de/werkansicht?PPN=PPN3343671770&PHYSID=PHYS_0001
```

21.  英国图书馆URL（只生成dezoomify-rs.urls文件，生成后，请双击它下载）
```
http://www.bl.uk/manuscripts/Viewer.aspx?ref=or_6814!1_fs001r
```

22.  中国香港科技大学图书馆URL
```
https://lbezone.ust.hk/bib/b1129168
```

23.  中国台北故宫博物院-善本古籍URL （必须[参考PDF文档](/doc/pdf/05.%E4%BD%BF%E7%94%A8bookget%E4%B8%8B%E8%BD%BD%E5%8F%B0%E5%8C%97%E6%95%85%E5%AE%AB%E5%8D%9A%E7%89%A9%E9%99%A2%E5%96%84%E6%9C%AC%E5%8F%A4%E7%B1%8D.pdf)）

24.  日本国立历史民俗博物馆
```
单册URL:
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-1
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-173-1

多册URL，使用和“批量下载”相同格式，但是无需修改config.ini中配置。
如：第1-9册，第10-90册。用圆括号包围数字。   

https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-(1-90)
https://khirin-a.rekihaku.ac.jp/sohankanjo/h-173-(1-61)
```

25.  日本本市立米泽图书馆
```
https://www.library.yonezawa.yamagata.jp/dg/AA001_view.html
https://www.library.yonezawa.yamagata.jp/dg/AA002_view.html
```

26.  日本庆应义塾大学图书馆
```
https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1
```

27.  日本关西大学图书馆
```
https://www.iiif.ku-orcas.kansai-u.ac.jp/books/210185040#?page=1
```

28.  中国河南省洛阳市图书馆   
```
http://221.13.137.120:8090/productshow.php?cid=4&id=112
```
29.  中国浙江省温州市图书馆 - 瓯越记忆(自动下载相关资源分卷分册)   
```
https://oyjy.wzlib.cn/resource/?id=61e4c764505415b2e6921e5e
https://oyjy.wzlib.cn/resource/?id=62c56bb357de1ef36b1f5614
```
30.  巴伐利亚州立图书馆
```
https://ostasien.digitale-sammlungen.de/view/bsb11129280/1
```
31. 斯坦福大学图书馆
```
https://searchworks.stanford.edu/view/4182111   
```
32.  中国广东省深圳市图书馆-古籍
```
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=4&book_id=18269  
https://yun.szlib.org.cn/stgj2021/srchshowbook?type=1&book_id=18017
```

33. [familysearch.org 中國族譜收藏 1239-2014年](https://www.familysearch.org/search/collection/1787988)   
    注：此站点需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-C8KN?wc=3X27-MNY%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022419701&cc=1787988
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9SM-CRG9?wc=3X2Q-BZ7%3A1022211401%2C1021934502%2C1021937102%2C1021937602%2C1022421801&cc=1787988
```
[familysearch.org 家譜圖像](https://www.familysearch.org/records/images/)
```
https://www.familysearch.org/ark:/61903/3:1:3QS7-L9S9-WS92?view=explore&groupId=M94X-6HR
https://www.familysearch.org/records/images/image-details?rmsId=M94F-78D&jiapuOnly=true&surname=%E6%9C%B1&place=2013&showUnknown=true&page=1&pageSize=100&imageIndex=0
```
34. 中国广东省广州大典(http://gzdd.gzlib.gov.cn/Hrcanton/)    
    注：此站点需要cookie.txt，方法参考：[cookie.md](cookie.md)
```
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultDetail?BookId=GZDD022601004
http://gzdd.gzlib.gov.cn/Hrcanton/Search/ResultSummary?bookid=GZDD022601004&filename=GZDD022601004#
```

35. 國際敦煌項目(http://idp.nlc.cn/)    
    注：需先搜索关键词，例如`8210`，并且URL中含有`uid=xxxx`，短时间内有效，请在搜索结果后尽快下载。
```
http://idp.nlc.cn/database/oo_scroll_h.a4d?uid=47355195088;recnum=0;index=2
```