package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/spf13/cobra"
)

var (
	magnet string
)

var downloadCmd = &cobra.Command{
	Use: "download",
	Run: func(cmd *cobra.Command, args []string) {
		if magnet == "" {
			fmt.Println("请输入磁力链接")
			return
		}
		dd(magnet)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&magnet, "magnet", "m", "", "magnet address")
}

func dd(magnet string) {
	// 这一行替代了手动 signal.Notify + channel 的一堆模板代码
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop() // 确保退出时释放信号监听

	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = getDdPath()

	client, err := torrent.NewClient(cfg)
	if err != nil {
		fmt.Println("下载客户端初始化失败")
		os.Exit(1)
	}
	defer client.Close()

	t, err := client.AddMagnet(magnet)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 元信息获取：ctx 作为父 context，Ctrl+C 和超时都能中断
	fmt.Println("正在获取元信息...")
	metaCtx, metaCancel := context.WithTimeout(ctx, 2*time.Minute)
	defer metaCancel()

	select {
	case <-t.GotInfo():
	case <-metaCtx.Done():
		if ctx.Err() != nil {
			fmt.Println("\n用户取消，正在退出...")
			return
		}
		fmt.Println("获取元信息超时")
		return
	}

	t.DownloadAll()
	fmt.Printf("开始下载: %s\n", t.Info().Name)

	// 进度打印：也用 ctx 控制，Ctrl+C 时自动停
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				total := t.Info().TotalLength()
				if total > 0 {
					fmt.Printf("\r进度: %.2f%% | peers: %d",
						float64(t.BytesCompleted())/float64(total)*100,
						t.Stats().ActivePeers,
					)
				}
			}
		}
	}()

	// 等下载完成 或 Ctrl+C
	done := make(chan struct{})
	go func() {
		client.WaitAll()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("\n下载完成")
	case <-ctx.Done():
		fmt.Println("\n正在停止下载...")
		t.Drop()
		fmt.Println("已退出")
	}
}

func getDdPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return filepath.Join(home, "Downloads")
}
