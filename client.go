package aliyun

import (
	adbclient "github.com/alibabacloud-go/adb-20190315/v2/client"
	albclient "github.com/alibabacloud-go/alb-20200616/v2/client"
	amqpopenclient "github.com/alibabacloud-go/amqp-open-20191212/v2/client"
	bssclient "github.com/alibabacloud-go/bssopenapi-20171214/v3/client"
	cmsclient "github.com/alibabacloud-go/cms-20190101/v8/client"
	crclient "github.com/alibabacloud-go/cr-20181201/v2/client"
	csclient "github.com/alibabacloud-go/cs-20151215/v3/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	mongoclient "github.com/alibabacloud-go/dds-20151201/v4/client"
	dmsenterpriseclient "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	ecsclient "github.com/alibabacloud-go/ecs-20140526/v3/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v3/client"
	ossclient "github.com/alibabacloud-go/oss-20190517/v2/client"
	polardbclient "github.com/alibabacloud-go/polardb-20170801/v3/client"
	pvtzclient "github.com/alibabacloud-go/pvtz-20180101/v2/client"
	kvstoreclient "github.com/alibabacloud-go/r-kvstore-20150101/v3/client"
	ramclient "github.com/alibabacloud-go/ram-20150501/v2/client"
	rdsclient "github.com/alibabacloud-go/rds-20140815/v3/client"
	slbclient "github.com/alibabacloud-go/slb-20140515/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	vpcclient "github.com/alibabacloud-go/vpc-20160428/v6/client"
	wafclient "github.com/alibabacloud-go/waf-openapi-20190910/v2/client"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewClient(AccessKeyId, AccessKeySecret string) *Client {
	return &Client{
		AccessKeyId:     AccessKeyId,
		AccessKeySecret: AccessKeySecret,
	}
}

func (c *Client) CompleteConfig(region string) *openapi.Config {
	return &openapi.Config{
		AccessKeyId:     tea.String(c.AccessKeyId),
		AccessKeySecret: tea.String(c.AccessKeySecret),
		RegionId:        tea.String(region),
	}
}

// "github.com/alibabacloud-go/adb-20190315/v2/client"
func (c *Client) ADB(region string) (*adbclient.Client, error) {
	return adbclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/alb-20200616/v2/client"
func (c *Client) ALB(region string) (*albclient.Client, error) {
	return albclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/amqp-open-20191212/v2/client"
func (c *Client) AMQP(region string) (*amqpopenclient.Client, error) {
	return amqpopenclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/bssopenapi-20171214/v3/client"
func (c *Client) BSS(region string) (*bssclient.Client, error) {
	return bssclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/cms-20190101/v8/client"
func (c *Client) CMS(region string) (*cmsclient.Client, error) {
	return cmsclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/cr-20181201/v2/client"
func (c *Client) CR(region string) (*crclient.Client, error) {
	return crclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/cs-20151215/v3/client"
func (c *Client) CS(region string) (*csclient.Client, error) {
	return csclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/dms-enterprise-20181101/client"
func (c *Client) DMS(region string) (*dmsenterpriseclient.Client, error) {
	return dmsenterpriseclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/dds-20151201/v4/client"
func (c *Client) Mongo(region string) (*mongoclient.Client, error) {
	return mongoclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/ecs-20140526/v3/client"
func (c *Client) ECS(region string) (*ecsclient.Client, error) {
	return ecsclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/nas-20170626/v3/client"
func (c *Client) NAS(region string) (*nasclient.Client, error) {
	return nasclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/oss-20190517/v2/client"
func (c *Client) OSS(region string) (*ossclient.Client, error) {
	return ossclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/polardb-20170801/v3/client"
func (c *Client) PolarDB(region string) (*polardbclient.Client, error) {
	return polardbclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/pvtz-20180101/v2/client"
func (c *Client) PVTZ(region string) (*pvtzclient.Client, error) {
	return pvtzclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/r-kvstore-20150101/v3/client"
func (c *Client) KVStore(region string) (*kvstoreclient.Client, error) {
	return kvstoreclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/ram-20150501/v2/client"
func (c *Client) RAM(region string) (*ramclient.Client, error) {
	return ramclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/rds-20140815/v3/client"
func (c *Client) RDS(region string) (*rdsclient.Client, error) {
	return rdsclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/slb-20140515/v4/client"
func (c *Client) SLB(region string) (*slbclient.Client, error) {
	return slbclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/vpc-20160428/v6/client"
func (c *Client) VPC(region string) (*vpcclient.Client, error) {
	return vpcclient.NewClient(c.CompleteConfig(region))
}

// "github.com/alibabacloud-go/waf-openapi-20190910/v2/client"
func (c *Client) WAF(region string) (*wafclient.Client, error) {
	return wafclient.NewClient(c.CompleteConfig(region))
}
