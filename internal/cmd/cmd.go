package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"otter-dingtalk/internal/core"
	"otter-dingtalk/internal/global"
	"otter-dingtalk/internal/otter"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "otter-dingtalk 报警工具",
	Short: "otter-dingtalk",
	Long:  `otter-dingtalk`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Init()
		otter.MonitoringPatrol()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&global.GL_MYSQL_HOST, "mysql_host", "127.0.0.1:3306", "mysql地址")
	rootCmd.PersistentFlags().StringVar(&global.GL_MYSQL_DB, "mysql_db", "otter", "mysql数据库")
	rootCmd.PersistentFlags().StringVar(&global.GL_MYSQL_USER, "mysql_user", "otter", "mysql用户名")
	rootCmd.PersistentFlags().StringVar(&global.GL_MYSQL_PASS, "mysql_pass", "123456", "mysql密码")
	rootCmd.PersistentFlags().StringVar(&global.ZKADDR, "zk", "127.0.0.1:2181", "Zookeeper connection address")
	rootCmd.PersistentFlags().StringVar(&global.DINGTOKEN, "token", "", "钉钉机器人token")
	rootCmd.PersistentFlags().StringVar(&global.DINGSECRTE, "secret", "", "钉钉机器人secret")
	rootCmd.PersistentFlags().StringVar(&global.OTTER_URL, "otter-url", "", "otter 地址")
	rootCmd.PersistentFlags().StringVar(&global.OTTER_USERNAME, "otter-username", "", "otter 账号")
	rootCmd.PersistentFlags().StringVar(&global.OTTER_PASSWORD, "otter-password", "", "otter 密码")
	rootCmd.PersistentFlags().StringVar(&global.ENV, "env", "dev", "环境标识")
	rootCmd.PersistentFlags().StringArrayVarP(&global.CHANNEL, "channel", "c", []string{"28"}, "指定监控的channelId,all为全部")
	rootCmd.PersistentFlags().DurationVarP(&global.REFRESHTIME, "refresh", "", time.Second*60, "刷新间隔时间")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("secret")
}

func Execute() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
