package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
    return len(layers) * time
} 

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
    var sauce float64
    var noodles int
    for i := 0; i < len(layers); i++ {
        if layers[i] == "sauce" {
            sauce += 0.2
        } else if layers[i] == "noodles" {
            noodles += 50
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendRecipe []string, myRecipe []string) {
    myRecipe[len(myRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(amountsNeeded []float64, portions int) []float64 {
    var scaledQuantities []float64
    scaledQuantities = append(scaledQuantities, amountsNeeded...)
    for i := 0; i < len(scaledQuantities); i++ {
        scaledQuantities[i] *= float64(portions)/2
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
