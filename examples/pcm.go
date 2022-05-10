package main

import (
	pcm_pod_ali "github.com/JCCE-nudt/PCM/adaptor/pod/server/ali"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

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
