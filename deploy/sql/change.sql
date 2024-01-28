ALTER TABLE user
    ADD COLUMN signature varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个性签名',
    ADD COLUMN location_name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '地址名称',
    ADD COLUMN longitude DOUBLE PRECISION COMMENT '经度',
    ADD COLUMN latitude DOUBLE PRECISION COMMENT '纬度',
    ADD COLUMN total_prize int(0) NOT NULL DEFAULT 0 COMMENT '累计奖品',
    ADD COLUMN fans int(0) NOT NULL DEFAULT 0 COMMENT '粉丝数量',
    ADD COLUMN all_lottery int(0) NOT NULL DEFAULT 0 COMMENT '全部抽奖包含我发起的、我中奖的',
    ADD COLUMN initiation_record int(0) NOT NULL DEFAULT 0 COMMENT '发起抽奖记录',
    ADD COLUMN winning_record int(0) NOT NULL DEFAULT 0 COMMENT '中奖记录';