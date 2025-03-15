package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello User !!")

	//Get Nutritional score from data entered by the user

	nutritionalScore := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(600),
		Sugars:              SugarGram(100),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
		Sodium:              SodiumMilligram(10),
		Fruits:              FruitsPercent(50), //couldnt find it in Functional Spec
		Fiber:               FiberGram(8),
		Protein:             ProteinGram(2),
		//IsWater: Not added yet
	}, Food)

	fmt.Printf("The Nutritional Score is %d\n", nutritionalScore.Value)
	fmt.Printf("The NutriScore is %s\n", nutritionalScore.GetNutriScore())

}
