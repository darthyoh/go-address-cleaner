package cleaner

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	description, input, expected string
}{
	{
		"simple uppercase", "Test", "TEST",
	},
	{
		"more uppercase", "TeSting", "TESTING",
	},
	{
		"unchanged number", "12345", "12345",
	},
	{
		"unchanged letters", "ablation", "ABLATION",
	},
	{
		"special chars", "1 St-Martin Street 75000 Paris - CEDEX 3", "1 ST MARTIN STREET 75000 PARIS CEDEX 3",
	},
	{
		"others stupid chars", "###{#}C_6-a", "C 6 A",
	},
	{
		"A variant", "à pâques päis", "A PAQUES PAIS",
	},
	{
		"C variant", "ça commence", "CA COMMENCE",
	},
	{
		"E variant", "éventuellement noël mèche être ou ne pas être", "EVENTUELLEMENT NOEL MECHE ETRE OU NE PAS ETRE",
	},
	{
		"I variant", "île maïs", "ILE MAIS",
	},
	{
		"O variant", "ôh öh ôh", "OH OH OH",
	},
	{
		"U variant", "où est oû", "OU EST OU",
	},
	{
		"Y variant", "ÿ est peu utile", "Y EST PEU UTILE",
	},
	{
		"Complète adresse", "3 rue de l'au delà", "3 RUE DE L AU DELA",
	},
	{
		"Complète adresse", "5 rue de l'île aux étourneaux, St-Martin ", "5 RUE DE L ILE AUX ETOURNEAUX ST MARTIN",
	},
	{
		"Complète adresse", "1 avenue des champs-élysées 75000 Paris", "1 AVENUE DES CHAMPS ELYSEES 75000 PARIS",
	},
}

func TestCleaner(t *testing.T) {
	for _, test := range testCases {
		if result := Clean(test.input); result != test.expected {
			fmt.Println([]rune(result))
			fmt.Println([]rune(test.expected))
			t.Fatalf("%s failed. get %s want %s", test.description, result, test.expected)
		}
	}
}
