package goecharts

import (
	"bytes"
	"io"
)

type Pie struct {
	InitOptions
	BaseOptions
	SeriesList

	HasXYAxis bool
}

//工厂函数，生成 `Pie` 实例
func NewPie() *Pie {
	pie := new(Pie)
	pie.HasXYAxis = false
	pie.initSeriesColors()
	return pie
}

func (pie *Pie) Add(name string, data map[string]interface{}, options ...interface{}) *Pie {
	type pieData struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}
	pd := make([]pieData, 0)
	for k, v := range data {
		pd = append(pd, pieData{k, v})
	}
	series := Series{Name: name, Type: pieType, Data: pd}
	series.setSingleSeriesOptions(options...)
	pie.SeriesList = append(pie.SeriesList, series)
	pie.setColor(options...)
	return pie
}

func (pie *Pie) SetGlobalConfig(options ...interface{}) *Pie {
	pie.BaseOptions.setRectGlobalConfig(options...)
	return pie
}

func (pie *Pie) verifyOpts() {
	pie.verifyInitOpt()
}

func (pie *Pie) Render(w ...io.Writer) {

	pie.insertSeriesColors(pie.appendColor)
	pie.verifyOpts()

	var b bytes.Buffer
	renderChart(pie, &b, "chart")
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}