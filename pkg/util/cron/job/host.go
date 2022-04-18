package job

import (
	"go_template/pkg/database"
	"go_template/pkg/model"
	"go_template/pkg/util/logger"
	"sync"
)

type RefreshHostInfo struct {
	hostService string
}

func NewRefreshHostInfo() *RefreshHostInfo {
	return &RefreshHostInfo{
		hostService: "TODO",
	}
}

func (r *RefreshHostInfo) Run() {
	var hosts []model.Test
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2) // 信号量
	database.DB.Find(&hosts)
	for _, host := range hosts {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			logger.Log.Info("TODO")
		}(host.Name)
	}
	wg.Wait()
}
