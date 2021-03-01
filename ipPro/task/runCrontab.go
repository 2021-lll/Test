package task

import (
	"fmt"
	"ipPro/collect"
	"os"
)


func Task()  {
	crontab := NewCrontab()

	//task1 := &testTask{}
	//if err := crontab.AddByID("1", "* * * * *",task1); err != nil{
	//	fmt.Println(err)
	//	os.Exit(-1)
	//}

	taskFunc := func() {
		collect.GetToken()
		fmt.Println("获取token成功")
	}
	if err := crontab.AddByFunc("2", "* * * * *",  taskFunc); err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	crontab.Start()
	select {

	}

}
