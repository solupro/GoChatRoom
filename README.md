GoChatRoom
==========

这是一个基于`WebSocket`使用`Go语言`编写的聊天室.  
代码参考了这篇文章[http://gary.beagledreams.com/page/go-websocket-chat.html](http://gary.beagledreams.com/page/go-websocket-chat.html)  
`go build` 编译以后通过 `./GoChatRoom -h` 查看帮助  

	Usage of ./GoChatRoom:
	  -addr=":8177": http service address
	  -mf=2: message's frequency(>second/msg)
	  -size=50: max size of the message's queue
	
预览地址:[http://gochatroom.coding.io/](http://gochatroom.coding.io/)
