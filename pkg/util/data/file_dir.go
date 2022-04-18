package data

import (
	"go_template/pkg/constant"
	"os"

	"go.etcd.io/etcd/pkg/fileutil"
)

var initDirs = []string{
	constant.DefaultDataDir,
}

const phaseName = "create data dir"

type InitDataPhase struct{}

func (i *InitDataPhase) Init() error {
	for _, d := range initDirs {
		if !fileutil.Exist(d) {
			err := os.MkdirAll(d, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *InitDataPhase) PhaseName() string {
	return phaseName
}
