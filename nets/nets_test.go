package nets

import "time"

func socket(){
	go Connect()
	time.Sleep(time.Second*1)
	Dial()
}

