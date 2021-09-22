package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	SensitiveWords []string

	MessageQueueLen = 1024
)

func initConfig() {
	viper.SetConfigName("chatroom")
	viper.AddConfigPath(RootDir + "/config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	SensitiveWords = viper.GetStringSlice("sensitive")
	MessageQueueLen = viper.GetInt("message-queue")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 基本上就是我之前想要实现的功能， 开启一个文件夹的监控功能
		// 然后通过 readinconfig 之后在对所有的对象进行重新的赋值操作
		// 然后就可以不用重启就改变文件的
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("change the file is wrong ,please check", err)
		} else {
			SensitiveWords = viper.GetStringSlice("sensitive")
			MessageQueueLen = viper.GetInt("message-queue")
		}

	})
}
