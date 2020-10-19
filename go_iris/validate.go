package main

import (
	"errors"
	"fmt"
	"go_iris/common"
	"go_iris/encrypt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

// 设置集群地址，
var hostArray = []string{"127.0.0.1", "127.0.0.1"}

var localHost = "127.0.0.1"

var port = "8081"

var hashConsistent *common.Consistent

// 存放控制信息
type AccessControl struct {
	sourcesArray map[int]interface{}
	sync.RWMutex
}

// 创建全局变量
var accessControl = &AccessControl{sourcesArray: make(map[int]interface{})}

// 获取设定的数据
func (a *AccessControl) GetNewRecord(uid int) interface{} {
	a.RWMutex.RLock()
	defer a.RWMutex.RUnlock()
	return a.sourcesArray[uid]
}

// 设置记录
func (a *AccessControl) SetNewRecord(uid int, data interface{}) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	a.sourcesArray[uid] = data
}

func (a *AccessControl) GetDistributedRight(req *http.Request) bool {
	// 获取用户id
	uid, err := req.Cookie("uid")
	if err != nil {
		return false
	}

	// 根据一致性hash算法，判断出具体的机器
	hostRequest, err := hashConsistent.Get(uid.Value)
	if err != nil {
		return false
	}

	//判断是否为本机
	if hostRequest == localHost {
		// 执行本机的数据校验
		return a.GetDataFromMap(uid.Value)
	} else {
		// 不是本机则充当转发
		return GetDataFromOtherMap(hostRequest, req)
	}
}

func (a *AccessControl) GetDataFromMap(uid string) bool {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	data := a.GetNewRecord(uidInt)
	return data == nil
}

func GetDataFromOtherMap(host string, request *http.Request) bool {
	uidString, err := request.Cookie("uid")
	if err != nil {
		return false
	}

	signStr, err := request.Cookie("sign")
	if err != nil {
		return false
	}

	// 模拟接口访问
	client := &http.Client{}
	req2, err := http.NewRequest("GET", "http://"+host+":"+port+"/check", nil)
	if err != nil {
		return false
	}

	// 手动指定cookie
	cookieUid := &http.Cookie{Name: "uid", Value: uidString.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: signStr.Value, Path: "/"}
	req2.AddCookie(cookieUid)
	req2.AddCookie(cookieSign)

	response, err := client.Do(req2)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}

	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

func Check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check")
}

func Auth(w http.ResponseWriter, r *http.Request) error {
	// auth ?

	err := CheckUserInfo(r)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserInfo(r *http.Request) error {
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		return errors.New("用户UID Cookie 获取失败！")
	}

	signCookie, err := r.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串 Cookie 获取失败！")
	}

	signByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("用户加密串 解密失败！")
	}

	if checkInfo(uidCookie.Value, string(signByte)) {
		return nil
	}

	return errors.New("身份校验失败")
}

// 自定义逻辑判断
func checkInfo(checkStr string, signStr string) bool {
	return checkStr == signStr
}

func main() {
	// 负载均衡器设置
	// 采用一致性hash算法
	hashConsistent = common.NewConsistent()
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}

	// 过滤器
	filter := common.NewFilter()

	// 拦截器注册
	filter.RegisterFilterUri("/check", Auth)

	// 启动服务
	http.HandleFunc("/check", filter.Handle(Check))

	http.ListenAndServe(":0883", nil)

}
