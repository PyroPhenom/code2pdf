# code2pdf প্রজেক্টকে PDF-এ রূপান্তর করার টুল
code2pdf হল একটি প্রোগ্রাম যা নির্দিষ্ট একটি ফোল্ডারের মধ্যে থাকা সব টেক্সট-টাইপ ফাইলগুলিকে একটি একক PDF ফাইলে রূপান্তর করে। এটি PDF-এর শুরুতে ডিরেক্টরি স্ট্রাকচার ডায়াগ্রাম অ্যাটাচ করে এবং প্রতিটি ফাইলের বিষয়বস্তু পরবর্তীতে অ্যাপেন্ড করে।

### দ্রুত সূচনা
1. আপনার প্রজেক্ট ফোল্ডার ডিরেক্টরির মধ্যে code2pdf টুলটি রাখুন।
2. code2pdf চালু করুন এবং পানির একটি সিপ নিন।
3. ফলাফল দেখতে সৃষ্ট 1.pdf ফাইলটি খুলুন।

### প্যারামিটারসমূহ
-addfiles string

যদি তৈরি PDF-তে কিছু ফাইল অনুপস্থিত থাকে, তাহলে ম্যানুয়ালি এই প্যারামিটার ব্যবহার করে যোগ করুন। উদাহরণস্বরূপ, -addfiles=".abc|.cba"

-dir string

আপনি যে ফোল্ডারটি PDF-এ রূপান্তর করতে চান সেটি নির্দিষ্ট করুন।

-fontfile string

ফন্ট ফাইল পাথ সেট করুন (ডিফল্ট "C:\Windows\Fonts\simfang.ttf")।

-fontsize int

PDF-এ ফন্ট সাইজ সেট করুন (ডিফল্ট 14)।

-fonttype string


ব্যবহার করার জন্য ফন্ট ফাইলে ফন্ট ফর্ম্যাট নির্দিষ্ট করুন (ডিফল্ট "fangsong")।

-pdfname string

রূপান্তরিত PDF এর ফাইল নাম সেট করুন (ডিফল্ট "1.pdf")।
