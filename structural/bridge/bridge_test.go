package bridge

import (
	"strings"
	"testing"
)

func TestPrinterAPI1(t *testing.T) {
	apiv1 := PrinterImpl1{}
	err := apiv1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")
	if err != nil {
		expectedErrMessage := "You need pass an io.writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrMessage) {
			t.Errorf("Error Message was not correct.\nActual : %s\nExpected: %s\n", err.Error(), expectedErrMessage)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{
		&testWriter,
	}
	expectedMessage := "Hello"
	err = api2.PrintMessage((expectedMessage))
	if err != nil {
		t.Errorf("error trying to API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}

}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writter"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	if normal.Msg != testWriter.Msg {
		t.Fatalf("API2 did not write correctly on the io.Writer. \nActual: %s\nExpected: %s\n", testWriter.Msg, normal.Msg)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packet: Hello io.Writer"
	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}
	err = packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
