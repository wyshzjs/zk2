package pkg

import (
	"github.com/smartwalle/alipay/v3"
)

func Pay(subject, outTradeNo, totalAmount string) (string, error) {
	client, err := alipay.New(appId, privateKey, false)
	if err != nil {
		return "", err
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://111.231.76.156:9200"
	p.ReturnURL = "https://gorm.io"
	p.Subject = subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		return "", nil
	}

	return url.String(), nil
}

var (
	privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCePD0sQIXlI2vZVLWjE7sfq0Xrcp0CFog5dalodv+ecfVjL0GUW3zWCp5t5ju8UvITGSR7mwBkG2RSmhOvM9okW0e5/+JwvJTaMBeKPMzWN3QdzVI53nG9094CkFJJQsUdfiHE+ECxIZVMAM1RblJ4CxnXsMeIA+jhQ7TTafRCqUCf1238tuu7bX5m3TjIWWImj0YKNJ0HegSzWwCBelS34Hqr/bu4E0T7Mi5wHMq6jDVi/vXfLmq1OKHpZfd+bL3ZqBJGZXrp6/BwBJ27LdwOxjN5f/6QqTfZWzJgD14qMf7HWCQIgdcgkvht1rADWxPh2fLI186ZVCJe4u13WRtNAgMBAAECggEAA3UdgRIH2vHjMWSAc4fC1vZPGM3wUFNdtqOzh+zRuQyUaA5oSG2XVLX2sEAPnhXR5EjwWyMP5yASGRv0widH5PN5pb7Jm90DA6F5YjPh5exgRSZvXYRWx+cROs0cfy1OUU6oPj2ObFBnirmZwuW+N5FGh0f9lCNBu2pXUSdHhw1S6AOEvnV9RVoQO9V1Z1mAmdkiYbOHQFrJVV7XSjPGl8B8h8cjIq0Hnx7hPPASM6TNIbwHhCEz6iwzcvRAPqdQqRs8ac9nBONnp++QkcGdAHyuBlwjKRbRAkOfNpDfUbGrKxarf1Td0R9zdjHawTQM/Ezbw/6cilEllPQ04D9UoQKBgQDNfOuSKYgT9U7mfiJ/V3CDw9k6JAFeD1H9ixzWe6QC5cl8CLOg/7nwOWbhQcm7OR5BAzPv4qqhSqjA8r7TfaeVyd2EB5SnBl3OUybPISAawUK8Gb+plsnFql4f0qh7Cs5BpWFpCqGJYR1kNyyfiSv+/7KAG+MXLCzbiOGzbFuY3wKBgQDFIcXS4/Ib6f4hlMMVitklS4SqJBIuixlsq8l8yTjb6pE5OjtIKU16oSsPAFM24tvwWG9E4qSsNaAJKMMsEJSnRXDmA38o7pAuYxV+Rra8owXgqbvFBmQTmWxch6p7sRcSF5++NZZx3I0SfTjpUOwv/m2Dr2QsiL12XvstzivVUwKBgH8iEZE9ls5aUDV0xp0omRk8u5IoIpxyOBT405+zoVaVmqHdtF0guoZwzUECdfAat6gbdi+tekoOKfgud3fb3ypiWWXNVuA7pR/4/CLcaMOE3GgKOjWPcEiWTmO77rcpiVC+JFVb2XNonUmahUYzFQBQbLysBEUa9Knit1lN6iHbAoGAexddiROgZ+GyN5HGRrtOX8VJ71tapwlgNMUHqQ2UnW1YhN1hiuuPX5UnZ9SSIWGCaA8wW1kR7h4F8fIr/GXWDKSFnHpuekaMbQqdrx3EmQFKeEBD6QDWmI5iZVypDVdEyp1Jt+GKMUrAI91hHf7TnBxIkcOAz3YKp9EjEjT8ldUCgYAjMDNOQY9QfIU4AvMGnsYNCeU6sEZwXSdd6+zaZMblUbzZwzxFVXsTD6tkevG2XlC68ZDhRpn/OB7a23ROwY9TNuN2LNfTUvkZb9UUaA/FgdCILDPM4Wjpc2B8KlhQuX9pdoZ+9DSGfPlqb1CNQbcpk64E6YsgxiBwwDL8Ie+OXA=="
	appId      = "9021000137646369"
)
