package main

import (
	"database/sql"
	"time"
)

type Goods struct {
	GoodsId        int64        `db:"goods_id"`
	Title          string       `db:"title"`
	SubTitle       string       `db:"sub_title"`
	OriginalCost   float64      `db:"original_cost"`
	CurrentPrice   float64      `db:"current_price"`
	Discount       float64      `db:"discount"`
	IsFreeDelivery int32        `db:"is_free_delivery"`
	CategoryId     int64        `db:"category_id"`
	LastUpdateTime sql.NullTime `db:"last_update_time"`
}

type GoodsCover struct {
	GcId       int64  `db:"gc_id"`
	GoodsId    int64  `db:"goods_id"`
	GcPicUrl   string `db:"gc_pic_url"`
	GcThumbUrl string `db:"gc_thumb_url"`
	GcOrder    int64  `db:"gc_order"`
}

type GoodsDetail struct {
	GdId     int64  `db:"gd_id"`
	GoodsId  int64  `db:"goods_id"`
	GdPicUrl string `db:"gd_pic_url"`
	GdOrder  int32  `db:"gd_order"`
}

type GoodsParam struct {
	GdId         int64  `db:"gp_id"`
	GpParamName  string `db:"gp_param_name"`
	GpParamValue string `db:"gp_param_value"`
	GoodsId      int64  `db:"goods_id"`
	GdOrder      int32  `db:"gp_order"`
}
type PromotionSecKill struct {
	PsId         int64     `db:"ps_id"`
	GoodsId      int64     `db:"goods_id"`
	PsCount      int64     `db:"ps_count"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
	Status       int32     `db:"status"`
	CurrentPrice float64   `db:"current_price"`
	Version      int64     `db:"version"`
}

type SuccessKilled struct {
	GoodsId    int64     `db:"goods_id"`
	UserId     int64     `db:"user_id"`
	State      int16     `db:"state"`
	CreateTime time.Time `db:"create_time"`
}
