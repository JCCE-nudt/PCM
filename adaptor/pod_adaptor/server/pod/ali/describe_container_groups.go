package ali

import (
	"github.com/JCCE-nudt/PCM/adaptor/pod_adaptor/server/pod"
	"github.com/JCCE-nudt/PCM/lan_trans/idl/pbpod"
	"github.com/JCCE-nudt/PCM/lan_trans/idl/pbtenant"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeContainerGroups invokes the eci.DescribeContainerGroups API synchronously
// api document: https://help.aliyun.com/api/eci/describecontainergroups.html
func DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {

	provider := pbtenant.CloudProvider(request.ProviderId)
	containerGroups := make([]DescribeContainerGroupsContainerGroup0, 0)
	//trans Ali request to  PCM request
	requestPCM := &pbpod.ListPodReq{
		RequestSource: "ali",
		Provider:      provider,
		Namespace:     "pcm",
	}

	resp, err := pod.ListPod(nil, requestPCM)

	//trans PCM response pod set to Ali ContainerGroup set
	for k := range resp.Pods {
		podId := resp.Pods[k].PodId
		podName := resp.Pods[k].PodName

		containerGroup := new(DescribeContainerGroupsContainerGroup0)
		containerGroup.ContainerGroupName = podName
		containerGroup.ContainerGroupId = podId

		containerGroups = append(containerGroups, *containerGroup)

	}

	response = &DescribeContainerGroupsResponse{
		BaseResponse:    nil,
		RequestId:       "",
		NextToken:       "",
		TotalCount:      0,
		ContainerGroups: containerGroups,
	}

	return response, nil
}

// DescribeContainerGroupsRequest is the request struct for api DescribeContainerGroups
type DescribeContainerGroupsRequest struct {
	*requests.RpcRequest
	/*********PCM param************/
	RequestSource string `position:"Query" name:"RequestSource"`
	ProviderId    int32  `position:"Query" name:"ProviderId"`
	AccountName   string `position:"Query" name:"AccountName"`
	Namespace     string `position:"Query" name:"Namespace"`
	/*********PCM param************/
	OwnerId              requests.Integer              `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	RegionId             string                        `position:"Query" name:"RegionId"`
	ZoneId               string                        `position:"Query" name:"ZoneId"`
	VSwitchId            string                        `position:"Query" name:"VSwitchId"`
	NextToken            string                        `position:"Query" name:"NextToken"`
	Limit                requests.Integer              `position:"Query" name:"Limit"`
	Tag                  *[]DescribeContainerGroupsTag `position:"Query" name:"Tag" type:"Repeated"`
	ContainerGroupIds    string                        `position:"Query" name:"ContainerGroupIds"`
	ContainerGroupName   string                        `position:"Query" name:"ContainerGroupName"`
	Status               string                        `position:"Query" name:"Status"`
	VkClientVersion      string                        `position:"Query" name:"VkClientVersion"`
	ResourceGroupId      string                        `position:"Query" name:"ResourceGroupId"`
	WithEvent            requests.Boolean              `position:"Query" name:"WithEvent"`
}

type DescribeContainerGroupsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeContainerGroupsResponse is the response struct for api DescribeContainerGroups
type DescribeContainerGroupsResponse struct {
	*responses.BaseResponse
	RequestId       string                                   `json:"RequestId" xml:"RequestId"`
	NextToken       string                                   `json:"NextToken" xml:"NextToken"`
	TotalCount      int                                      `json:"TotalCount" xml:"TotalCount"`
	ContainerGroups []DescribeContainerGroupsContainerGroup0 `json:"ContainerGroups" xml:"ContainerGroups"`
}

type DescribeContainerGroupsContainerGroup0 struct {
	ContainerGroupId   string                                     `json:"ContainerGroupId" xml:"ContainerGroupId"`
	ContainerGroupName string                                     `json:"ContainerGroupName" xml:"ContainerGroupName"`
	RegionId           string                                     `json:"RegionId" xml:"RegionId"`
	ZoneId             string                                     `json:"ZoneId" xml:"ZoneId"`
	Memory             float32                                    `json:"Memory" xml:"Memory"`
	Cpu                float32                                    `json:"Cpu" xml:"Cpu"`
	VSwitchId          string                                     `json:"VSwitchId" xml:"VSwitchId"`
	SecurityGroupId    string                                     `json:"SecurityGroupId" xml:"SecurityGroupId"`
	RestartPolicy      string                                     `json:"RestartPolicy" xml:"RestartPolicy"`
	IntranetIp         string                                     `json:"IntranetIp" xml:"IntranetIp"`
	Status             string                                     `json:"Status" xml:"Status"`
	InternetIp         string                                     `json:"InternetIp" xml:"InternetIp"`
	CreationTime       string                                     `json:"CreationTime" xml:"CreationTime"`
	SucceededTime      string                                     `json:"SucceededTime" xml:"SucceededTime"`
	EniInstanceId      string                                     `json:"EniInstanceId" xml:"EniInstanceId"`
	InstanceType       string                                     `json:"InstanceType" xml:"InstanceType"`
	ExpiredTime        string                                     `json:"ExpiredTime" xml:"ExpiredTime"`
	FailedTime         string                                     `json:"FailedTime" xml:"FailedTime"`
	RamRoleName        string                                     `json:"RamRoleName" xml:"RamRoleName"`
	Ipv6Address        string                                     `json:"Ipv6Address" xml:"Ipv6Address"`
	VpcId              string                                     `json:"VpcId" xml:"VpcId"`
	Discount           int                                        `json:"Discount" xml:"Discount"`
	ResourceGroupId    string                                     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags               []DescribeContainerGroupsLabel1            `json:"Tags" xml:"Tags"`
	Events             []DescribeContainerGroupsEvent1            `json:"Events" xml:"Events"`
	Containers         []DescribeContainerGroupsContainer1        `json:"Containers" xml:"Containers"`
	Volumes            []DescribeContainerGroupsVolume1           `json:"Volumes" xml:"Volumes"`
	InitContainers     []DescribeContainerGroupsContainer1        `json:"InitContainers" xml:"InitContainers"`
	HostAliases        []DescribeContainerGroupsHostAliase1       `json:"HostAliases" xml:"HostAliases"`
	DnsConfig          DescribeContainerGroupsDnsConfig1          `json:"DnsConfig" xml:"DnsConfig"`
	EciSecurityContext DescribeContainerGroupsEciSecurityContext1 `json:"EciSecurityContext" xml:"EciSecurityContext"`
}

type DescribeContainerGroupsLabel1 struct {
	Key   string `json:"Key" xml:"Key"`
	Value string `json:"Value" xml:"Value"`
}

type DescribeContainerGroupsEvent1 struct {
	Count          int    `json:"Count" xml:"Count"`
	Type           string `json:"Type" xml:"Type"`
	Name           string `json:"Name" xml:"Name"`
	Message        string `json:"Message" xml:"Message"`
	FirstTimestamp string `json:"FirstTimestamp" xml:"FirstTimestamp"`
	LastTimestamp  string `json:"LastTimestamp" xml:"LastTimestamp"`
	Reason         string `json:"Reason" xml:"Reason"`
}

type DescribeContainerGroupsContainer1 struct {
	Name            string                                   `json:"Name" xml:"Name"`
	Image           string                                   `json:"Image" xml:"Image"`
	Memory          float32                                  `json:"Memory" xml:"Memory"`
	Cpu             float32                                  `json:"Cpu" xml:"Cpu"`
	RestartCount    int                                      `json:"RestartCount" xml:"RestartCount"`
	WorkingDir      string                                   `json:"WorkingDir" xml:"WorkingDir"`
	ImagePullPolicy string                                   `json:"ImagePullPolicy" xml:"ImagePullPolicy"`
	Ready           bool                                     `json:"Ready" xml:"Ready"`
	Gpu             int                                      `json:"Gpu" xml:"Gpu"`
	Stdin           bool                                     `json:"Stdin" xml:"Stdin"`
	StdinOnce       bool                                     `json:"StdinOnce" xml:"StdinOnce"`
	Tty             bool                                     `json:"Tty" xml:"Tty"`
	VolumeMounts    []DescribeContainerGroupsVolumeMount2    `json:"VolumeMounts" xml:"VolumeMounts"`
	Ports           []DescribeContainerGroupsPort2           `json:"Ports" xml:"Ports"`
	EnvironmentVars []DescribeContainerGroupsEnvironmentVar2 `json:"EnvironmentVars" xml:"EnvironmentVars"`
	Commands        []string                                 `json:"Commands" xml:"Commands"`
	Args            []string                                 `json:"Args" xml:"Args"`
	PreviousState   DescribeContainerGroupsPreviousState2    `json:"PreviousState" xml:"PreviousState"`
	CurrentState    DescribeContainerGroupsCurrentState2     `json:"CurrentState" xml:"CurrentState"`
	ReadinessProbe  DescribeContainerGroupsReadinessProbe2   `json:"ReadinessProbe" xml:"ReadinessProbe"`
	LivenessProbe   DescribeContainerGroupsLivenessProbe2    `json:"LivenessProbe" xml:"LivenessProbe"`
	SecurityContext DescribeContainerGroupsSecurityContext2  `json:"SecurityContext" xml:"SecurityContext"`
}

type DescribeContainerGroupsVolumeMount2 struct {
	Name      string `json:"Name" xml:"Name"`
	MountPath string `json:"MountPath" xml:"MountPath"`
	ReadOnly  bool   `json:"ReadOnly" xml:"ReadOnly"`
}

type DescribeContainerGroupsPort2 struct {
	Port     int    `json:"Port" xml:"Port"`
	Protocol string `json:"Protocol" xml:"Protocol"`
}

type DescribeContainerGroupsEnvironmentVar2 struct {
	Key       string                            `json:"Key" xml:"Key"`
	Value     string                            `json:"Value" xml:"Value"`
	ValueFrom DescribeContainerGroupsValueFrom3 `json:"ValueFrom" xml:"ValueFrom"`
}

type DescribeContainerGroupsValueFrom3 struct {
	FieldRef DescribeContainerGroupsFieldRef4 `json:"FieldRef" xml:"FieldRef"`
}

type DescribeContainerGroupsFieldRef4 struct {
	FieldPath string `json:"FieldPath" xml:"FieldPath"`
}

type DescribeContainerGroupsPreviousState2 struct {
	State        string `json:"State" xml:"State"`
	DetailStatus string `json:"DetailStatus" xml:"DetailStatus"`
	ExitCode     int    `json:"ExitCode" xml:"ExitCode"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	FinishTime   string `json:"FinishTime" xml:"FinishTime"`
	Reason       string `json:"Reason" xml:"Reason"`
	Message      string `json:"Message" xml:"Message"`
	Signal       int    `json:"Signal" xml:"Signal"`
}

type DescribeContainerGroupsCurrentState2 struct {
	State        string `json:"State" xml:"State"`
	DetailStatus string `json:"DetailStatus" xml:"DetailStatus"`
	ExitCode     int    `json:"ExitCode" xml:"ExitCode"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	FinishTime   string `json:"FinishTime" xml:"FinishTime"`
	Reason       string `json:"Reason" xml:"Reason"`
	Message      string `json:"Message" xml:"Message"`
	Signal       int    `json:"Signal" xml:"Signal"`
}

type DescribeContainerGroupsReadinessProbe2 struct {
	InitialDelaySeconds int                               `json:"InitialDelaySeconds" xml:"InitialDelaySeconds"`
	PeriodSeconds       int                               `json:"PeriodSeconds" xml:"PeriodSeconds"`
	TimeoutSeconds      int                               `json:"TimeoutSeconds" xml:"TimeoutSeconds"`
	SuccessThreshold    int                               `json:"SuccessThreshold" xml:"SuccessThreshold"`
	FailureThreshold    int                               `json:"FailureThreshold" xml:"FailureThreshold"`
	Execs               []string                          `json:"Execs" xml:"Execs"`
	HttpGet             DescribeContainerGroupsHttpGet3   `json:"HttpGet" xml:"HttpGet"`
	TcpSocket           DescribeContainerGroupsTcpSocket3 `json:"TcpSocket" xml:"TcpSocket"`
}

type DescribeContainerGroupsHttpGet3 struct {
	Path   string `json:"Path" xml:"Path"`
	Port   int    `json:"Port" xml:"Port"`
	Scheme string `json:"Scheme" xml:"Scheme"`
}

type DescribeContainerGroupsTcpSocket3 struct {
	Host string `json:"Host" xml:"Host"`
	Port int    `json:"Port" xml:"Port"`
}

type DescribeContainerGroupsLivenessProbe2 struct {
	InitialDelaySeconds int                               `json:"InitialDelaySeconds" xml:"InitialDelaySeconds"`
	PeriodSeconds       int                               `json:"PeriodSeconds" xml:"PeriodSeconds"`
	TimeoutSeconds      int                               `json:"TimeoutSeconds" xml:"TimeoutSeconds"`
	SuccessThreshold    int                               `json:"SuccessThreshold" xml:"SuccessThreshold"`
	FailureThreshold    int                               `json:"FailureThreshold" xml:"FailureThreshold"`
	Execs               []string                          `json:"Execs" xml:"Execs"`
	HttpGet             DescribeContainerGroupsHttpGet3   `json:"HttpGet" xml:"HttpGet"`
	TcpSocket           DescribeContainerGroupsTcpSocket3 `json:"TcpSocket" xml:"TcpSocket"`
}

type DescribeContainerGroupsSecurityContext2 struct {
	ReadOnlyRootFilesystem bool                               `json:"ReadOnlyRootFilesystem" xml:"ReadOnlyRootFilesystem"`
	RunAsUser              int64                              `json:"RunAsUser" xml:"RunAsUser"`
	Capability             DescribeContainerGroupsCapability3 `json:"Capability" xml:"Capability"`
}

type DescribeContainerGroupsCapability3 struct {
	Adds []string `json:"Adds" xml:"Adds"`
}

type DescribeContainerGroupsVolume1 struct {
	Type                              string                                                     `json:"Type" xml:"Type"`
	Name                              string                                                     `json:"Name" xml:"Name"`
	NFSVolumePath                     string                                                     `json:"NFSVolumePath" xml:"NFSVolumePath"`
	NFSVolumeServer                   string                                                     `json:"NFSVolumeServer" xml:"NFSVolumeServer"`
	NFSVolumeReadOnly                 bool                                                       `json:"NFSVolumeReadOnly" xml:"NFSVolumeReadOnly"`
	DiskVolumeDiskId                  string                                                     `json:"DiskVolumeDiskId" xml:"DiskVolumeDiskId"`
	DiskVolumeFsType                  string                                                     `json:"DiskVolumeFsType" xml:"DiskVolumeFsType"`
	FlexVolumeDriver                  string                                                     `json:"FlexVolumeDriver" xml:"FlexVolumeDriver"`
	FlexVolumeFsType                  string                                                     `json:"FlexVolumeFsType" xml:"FlexVolumeFsType"`
	FlexVolumeOptions                 string                                                     `json:"FlexVolumeOptions" xml:"FlexVolumeOptions"`
	ConfigFileVolumeConfigFileToPaths []DescribeContainerGroupsConfigFileVolumeConfigFileToPath2 `json:"ConfigFileVolumeConfigFileToPaths" xml:"ConfigFileVolumeConfigFileToPaths"`
}

type DescribeContainerGroupsConfigFileVolumeConfigFileToPath2 struct {
	Content string `json:"Content" xml:"Content"`
	Path    string `json:"Path" xml:"Path"`
}

type DescribeContainerGroupsHostAliase1 struct {
	Ip        string   `json:"Ip" xml:"Ip"`
	Hostnames []string `json:"Hostnames" xml:"Hostnames"`
}

type DescribeContainerGroupsDnsConfig1 struct {
	Options     []DescribeContainerGroupsOption2 `json:"Options" xml:"Options"`
	NameServers []string                         `json:"NameServers" xml:"NameServers"`
	Searches    []string                         `json:"Searches" xml:"Searches"`
}

type DescribeContainerGroupsOption2 struct {
	Name  string `json:"Name" xml:"Name"`
	Value string `json:"Value" xml:"Value"`
}

type DescribeContainerGroupsEciSecurityContext1 struct {
	Sysctls []DescribeContainerGroupsSysctl2 `json:"Sysctls" xml:"Sysctls"`
}

type DescribeContainerGroupsSysctl2 struct {
	Name  string `json:"Name" xml:"Name"`
	Value string `json:"Value" xml:"Value"`
}

// CreateDescribeContainerGroupsRequest creates a request to invoke DescribeContainerGroups API
func CreateDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
	request = &DescribeContainerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "DescribeContainerGroups", "eci", "openAPI")
	return
}

// CreateDescribeContainerGroupsResponse creates a response to parse from DescribeContainerGroups response
func CreateDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
	response = &DescribeContainerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
