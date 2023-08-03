# Ferramenta code2pdf para conversão de projeto em PDF
O code2pdf é um programa que converte todos os arquivos de texto em uma pasta especificada em um único arquivo PDF. Ele anexa o diagrama da estrutura de diretórios no início do PDF e acrescenta o conteúdo de cada arquivo depois.

### Início Rápido
1. Coloque a ferramenta code2pdf dentro do diretório da pasta do seu projeto.
2. Execute o code2pdf e tome um gole de água.
3. Abra o arquivo gerado 1.pdf para ver os resultados.

### Parâmetros
-addfiles string

Se o PDF gerado estiver sem alguns arquivos, adicione-os manualmente usando este parâmetro. Por exemplo, -addfiles=".abc|.cba"

-dir string

Especifique a pasta que você deseja converter para PDF.

-fontfile string

Defina o caminho para o arquivo de fonte (padrão "C:\\Windows\\Fonts\\simfang.ttf").

-fontsize int

Defina o tamanho da fonte no PDF (padrão 14).

-fonttype string

Especifique o formato da fonte no arquivo de fonte a ser usado (padrão "fangsong").

-pdfname string

Defina o nome do arquivo PDF convertido (padrão "1.pdf").