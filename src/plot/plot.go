package plot


import (

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"image/color"
	"log"
	"strconv"
	"tool"
)

func PlotSingle(title string,xl string,yl string,ytext []float64,count int){
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X, _ = strconv.ParseFloat(strconv.Itoa(i),64)
			pts[i].Y = ytext[i]
		}
		return pts
	}
	n := count

	linePointsData := randomPoints(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = title
	p.X.Label.Text = xl
	p.Y.Label.Text = yl
	//p.Add(plotter.NewGrid())
	lpLine, err := plotter.NewHistogram(linePointsData,n)
	//lpLine, err := plotter.NewLine(linePointsData)
	if err != nil {
		log.Panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpLine.LineStyle.Width = 2
	//lpPoints.Shape = draw.CircleGlyph{}
	//lpPoints.Color = color.RGBA{R: 255, A: 255}
	p.Add(lpLine)
	p.Legend.Add(title, lpLine)
	p.Legend.Top = true
	picFile := title + "_" + tool.GetNowTime() + ".png"
	err = p.Save(800, 450, picFile)
	if err != nil {
		log.Panic(err)
	}
}

func PlotCpu(title string,xl string,yl string,usr []float64,sys []float64,idle []float64,iowait []float64,count int) {
	usrpoint := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X, _ = strconv.ParseFloat(strconv.Itoa(i),64)
			pts[i].Y = usr[i]
		}
		return pts
	}
	syspoint := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X, _ = strconv.ParseFloat(strconv.Itoa(i),64)
			pts[i].Y = sys[i]
		}
		return pts
	}
	idlepoint := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X, _ = strconv.ParseFloat(strconv.Itoa(i),64)
			pts[i].Y = idle[i]
		}
		return pts
	}
	iowaitpoint := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			pts[i].X, _ = strconv.ParseFloat(strconv.Itoa(i),64)
			pts[i].Y = iowait[i]
		}
		return pts
	}
	n := count
	usrdata := usrpoint(n)
	sysdata := syspoint(n)
	idledata := idlepoint(n)
	iowaitdata := iowaitpoint(n)
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = title
	p.X.Label.Text = xl
	p.Y.Label.Text = yl
	p.Add(plotter.NewGrid())
	usrLine, err := plotter.NewLine(usrdata)
	sysline, err := plotter.NewLine(sysdata)
	idleline, err := plotter.NewLine(idledata)
	iowaitline, err := plotter.NewLine(iowaitdata)
	if err != nil {
		log.Panic(err)
	}
	usrLine.Color = color.RGBA{G: 255, A: 255}
	usrLine.LineStyle.Width = 2
	sysline.Color = color.RGBA{R: 255, A: 255}
	sysline.LineStyle.Width = 2
	idleline.Color = color.RGBA{B: 255, A: 255}
	idleline.LineStyle.Width = 2
	iowaitline.Color = color.RGBA{G: 128, A: 255}
	iowaitline.LineStyle.Width = 2
	p.Add(usrLine)
	p.Legend.Add("usr", usrLine)

	p.Add(sysline)
	p.Legend.Add("sys", sysline)

	p.Add(idleline)
	p.Legend.Add("idle", idleline)

	p.Add(iowaitline)
	p.Legend.Add("iowait", iowaitline)
	p.Legend.Top = true
	picFile := title + "_" + tool.GetNowTime() + ".png"
	err = p.Save(800, 450, picFile)
	if err != nil {
		log.Panic(err)
	}
}