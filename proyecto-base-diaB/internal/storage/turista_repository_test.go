package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestGuardar_Turista(t *testing.T) {
	repo := NewTuristaMemoria()

	turistaBase := models.Turista{
		ID: 1, Nombre: "Carlos", Nacionalidad: "Peruano",
		IdiomaPreferido: "es",
	}
	if err := repo.Guardar(turistaBase); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	casos := []struct {
		nombre    string
		entrada   models.Turista
		esperaErr error
	}{
		{
			nombre: "Turista Guardado",
			entrada: models.Turista{
				ID: 81, Nombre: "Carlos Tres", Nacionalidad: "Peruano",
				IdiomaPreferido: "es",
			},
			esperaErr: nil,
		},
		{
			nombre: "Turista no guardado",
			entrada: models.Turista{
				ID: 82, Nombre: "", Nacionalidad: "Peruano",
				IdiomaPreferido: "",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "Idioma de turista no válido",
			entrada: models.Turista{
				ID: 83, Nombre: "JuansX", Nacionalidad: "peruano",
				IdiomaPreferido: "ch",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.Turista{
				ID: 1, Nombre: "Carlos Segundo", Nacionalidad: "Peruano",
				IdiomaPreferido: "es",
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := repo.Guardar(c.entrada)

			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(%q): esperaba error=%v, obtuvo error=%v",
					c.entrada.Nombre, c.esperaErr, err)
			}
		})
	}
}

func TestBuscarPorID_TuristaExiste(t *testing.T) {
	repo := NewTuristaMemoria()

	// Arrange: creamos y guardamos un negocio.
	esperado := models.Turista{
		ID: 1, Nombre: "Carlos", Nacionalidad: "Peruano",
		IdiomaPreferido: "es",
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}
	// Act: buscamos el negocio por su ID.
	obtenido, err := repo.BuscarPorID(1)
	// Assert: no debe haber error y debe coincidir con lo guardado.
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if obtenido.ID != esperado.ID {
		t.Errorf("ID: esperaba %d, obtuvo %d", esperado.ID, obtenido.ID)
	}
	if obtenido.Nombre != esperado.Nombre {
		t.Errorf("Nombre: esperaba %q, obtuvo %q", esperado.Nombre, obtenido.Nombre)
	}

}

func TesTListar_TuristaEsta(t *testing.T) {
	repo := NewTuristaMemoria()
	p := models.Turista{ID: 1, Nombre: "Carlos", Nacionalidad: "Peruano",
		IdiomaPreferido: "es",
	}
	err := repo.Guardar(p)
	if err != nil {
		t.Errorf("No esperaba error %v", err)
	}
	if len(repo.Listar()) != 1 {
		t.Error("Esperaba error")
	}
}
