package main

import "math/rand"

// Predefined routes for drivers (Mapped to New Delhi, India)
// Center Point: Connaught Place
var PredefinedRoutes = [][][]float64{
	{
		// Route 1: Inner Circle to Janpath
		{28.6328, 77.2195},
		{28.6315, 77.2201},
		{28.6298, 77.2212},
		{28.6275, 77.2218},
	},
	{
		// Route 2: Minto Road toward Deen Dayal Upadhyaya Marg
		{28.6342, 77.2225},
		{28.6355, 77.2238},
		{28.6361, 77.2245},
		{28.6355, 77.2238},
		{28.6370, 77.2252},
		{28.6385, 77.2268},
		{28.6391, 77.2274},
		{28.6382, 77.2295},
		{28.6365, 77.2289},
		{28.6350, 77.2285},
		{28.6358, 77.2305},
	},
	{
		// Route 3: Sansad Marg (Parliament Street) loop
		{28.6295, 77.2152},
		{28.6268, 77.2145},
		{28.6255, 77.2141},
		{28.6252, 77.2162},
		{28.6258, 77.2185},
		{28.6262, 77.2205},
		{28.6245, 77.2201},
		{28.6232, 77.2198},
	},
	{
		// Route 4: Reverse Path (Patel Chowk back towards Tolstoy Marg)
		{28.6232, 77.2198},
		{28.6245, 77.2201},
		{28.6262, 77.2205},
		{28.6258, 77.2185},
		{28.6252, 77.2162},
		{28.6268, 77.2145},
		{28.6295, 77.2152},
		{28.6255, 77.2141},
	},
}

func GenerateRandomPlate() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plate := ""
	for i := 0; i < 3; i++ {
		plate += string(letters[rand.Intn(len(letters))])
	}

	return plate
}
