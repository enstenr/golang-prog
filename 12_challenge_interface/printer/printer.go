package printer

import (
	"fmt"
	"os"
	"text/tabwriter"
	"tempService/models"
)

type Printer struct{
	w *tabwriter.Writer
}

func New() *Printer{
	w:=tabwriter.NewWriter(os.Stdout,3,0,3,' ',tabwriter.TabIndent)
	return &Printer{
		w:w,
	}
}

func (p *Printer)CityHeader(){
	fmt.Fprintln(p.w,"Name\tTempc\tTempF\tBeach Vacation Ready ? \tMountain Vacation Ready? ")
}


func (p *Printer)CityDetails(c models.CityTemp){
	fmt.Fprintf(p.w,"%v\t%v\t%v\t%v\t%v\n",c.Name(),c.Tempc(),c.TempF(),c.BeachVacationReady(),c.SkiVacationReady())
}

func (p *Printer)CleanUp(){
	p.w.Flush()
}