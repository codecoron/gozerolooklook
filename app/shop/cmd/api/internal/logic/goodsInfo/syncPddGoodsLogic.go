package goodsInfo

import (
	"context"
	"fmt"
	"github.com/liunian1004/pdd"
	"looklook/app/shop/model"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/shop/cmd/api/internal/svc"
)

type SyncPddGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncPddGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncPddGoodsLogic {
	return &SyncPddGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncPddGoodsLogic) SyncPddGoods() error {
	p := pdd.NewPdd(&pdd.Config{
		ClientId:     "c39367656d5d436baeffc8b160d6ce68",
		ClientSecret: "448c72ef863e41a03439ff9120ca98b7c35e7b7b",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	fmt.Println("11111111111")
	// 初始化多多客相关 API 调用
	d := p.GetDDK()
	fmt.Println("2222222222222")

	search, err := d.GoodsSearch()
	if err != nil {
		fmt.Println("报错了")
	}
	//得到查找到的商品列表
	goodslist := search.GoodsList
	//for-range循环遍历列表的每个商品
	for _, pddGoods := range goodslist {
		//将商品的信息写入到数据库中 goods.字段名
		goods := new(model.Goods)
		goods.GoodsId = strconv.Itoa(pddGoods.GoodsId)
		goods.GoodsDesc = pddGoods.GoodsDesc
		_, err := l.svcCtx.GoodsModel.Insert(l.ctx, goods)
		if err != nil {
			return err
		}
	}

	//resp, err := json.Marshal(search)
	//fmt.Println(string(resp))

	//根据商品id获取具体信息
	//报错 商品id
	//res, err := d.GoodsDetail(329746189863)
	//if err != nil {
	//	fmt.Println("获取信息失败")
	//}
	//resp, err := json.Marshal(res)
	//fmt.Println(string(resp))
	return err
}
