package main

import (
	"flag"
	pcm_pod_ali "github.com/JCCE-nudt/PCM/adaptor/pod_adaptor/server/pod/ali"
	pcm_service "github.com/JCCE-nudt/PCM/adaptor/pod_adaptor/service"
	"github.com/JCCE-nudt/PCM/common/tenanter"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"os"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient() (pod pcm_service.Poder, _err error) {
	var configFile string
	flag.StringVar(&configFile, "conf", "configs/config.yaml", "config.yaml")
	flag.Parse()
	defer glog.Flush()

	if err := tenanter.LoadCloudConfigsFromFile(configFile); err != nil {
		if !errors.Is(err, tenanter.ErrLoadTenanterFileEmpty) {
			glog.Fatalf("LoadCloudConfigsFromFile error %+v", err)
		}
		glog.Warningf("LoadCloudConfigsFromFile empty file path %s", configFile)
	}

	glog.Infof("load tenant from file finished")
	tenanters, _ := tenanter.GetTenanters(0)
	region, _ := tenanter.NewRegion(0, 6)

	pod, _ = pcm_service.NewPodClient(0, region, tenanters[0])

	return
}

func _main(args []*string) (_err error) {

	//查询
	//describeContainerGroupsRequest := pcm_pod_ali.DescribeContainerGroupsRequest{
	//	RegionId:   "cn-hangzhou",
	//	ProviderId: 1,
	//}
	//
	//// 复制代码运行请自行打印 API 的返回值
	//resp, _err := pcm_pod_ali.DescribeContainerGroups(&describeContainerGroupsRequest)
	//println(resp.ContainerGroups[0].ContainerGroupId)

	//创建
	container := pcm_pod_ali.CreateContainerGroupContainer{
		Image:  "nginx:latest",
		Name:   "pcm-sdk-ali",
		Cpu:    "1",
		Memory: "2",
	}
	containers := make([]pcm_pod_ali.CreateContainerGroupContainer, 0)
	containers = append(containers, container)

	createContainerGroupsRequest := pcm_pod_ali.CreateContainerGroupRequest{
		RpcRequest: nil,
		//ali:cn-hangzhou  tc:ap-shanghai  hw:  k8s
		RegionId:           "cn-hangzhou",
		ProviderId:         3,
		Namespace:          "pcm",
		ContainerGroupName: "sdk-alitohuawei",
		Container:          &containers,
	}

	// 复制代码运行请自行打印 API 的返回值
	resp, _err := pcm_pod_ali.CreateContainerGroup(&createContainerGroupsRequest)

	println(resp.ContainerGroupId)
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
