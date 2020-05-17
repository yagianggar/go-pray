package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	pkpu "github.com/yagianggar/go-pray/service"
	"log"
	"strconv"
	"time"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := pkpu.NewService("http://jadwalsholat.pkpu.or.id/monthly.php?id=83")
	ss := p.GetPrayingSchedule()

	//today := time.Now()
	//currentDate, err := strconv.Atoi(today.Format("02"))
	//if err != nil {
	//	panic(fmt.Sprintf("Error : %s", err))
	//}

	x, y := ui.TerminalDimensions()

	header := widgets.NewParagraph()
	header.Title = "Your Area"
	header.Text = fmt.Sprintf("Jakarta | [%s](fg:green)", time.Now().Format("02-01-2006 15:04:05"))
	header.BorderStyle.Fg = ui.ColorCyan
	header.TitleStyle.Fg = ui.ColorWhite
	header.TextStyle.Fg = ui.ColorCyan
	header.SetRect(0, 0, x, 3)

	table1 := widgets.NewTable()
	var schedules = [][]string{
		[]string{"Date", "Shubuh", "Dzuhur", "Ashr", "Maghrib", "Isya"},
	}

	for _, v := range ss {
		schedules = append(schedules, []string{v.Date, v.Shubuh, v.Dzuhur, v.Ashar, v.Maghrib, v.Isya})
	}

	table1.Rows = schedules

	table1.FillRow = true
	table1.RowStyles[0] = ui.NewStyle(ui.ColorCyan, ui.ColorClear, ui.ModifierBold)
	//table1.RowStyles[currentDate] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	table1.TextStyle = ui.NewStyle(ui.ColorCyan)
	table1.BorderStyle.Fg = ui.ColorCyan
	table1.RowSeparator = false
	table1.Title = "PKPU Prayer Times"
	table1.TitleStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(0, 3, x, y)

	ui.Render(header, table1)

	ticker := time.NewTicker(time.Second).C

	uiEvents := ui.PollEvents()

	i := 0
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "f":
			case "<Escape>":
				return
			case "<Resize>":
				updatedX, updatedY := ui.TerminalDimensions()
				header.SetRect(0, 0, updatedX, 3)
				table1.SetRect(0, 3, updatedX, updatedY)

				ui.Clear()
				ui.Render(header, table1)
			case "<Up>":
				if i != 0 {
					i--
				}
				if i != 0 {
					table1.RowStyles[i] = ui.NewStyle(ui.ColorWhite, ui.ColorGreen, ui.ModifierBold)
					table1.RowStyles[i+1] = ui.NewStyle(ui.ColorCyan, ui.ColorClear)
				}
				ui.Clear()
				ui.Render(header, table1)
			case "<Down>":
				i++
				table1.RowStyles[i] = ui.NewStyle(ui.ColorWhite, ui.ColorGreen, ui.ModifierBold)
				if i != 1 {
					table1.RowStyles[i-1] = ui.NewStyle(ui.ColorCyan, ui.ColorClear)
				}
				ui.Clear()
				ui.Render(header, table1)
			}
		case <-ticker:
			today := time.Now()
			currentDate, err := strconv.Atoi(today.Format("02"))
			if err != nil {
				panic(fmt.Sprintf("Error : %s", err))
			}

			header.Text = fmt.Sprintf("Jakarta | [%s](fg:green)", time.Now().Format("02-01-2006 15:04:05"))
			if currentDate != 1 {
				table1.RowStyles[currentDate-1] = ui.NewStyle(ui.ColorCyan, ui.ColorClear)
			}
			table1.RowStyles[currentDate] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
			ui.Render(header, table1)
		}
	}
}
