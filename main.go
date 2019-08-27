package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strconv"
	"time"
)

type Time time.Time

/**
 * Create the header for the calendar
 * Includes navigation buttons, date and time
 */
func calendarHeader() Widget{
	//get curr time
	mins := time.Now().Minute()
	hour := time.Now().Hour()
	pm := false

	//convert to 24 hour time
	if hour > 12{
		hour -= 12
		pm = true
	}

	//all widgets:
	widgets := make([]Widget, 0)

	title := TextLabel{
		Text: "this is the header",
		Background: SolidColorBrush{Color: walk.RGB(30, 30, 255)},
		MaxSize: Size{200, 25},
	}
	//
	left := PushButton{
		Text: "←",
		MaxSize: Size{25, 25},
	}
	//
	right := PushButton{
		Text: "→",
		MaxSize: Size{25, 25},
	}
	//
	clock := TextLabel{
		TextColor: walk.RGB(255,255,255),
		MinSize:Size{40, 25},
	}
	if pm {
		clock.Text = strconv.Itoa(hour) + " : " + strconv.Itoa(mins) +"pm"
	} else {
		clock.Text = strconv.Itoa(hour) + " : " + strconv.Itoa(mins) +"am"
	}
	//
	widgets = append(widgets, title)
	widgets = append(widgets, left)
	widgets = append(widgets, right)
	widgets = append(widgets, clock)

	header := HSplitter{
		Background: SolidColorBrush{Color: walk.RGB(220, 0, 0)},
		MaxSize: Size{600, 25},
		Children: widgets,
	}

	return header
}


/**
 * Calculates the first day of the month in integer form
 */
func firstDayOfTheMonth() int{
	currDayName := time.Now().Weekday();
	currDayValue := time.Now().Day();

	dayPos := 0
	switch currDayName {
		case time.Monday: dayPos = 1
		case time.Tuesday: dayPos = 2
		case time.Wednesday: dayPos = 3
		case time.Thursday: dayPos = 4
		case time.Friday: dayPos = 5
		case time.Saturday: dayPos = 6
		case time.Sunday: dayPos = 7
	}

	return ((currDayValue - dayPos) % 7) - 1
}


/**
 * creates a slice of 42 widgets, which correspond to the date
 */
func makeDayLabels(length int, day int) []Widget {
	arr := make([]Widget, 0)

	firstDay := firstDayOfTheMonth()

	headers := makeDayHeaderLabels()
	for i := 0; i < 7; i++{
		arr = append(arr, headers[i])
	}

	//append actual days
	for i := 0; i < 42; i++ {
		w := TextLabel{
			TextAlignment: Alignment2D(AlignCenter),
			MaxSize: Size{ 60, 40},
			MinSize: Size{ 60, 40},
		}

		if i == day + firstDay - 1{ //todays date
			w.Text = "Day: " + strconv.Itoa(i - firstDay + 1)
			w.Background = SolidColorBrush{Color: walk.RGB(30, 150, 255)}
		} else if i >= firstDay && i <= length + firstDay{ //actual dates
			w.Text = "Day: " + strconv.Itoa(i - firstDay + 1)
			w.Background = SolidColorBrush{Color: walk.RGB(220, 220, 220)}
		} else { //not a date
			w.Background = SolidColorBrush{Color: walk.RGB(0, 220, 0)}
		}

		arr = append(arr, w)
	}

	return arr
}


/**
 * create a slice of the 7 day header widgets - used in MakeDayLabels
 */
func makeDayHeaderLabels() []Widget {
	arr := make([]Widget, 0)

	for i := 0; i < 7; i++ {
		w := TextLabel{
			TextAlignment: Alignment2D(AlignCenter),
			Background: SolidColorBrush{Color: walk.RGB(180, 180, 180)},
			MaxSize: Size{ 60, 40},
			MinSize: Size{ 60, 40},
		}

		switch i {
			case 0: w.Text = "Monday"
			case 1: w.Text = "Tuesday"
			case 2: w.Text = "Wednesday"
			case 3: w.Text = "Thursday"
			case 4: w.Text = "Friday"
			case 5: w.Text = "Saturday"
			case 6: w.Text = "Sunday"
		}

		arr = append(arr, w)
	}

	return arr
}


/**
 * calls for drawing the window
 */
func drawWindow(widgets []Widget){
	//main window
	if _, err := (MainWindow{
		Title:   "Calendar",
		//MinSize: Size{500, 450},

		Background:SolidColorBrush{Color: walk.RGB(30, 30, 30)},
		Layout: VBox{
			//Rows: 2,
			//Alignment: Alignment2D(AlignCenter), //todo
		},
		Children: widgets,

	}.Run()); err != nil { //error log
		log.Fatal(err)
	}
}


func main(){

	//month vals
	monthLengths := [12]int{31,28,31,30,31,30,31,31,30,31,30,31}
	currMonth := time.Now().String()
	currMonthValue, _ := strconv.Atoi(currMonth[5:7])
	currDay := time.Now().Day()


	//children widgets
	widgets := make([]Widget, 0)

	//actual title
	title := calendarHeader()
	widgets = append(widgets, title)

	//grid of days
	days := Composite{
		Background: SolidColorBrush{Color: walk.RGB(0, 220, 220)},
		Layout: Grid{
			Alignment:AlignHCenterVCenter, //todo
			Columns: 7,
		},
		Children: makeDayLabels(monthLengths[currMonthValue], currDay),
	}
	widgets = append(widgets, days)

	//display the window
	drawWindow(widgets)
}
