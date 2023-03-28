package main

import (
	"encoding/xml"
	"testing"
)

func TestXml(t *testing.T) {
	testStr := []string{
		`<xml><ToUserName><![CDATA[gh_b78324305512]]></ToUserName>
<FromUserName><![CDATA[oNmT-0q5h-NTPQByNiGj1vVztgDU]]></FromUserName>
<CreateTime>1680002497</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[普通文本222]]></Content>
<MsgId>24051853658678060</MsgId>
</xml>`,
		`<person><name>John</name><age>30</age></person>`,
		"",
	}

	for _, str := range testStr {
		msg := &xmlMsg{}
		err := xml.Unmarshal([]byte(str), msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("msg=%+v", msg)
	}
}
