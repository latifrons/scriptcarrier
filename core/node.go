package core

import (
	"github.com/latifrons/scriptcarrier/rpc"
	"github.com/latifrons/scriptcarrier/tools"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type Node struct {
	FolderConfig tools.FolderConfig
	components   []Component
}

func (n *Node) Setup() {
	server := &rpc.RpcServer{
		C: &rpc.RpcController{
			FolderConfig:               n.FolderConfig,
			ReturnDetailedErrorMessage: viper.GetBool("debug.return_detailed_error"),
			Mode:                       viper.GetString("common.mode"),
			AllowOrigins:               strings.Split(viper.GetString("rpc.allow_origins"), ","),
		},
		Port: viper.GetString("rpc.port"),
	}
	server.InitDefault()

	n.components = append(n.components, server)
}

func (n *Node) Start() {
	for _, component := range n.components {
		logrus.Infof("Starting %s", component.Name())
		component.Start()
		logrus.Infof("Started: %s", component.Name())

	}
	//n.heightEventChan <- 10943851
	logrus.Info("Node Started")

}

func (n *Node) Stop() {
	for i := len(n.components) - 1; i >= 0; i-- {
		comp := n.components[i]
		logrus.Infof("Stopping %s", comp.Name())
		comp.Stop()
		logrus.Infof("Stopped: %s", comp.Name())
	}
	logrus.Info("Node Stopped")
}
