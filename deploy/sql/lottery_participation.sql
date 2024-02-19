-- auto-generated definition
CREATE TABLE lottery_participation
(
    id         BIGINT AUTO_INCREMENT COMMENT '主键'
        PRIMARY KEY,
    lottery_id INT     NOT NULL COMMENT '参与的抽奖的id',
    user_id    INT     NOT NULL COMMENT '用户id',
    is_won     TINYINT NOT NULL COMMENT '中奖了吗？',
    prize_id   BIGINT  NOT NULL COMMENT '中奖id',
    del_state tinyint NOT NULL DEFAULT '0',
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT index_lottery_user UNIQUE (lottery_id, user_id)
)
    COMMENT '参与抽奖' COLLATE = utf8mb4_general_ci;

