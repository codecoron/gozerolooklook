package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/api/internal/config"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}

// 这里用到了什么设计模式
// 这里用到了抽象工厂模式
// 抽象工厂模式是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。 抽象工厂定义了一系列用于创建产品家族的方法， 每个方法均返回不同抽象产品。 调用者需要知道所使用对象的具体类， 但不需要了解对象的创建过程。
// 抽象工厂模式的关键是抽象工厂和具体工厂都实现同一个接口，这样具体工厂才能在任何使用抽象工厂的地方都能使用。
// 抽象工厂模式的优点：
// 1. 你可以确保同一工厂生产的产品相互兼容。
// 2. 你可以避免客户端代码和具体产品类之间的绑定。
// 3. 你可以在运行时改变具体工厂。
// 4. 单一职责原则。 你可以将产品生成代码抽取到同一位置， 使得代码更容易维护。
// 5. 开闭原则。 向应用程序中添加新产品变体时， 无需修改现有代码。
// 抽象工厂模式的缺点：
// 1. 由于该模式的关键是创建多个相关的产品， 因此其主要缺点在于产品族的扩展将很困难。
// 2. 产品族的横向扩展会导致大量的新类。
// 3. 产品族的纵向扩展会导致大量的新工厂。
// 4. 由于抽象工厂模式中的产品是抽象的， 因此客户端无法直接实例化它们。 这可能会导致一些额外的工作。
// 抽象工厂模式的应用场景：
// 1. 当你的代码需要与多个不同系列的相关产品交互时， 可以使用抽象工厂模式。
// 2. 抽象工厂提供了一种将一组具有共同主题的单个工厂封装起来的方式。 该模式有助于减少与单个工厂关联的代码量， 使得应用程序更易于交换产品系列。
// 3. 如果你发现自己编写大量的重复代码以创建不同系列的产品， 那么你可能需要使用抽象工厂模式。
// 4. 抽象工厂模式有助于将产品的实现从客户端代码中解耦， 使其更易于与特定产品系列的具体实现进行交互。
// 5. 抽象工厂模式有助于提供一种封装机制， 使得客户端代码无需了解产品的创建过程。
// 代码中的抽象工厂模式的实现：
// 1. ServiceContext 是抽象工厂
// 2. NewServiceContext 是具体工厂
// 3. Config 是抽象产品
// 4. UsercenterRpc 和 LotteryRpc 是具体产品
// 5. 抽象工厂模式的实现中，抽象工厂和具体工厂都实现了同一个接口 ServiceContext
