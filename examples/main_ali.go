//// This file is auto-generated, don't edit it. Thanks.
package main

//
//import (
//	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
//	eci20180808 "github.com/alibabacloud-go/eci-20180808/v2/client"
//	"github.com/alibabacloud-go/tea/tea"
//	"os"
//)
//
///**
// * 使用AK&SK初始化账号Client
// * @param accessKeyId
// * @param accessKeySecret
// * @return Client
// * @throws Exception
// */
//func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *eci20180808.Client, _err error) {
//	config := &openapi.Config{
//		// 您的AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 您的AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//	}
//	// 访问的域名
//	config.Endpoint = tea.String("eci.aliyuncs.com")
//	_result = &eci20180808.Client{}
//	_result, _err = eci20180808.NewClient(config)
//	return _result, _err
//}
//
//func _main(args []*string) (_err error) {
//	client, _err := CreateClient(tea.String(""), tea.String(""))
//	if _err != nil {
//		return _err
//	}
//
//	describeContainerGroupsRequest := &eci20180808.DescribeContainerGroupsRequest{
//		RegionId: tea.String("cn-hangzhou"),
//	}
//	// 复制代码运行请自行打印 API 的返回值
//	resp, _err := client.DescribeContainerGroups(describeContainerGroupsRequest)
//	println(*resp.Body.ContainerGroups[0].ContainerGroupId)
//	if _err != nil {
//		return _err
//	}
//	return _err
//}
//
//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}
