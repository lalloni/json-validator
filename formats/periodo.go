package formats

import (
	"fmt"

	"github.com/lalloni/afip/periodo"
)

var PeriodoDiario = FormatCheckerFunc(checkPeriodoDiario)
var PeriodoMensual = FormatCheckerFunc(checkPeriodoMensual)
var PeriodoAnual = FormatCheckerFunc(checkPeriodoAnual)

func checkPeriodoDiario(in interface{}) bool {
	ok, _, _, _ := periodo.Parse(periodo.Diario, fmt.Sprint(in))
	return ok
}

func checkPeriodoMensual(in interface{}) bool {
	ok, _, _, _ := periodo.Parse(periodo.Mensual, fmt.Sprint(in))
	return ok
}

func checkPeriodoAnual(in interface{}) bool {
	ok, _, _, _ := periodo.Parse(periodo.Anual, fmt.Sprint(in))
	return ok
}
