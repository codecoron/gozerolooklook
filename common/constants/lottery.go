package constants

const (
	_ = iota
	AnnounceTypeTimeLottery
	AnnounceTypePeopleLottery
)

const (
	IsSelectedLottery  = 1
	NotSelectedLottery = 0
)

const (
	LotteryTodayParticipantsCount    = 5
	LotteryThisWeekParticipantsCount = 10
)

// ------ 抽奖 打卡任务 --------

// 任务类型
const (
	_                             = iota
	ExperienceMiniPrograms        // 体验小程序
	BrowseOfficialAccountArticles // 浏览公众号文章
	BrowseImage                   // 浏览图片
	BrowseVideo                   // 浏览视频号视频
)

// 任务秒数
const (
	ExperienceMiniProgramsSeconds        = 15
	BrowseOfficialAccountArticlesSeconds = 15
	BrowseImageSeconds                   = 6
	BrowseVideoSeconds                   = 15
)

// 任务类型文案
const (
	ExperienceMiniProgramsText        = "体验小程序%d秒"
	BrowseOfficialAccountArticlesText = "浏览指定公众号文章%d秒"
	BrowseImageText                   = "浏览图片%d秒"
	BrowseVideoText                   = "浏览视频号视频%d秒"
)
