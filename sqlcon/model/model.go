package model

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);"`
	Email     string `gorm:"varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Authority int    `gorm:"not null"`
}

type Announce struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50);not null"`
	Content string `gorm:"type:varchar(200);"`
	Url     string `gorm:"type:varchar(100);"`
}

type AnnounceUser struct {
	gorm.Model
	Aid    uint `gorm:"not null"`
	Uid    uint `gorm:"not null"`
	Status int  `gorm:"default:0"` //已读状态
}
type Carousel struct {
	gorm.Model
	Img string `gorm:"size:255;"`
	Url string `gorm:"type:varchar(100)"`
}

type Collection struct {
	gorm.Model
	Cover string `gorm:"size:255;not null"`
	Title string `gorm:"type:varchar(50);not null;"`
	Desc  string `gorm:"size:255;"`
	Uid   uint   // 作者
}

type Comment struct {
	gorm.Model
	Vid     uint   `gorm:"not null;index"`             //视频ID
	Content string `gorm:"type:varchar(255);not null"` //内容
	Uid     uint   `gorm:"not null"`                   //用户
	//回复数,3.2版本新增,用于获取评论列表V2接口
	ReplyCount int `gorm:"default:0"`
}

type Danmaku struct {
	gorm.Model
	Vid   uint   `gorm:"not null;index"`
	Time  uint   `gorm:"not null"`  //时间
	Type  int    `gorm:"default:0"` //类型0滚动;1顶部;2底部
	Color string `gorm:"type:varchar(10);default:'#fff'"`
	Text  string `gorm:"type:varchar(100);not null"`
	Uid   uint   `gorm:"not null"`
}

type Follow struct {
	gorm.Model
	Uid uint `gorm:"not null"`
	Fid uint `gorm:"not null"`
}

type Review struct {
	gorm.Model
	Vid     uint   `gorm:"not null;index"` //视频ID
	Video   Video  `gorm:"ForeignKey:id;AssociationForeignKey:vid"`
	Status  int    `gorm:"not null"`          //审核状态
	Remarks string `gorm:"type:varchar(20);"` //备注
}

type User struct {
	gorm.Model
	Avatar   string    `gorm:"size:255;"`
	Name     string    `gorm:"type:varchar(20);not null"`
	Email    string    `gorm:"varchar(20);not null;index"`
	Password string    `gorm:"size:255;not null"`
	Gender   int       `gorm:"default:0"`
	Birthday time.Time `gorm:"default:'1970-01-01'"`
	Sign     string    `gorm:"varchar(50);default:'这个人很懒，什么都没有留下'"`
}
type Video struct {
	gorm.Model
	Title        string  `gorm:"type:varchar(50);not null;index"`
	Cover        string  `gorm:"size:255;not null"`
	Video        string  `gorm:"size:255"`
	VideoType    string  `gorm:"varchar(5)"`
	Introduction string  `gorm:"varchar(100);default:'什么也没有'"` //视频简介
	Uid          uint    `gorm:"not null;index"`
	Author       User    `gorm:"ForeignKey:id;AssociationForeignKey:uid"`
	Original     bool    `gorm:"not null"`      //是否为原创
	Weights      float32 `gorm:"default:0"`     //视频权重(目前还没使用)
	Clicks       int     `gorm:"default:0"`     //点击量
	Review       bool    `gorm:"default:false"` //是否审查通过
}

type VideoCollection struct {
	gorm.Model
	Vid          uint
	CollectionId uint
}

type Interactive struct {
	gorm.Model
	Uid     uint `gorm:"not null"`
	Vid     uint `gorm:"not null"`
	Collect bool `gorm:"default:false"` //是否收藏
	//like和SQL的关键词冲突了，查询时需要写成`like`
	Like  bool  `gorm:"default:false"` //是否点赞
	Video Video `gorm:"ForeignKey:id;AssociationForeignKey:vid"`
}
type Message struct {
	gorm.Model
	Uid     uint   `gorm:"not null;"` //用户ID
	Fid     uint   `gorm:"not null;"` //关联ID
	FromId  uint   `gorm:"not null;"` // 发送者
	ToId    uint   `gorm:"not null;"` // 接受者
	Content string `gorm:"size:255;"`
	Status  int    `gorm:"default:0"` //已读状态
}
type Reply struct {
	gorm.Model
	Cid       uint   `gorm:"not null;index"`             //评论的ID
	Content   string `gorm:"type:varchar(255);not null"` //内容
	Uid       uint   `gorm:"not null"`                   //用户
	ReplyUid  uint   //回复的人的uid
	ReplyName string `gorm:"type:varchar(20);"` //回复的人的昵称
}
