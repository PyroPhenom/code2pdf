# code2pdf Projekt zur PDF-Konvertierung
code2pdf ist ein Programm, das alle Textdateien in einem angegebenen Ordner in eine PDF-Datei konvertiert. Es fügt am Anfang der PDF ein Verzeichnisstrukturdiagramm hinzu und hängt anschließend den Inhalt jeder Datei an.

### Schnellstart
Platzieren Sie das code2pdf-Tool innerhalb des Ordners Ihres Projektverzeichnisses.
Führen Sie code2pdf aus und nehmen Sie einen Schluck Wasser.
Öffnen Sie die generierte Datei 1.pdf, um die Ergebnisse anzuzeigen.

### Parameter
-addfiles string

Fügen Sie fehlende Dateien manuell mit diesem Parameter hinzu, falls die generierte PDF unvollständig ist. Zum Beispiel: -addfiles=".abc|.cba"

-dir string

Geben Sie den Ordner an, den Sie in eine PDF konvertieren möchten.

-fontfile string

Legen Sie den Pfad zur Schriftartdatei fest (Standard: "C:\Windows\Fonts\simfang.ttf").

-fontsize int

Legen Sie die Schriftgröße in der PDF fest (Standard: 14).

-fonttype string

Geben Sie das Schriftformat in der zu verwendenden Schriftartdatei an (Standard: "fangsong").

-pdfname string

Legen Sie den Dateinamen der konvertierten PDF fest (Standard: "1.pdf").
