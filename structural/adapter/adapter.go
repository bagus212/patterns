package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (pr *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (pa *PrinterAdapter) PrintStored() (newMsg string) {
	if pa.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", pa.Msg)
		newMsg = pa.OldPrinter.Print(newMsg)
	} else {
		newMsg = pa.Msg
	}
	return
}
