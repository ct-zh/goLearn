package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "testfamily.busi.inkept.cn", Path: "/ws", RawQuery: "user_id=29251560&document_id=88&ticket=STXlYdayqvbEUPmYNpDdVoBzaEJlcaMNAuC"}
	//u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/echo"}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	go func() {

	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(str1))
			if err != nil {
				log.Println("write:", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, []byte(str2))
			if err != nil {
				log.Println("write:", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, []byte(str3))
			if err != nil {
				log.Println("write:", err)
				return
			}
			err = c.WriteMessage(websocket.TextMessage, []byte(str4))
			if err != nil {
				log.Println("write:", err)
				return
			}
			log.Printf("send once in [%s]", time.Now().Format("2006-01-02 15:04:05"))

		case <-interrupt:
			log.Println("interrupt")

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

var (
	str1 = `{"user_id":"29251560","document_id":"88","busi_content":"{\"type\":\"update\",\"record\":[{\"field\":\"0.kr_info.2.raw\",\"value\":{\"blocks\":[{\"key\":\"bkboo\",\"text\":\"在产品UI方向沉淀1-2套完整的资源库，提升对产品UI的支持能力，能快速完成对产品的支持，以及马甲包的产出\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[{\"offset\":0,\"length\":53,\"style\":\"BOLD\"}],\"entityRanges\":[{\"offset\":3,\"length\":37,\"key\":1}],\"data\":{}}],\"entityMap\":{\"1\":{\"type\":\"COMMENT\",\"mutability\":\"MUTABLE\",\"data\":{\"0\":{\"anchorKey\":\"bkboo\",\"start\":3,\"end\":40,\"selectedText\":\"UI方向沉淀1-2套完整的资源库，提升对产品UI的支持能力，能快速完成对产\",\"key\":\"fo5og\",\"resolved\":0,\"uid\":29251560}}}}},\"oIndex\":0}],\"card_id\":\"88-12355,12356,12357-3kmgj-748dk\",\"commentIds\":{\"cids\":\"1004,1176,uid-29251560,uid-287,1007,1204,1211\",\"dcids\":\"\"},\"version\":17}","unique":"","busi_type":"updateOkrCard"}`
	str2 = `{"user_id":"29251560","document_id":"88","busi_content":"{\"raw\":{\"blocks\":[{\"key\":\"2a6bl\",\"text\":\"====================\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"2qnfi\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":1}],\"data\":{}},{\"key\":\"bf7ar\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"6ahl5\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":2}],\"data\":{}},{\"key\":\"aorqo\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"3kmgj\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":3}],\"data\":{}},{\"key\":\"dr3k4\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"i0oh\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":4}],\"data\":{}},{\"key\":\"7cs3f\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"ai36\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":5}],\"data\":{}},{\"key\":\"56cc0\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"cfigd\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":6}],\"data\":{}},{\"key\":\"58fvt\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"4411f\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":7}],\"data\":{}},{\"key\":\"bmhpr\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"fb856\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":8}],\"data\":{}},{\"key\":\"9bsq1\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"2vbd2\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":9}],\"data\":{}},{\"key\":\"7lb07\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"bn3l3\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":10}],\"data\":{}},{\"key\":\"lptt\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"16n0r\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":11}],\"data\":{}},{\"key\":\"95jof\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"8s2l2\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":12}],\"data\":{}},{\"key\":\"5c5jm\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"6t4dh\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":13}],\"data\":{}},{\"key\":\"f040p\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"3ivgf\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":14}],\"data\":{}},{\"key\":\"8fpa0\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"23nud\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":15}],\"data\":{}},{\"key\":\"469v5\",\"text\":\"\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}},{\"key\":\"85qf5\",\"text\":\" \",\"type\":\"atomic\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[{\"offset\":0,\"length\":1,\"key\":16}],\"data\":{}},{\"key\":\"dm0t8\",\"text\":\"6\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[],\"entityRanges\":[],\"data\":{}}],\"entityMap\":{\"1\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"11\",\"name\":\"李一文\",\"email\":\"liyw@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11730\":1,\"11790\":1,\"11795\":1}},\"card_id\":\"88-11730,11790,11795-2qnfi-bpnef\",\"cids\":\"1034,1000,1025,1026,uid-88,1027,1021,1035,1016,1022,1190,1189,1188,1266,1031,uid-12,1187,1237,1201,1011,1008\",\"dcids\":\"1018\"}},\"2\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"12\",\"name\":\"朱泽瑞\",\"email\":\"zhuzerui@inke.cn\",\"portrait\":\"http://m4a.inke.cn/MTYyMzkyNDI1MjU5NiM1NTQjanBn.jpg\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11984\":1,\"11985\":1}},\"card_id\":\"88-11984,11985-6ahl5-9mmrt\",\"cids\":\"\",\"dcids\":\"\"}},\"3\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"15\",\"name\":\"李海平\",\"email\":\"lihaiping@inke.cn\",\"portrait\":\"http://img.ikstatic.cn/MTU4NTYzMzQ3MDM2MCMgODEjcG5n.png\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12355\":1,\"12356\":1,\"12357\":1}},\"card_id\":\"88-12355,12356,12357-3kmgj-748dk\",\"cids\":\"1004,1176,uid-29251560,uid-287,1007,1204,1211\",\"dcids\":\"\"}},\"4\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"18\",\"name\":\"刘迎春\",\"email\":\"liuyingchun@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12122\":1,\"12197\":1,\"12203\":1,\"12570\":1}},\"card_id\":\"88-12122,12197,12203,12570-i0oh-1hq3m\",\"cids\":\"\",\"dcids\":\"\"}},\"5\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"71\",\"name\":\"朱韬奋\",\"email\":\"zhutaofen@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11352\":1,\"12777\":1}},\"cids\":\"\",\"dcids\":\"\",\"card_id\":\"88-11352,12777-ai36-btlfl\"}},\"6\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"87\",\"name\":\"刘名运\",\"email\":\"liumy@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11979\":1,\"11981\":1,\"11983\":1}},\"card_id\":\"88-11979,11981,11983-cfigd-ees86\",\"cids\":\"\",\"dcids\":\"\"}},\"7\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"88\",\"name\":\"殷燃\",\"email\":\"yinran@inke.cn\",\"portrait\":\"http://m4a.inke.cn/MTYxNTM2ODU4MzA2OSM5NjcjcG5n.png\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12281\":1}},\"card_id\":\"88-12281-4411f-fp5n6\",\"cids\":\"\",\"dcids\":\"\"}},\"8\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"246\",\"name\":\"文小清\",\"email\":\"wenxq@inke.cn\",\"portrait\":\"http://img.ikstatic.cn/MTU4NDI3MDY2OTU4MyM2MDcjcG5n.png\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12109\":1,\"12110\":1}},\"card_id\":\"88-12109,12110-fb856-dt340\",\"cids\":\"\",\"dcids\":\"\"}},\"9\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"287\",\"name\":\"唐昆侯\",\"email\":\"tangkunhou@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12075\":1,\"12175\":1}},\"card_id\":\"88-12075,12175-2vbd2-fdkfj\",\"cids\":\"\",\"dcids\":\"\"}},\"10\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"263082\",\"name\":\"周渝雄\",\"email\":\"zhouyuxiong@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11844\":1,\"11845\":1,\"12047\":1}},\"card_id\":\"88-11844,11845,12047-bn3l3-78c1n\",\"cids\":\"\",\"dcids\":\"\"}},\"11\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"263288\",\"name\":\"崔世杰\",\"email\":\"cuishijie@inke.cn\",\"portrait\":\"http://m4a.inke.cn/MTYxNTM0MzQwMTMyMSM1NTgjcG5n.png\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12211\":1,\"12212\":1,\"12213\":1}},\"card_id\":\"88-12211,12212,12213-16n0r-duhpf\",\"cids\":\"\",\"dcids\":\"\"}},\"12\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"263347\",\"name\":\"叶溱\",\"email\":\"yezhen@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11846\":1}},\"card_id\":\"88-11846-8s2l2-61n99\",\"cids\":\"\",\"dcids\":\"\"}},\"13\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"10110095\",\"name\":\"黄江\",\"email\":\"huangjiang@inke.cn\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"12147\":1,\"12148\":1}},\"card_id\":\"88-12147,12148-6t4dh-7a70e\",\"cids\":\"\",\"dcids\":\"\"}},\"14\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"29244547\",\"name\":\"刘柱彤\",\"email\":\"liuzhutong@inke.cn\",\"portrait\":\"https://static-legacy.dingtalk.com/media/lADPDgQ9qUMYqzPNAtjNAuo_746_728.jpg\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11700\":1,\"11703\":1,\"11705\":1,\"12467\":1}},\"card_id\":\"88-11700,11703,11705,12467-3ivgf-en33m\",\"cids\":\"\",\"dcids\":\"\"}},\"15\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"29252132\",\"name\":\"陈娜\",\"email\":\"chenna@inke.cn\",\"portrait\":\"http://m4a.inke.cn/MTYyODIxMzQ2ODA2MCM2NDkjanBn.jpg\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11741\":1,\"11743\":1,\"11744\":1,\"12155\":1,\"12778\":1}},\"cids\":\"\",\"dcids\":\"\",\"card_id\":\"88-11741,11743,11744,12155,12778-23nud-du5g5\"}},\"16\":{\"type\":\"OKR\",\"mutability\":\"IMMUTABLE\",\"data\":{\"initUser\":\"\",\"userInfo\":{\"user_id\":\"29252298\",\"name\":\"彭丹\",\"email\":\"pengdan@inke.cn\",\"portrait\":\"https://static-legacy.dingtalk.com/media/lADOasKy1c0C7s0C7g_750_750.jpg\",\"code\":0,\"quarterText\":\"2021年9月-10月\",\"okrData\":{\"11624\":1,\"11625\":1,\"11646\":1,\"11647\":1}},\"card_id\":\"88-11624,11625,11646,11647-85qf5-1ted2\",\"cids\":\"\",\"dcids\":\"\"}}}},\"comment_ids\":[\"1034\",\"1000\",\"1025\",\"1026\",\"uid-88\",\"1027\",\"1021\",\"1035\",\"1016\",\"1022\",\"1190\",\"1189\",\"1188\",\"1266\",\"1031\",\"uid-12\",\"1187\",\"1237\",\"1201\",\"1011\",\"1008\",\"1004\",\"1176\",\"uid-29251560\",\"uid-287\",\"1007\",\"1204\",\"1211\"],\"done_comment_ids\":[\"1018\"],\"version\":237}","unique":"","busi_type":"updatecontent"}`
	str3 = `{"user_id":"29251560","document_id":"88","busi_content":"{\"type\":\"update\",\"record\":[{\"field\":\"0.kr_info.2.raw\",\"value\":{\"blocks\":[{\"key\":\"bkboo\",\"text\":\"在产品UI方向沉淀1-2套完整的资源库，提升对产品UI的支持能力，能快速完成对产品的支持，以及马甲包的产出\",\"type\":\"unstyled\",\"depth\":0,\"inlineStyleRanges\":[{\"offset\":0,\"length\":53,\"style\":\"BOLD\"}],\"entityRanges\":[{\"offset\":3,\"length\":37,\"key\":1}],\"data\":{}}],\"entityMap\":{\"1\":{\"type\":\"COMMENT\",\"mutability\":\"MUTABLE\",\"data\":{\"0\":{\"anchorKey\":\"bkboo\",\"start\":3,\"end\":40,\"selectedText\":\"UI方向沉淀1-2套完整的资源库，提升对产品UI的支持能力，能快速完成对产\",\"key\":\"fo5og\",\"resolved\":0,\"comment_id\":1267}}}}},\"oIndex\":0}],\"card_id\":\"88-12355,12356,12357-3kmgj-748dk\",\"commentIds\":{\"cids\":\"1004,1176,1267,uid-287,1007,1204,1211\",\"dcids\":\"\"},\"version\":18}","unique":"","busi_type":"updateOkrCard"}`
	str4 = `{"user_id":"29251560","document_id":"88","busi_content":"{\"type\":\"add-cross\",\"updateTime\":\"2021-09-10 17:30:06\",\"id\":1267,\"content\":\"测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试\",\"mark\":{\"anchorKey\":\"bkboo\",\"start\":3,\"end\":40,\"selectedText\":\"UI方向沉淀1-2套完整的资源库，提升对产品UI的支持能力，能快速完成对产\",\"key\":\"fo5og\",\"resolved\":0,\"comment_id\":1267},\"user\":{\"user_id\":\"29251560\",\"name\":\"曹庭\"}}","unique":"","busi_type":"UPDATE_CROSS"}`
)
