# code2pdf Herramienta de conversión de proyecto a PDF
code2pdf es un programa que convierte todos los archivos de texto en una carpeta especificada en un solo archivo PDF. Adjunta el diagrama de estructura de directorios al inicio del PDF y agrega el contenido de cada archivo después.

### Inicio Rápido
1. Coloque la herramienta code2pdf dentro del directorio de la carpeta de su proyecto.
2. Ejecute code2pdf y tome un sorbo de agua.
3. Abra el archivo generado 1.pdf para ver los resultados.

### Parámetros
-addfiles string

Si al PDF generado le faltan algunos archivos, agréguelos manualmente usando este parámetro. Por ejemplo, -addfiles=".abc|.cba"

-dir string

Especifique la carpeta que desea convertir a PDF.

-fontfile string

Establezca la ruta al archivo de fuente (predeterminado "C:\Windows\Fonts\simfang.ttf").

-fontsize int

Establezca el tamaño de fuente en el PDF (predeterminado 14).

-fonttype string

Especifique el formato de fuente en el archivo de fuente a utilizar (predeterminado "fangsong").

-pdfname string

Establezca el nombre del archivo PDF convertido (predeterminado "1.pdf").