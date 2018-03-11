package neuralNetwork

import (
	"github.com/fxsjy/gonn/gonn"
	"log"
)



func Test() {
	//nn := getNN()
	//if nn == nil {
		nn := createNN()
		log.Println("creating nn")
	//}

	trainNN(nn)
	gonn.DumpNN("testNN", nn)


}


func createNN() *gonn.NeuralNetwork {
	return gonn.DefaultNetwork(4, 16, 1, false)
}


func trainNN(nn *gonn.NeuralNetwork) {
	/*input := [][]interface{} {
		{"repeat", "Something that occurs or is done again.", 6},
		{"fraction", "A numerical quantity that is not a whole number (e.g. 1/2, 0.5).", 5},
		{"neck", "The part of a person's or animal's body connecting the head to the rest of the body.", 12},
		{"rebellion", "An act of armed resistance to an established government or leader", 2},
		{"equinox", "The time or date (twice each year) at which the sun crosses the celestial equator, when day and night are of equal length (about 22 September and 20 March).", 2},
	}*/

	input := [][]float64 {
		{6, 32, 6, 3068},
		{8, 53, 5, 5374},
		{7, 127, 2, 52314},
	}

	// Теперь создаём "цели" - те результаты, которые нужно получить
	target := [][]float64 {
		{0.3},
		{0.5},
		{0.9},
	}

	nn.Train(input, target, 100000)

	out := nn.Forward([]float64{12, 49, 3, 13958})
	log.Println("res:", getResult(out))
}


func getNN() *gonn.NeuralNetwork{
	return gonn.LoadNN("testNN")
}


func getResult(output []float64) float64 {

	var (
		max float64 = -99999
	)


	for _, value := range output {
		if value > max {
			max = value
		}
	}

	return max
}