# code2pdf प्रोजेक्ट को PDF में कनवर्ट करने का टूल
code2pdf एक प्रोग्राम है जो निर्दिष्ट फ़ोल्डर के अंदर के सभी टेक्स्ट-टाइप फ़ाइलों को एकल PDF फ़ाइल में कनवर्ट करता है। यह PDF के शुरुआत में डायरेक्टरी स्ट्रक्चर डायग्राम अटैच करता है और उसके बाद प्रत्येक फ़ाइल की कंटेंट ऐपेंड करता है।

### त्वरित प्रारम्भ
1. अपने प्रोजेक्ट फ़ोल्डर डायरेक्टरी के अंदर code2pdf टूल रखें।
2. code2pdf चलाएँ और पानी का एक सिप लें।
3. परिणाम देखने के लिए बनी 1.pdf फ़ाइल खोलें।

### पैरामीटर्स
-addfiles string

अगर बनी PDF में कुछ फ़ाइलें गायब हैं, तो इस पैरामीटर की मदद से मैन्युअली उन्हें ऐड करें। उदाहरण के लिए, -addfiles=".abc|.cba"

-dir string

वो फ़ोल्डर स्पेसिफ़ाई करें जिसे आप PDF में कनवर्ट करना चाहते हैं।

-fontfile string

फ़ॉन्ट फ़ाइल का पथ सेट करें (डिफ़ॉल्ट "C:\Windows\Fonts\simfang.ttf")।

-fontsize int

PDF में फ़ॉन्ट साइज़ सेट करें (डिफ़ॉल्ट 14)।

-fonttype string

उपयोग करने के लिए फ़ॉन्ट फ़ाइल में फ़ॉन्ट फॉर्मैट स्पेसिफ़ाई करें (डिफ़ॉल्ट "fangsong")।

-pdfname string

कनवर्ट होने वाली PDF का फ़ाइल नाम सेट करें (डिफ़ॉल्ट "1.pdf")।