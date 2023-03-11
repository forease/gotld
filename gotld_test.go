package gotld

import (
	//"net"
	"fmt"
	"testing"
)

var testUrls []string = []string{"www.google.com.hk", "www.forease.net", "com",
	"www.forease.com.cn", "www.ritto.shiga.jp", "ritto.shiga.jp", "jonsen.yang"}

func TestGetTld(t *testing.T) {
	//url := "www.google.com.hk"
	for _, url := range testUrls {
		ss, dd, tld := GetSubdomain(url, 2)
		fmt.Println(ss, dd, tld)
		tldItem, domain, err := GetTld(url)
		if nil != err {
			t.Error("Failed get TLD:" + err.Error())
			return
		}
		t.Logf("%s: %v, %s\n", url, tldItem.Tld, domain)
	}

	t.Fail()
}

func TestGetVersion(t *testing.T) {

	t.Logf("%s\n", GetVersion())

	t.Fail()
}

func BenchmarkGetTld(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetTld("www.aaa.bbb.ccc.ddd.forease.com.cn")
	}
}

func BenchmarkGetSubdomain(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetSubdomain("www.aaa.bbb.ccc.ddd.forease.com.cn", 0)
	}
}
