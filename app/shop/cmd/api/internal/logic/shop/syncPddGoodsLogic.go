package shop

import (
	"context"
	"fmt"
	"github.com/liunian1004/pdd"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/model"
	"strconv"
	"time"
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
	goodsList := search.GoodsList
	goods := new(model.Goods)
	//for-range循环遍历列表的每个商品
	for _, pddGoods := range goodsList {
		//将商品的信息写入到数据库中 goods.字段名
		goods.GoodsId = strconv.Itoa(pddGoods.GoodsId)
		goods.CategoryId = int64(pddGoods.CategoryId)
		//在包中为min_normal_price
		goods.Precoupon_Price = float64(pddGoods.MinNormalPrice)
		//先设置为券前价-券面值
		goods.Aftercoupon_Price = float64(pddGoods.MinGroupPrice - pddGoods.CouponDiscount)
		goods.GoodsDesc = pddGoods.GoodsDesc
		//心愿值自己设置为1
		goods.WishPoints = 1
		goods.CouponStartTime = int64(pddGoods.CouponStartTime)
		goods.CouponEndTime = int64(pddGoods.CouponEndTime)
		goods.CouponDiscount = int64(pddGoods.CouponDiscount)
		goods.CouponRemainQuantity = int64(pddGoods.CouponRemainQuantity)
		//如果数据库已经有数据了,可以先注释掉
		_, err := l.svcCtx.GoodsModel.Insert(l.ctx, goods)
		if err != nil {
			l.Logger.Error("数据库插入数据失败")
		}
		goodsCategory := new(model.GoodsCategory)
		goodsCategory.CategoryId = int64(pddGoods.CategoryId)
		goodsCategory.CategoryName = pddGoods.CategoryName
		//如果数据库已经有数据了,可以先注释掉
		l.svcCtx.GoodsCategoryModel.Insert(l.ctx, goodsCategory)
	}

	//pddOrigingoods, err := syncLogic.getSyncPddGoods() //先获取一次数据,并把数据写库
	// 启动定时任务
	ticker := time.NewTicker(30 * time.Second) // 每隔一段时间执行一次同步任务
	defer ticker.Stop()                        // 在程序结束时停止定时器

	//if err != nil {
	//	fmt.Printf("Failed to sync PDD goods: %v\n", err)
	//} else {
	//	fmt.Println("Sync PDD goods successfully.")
	//}
	for {
		time.Sleep(30 * time.Second)
		// 在定时器触发时调用 SyncPddGoods 方法
		//只是获取数据并赋值

		syncSearch, err := d.GoodsSearch()
		if err != nil {
			fmt.Println("报错了")
		}

		//得到查找到的商品列表
		syncGoodsList := syncSearch.GoodsList
		syncGoods := new(model.Goods)
		//for-range循环遍历列表的每个商品
		for _, syncPddGoods := range syncGoodsList {
			//将商品的信息写入到数据库中 goods.字段名
			syncGoods.GoodsId = strconv.Itoa(syncPddGoods.GoodsId)

			if syncGoods.GoodsId != goods.GoodsId {
				l.svcCtx.GoodsModel.Update(l.ctx, syncGoods)
			}

			syncGoods.CategoryId = int64(syncPddGoods.CategoryId)
			//在包中为min_normal_price
			syncGoods.Precoupon_Price = float64(syncPddGoods.MinNormalPrice)
			//先设置为券前价-券面值
			syncGoods.Aftercoupon_Price = float64(syncPddGoods.MinGroupPrice - syncPddGoods.CouponDiscount)
			syncGoods.GoodsDesc = syncPddGoods.GoodsDesc
			//心愿值自己设置为1
			syncGoods.WishPoints = 1
			syncGoods.CouponStartTime = int64(syncPddGoods.CouponStartTime)
			syncGoods.CouponEndTime = int64(syncPddGoods.CouponEndTime)
			syncGoods.CouponDiscount = int64(syncPddGoods.CouponDiscount)
			syncGoods.CouponRemainQuantity = int64(syncPddGoods.CouponRemainQuantity)

			syncGoodsCategory := new(model.GoodsCategory)
			syncGoodsCategory.CategoryId = int64(syncPddGoods.CategoryId)
			syncGoodsCategory.CategoryName = syncPddGoods.CategoryName

			goods = syncGoods
		}
		////如果说这里pddOrigingoods和pddSyncGoods比较后不同，就把数据插入或者更新
		////否则不用做任何的变化
		//if pddSyncGoods.GoodsId != pddOrigingoods.GoodsId {
		//	l.svcCtx.GoodsModel.Insert(l.ctx, pddSyncGoods)
		//}
		//todo 其他的字段信息比较

	}
	return nil
}
