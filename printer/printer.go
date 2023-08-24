package printer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/logrusorgru/aurora"
)

type Message struct {
	Description string
	Value       interface{}
	ValueList   []interface{}
	ValueMap    map[string]interface{}
}

// Store OutputType for reference when posting messages
type Output struct {
	// OutputType   OutputType
	JsonOutput     bool
	CsvOutput      bool
	HtmlOutput     bool
	MarkDownOutput bool
	DebugOut       bool
	Messages       []*Message
	TableHeaders   []interface{}
	TableRows      [][]interface{}
}

func New(jsonOutput bool, csvOutput bool, htmlOutput bool, markDownOutput bool, debugOut bool) *Output {
	o := &Output{
		JsonOutput:     jsonOutput,
		CsvOutput:      csvOutput,
		HtmlOutput:     htmlOutput,
		MarkDownOutput: markDownOutput,
		DebugOut:       debugOut,
	}

	return o
}

// Error message to user
func (o *Output) AddErrorMessage(description string, value interface{}, exit bool) {
	// only output for default
	if o.JsonOutput {
		return
	}
	fmt.Println(aurora.Red("[ERROR]\t"), aurora.Red(description), value)
	// Option to exit program
	if exit {
		os.Exit(1)
	}
}

// Warning Message to user
func (o *Output) AddWarningMessage(value string) {
	// only output for default
	if o.JsonOutput {
		return
	}
	fmt.Println(aurora.Yellow("[WARNING]\t"), value)
}

// Info Message to user
func (o *Output) AddInfoMessage(value string) {
	// only output for default
	if o.JsonOutput {
		return
	}
	fmt.Println(aurora.BgBrightGreen("[INFO]\t"), aurora.Green(value))
}

// User Prompt Message to user
func (o *Output) AddUserInputMessage(value string) {
	// only output for default
	if o.JsonOutput {
		return
	}
	fmt.Println(aurora.Yellow("[User Input]\t"), aurora.Yellow(value))
}

// Debug message to user
func (o *Output) AddDebugMessage(description string, value interface{}) {
	// only output for default
	if o.JsonOutput {
		return
	}
	if o.DebugOut {
		fmt.Println("[DEBUG]\t", description, value)
	}
}

// Adds a message that will not print until Close() is called, to print and align
// Single primitive Value (string, int, bool)
func (o *Output) AddDeferredMessageRead(description string, value interface{}) {
	o.Messages = append(o.Messages, &Message{description, value, nil, nil})
}

// Adds a message that will not print until Close() is called, to print and align
// List of primitive Values (string, int, bool)
func (o *Output) AddDeferredListMessageRead(description string, value []interface{}) {
	o.Messages = append(o.Messages, &Message{description, nil, value, nil})
}

// Adds a message that will not print until Close() is called, to print and align
// Map Value, string -> primitive (string, int, bool)
func (o *Output) AddDeferredMapMessageRead(description string, value map[string]interface{}) {
	o.Messages = append(o.Messages, &Message{description, nil, nil, value})
}

// create a row from an array of interfaces, required since table.Row{} uses a variadic
// should work for any type
func createRow(items []interface{}) table.Row {
	h := make([]interface{}, len(items))
	copy(h, items)
	return h
}

// print
func (o Output) closeMessagesDefault() {
	maxLength := 0
	// determine spacing based on largest message description (left justify)
	for _, k := range o.Messages {
		if len(k.Description) > maxLength {
			maxLength = len(k.Description)
		}
	}

	for _, k := range o.Messages {
		if k.Value != nil {
			fmt.Println(fmt.Sprintf("%-*s", maxLength+1, aurora.Bold(k.Description+":")), aurora.Blue(k.Value))
		}
		if k.ValueList != nil {
			fmt.Printf("%-*s\n", maxLength+1, aurora.Bold(k.Description+":"))
			for _, v := range k.ValueList {
				// fmt.Println(fmt.Sprintf("%-*s", maxLength+1, ""), aurora.Blue(v))
				fmt.Printf("  %s\n", aurora.Blue(v))
			}
		}
		if k.ValueMap != nil {
			fmt.Printf("%-*s\n", maxLength+1, aurora.Bold(k.Description+":"))

			// determine spacing based on largest key in map (left justify)
			maxLengthMap := 0
			for k := range k.ValueMap {
				if len(k) > maxLengthMap {
					maxLengthMap = len(k)
				}
			}
			for k, v := range k.ValueMap {
				fmt.Println(fmt.Sprintf("  %-*s", maxLengthMap+1, aurora.Bold(k+":")), aurora.Blue(v))
			}
		}
	}
}

func (o Output) closeMessagesJson() {
	tempMap := make(map[string]interface{}, len(o.Messages))
	for _, k := range o.Messages {
		if k.Value != nil {
			tempMap[k.Description] = k.Value
		}
		if k.ValueList != nil {
			tempMap[k.Description] = k.ValueList
		}
		if k.ValueMap != nil {
			tempMap[k.Description] = k.ValueMap
		}
	}

	b, _ := json.Marshal(tempMap)
	fmt.Println(string(b))
}

func (o Output) closeTableDefault() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	t.AppendHeader(createRow(o.TableHeaders))
	for _, i := range o.TableRows {
		t.AppendRow(createRow(i))
	}
	t.SetStyle(table.StyleRounded)

	if o.CsvOutput {
		t.RenderCSV()
	} else if o.MarkDownOutput {
		t.RenderMarkdown()
	} else if o.HtmlOutput {
		t.RenderHTML()
	} else {
		t.Render()
	}

}

// Craziness... Create an array of maps that can then be Marshalled to JSON
func (o Output) closeTableJson() {
	tempList := make([]map[string]interface{}, len(o.TableRows))
	for index1, v1 := range o.TableRows {
		tempMap := make(map[string]interface{}, len(o.TableHeaders))
		for index2, v2 := range o.TableHeaders {
			if index2 < len(v1) { // quick check for the case where there are more header values than elements in the row
				tempMap[v2.(string)] = v1[index2]
			}
		}
		tempList[index1] = tempMap
	}

	b, _ := json.Marshal(tempList)
	fmt.Println(string(b))
}

func (o *Output) Close() {
	if o == nil { // in the case this has not been initialized yet
		return
	}
	if len(o.Messages) > 0 {
		if o.JsonOutput {
			o.closeMessagesJson()
		} else {
			o.closeMessagesDefault()
		}
		o.Messages = o.Messages[:0]
	}

	if len(o.TableHeaders) > 0 {
		if o.JsonOutput {
			o.closeTableJson()
		} else {
			o.closeTableDefault()
		}
		o.TableHeaders = o.TableHeaders[:0]
		o.TableRows = o.TableRows[:0]
	}
}

// Add headers for table, used for a list of items
func (o *Output) AddTableHeaders(headers ...interface{}) {
	o.TableHeaders = headers
}

// Add rows for at table, will be matched to headers set in AddTableHeaders, used for a list of items
func (o *Output) AddTableRows(rows ...interface{}) {
	o.TableRows = append(o.TableRows, rows)
}
