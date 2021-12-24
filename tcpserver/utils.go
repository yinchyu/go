package main

func SendRest(){
	ln, err := net.ListenTCP("tcp4", nil)
	var t testing.T
	if err != nil {
		t.Fatal(err)
	}

	defer ln.Close()
	done := make(chan struct{})
	accepted := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			t.Error(err)
		}

		close(accepted)

		io.ReadAll(conn)
	}()

	conn, err := net.DialTCP("tcp4", nil, ln.Addr().(*net.TCPAddr))
	if err != nil {
		t.Fatal(err)
	}
	// This makes sure a TCP RST is sent when Close is called.
	if err := conn.SetLinger(0); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s <-> %s\n", conn.LocalAddr(), conn.RemoteAddr())
	go func() {
		for {
			b := make([]byte, 1+rand.Intn(1000))
			_, err := conn.Write(b)
			if err != nil {
				return
			}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
		}
	}()
	// 连接上服务端就继续执行代码
	<-accepted
	time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	// 发送rest 终止报文// readall 读取到EOF 自动的返回，传递done 信号
	if err := conn.Close(); err != nil {
		t.Fatal(err)
	}

	select {
		// 
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("remote conn didn't close")
	}

}