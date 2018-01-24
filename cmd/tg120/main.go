package main

import (
	//"fmt"
	//"math"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/yjiong/go_tg120/internal/device"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)
	//r := device.HEELIGHT{}
	//r := device.QDSL_SM510{}
	//r := device.TEST_GO{}
	//r := device.RSBAS{}
	//r := device.FUJITSU{}
	r := device.TOSHIBA{}
	//r := device.DTSD422{}
	tval, _ := r.NewDev("toshiba", map[string]string{
		//"devaddr": "3300027014",
		"devaddr":   "2",
		"commif":    "rs485-1",
		"IndoorNum": "8",
		//"mtype":    "VRF",
		//"sub_addr": "1",
		//"BaudRate": "19200",
		//"DataBits": "8",
		//"StopBits": "1",
	})
	elem := map[string]interface{}{
	//"_varname":  "运行模式设置",
	//"_varvalue": "送风",
	//"_varname":  "运行开关设置",
	//"_varvalue": "运行",
	//"_varname":  "设置温度设定值",
	//"_varvalue": "28",
	//"_varname":  "气流设置",
	//"_varvalue": "中",
	//"_varname":  "垂直空气方向位置状态",
	//"_varvalue": "位置4",
	//"_varname":  "水平空气方向位置状态",
	//"_varvalue": "位置2",
	//"_varname":  "遥控器运行禁止设置",
	//"_varvalue": "允许",
	//"_varname":  "过滤网标志重置",
	//"_varvalue": "重置",
	//"_varname":  "经济运行模式设置",
	//"_varvalue": "节能运行",
	//"_varname":  "防冻液运行设置",
	//"_varvalue": "释放",
	//"_varname":  "制冷/干燥温度上限设置",
	//"_varvalue": "16",
	//"_varname":  "制冷/干燥温度下限设置",
	//"_varvalue": "16",
	//"_varname":  "加热温度上限设置",
	//"_varvalue": "16",
	//"_varname": "加热温度下限设置",
	//"_varvalue": "16",
	//"_varname":  "自动温度上限设置",
	//"_varvalue": "16",
	//"_varname":  "自动温度下限设置",
	//"_varvalue": "6",
	//"_varname":  "外部关热设置",
	//"_varvalue": "关热",
	//"_varname":  "紧急停止",
	//"_varvalue": "紧急停止请求",
	//"_varname":  "室外机低噪音运行设置",
	//"_varvalue": "性能优先3",
	//"_varname":  "室外机额定容量节省指令",
	//"_varvalue": "70%",
	}
	//tval.RWDevValue("w", nil)
	for {

		if ret, err := tval.RWDevValue("r", elem); err != nil {
			log.Debugf("error=%s", err)
		} else {
			jret, _ := json.Marshal(ret)
			jsret, _ := simplejson.NewJson(jret)
			prettyret, _ := jsret.EncodePretty()
			log.Debugln(ret)
			log.Debugln(string(prettyret))
		}
		time.Sleep(0 * time.Second)
		break
	}
}
