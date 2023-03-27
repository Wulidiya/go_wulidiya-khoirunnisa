# 20_Section_Intro Echo Golang
Library golang:
1.	echo 
2.	Go kit
3.	Logrus
4.	gorm.io
5.	cobra
# echo
- Untuk membantu pengembangan web, memiliki kecepatan tinggi, extensible, minimalist, dan web framework. 
- Cara melakukan routing menggunakan framework echo dengan method get pada path end point "/products" terhadap ProductController adalah GET("/product", ProductController)
- Cara installasi go dan update yaitu dengan go get github.com/labstack/echo/v4
- Cara menerima id ketika membuat Rest API untuk detail product yaitu melalui URL param
- Keunggulan dari framework echo adalah: 
1. Dapat melakukan optimize router sehingga menjadi lebih cepat
2. Mempunyai scalable yang bagus untuk traffic besar
3. Dapat melakukan databinding
- Cara mengirim response dalam bentuk JSON dengan menggunakan framework echo dimana terdapat echo.Context bernama ctx yaitu dengan ctx.JSON

 
