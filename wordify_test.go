package wordify_test

import (
	"testing"

	"github.com/psotou/wordify"
)

func TestWordify(t *testing.T) {
	t.Run("Números menores a 1_000", func(t *testing.T) {
		tests := []struct {
			Case  string
			Input int
			Want  string
		}{
			{
				Case:  "número menor a 30",
				Input: -27,
				Want:  "menos veintisiete",
			},
			{
				Case:  "número mayor o igual a 30 y menores que 100",
				Input: -83,
				Want:  "menos ochenta y tres",
			},
			{
				Case:  "número mayor o igual a 30 y menores que 100 que terminen en 0",
				Input: 60,
				Want:  "sesenta",
			},
			{
				Case:  "número entre 100 y 999",
				Input: 874,
				Want:  "ochocientos setenta y cuatro",
			},
			{
				Case:  "número entre 100 y 999",
				Input: 501,
				Want:  "quinientos uno",
			},
			{
				Case:  "número entre 100 y 999",
				Input: 500,
				Want:  "quinientos",
			},
			{
				Case:  "número menor a 10",
				Input: 4,
				Want:  "cuatro",
			},
		}

		for _, tc := range tests {
			got := wordify.Int(tc.Input)
			if got != tc.Want {
				t.Errorf("\ngot : %v\nwant: %v\n", got, tc.Want)
			}
		}
	})

	t.Run("Números entre 1_000 y 1_000_000", func(t *testing.T) {
		tests := []struct {
			Case  string
			Input int
			Want  string
		}{
			{
				Case:  "número menor a 30_000",
				Input: 12_150,
				Want:  "doce mil ciento cincuenta",
			},
			{
				Case:  "número mayor o igual a 30_000 y menores que 100_000",
				Input: 73_233,
				Want:  "setenta y tres mil doscientos treinta y tres",
			},
			{
				Case:  "número mayor o igual a 30_000 y menores que 100_000 que terminen en 0",
				Input: 81_230,
				Want:  "ochenta y un mil doscientos treinta",
			},
			{
				Case:  "número entre 100_000 y 999_999",
				Input: 731_909,
				Want:  "setecientos treinta y un mil novecientos nueve",
			},
			{
				Case:  "número menor a 10_000",
				Input: 3_242,
				Want:  "tres mil doscientos cuarenta y dos",
			},
		}

		for _, tc := range tests {
			got := wordify.Int(tc.Input)
			if got != tc.Want {
				t.Errorf("\ngot : %v\nwant: %v\n", got, tc.Want)
			}
		}
	})
	t.Run("Números entre 1_000_000 y 1_000_000_000", func(t *testing.T) {
		tests := []struct {
			Case  string
			Input int
			Want  string
		}{
			{
				Case:  "número menor a 30_000_000",
				Input: 13_221_820,
				Want:  "trece millones doscientos veintiún mil ochocientos veinte",
			},
			{
				Case:  "número mayor o igual a 30_000_000 y menores que 100_000_000",
				Input: 45_100_001,
				Want:  "cuarenta y cinco millones cien mil uno",
			},
			{
				Case:  "número mayor o igual a 30_000_000 y menores que 100_000_000 que terminen en 0",
				Input: 70_901_000,
				Want:  "setenta millones novecientos un mil",
			},
			{
				Case:  "número entre 100_000_000 y 1_000_000_000",
				Input: 601_001_856,
				Want:  "seiscientos un millones mil ochocientos cincuenta y seis",
			},
			{
				Case:  "número entre 100_000_000 y 1_000_000_000",
				Input: 721_001_856,
				Want:  "setecientos veintiún millones mil ochocientos cincuenta y seis",
			},
			{
				Case:  "número menor a 10_000_000",
				Input: 1_001_221,
				Want:  "un millón mil doscientos veintiuno",
			},
		}

		for _, tc := range tests {
			got := wordify.Int(tc.Input)
			if got != tc.Want {
				t.Errorf("\ngot : %v\nwant: %v\n", got, tc.Want)
			}
		}
	})

	t.Run("Números entre 1_000_000_000 y 1_000_000_000_000", func(t *testing.T) {
		tests := []struct {
			Case  string
			Input int
			Want  string
		}{
			{
				Case:  "número menor a 30_000_000_000",
				Input: 1_001_823_000,
				Want:  "mil un millón ochocientos veintitrés mil",
			},
			{
				Case:  "número mayor o igual a 30_000_000_000 y menores que 100_000_000_000",
				Input: 91_021_232_020,
				Want:  "noventa y un mil veintiún millones doscientos treinta y dos mil veinte",
			},
			{
				Case:  "número mayor o igual a 30_000_000_000 y menores que 100_000_000_000 que terminen en 0",
				Input: 43_070_901_000,
				Want:  "cuarenta y tres mil setenta millones novecientos un mil",
			},
			{
				Case:  "número entre 100_000_000_000 y 1_000_000_000_000",
				Input: 200_002_921_807,
				Want:  "doscientos mil dos millones novecientos veintiún mil ochocientos siete",
			},
			{
				Case:  "número menor a 10_000_000_000",
				Input: 6_013_020_231,
				Want:  "seis mil trece millones veinte mil doscientos treinta y uno",
			},
		}

		for _, tc := range tests {
			got := wordify.Int(tc.Input)
			if got != tc.Want {
				t.Errorf("\ngot : %v\nwant: %v\n", got, tc.Want)
			}
		}
	})
}
