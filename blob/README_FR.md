# code2pdf Outil de conversion de projet en PDF
code2pdf est un programme qui convertit tous les fichiers texte dans un dossier spécifié en un seul fichier PDF. Il attache le diagramme de structure du répertoire au début du PDF et ajoute ensuite le contenu de chaque fichier.

### Démarrage rapide
Placez l'outil code2pdf dans le répertoire de votre dossier de projet.
Exécutez code2pdf et prenez une gorgée d'eau.
Ouvrez le fichier généré 1.pdf pour voir les résultats.

### Paramètres
-addfiles string

Si le PDF généré est incomplet, ajoutez manuellement les fichiers manquants à l'aide de ce paramètre. Par exemple : -addfiles=".abc|.cba"

-dir string

Spécifiez le dossier à convertir en PDF.

-fontfile string

Définissez le chemin d'accès au fichier de police (par défaut "C:\Windows\Fonts\simfang.ttf").

-fontsize int

Définissez la taille de police dans le PDF (par défaut 14).

-fonttype string

Spécifiez le format de police dans le fichier de police à utiliser (par défaut "fangsong").

-pdfname string

Définissez le nom du fichier PDF converti (par défaut "1.pdf").
