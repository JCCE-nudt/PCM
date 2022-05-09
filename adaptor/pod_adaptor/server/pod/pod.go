package pod

import (
	"context"
	"fmt"
	"github.com/JCCE-nudt/PCM/adaptor/pod_adaptor/server/pod/ali"
	"sync"

	"github.com/JCCE-nudt/PCM/adaptor/pod_adaptor/service"
	"github.com/JCCE-nudt/PCM/common/tenanter"
	"github.com/JCCE-nudt/PCM/lan_trans/idl/pbpod"
	"github.com/JCCE-nudt/PCM/lan_trans/idl/pbtenant"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// GetPodRegion get the available region for pod
func GetPodRegion(ctx context.Context, req *pbpod.GetPodRegionReq) (resp *pbpod.GetPodRegionResp, err error) {
	var (
		regionInit tenanter.Region
		regions    []*pbtenant.Region
	)

	switch req.GetProvider() {
	case pbtenant.CloudProvider_ali:
		regionInit, _ = tenanter.NewRegion(req.GetProvider(), 2)
	case pbtenant.CloudProvider_tencent:
		regionInit, _ = tenanter.NewRegion(req.GetProvider(), 5)
	case pbtenant.CloudProvider_huawei:
		regionInit, _ = tenanter.NewRegion(req.GetProvider(), 5)
	}
	tenanters, err := tenanter.GetTenanters(req.GetProvider())
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	for _, tenant := range tenanters {

		pod, err := poder.NewPodClient(req.GetProvider(), regionInit, tenant)
		if err != nil {
			return nil, errors.WithMessage(err, "NewPodClient error")
		}
		request := &pbpod.GetPodRegionReq{
			Provider: req.GetProvider(),
		}
		resp, err := pod.GetPodRegion(ctx, request)
		if err != nil {
			return nil, errors.Wrap(err, "GetPodRegion error")
		}
		for _, region := range resp.GetRegions() {
			regions = append(regions, region)
		}
	}

	return &pbpod.GetPodRegionResp{Regions: regions}, nil
}

func CreatePods(ctx context.Context, req *pbpod.CreatePodsReq) (*pbpod.CreatePodsResp, error) {
	var (
		wg         sync.WaitGroup
		requestIds = make([]string, 0)
	)
	wg.Add(len(req.CreatePodReq))
	c := make(chan string)
	for k := range req.CreatePodReq {
		reqPod := req.CreatePodReq[k]
		go func() {
			defer wg.Done()
			resp, err := CreatePod(ctx, reqPod)
			if err != nil || resp == nil {
				fmt.Println(errors.Wrap(err, "Batch pod creation error"))
				return
			}
			c <- resp.RequestId
		}()

	}
	go func() {
		defer close(c)
		wg.Wait()
	}()
	isFinished := false
	if len(requestIds) > 0 {
		isFinished = true
	}

	for v := range c {
		requestIds = append(requestIds, v)
	}

	return &pbpod.CreatePodsResp{
		Finished:  isFinished,
		RequestId: requestIds,
	}, nil
}

func CreatePod(ctx context.Context, req *pbpod.CreatePodReq) (*pbpod.CreatePodResp, error) {
	var (
		pod poder.Poder
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	region, err := tenanter.NewRegion(req.Provider, req.RegionId)
	if err != nil {
		return nil, errors.WithMessagef(err, "provider %v regionId %v", req.Provider, req.RegionId)
	}

	for _, tenant := range tenanters {
		if req.AccountName == "" || tenant.AccountName() == req.AccountName {
			if pod, err = poder.NewPodClient(req.Provider, region, tenant); err != nil {
				return nil, errors.WithMessage(err, "NewPodClient error")
			}
			break
		}
	}

	return pod.CreatePod(ctx, req)
}

func DeletePod(ctx context.Context, req *pbpod.DeletePodReq) (*pbpod.DeletePodResp, error) {
	var (
		pod poder.Poder
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	region, err := tenanter.NewRegion(req.Provider, req.RegionId)
	if err != nil {
		return nil, errors.WithMessagef(err, "provider %v regionId %v", req.Provider, req.RegionId)
	}

	for _, tenant := range tenanters {
		if req.AccountName == "" || tenant.AccountName() == req.AccountName {
			if pod, err = poder.NewPodClient(req.Provider, region, tenant); err != nil {
				return nil, errors.WithMessage(err, "NewPodClient error")
			}
			break
		}
	}

	return pod.DeletePod(ctx, req)
}

func UpdatePod(ctx context.Context, req *pbpod.UpdatePodReq) (*pbpod.UpdatePodResp, error) {
	var (
		pod poder.Poder
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	region, err := tenanter.NewRegion(req.Provider, req.RegionId)
	if err != nil {
		return nil, errors.WithMessagef(err, "provider %v regionId %v", req.Provider, req.RegionId)
	}

	for _, tenant := range tenanters {
		if req.AccountName == "" || tenant.AccountName() == req.AccountName {
			if pod, err = poder.NewPodClient(req.Provider, region, tenant); err != nil {
				return nil, errors.WithMessage(err, "NewPodClient error")
			}
			break
		}
	}

	return pod.UpdatePod(ctx, req)
}

func ListPodDetail(ctx context.Context, req *pbpod.ListPodDetailReq) (*pbpod.ListPodDetailResp, error) {
	var (
		pod poder.Poder
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	region, err := tenanter.NewRegion(req.Provider, req.RegionId)
	if err != nil {
		return nil, errors.WithMessagef(err, "provider %v regionId %v", req.Provider, req.RegionId)
	}

	for _, tenant := range tenanters {
		if req.AccountName == "" || tenant.AccountName() == req.AccountName {
			if pod, err = poder.NewPodClient(req.Provider, region, tenant); err != nil {
				return nil, errors.WithMessage(err, "NewPodClient error")
			}
			break
		}
	}

	return pod.ListPodDetail(ctx, req)
}

// CreateContainerGroup invokes the eci.CreateContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/createcontainergroup.html
func CreateContainerGroup(request *ali.CreateContainerGroupRequest) (response *ali.CreateContainerGroupResponse, err error) {

	provider := pbtenant.CloudProvider(request.ProviderId)
	tenanters, err := tenanter.GetTenanters(provider)
	regionId, err := tenanter.GetAliRegionId(request.RegionId)
	container := *request.Container
	containerImage := container[0].Image
	containerName := container[0].Name
	containerPod := container[0].Cpu
	memoryPod := container[0].Memory

	requestPCM := &pbpod.CreatePodReq{
		Provider:        provider,
		AccountName:     tenanters[0].AccountName(),
		PodName:         request.ContainerGroupName,
		RegionId:        regionId,
		ContainerImage:  containerImage,
		ContainerName:   containerName,
		CpuPod:          string(containerPod),
		MemoryPod:       string(memoryPod),
		SecurityGroupId: "sg-6qlun7hd",
		SubnetId:        "subnet-mnwfg2fk",
		VpcId:           "vpc-rkwt40g5",
		Namespace:       "pcm",
	}

	resp, err := CreatePod(nil, requestPCM)

	response = &ali.CreateContainerGroupResponse{
		BaseResponse:     nil,
		RequestId:        resp.RequestId,
		ContainerGroupId: resp.PodId,
	}

	return response, nil
}

// DeleteContainerGroup invokes the eci.DeleteContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/deletecontainergroup.html
func DeleteContainerGroup(request *ali.DeleteContainerGroupRequest) (response *ali.DeleteContainerGroupResponse, err error) {

	provider := pbtenant.CloudProvider(request.ProviderId)
	regionId, err := tenanter.GetAliRegionId(request.RegionId)
	podId := request.ContainerGroupId
	podName := request.ContainerGroupName

	requestPCM := &pbpod.DeletePodReq{
		Provider:    provider,
		AccountName: request.AccountName,
		PodId:       podId,
		PodName:     podName,
		Namespace:   request.Namespace,
		RegionId:    regionId,
	}

	resp, err := DeletePod(nil, requestPCM)

	response = &ali.DeleteContainerGroupResponse{
		BaseResponse: nil,
		RequestId:    resp.RequestId,
	}

	return response, err

}

// DescribeContainerGroups invokes the eci.DescribeContainerGroups API synchronously
// api document: https://help.aliyun.com/api/eci/describecontainergroups.html
func DescribeContainerGroups(request *ali.DescribeContainerGroupsRequest) (response *ali.DescribeContainerGroupsResponse, err error) {

	provider := pbtenant.CloudProvider(request.ProviderId)
	containerGroups := make([]ali.DescribeContainerGroupsContainerGroup0, 0)
	requestPCM := &pbpod.ListPodReq{
		Provider:  provider,
		Namespace: "pcm",
	}

	resp, err := ListPod(nil, requestPCM)

	//trans PCM response pod set to Ali ContainerGroup set
	for k := range resp.Pods {
		podId := resp.Pods[k].PodId
		podName := resp.Pods[k].PodName

		containerGroup := new(ali.DescribeContainerGroupsContainerGroup0)
		containerGroup.ContainerGroupName = podName
		containerGroup.ContainerGroupId = podId

		containerGroups = append(containerGroups, *containerGroup)

	}

	response = &ali.DescribeContainerGroupsResponse{
		BaseResponse:    nil,
		RequestId:       "",
		NextToken:       "",
		TotalCount:      0,
		ContainerGroups: containerGroups,
	}

	return response, nil
}

// UpdateContainerGroup invokes the eci.UpdateContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/updatecontainergroup.html
func UpdateContainerGroup(request *ali.UpdateContainerGroupRequest) (response *ali.UpdateContainerGroupResponse, err error) {

	provider := pbtenant.CloudProvider(request.ProviderId)
	regionId, err := tenanter.GetAliRegionId(request.RegionId)
	containers := *request.Container

	requestPCM := &pbpod.UpdatePodReq{
		Provider:       provider,
		AccountName:    request.AccountName,
		PodId:          request.ContainerGroupId,
		PodName:        request.ContainerGroupName,
		Namespace:      request.Namespace,
		RegionId:       regionId,
		ContainerImage: containers[0].Image,
		ContainerName:  containers[0].Name,
		CpuPod:         string(containers[0].Cpu),
		MemoryPod:      string(containers[0].Memory),
		RestartPolicy:  request.RestartPolicy,
		Labels:         "",
	}

	resp, err := UpdatePod(nil, requestPCM)

	response = &ali.UpdateContainerGroupResponse{
		BaseResponse: nil,
		RequestId:    resp.RequestId,
	}

	return response, err
}

func ListPod(ctx context.Context, req *pbpod.ListPodReq) (*pbpod.ListPodResp, error) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		pods  []*pbpod.PodInstance
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}
	//get the available region for product
	reqPodRegion := &pbpod.GetPodRegionReq{Provider: req.GetProvider()}
	respPodRegion, err := GetPodRegion(ctx, reqPodRegion)
	if err != nil {
		return nil, errors.WithMessage(err, "getPodRegion error")
	}

	wg.Add(len(tenanters) * len(respPodRegion.Regions))
	for _, t := range tenanters {
		for _, region := range respPodRegion.Regions {
			go func(tenant tenanter.Tenanter, region tenanter.Region) {
				defer wg.Done()
				pod, err := poder.NewPodClient(req.Provider, region, tenant)
				if err != nil {
					glog.Errorf("New Pod Client error %v", err)
					return
				}

				request := &pbpod.ListPodDetailReq{
					Provider:    req.Provider,
					AccountName: tenant.AccountName(),
					RegionId:    region.GetId(),
					Namespace:   req.Namespace,
					PageNumber:  1,
					PageSize:    100,
					NextToken:   "",
				}
				for {
					resp, err := pod.ListPodDetail(ctx, request)
					if err != nil {
						glog.Errorf("ListDetail error %v", err)
						return
					}
					mutex.Lock()
					pods = append(pods, resp.Pods...)
					mutex.Unlock()
					if resp.Finished {
						break
					}
					request.PageNumber, request.PageSize, request.NextToken = resp.PageNumber, resp.PageSize, resp.NextToken
				}
			}(t, region)

		}
	}
	wg.Wait()

	return &pbpod.ListPodResp{Pods: pods}, nil
}

func ListPodAll(ctx context.Context) (*pbpod.ListPodResp, error) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		pods  []*pbpod.PodInstance
	)

	wg.Add(len(pbtenant.CloudProvider_name))
	for k := range pbtenant.CloudProvider_name {
		go func(provider int32) {
			defer wg.Done()

			//针对私有K8S集群，调用listAll时默认只查询ListPodDetailReq namespace下的pod
			if provider == 3 {
				resp, err := ListPod(ctx, &pbpod.ListPodReq{Provider: pbtenant.CloudProvider(provider), Namespace: "pcm"})
				if err != nil {
					glog.Errorf("List error %v", err)
					return
				}
				mutex.Lock()
				pods = append(pods, resp.Pods...)
				mutex.Unlock()
			} else {
				resp, err := ListPod(ctx, &pbpod.ListPodReq{Provider: pbtenant.CloudProvider(provider)})
				if err != nil {
					glog.Errorf("List error %v", err)
					return
				}
				mutex.Lock()
				pods = append(pods, resp.Pods...)
				mutex.Unlock()
			}

		}(k)
	}

	wg.Wait()

	return &pbpod.ListPodResp{Pods: pods}, nil
}
