package main

import (
"io"
"log"
"os"
"sync"
)


var (
errorlog=log.New(os.Stdout,"\033[31m[error]\033[0m",log.LstdFlags|log.Lshortfile)

infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex

)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)
