package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	sqlTpl = "CREATE TABLE `tb_auth_%s_%s` (\n  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增唯一ID',\n  `scope` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '授权范围,目前只有0,活动粒度授权',\n  `actId` varchar(64) NOT NULL COMMENT '活动id',\n  `uid` varchar(64) NOT NULL COMMENT '渠道侧用户ID',\n  `openid` varchar(64) NOT NULL COMMENT '游戏账号openid/gopenid',\n  `sArea` varchar(8) NOT NULL DEFAULT '' COMMENT '大区',\n  `sPlatId` varchar(8) NOT NULL DEFAULT '' COMMENT '平台ID',\n  `sPartition` varchar(8) NOT NULL DEFAULT '' COMMENT '小区ID',\n  `sRoleId` varchar(64) NOT NULL DEFAULT '' COMMENT '角色ID',\n  `srcPartition` varchar(8) NOT NULL DEFAULT '' COMMENT '出生服小区ID，用于转区',\n  `sAmsNewRoleId` varchar(80) NOT NULL DEFAULT '' COMMENT '唯一角色ID，用于转区',\n  `sExtend` varchar(512) NOT NULL DEFAULT '' COMMENT '绑定的扩展信息，JSON格式',\n  `data` text COMMENT '绑定数据，JSON 格式',\n  `bindStatus` tinyint(3) unsigned NOT NULL COMMENT '绑定状态：1-已绑定，2-已解绑',\n  `bindTime` datetime NOT NULL COMMENT '绑定时间',\n  `bindSerial` char(50) NOT NULL COMMENT '绑定时的流水号',\n  `bindCount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '绑定次数',\n  `bindFrom` varchar(32) NOT NULL DEFAULT '' COMMENT '绑定来源',\n  `bindPhone` varchar(32) NOT NULL DEFAULT '' COMMENT '绑定的手机号',\n  `unbindTime` datetime DEFAULT NULL COMMENT '解绑时间',\n  `unbindSerial` char(50) NOT NULL DEFAULT '' COMMENT '解绑时的流水号',\n  `unbindCount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '解绑次数',\n  PRIMARY KEY (`id`) COMMENT '主键',\n  KEY `game_channel_openid` (`openid`) COMMENT '索引：业务-渠道-openid',\n  KEY `channel_uid` (`uid`) COMMENT '索引：渠道-uid'\n) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='渠道用户和游戏角色的绑定关系表';"
)

func GenerateSQL(gameName string) {
	url := "https://game.gtimg.cn/images/cbd/tokenlink/channels/info.json"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close() // 确保最后关闭resp.Body

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
		return
	}

	// 读取响应体内容到内存
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}

	channelInfo := make(map[string]interface{}, 0)
	json.Unmarshal(body, &channelInfo)

	fileName := "./create.sql"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()

	for k, _ := range channelInfo {
		sql := fmt.Sprintf(sqlTpl, gameName, k)
		file.WriteString(sql + "\n")
	}
}
