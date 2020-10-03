package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-echarts/go-echarts/charts"

	// postgres
	_ "github.com/lib/pq"
)

var insts []string

// KLineController KLine
type KLineController struct {
	beego.Controller
}

func init() {
	// 注册模型,带shcema无效
	// orm.RegisterModelWithPrefix("future", new(Min))
	// orm.RegisterModel(new(Min))
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 注册数据库
	pgConfig := "postgres://postgres:123456@172.19.129.98:15432/postgres?sslmode=disable"
	if tmp := os.Getenv("pgConfig"); tmp != "" {
		pgConfig = tmp
	}

	err := orm.RegisterDataBase("default", "postgres", pgConfig) // "user=postgres password=123456 dbname=postgres host=172.19.129.98 port=15432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		beego.Error(err)
	} else {
		// 取合约列表,需优化**********************************
		// o := orm.NewOrm()
		// o.Using("default")
		// var maps []orm.Params
		// num, err := o.Raw("select distinct \"Instrument\" from future.future_min").Values(&maps)
		// if err != nil {
		// 	beego.Error("Returned Rows Num: %s, %s", num, err)
		// } else {
		// 	for _, item := range maps {
		// 		insts = append(insts, item["Instrument"].(string))
		// 	}
		// 	sort.Strings(insts)
		// }

		cal, err := os.Open("instrument.csv")
		defer cal.Close()
		if err != nil {
			beego.Error(err)
		}
		reader := csv.NewReader(cal)
		lines, err := reader.ReadAll()
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}
			if line[9] > "20190101" {
				insts = append(insts, line[0])
			}
		}
		sort.Strings(insts)
	}
}

// Show 显示K线
func (c *KLineController) Show() {
	inst := c.Ctx.Input.Param(":instrument")
	o := orm.NewOrm()
	o.Using("default")
	var maps []orm.Params
	c.Data["Instruments"] = insts

	if len(inst) > 0 {
		// var mins []*Min
		num, err := o.Raw(fmt.Sprintf(`select * from future.future_min where "Instrument"='%s' order by "DateTime"`, inst)).Values(&maps)
		if err != nil {
			beego.Warning("Returned Rows Num: %s, %s", num, err)
		} else {
			var date []string
			var data [][4]float32
			for _, item := range maps {
				date = append(date, item["DateTime"].(string))
				o, _ := strconv.ParseFloat(item["Open"].(string), 32)
				c, _ := strconv.ParseFloat(item["Close"].(string), 32)
				l, _ := strconv.ParseFloat(item["Low"].(string), 32)
				h, _ := strconv.ParseFloat(item["High"].(string), 32)
				data = append(data, [4]float32{float32(o), float32(c), float32(l), float32(h)})
			}
			c.Data["Dates"] = date
			c.Data["Bars"] = data
			c.Data["title"] = inst
			c.TplName = "kline.tpl"
		}
		// qs := o.QueryTable("future_min")
		// num, err := qs.Filter("Instrument", inst).All(&mins)
		// beego.Warning("Returned Rows Num: %s, %s", num, err)
	} else {
		c.TplName = "kline.tpl"
	}
}

// Get 显示K线
func (c *KLineController) Get() {
	inst := c.GetString("inst")
	fmt.Printf(inst)

	kd = append(kd, klineData{date: "20200927", data: [4]float32{10, 12, 15, 18}})

	k := klineStyle()
	page := charts.NewPage().Add(k)
	page.Render(c.Ctx.ResponseWriter)
	f, _ := os.Create("kline.html")
	page.Render(f)
	// page := charts.NewPage(charts.RouterOpts{URL: "/bar", Text: "Bar-(柱状图)"})
	// page.Add(
	// 	// klineStyle(),
	// 	k,
	// )
	// page.Render(c.Ctx.ResponseWriter)
}

type klineData struct {
	date string
	data [4]float32
}

// 数据
var kd []klineData

func klineStyle() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}
	kline.AddXAxis(x)
	kline.AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TooltipOpts{Show: true, Trigger: "axis"}, //TriggerOn: "mousemove | click"
		// charts.TitleOpts{Title: "Kline-不同风格"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		// charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	kline.SetSeriesOptions(
		charts.MPNameTypeItem{Name: "highest value", Type: "max", ValueDim: "highest"},
		charts.MPNameTypeItem{Name: "lowest value", Type: "min", ValueDim: "lowest"},
		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
		charts.ItemStyleOpts{
			Color: "#ec0000", Color0: "#00da3c", BorderColor: "#8A0000", BorderColor0: "#008F28"},
	)
	return kline
}
