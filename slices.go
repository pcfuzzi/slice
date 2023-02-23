package main

import "fmt"

func main() {
    slices1 := []int{5, 8, 10, 11, 15}
    slices2 := []int{5, 8, 10, 11, 15}
	slices3 := []int{5, 8, 10, 11, 18}

   // equal := compareSlice(slices1, slices2)

   // fmt.Printf("Slices are equal: %t\n", equal)

   fmt.Println(compareSlice(slices1, slices2))
   fmt.Println(compareSlice(slices1, slices3))
}

// compareSlice vergleicht 2 slices und gibt true zurück, wenn sie gleich sind,
// oder es gibt false, wenn sie unterschiedlich sind
func compareSlice(slices1, slices2 []int) bool {
    // erster Vergleich: haben beide Slices die selbe Länge?
    if len(slices1) != len(slices2) {
        return false
    }

    // zweiter Vergleich: sind die Inhalte identisch?
    for i := range slices1 {
        if slices1[i] != slices2[i] {
            return false
        }
    }

    return true
}
