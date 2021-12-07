package main
import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)
var basedir ="D:\\train_dataset\\new"
// 一个conunt 没有增加引发的问题， 导致数据全部丢失，因为所有的数据文件都写到了一个里边，所以说测试也是很重要的，特别是涉及到文件的操作，一定要做好备份操作

var counter =0
func CreateMove(name string, context string){
	newfilename:=fmt.Sprintf("%04d",counter)
	file,err:=os.Create(filepath.Join(basedir,"ycy",newfilename+".lab"))
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()
	file.Write([]byte(context))
	splitname:=strings.Split(name,"_")
	//Mycy_C71001.pcm.wav
	fmt.Println(name,splitname)
	//wavname:="Mycy_"+strings.Join(splitname,"")+".pcm.wav"
	wavname:=strings.Join(splitname,"")+".pcm.wav"
	oldpath:=filepath.Join(basedir,	strings.ToUpper(splitname[0]),wavname)
	newpath:=filepath.Join(basedir,	"ycy",newfilename+".wav")
	err=os.Rename(oldpath,newpath)
	if err != nil {
		log.Println(err)
	}
}
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	HZGB2312 = Charset("HZGB2312")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case HZGB2312:
		var decodeBytes,_=simplifiedchinese.HZGB2312.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
func Process(){
	txtdir:="D:\\train_dataset\\new\\txt"
	entry,err:=os.ReadDir(txtdir)
	if err != nil {
		log.Println(err)
	}
	for _,file:=range entry{
		if !file.IsDir(){
			files,err:=os.Open(path.Join(txtdir,file.Name()))
			if err != nil {
				log.Println(err)
			}
			//data,_:=io.ReadAll(files)
			//utf8data:=ConvertByte2String(data,GB18030)
			//fmt.Println(utf8data)
			buff:=bufio.NewReader(files)
			for {
				line, _, eof := buff.ReadLine()
				if eof == io.EOF {
					break
				}
				counter++
				newline:=ConvertByte2String(line,GB18030)
				lineparts:=strings.Split(newline,"\t")
				linename:=lineparts[0]
				linecontext:=lineparts[1:]
				fmt.Println(file.Name(),linename,lineparts)
				CreateMove(linename,strings.Join(linecontext,""))
				//os.Exit(0)
			}
		}
	}
}
