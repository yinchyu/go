package filetype

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {
	f, err := os.Open("test.html")
	//f, err := os.Open("C:\\Users\\Administrator\\Desktop\\Wildlife.wmv")
	if err != nil {
		t.Logf("open error: %v", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	fmt.Println(fSrc)
	t.Log(GetFileType(fSrc[:10]))
}
