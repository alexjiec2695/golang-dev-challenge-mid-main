package download_test

import (
	"golang-dev-challenge-mid/pkg/download"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tt := []struct {
		name        string
		path        string
		error       bool
		expectError string
		nameFile    string
	}{
		{
			name:     "test success",
			path:     "http://bibliotecadigital.ilce.edu.mx/Colecciones/ObrasClasicas/_docs/Corazon_Amicis.pdf",
			error:    false,
			nameFile: "Corazon_Amicis.pdf",
		},
		{
			name:     "test success",
			path:     "http://bibliotecadigital.ilce.edu.mx/Colecciones/ObrasClasicas/_docs/10Cuentos_LasMilNoches.pdf",
			error:    false,
			nameFile: "10Cuentos_LasMilNoches.pdf",
		},
		{
			name:     "test success",
			path:     "http://bibliotecadigital.ilce.edu.mx/Colecciones/ObrasClasicas/_docs/Asesinato.pdf",
			error:    false,
			nameFile: "Asesinato.pdf",
		},
		{
			name:     "test success",
			path:     "http://bibliotecadigital.ilce.edu.mx/Colecciones/ObrasClasicas/_docs/JardinCerezos.pdf",
			error:    false,
			nameFile: "JardinCerezos.pdf",
		},
		{
			name:        "test failed parsing url",
			path:        "http://bibliotecadigital.ilce.edu.mx",
			error:       true,
			expectError: "URL is empty",
		},
		{
			name:        "creating file",
			path:        "http://bibliotecadigital.ilce.edu.mx//",
			error:       true,
			expectError: "creating file: open : The system cannot find the file specified.",
		},

		{
			name:        "calling document",
			path:        "http://bibliotecadigital.ilce.edu.mx/Corazon_Amicis.pdf",
			error:       true,
			expectError: "calling document: 404",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			err := download.Run(tc.path)

			if tc.error {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectError)
			} else {
				assert.FileExists(t, tc.nameFile)
				assert.NoError(t, err)
				if err = os.Remove(tc.nameFile); err != nil {
					t.FailNow()
				}
			}

		})
	}
}
