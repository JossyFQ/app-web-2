package storage

import (
	//"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func setupRepos(t *testing.T) (TuristaRepository, NegocioRepository, *CheckInMemoria) {
	turistas, negocios, checkins := setupRepos(t)
	checkins.Guardar(models.CheckIn{
		ID: 1, TuristaID: 1, NegocioID: 1,
		Fecha: "2026-04-10", Calificacion: 5,
	})
	checkins.Guardar(models.CheckIn{
		ID: 2, TuristaID: 1, NegocioID: 1,
		Fecha: "2026-04-11", Calificacion: 4,
	})
	visitas, err := checkins.BuscarPorTurista(1)
	if err != nil {
		t.Errorf("no esperaba error: %v", err)
	}
	if len(visitas) != 2 {
		t.Errorf("esperaba 2 visitas, obtuvo %d", len(visitas))
	}

	casos := []struct {
		nombre string
		//idCheckin   int
		entrada     models.CheckIn
		errEsperado error
	}{
		{
			nombre: "Caso feliz checkin valido",
			//idCheckin: 1,
			entrada: models.CheckIn{
				ID: 1, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-04-11", Calificacion: 4,
			},
			errEsperado: nil,
		}, {
			nombre: "Fecha vacia",
			//idCheckin: 2,
			entrada: models.CheckIn{
				ID: 2, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-14-11", Calificacion: 4,
			},
			errEsperado: errs.ErrDatosInvalidos,
		}, {
			nombre: "Calificacion igual 0",
			//idCheckins: 3,
			entrada: models.CheckIn{
				ID: 1, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-04-11", Calificacion: 0,
			},
			errEsperado: errs.ErrDatosInvalidos,
		}, {
			nombre: "Calificacion igual 6",
			//idCheckins: 4,
			entrada: models.CheckIn{
				ID: 1, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-04-11", Calificacion: 6,
			},
			errEsperado: errs.ErrDatosInvalidos,
		}, {
			nombre: "Turista id que no existe en el repo del Turista",
			//idCheckins: 5,
			entrada: models.CheckIn{
				ID: 022, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-04-11", Calificacion: 4,
			},
			errEsperado: errs.ErrDatosInvalidos,
		}, { ///////////
		}, {
			nombre: "IdCheckin",
			//idCheckins: 5,
			entrada: models.CheckIn{
				ID: 97, TuristaID: 1, NegocioID: 1,
				Fecha: "2026-04-11", Calificacion: 4,
			},
			errEsperado: errs.ErrDatosInvalidos,
		},
	}
}
func TestCheckInMemoria_BuscarPorTurista(t *testing.T) {
	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	checkins := NewCheckInMemoria(turistas, negocios)

	turistas.Guardar(models.Turista{
		ID: 1, Nombre: "John", Nacionalidad: "USA", IdiomaPreferido: "en",
	})
	negocios.Guardar(models.Negocio{
		ID: 1, Nombre: "Café", Tipo: "restaurante",
		Ciudad: "Manta", IdiomasHablados: []string{"es", "en"},
	})
	checkins.Guardar(models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1,
		Fecha: "06/05/2026", Calificacion: 5},
	)

}

/*
func TestCheckInMemoria_Buscar() {
	turistas, negocios, checkins := setupRepos(t)
	checkins.Guardar(models.CheckIn{
		ID: 1, TuristaID: 1, NegocioID: 1,
		Fecha: "2026-04-10", Calificacion: 5,
	})
	checkins.Guardar(models.CheckIn{
		ID: 2, TuristaID: 1, NegocioID: 1,
		Fecha: "2026-04-11", Calificacion: 4,
	})
}
/**/
