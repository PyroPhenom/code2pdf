# code2pdf Project to PDF Conversion Tool
code2pdf is a program that converts all text-type files within a specified folder into a single PDF file. It attaches the directory structure diagram at the beginning of the PDF and appends the contents of each file afterward.

[中文](blob/README_CN.md) | [عربي](blob/README_AR.md) | [বাংলা](blob/README_BD.md) | [Deutsch](blob/README_DE.md) | [Français](blob/README_FR.md) | [हिंदी](blob/README_IN.md) | [日本](blob/README_JP.md) | [한국인](blob/README_KR.md) | [Português](blob/README_PT.md) | [Русский](blob/README_RU.md)

### Quick Start
1. Place the code2pdf tool inside your project folder directory.
2. Run code2pdf and take a sip of water.
3. open the generated file 1.pdf to view the results.

### Parameters

Usage of code2pdf:

-addfiles string

If the generated PDF is missing some files, manually add them using this parameter. For example, -addfiles=".abc|.cba"

-dir string

Specify the folder you want to convert to PDF.

-fontfile string

Set the font file path (default "C:\Windows\Fonts\simfang.ttf").

-fontsize int

Set the font size in the PDF (default 14).

-fonttype string

Specify the font format in the font file to use (default "fangsong").

-pdfname string

Set the file name of the converted PDF (default "1.pdf").
