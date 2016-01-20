练习使用一下beego 搭建了一个非常简单的服务。

在网页输入框内输入go的代码，点击编译，会在网页上显示编译后的结果。

如果有人想体验go语言编程，但是没有环境，那么可以使用这个简单的程序来让他在网页上感受go编程。



go get github.com/astaxie/beego

go get github.com/gitxiaolin/buildonline

//linux:
cd $GOPATH/src/github.com/gitxiaolin/buildonline

//windows 自行寻找目录gopath下src/github.com/gitxiaolin/buildonline

bee run或者go run main.go

打开浏览器，http://localhost:8080/build
注：linux下可能路径会有些问题，请自行修正
