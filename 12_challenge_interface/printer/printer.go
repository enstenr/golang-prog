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
	fmt.Fprintln(p.w,"Name\tTempc\tTempF")
}


func (p *Printer)CityDetails(c *models.City){
	fmt.Fprintln(p.w,"%vt%v\t%v",c.Name,c.Tempc,c.TempF)
}

func (p *Printer)CleanUp(){
	p.w.Flush()
}