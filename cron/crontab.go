package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)
type task struct {
	Taskname         string `json:"taskname"`
	Intervaltime     string `json:"intervaltime"`
	Workingdirectory string `json:"workingdirectory"`
	Execcommand      string `json:"execcommand"`
	terminal         *exec.Cmd
	bufin            *bytes.Buffer
	bufout           *bytes.Buffer
}
type tasklist struct {
	Tasklist  []task `json:"tasklist"`
	Checktime int    `json:"checktime"`
}
func readfile(filepath string) tasklist {
	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read file failed", err)
	}
	allltask := tasklist{}
	err = json.Unmarshal(filedata, &allltask)
	if err != nil {
		fmt.Println("read file failed", err)
		if len(allltask.Tasklist) == 0 {
			fmt.Println("list error")
		}
	}
	return allltask
}
func CheckProcessAlive(taskmap *map[string]task, checktime int) {
	timetricker := time.Tick(time.Second * time.Duration(checktime))
	for timer := range timetricker {
		fmt.Println("==================check alive==================")
		for _, taskname := range *taskmap {
			pid := taskname.terminal.Process.Pid
			pn, err := process.NewProcess(int32(pid))
			if err == nil {
				// 获取command
				command, err := pn.Cmdline()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("command line:",command)
				}
				fmt.Println(timer, taskname.Taskname, pid, "is alive")
			} else {
				fmt.Println(timer, taskname.Taskname, pid, "is dead", err)
				taskname.StartProcess(taskmap)
			}

		}
	}

}
func (processd *task) KillProcess() {
	pid := processd.terminal.Process.Pid
	cmd := exec.Command("taskkill", "/f", "/pid", strconv.Itoa(pid))
	err := cmd.Start()
	if err != nil {
		fmt.Println(time.Now(), pid, "kill process failed", err)
	} else {
		fmt.Println(time.Now(), pid, "kill process succeed")
	}

}
func (process *task) StartProcess(alltast *map[string]task) {
	list:=strings.Split(process.Execcommand," ")
	args:=make([]string,0,len(list))
	for index:=range list{
		if list[index]!=""{
			args=append(args,list[index])
		}
	}
	cmd := exec.Command(args[0],args[1:]...)
	cmd.Dir = process.Workingdirectory
	process.terminal = cmd
	err := process.terminal.Start()
	//  对map 进行加锁
	mux.Lock()
	(*alltast)[process.Taskname] = *process
	mux.Unlock()
	if err != nil {
		fmt.Println(time.Now(), process.Taskname, "start process failed", err)
	}
	fmt.Println(time.Now(), process.Taskname, "start process succeed")
}
func (processd *task) Run() {
	t := taskmap[processd.Taskname]
	fmt.Println("enter timmer ,kill process ............", t.terminal.Process.Pid)
	t.KillProcess()
}
// map 需要定义成全局变量然后方便访问,对map 进行上锁操作，避免写入数据的过程中被读取数据
var mux sync.Mutex
var taskmap map[string]task
func main() {
	taskmap = make(map[string]task)
	alltast := readfile("./task.json")
	timmer := cron.New()
	for index, _ := range alltast.Tasklist {
		fmt.Println(time.Now(), "new task, create task", alltast.Tasklist[index].Taskname)
		// 直接通过索引访问不是通过值访问所以注册的时候不影响
		alltast.Tasklist[index].StartProcess(&taskmap)
	}
	for key, _ := range taskmap {
		temptask := taskmap[key]
		id, err := timmer.AddJob(temptask.Intervaltime, &temptask)
		if err != nil {
			fmt.Println(time.Now(), "timmer error", err, id)
		}
		timmer.Start()
		defer timmer.Stop()
	}
	CheckProcessAlive(&taskmap, alltast.Checktime)
}
