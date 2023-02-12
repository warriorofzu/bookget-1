# bookget
bookget 数字图书馆下载工具

#### IIIF自动检测
IIIF是一种业界标准，凡使用此标准的网站，都有提供IIIF Manifest 链接，如哈佛大学、牛津大学等图书馆。适用性更广，理论上所有支持IIIF的图书馆都可以下载。

#### 使用方法：
提示：自v0.2.7版开始已自动识别URL，无需修改配置文件或命令参数。     

复制包含IIIF Manifest链接的网页URL，或者直接复制IIIF Manifest的URL，
粘贴到urls.txt中保存，运行bookget即可自动检测识别。

如果检测成功，你会看到类似以下内容的提示。
```
2022/03/27 08:00:54 Auto Detect 0001  https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1
2022/03/27 08:00:55 Get 0001  https://dcollections.lib.keio.ac.jp/sites/default/files/iiif/KAN/110X-24-1/manifest.json
2022/03/27 08:00:56 A total of 46 pages.
2022/03/27 08:00:56 Save as  D:\src\bookget\Downloads\book.110X-24-1\dezoomify-rs.urls.bat  (5.38 KB)
2022/03/27 08:00:56 Get 0001  https://iiif.lib.keio.ac.jp/iipsrv/KAN/110X-24-1/tif/001.tif/full/full/0/default.jpg
```

附：已通过测试的URL如下
第一种：IIIF Manifest 的URL（推荐手动找URL）
```
https://iiif.lib.harvard.edu/manifests/drs:53262215
https://digicoll.lib.berkeley.edu/nanna/iiif/91514/manifest
https://khirin-a.rekihaku.ac.jp/iiif/rekihaku/H-173-1/manifest.json
https://khirin-a.rekihaku.ac.jp/manifests/sohan_shiki/H-172-01.json
https://dcollections.lib.keio.ac.jp/sites/default/files/iiif/KAN/110X-24-1/manifest.json
https://iiif.bodleian.ox.ac.uk/iiif/manifest/310cb04e-6bce-44e3-85b5-03417c9644a8.json
https://api.digitale-sammlungen.de/iiif/presentation/v2/bsb11129280/manifest
https://figgy.princeton.edu/concern/scanned_resources/e5313f5e-f2fc-4bdd-a894-7cffac271dfd/manifest
```
第二种：包含IIIF Manifest链接的网页URL
```
https://dcollections.lib.keio.ac.jp/ja/kanseki/110x-24-1  
https://khirin-a.rekihaku.ac.jp/sohanshiki/h-172-1
https://khirin-a.rekihaku.ac.jp/sohankanjo/h-173-1 
```

