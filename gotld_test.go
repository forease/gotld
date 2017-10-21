package gotld

import (
	//"net"
	"testing"
)

var testUrls []string = []string{"www.google.com.hk", "www.forease.net", "com",
	"www.forease.com.cn", "www.ritto.shiga.jp", "ritto.shiga.jp", "ymm"}

func TestGetTld(t *testing.T) {
	//url := "www.google.com.hk"
	for _, url := range testUrls {
		tld, domain, err := GetTld(url)
		if nil != err {
			t.Error("Failed get TLD:" + err.Error())
			return
		}
		t.Logf("%s: %v, %s\n", url, tld.Tld, domain)
	}

	t.Fail()
}

func TestGetVersion(t *testing.T) {

	t.Logf("%s\n", GetVersion())

	t.Fail()
}
