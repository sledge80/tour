package timer

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var TimeCmd *cobra.Command

func init() {
	// 主命令
	var use = "time"
	var short = "时间计算"
	var long = "时间计算"
	var run = func(cmd *cobra.Command, args []string) {
	}

	TimeCmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   run,
	}
	// 子命令
	addChildCmd()
}

func addChildCmd() {
	var calculateTime string
	var duration string

	var calculateTimeCmd = &cobra.Command{
		Use:   "calc",
		Short: "计算所需时间",
		Long:  "计算所需时间",
		Run: func(cmd *cobra.Command, args []string) {
			var currentTimer time.Time
			var layout = "2006-01-02 15:04:05"
			if calculateTime == "" {
				currentTimer = GetNowTime()
			} else {
				var err error
				space := strings.Count(calculateTime, " ")
				if space == 0 {
					layout = "2006-01-02"
				}
				currentTimer, err = time.Parse(layout, calculateTime)
				if err != nil {
					t, _ := strconv.Atoi(calculateTime)
					currentTimer = time.Unix(int64(t), 0)
				}
			}
			t, err := GetCalculateTime(currentTimer, duration)
			if err != nil {
				log.Fatalf("timer.GetCalculateTime err: %v", err)
			}

			log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
		},
	}
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", ` 需要计算的时间，有效单位为时间戳或已格式化后的时间 `)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", ` 持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
	TimeCmd.AddCommand(calculateTimeCmd)

	nowTimeCmd := &cobra.Command{
		Use:   "now",
		Short: "当前时间",
		Long:  "当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := GetNowTime()
			log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		},
	}

	TimeCmd.AddCommand(nowTimeCmd)
}

// ------------------------------------常量定义---------------------------------------
const (
	Nanosecond  = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second      = 1000 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
)

// ------------------------------时间计算工具方法-------------------------------------

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (t time.Time, err error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	t = currentTimer.Add(duration)
	return
}
