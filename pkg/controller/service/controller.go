package service

import (
	"context"
	"fmt"
	"github.com/alibaba/kubeskoop/pkg/controller/diagnose"
	"github.com/alibaba/kubeskoop/pkg/controller/rpc"
	skoopContext "github.com/alibaba/kubeskoop/pkg/skoop/context"
	"github.com/alibaba/kubeskoop/pkg/skoop/utils"
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"sync"
)

type ControllerService interface {
	rpc.ControllerRegisterServiceServer
	GetAgentList() []*rpc.AgentInfo
	Capture(ctx context.Context, capture *CaptureArgs) (int, error)
	CaptureList(ctx context.Context) (map[int][]*CaptureTaskResult, error)
	WatchEvents() <-chan *rpc.Event
	Diagnose(ctx context.Context, args *skoopContext.TaskConfig) (int, error)
	DiagnoseList(ctx context.Context) ([]DiagnoseTaskResult, error)
	DownloadCaptureFile(ctx context.Context, id int) (string, int64, io.ReadCloser, error)
	PodList(ctx context.Context) ([]*Pod, error)
	NodeList(ctx context.Context) ([]*Node, error)
}

func NewControllerService() (ControllerService, error) {
	ctrl := &controller{
		diagnosor:   diagnose.NewDiagnoseController(),
		taskWatcher: sync.Map{},
	}
	var (
		restConfig *rest.Config
		err        error
	)
	if os.Getenv("KUBERNETES_SERVICE_HOST") == "" {
		restConfig, _, err = utils.NewConfig("~/.kube/config")
	} else {
		restConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, fmt.Errorf("error get incluster config, err: %v", err)
	}
	ctrl.k8sClient, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error create k8s client, err: %v", err)
	}
	return ctrl, nil
}

type controller struct {
	rpc.UnimplementedControllerRegisterServiceServer
	diagnosor   diagnose.DiagnoseController
	k8sClient   *kubernetes.Clientset
	taskWatcher sync.Map
}
