package utils

import (
	"fmt"
	"myproject/models"
	"os"

	"github.com/dop251/goja"
)

func GenerateMilsymbol(milsymbol models.Milsymbol) (string, error) {

	// Leer el archivo JavaScript de milsymbol
	jsFile, err := os.ReadFile("./web/pkg/utils/milsymbol.js")

	if err != nil {

		return "", fmt.Errorf("error al leer el archivo JavaScript: %v", err)
	}

	// Crear un nuevo runtime de goja
	vm := goja.New()

	// Ejecutar el código de milsymbol en goja
	_, err = vm.RunString(string(jsFile))
	if err != nil {
		return "", fmt.Errorf("error al ejecutar milsymbol: %v", err)
	}

	// Script de JavaScript para generar un símbolo basado en el código recibido
	script := fmt.Sprintf(`
        var ms = new ms.Symbol("%s", { size: %d });
        ms.asSVG();
    `, milsymbol.SymbolCode, milsymbol.Size)

	// Ejecutar el script
	result, err := vm.RunString(script)
	if err != nil {
		return "", fmt.Errorf("error al ejecutar el script: %v", err)
	}

	return result.String(), nil
}
