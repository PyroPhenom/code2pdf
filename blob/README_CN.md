# code2pdf 项目转pdf工具
code2pdf是一个能将指定文件夹内的所有文本类型的文件转换为单个pdf文件的程序, 它会在pdf的开头附上整个文件夹的结构图, 并在后面追加每个文件的内容.

### 快速开始
1. 将code2pdf放到你的项目文件夹目录内
2. 运行code2pdf, 然后喝口水
3. 运行当前目录内的1.pdf查看结果

### 参数
-addfiles string

如果生成的 PDF 缺少某些文件, 请使用此参数手动添加它们, 例如 -addfiles=".abc|.cba"

-dir string

您要转换为 pdf 的文件夹

-fontfile string

字体文件路径 (default "C:\Windows\Fonts\simfang.ttf")

-fontsize int

pdf中的字体大小 (default 14)

-fonttype string

使用字体文件中的哪种格式 (default "fangsong")

-pdfname string

转换后的 pdf 文件名 (default "1.pdf")
