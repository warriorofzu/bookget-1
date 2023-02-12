# bookget
bookget 数字图书馆下载工具

#### 通用批量下载(http/https链接)
因考虑到bookget不可能支持无穷数量的网站，特别提供通用批量下载功能。当然，这个功能在很多下载工具中都有了，bookget只是提供自动生成 0001/0002这样的顺序下载，以保证批量下载时文件名不乱。

#### 使用方法：
提示：自v0.2.7版开始已自动识别URL，无需修改配置文件或命令参数。   
例如：

```
$ bookget "https://ysts2.artron.net/books/book/BJRB-PM-11CP-HZMMY-001/XL/page_1.jpg"
$ bookget "https://ysts2.artron.net/books/book/BJRB-PM-11CP-HZMMY-001/XL/page_2.jpg"
$ bookget "https://ysts2.artron.net/books/book/BJRB-PM-11CP-HZMMY-001/XL/page_(1-36).jpg"

```
如果网址太多，可以编写一个文本文件 urls.txt，一行一个URL，回车换行。    
在终端内输入以下命令：
```
$ bookget -i urls.txt
```
注解：支持(01-100) 、(1-100)、(001-100)等格式通配符写法。


