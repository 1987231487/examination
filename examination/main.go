package main

import (
	"examination/routers"
)

func main(){

	r:=routers.InitRouter()
	r.Run(":10086")

}
