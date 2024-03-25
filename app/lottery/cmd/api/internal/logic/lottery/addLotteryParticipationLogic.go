package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(req *types.AddLotteryParticipationReq) (resp *types.AddLotteryParticipationResp, err error) {
	_, err = l.svcCtx.LotteryRpc.AddLotteryParticipation(l.ctx, &lottery.AddLotteryParticipationReq{
		UserId:    ctxdata.GetUidFromCtx(l.ctx),
		LotteryId: req.LotteryId,
	})
	return
}

// 这里用到了什么设计模式
// 这里用到了代理模式
// 代理模式是一种结构型设计模式， 让你能够提供对象的替代品或其占位符。 代理控制着对于原对象的访问， 并允许在将请求提交给对象前后进行一些处理。
// 代理模式的关键是代理类和真实类都实现同一个接口，这样代理类才能在任何使用真实类的地方都能使用。
// 代理模式的优点：
// 1. 你可以在客户端毫无察觉的情况下控制服务对象。
// 2. 如果客户端对服务对象的使用方式没有特殊要求， 你可以在不对服务对象做任何修改的情况下控制其行为。
// 3. 你可以在服务对象上方附加额外的功能， 比如缓存请求、 记录日志、 进行验证客户端请求等。
// 4. 你可以对服务对象进行保护， 使得客户端无法直接调用。
// 5. 代理模式符合开闭原则。
// 代理模式的缺点：
// 1. 代码可能会变得复杂， 因为需要新建许多类。
// 2. 服务响应可能会变慢， 因为在客户端和服务对象之间增加了一个代理。
// 3. 有些代理模式实现可能会造成请求的失败。
// 4. 代理模式会造成系统设计中类的数目增加。
// 5. 代理模式会增加系统的复杂度。
// 代理模式的应用场景：
// 1. 远程代理（Remote Proxy）为一个对象在不同的地址空间提供局部代表。
// 2. 虚拟代理（Virtual Proxy）根据需要创建开销很大的对象。
// 3. 保护代理（Protection Proxy）控制对原始对象的访问。
// 4. 智能指引（Smart Reference）取代了简单指针， 它在访问对象时执行一些附加操作。
// 5. 代理模式常用于实现延迟加载， 即先加载轻量级的代理对象， 真正需要时再加载真实对象。
// 6. 代理模式常用于实现日志记录， 以记录方法的调用情况。
// 7. 代理模式常用于实现安全控制， 以控制用户对方法的访问。
// 8. 代理模式常用于实现远程调用， 以实现远程方法调用。
// 9. 代理模式常用于实现事务控制， 以实现对方法的事务控制。
// 10. 代理模式常用于实现并发控制， 以实现对方法的并发控制。
// 代理模式的实现方式：
// 1. 保持一个引用使得代理可以访问实体。
// 2. 提供一个与Subject相同的接口， 这样代理就可以替代实体。
// 3. 控制对实体的存取， 可能需要创建、 删除实体。
// 4. 其他功能， 比如记录日志、 访问计数等。
// 代理模式的角色：
// 1. Subject（抽象主题角色）：声明了RealSubject和Proxy的共同接口， 这样在任何使用RealSubject的地方都可以使用Proxy。
// 2. RealSubject（真实主题角色）：定义了Proxy所代表的真实实体。
// 3. Proxy（代理主题角色）：保存一个引用使得代理可以访问实体。 提供一个与Subject相同的接口， 这样代理就可以替代实体。 控制对实体的存取， 可能需要创建、 删除实体。 其他功能， 比如记录日志、 访问计数等。
// 代码中的代理模式：
// 1. AddLotteryParticipationLogic是代理主题角色
// 2. svc.ServiceContext是真实主题角色
// 3. AddLotteryParticipationLogic实现了AddLotteryParticipation方法，这个方法的实现是调用了svc.ServiceContext的LotteryRpc.AddLotteryParticipation方法
// 4. AddLotteryParticipationLogic和svc.ServiceContext都实现了同一个接口
// 5. AddLotteryParticipationLogic保存了一个引用使得代理可以访问实体
// 6. AddLotteryParticipationLogic提供了一个与svc.ServiceContext相同的接口，这样代理就可以替代实体
// 7. AddLotteryParticipationLogic控制对svc.ServiceContext的存取
// 8. AddLotteryParticipationLogic提供了其他功能，比如记录日志
// 9. AddLotteryParticipationLogic实现了AddLotteryParticipation方法，这个方法的实现是调用了svc.ServiceContext的LotteryRpc.AddLotteryParticipation方法
