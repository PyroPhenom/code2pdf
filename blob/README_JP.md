# code2pdf プロジェクト PDF変換ツール
code2pdfは、指定されたフォルダ内のすべてのテキストファイルを1つのPDFファイルに変換するプログラムです。PDFの先頭にディレクトリ構造図が添付され、その後に各ファイルの内容が追加されます。

### クイックスタート
1. code2pdfツールをプロジェクトフォルダディレクトリに配置します。
2. code2pdfを実行し、作業中に水を飲みます。
3. 処理が完了したら、生成された1.pdfファイルを開いて結果を確認します。

### パラメーター
-addfiles string

生成されたPDFに一部のファイルが欠落している場合は、このパラメーターを使用して手動で追加します。例：-addfiles=".abc|.cba"

-dir string

PDFに変換したいフォルダを指定します。

-fontfile string

フォントファイルのパスを設定します（デフォルトは "C:\Windows\Fonts\simfang.ttf"）。

-fontsize int

PDFのフォントサイズを設定します（デフォルトは14）。

-fonttype string

使用するフォントファイルのフォント形式を指定します（デフォルトは "fangsong"）。

-pdfname string

変換されたPDFのファイル名を設定します（デフォルトは "1.pdf"）。