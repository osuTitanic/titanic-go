package templates

import (
	"fmt"
	"reflect"

	"github.com/CloudyKit/jet/v6"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

func formatNumber(a jet.Arguments) reflect.Value {
	a.RequireNumOfArguments("formatNumber", 1, 1)

	var result string
	var value any = a.Get(0).Interface()

	switch value := reflect.ValueOf(value); value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = printer.Sprintf("%d", value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = printer.Sprintf("%d", value.Uint())
	default:
		result = fmt.Sprint(value.Interface())
	}

	return reflect.ValueOf(result)
}
