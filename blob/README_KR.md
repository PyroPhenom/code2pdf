# code2pdf PDF 변환 도구
code2pdf는 지정된 폴더 내의 모든 텍스트 형식 파일을 하나의 PDF 파일로 변환하는 프로그램입니다. 디렉터리 구조 다이어그램을 PDF의 시작 부분에 첨부하고 각 파일의 내용을 그 뒤에 추가합니다.

### 빠른 시작
1. code2pdf 도구를 프로젝트 폴더 디렉터리 내에 배치합니다.
2. code2pdf를 실행하고 물 한 모금 마십시오.
3. 생성된 1.pdf 파일을 열어 결과를 확인합니다.

### 매개변수
-addfiles string

생성된 PDF에서 누락된 파일이 있는 경우 이 매개변수를 사용하여 수동으로 추가합니다. 예: -addfiles=".abc|.cba"

-dir string

PDF로 변환할 폴더를 지정합니다.

-fontfile string

폰트 파일 경로 설정(기본값 "C:\Windows\Fonts\simfang.ttf").

-fontsize int

PDF의 폰트 크기 설정(기본값 14).

-fonttype string

사용할 폰트 파일의 폰트 형식 지정(기본값 "fangsong").

-pdfname string

변환된 PDF 파일 이름 설정(기본값 "1.pdf")