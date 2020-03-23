package main

import . "github.com/egonelbre/exp/brutecheck/check"

func BenchLen1Ch(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if Len1Ch(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift015(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift015(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift016(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift016(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift021(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift021(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift026(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift026(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift032(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift032(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift130(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift130(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift132(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift132(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift135(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift135(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift143(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift143(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift240(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift240(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift246(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift246(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift304(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift304(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift351(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift351(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift402(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift402(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift404(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift404(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift414(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift414(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift501(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift501(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift503(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift503(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift505(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift505(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift510(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift510(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift513(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift513(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift525(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift525(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift602(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift602(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift603(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift603(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift604(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift604(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift605(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift605(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorXor_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorXor_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift016(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift016(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift031(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift031(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift240(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift240(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift302(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift302(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift351(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift351(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift400(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift400(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift401(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift401(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift405(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift405(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift413(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift413(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift500(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift500(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift505(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift505(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift511(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift511(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift512(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift512(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift601(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift601(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift603(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift603(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift605(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift605(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift606(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift606(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift616(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift616(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift620(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift620(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_XorAdd_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_XorAdd_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift015(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift015(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift026(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift026(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift135(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift135(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift230(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift230(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift301(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift301(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift304(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift304(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift412(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift412(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift510(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift510(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddXor_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddXor_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift022(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift022(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift130(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift130(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift133(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift133(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift230(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift230(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift302(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift302(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift400(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift400(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift413(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift413(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift500(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift500(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift511(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift511(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift601(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift601(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift602(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift602(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift620(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift620(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_AddAdd_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_AddAdd_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift032(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift032(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift143(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift143(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift031(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift031(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHash_OrAdd_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHash_OrAdd_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift015(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift015(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift016(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift016(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift021(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift021(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift026(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift026(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift032(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift032(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift130(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift130(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift132(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift132(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift135(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift135(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift143(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift143(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift240(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift240(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift246(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift246(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift304(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift304(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift351(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift351(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift402(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift402(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift404(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift404(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift414(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift414(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift501(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift501(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift503(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift503(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift505(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift505(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift510(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift510(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift513(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift513(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift525(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift525(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift602(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift602(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift603(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift603(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift604(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift604(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift605(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift605(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorXor_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorXor_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift016(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift016(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift031(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift031(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift240(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift240(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift302(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift302(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift351(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift351(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift400(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift400(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift401(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift401(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift405(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift405(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift413(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift413(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift500(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift500(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift505(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift505(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift511(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift511(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift512(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift512(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift601(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift601(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift603(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift603(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift605(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift605(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift606(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift606(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift616(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift616(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift620(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift620(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_XorAdd_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_XorAdd_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift015(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift015(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift026(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift026(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift135(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift135(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift230(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift230(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift301(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift301(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift304(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift304(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift361(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift412(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift412(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift510(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift510(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddXor_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddXor_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift022(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift022(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift024(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift024(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift025(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift025(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift130(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift130(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift133(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift133(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift230(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift230(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift302(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift302(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift303(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift303(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift400(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift400(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift403(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift403(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift413(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift413(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift500(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift500(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift511(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift511(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift601(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift601(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift602(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift602(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift620(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift620(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_AddAdd_Shift630(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_AddAdd_Shift630(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift032(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift032(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift034(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift034(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift035(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift035(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift043(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift043(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift045(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift045(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift046(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift046(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift143(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift143(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift145(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift145(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift146(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift146(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrXor_Shift260(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrXor_Shift260(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift031(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift031(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift033(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift033(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift041(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift041(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift042(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift042(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift050(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift050(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift051(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift051(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift060(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift060(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift061(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift061(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift062(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift062(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift073(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift073(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift150(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift150(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift161(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift161(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift250(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift250(word) {
				count++
			}
		}
	}
	return count
}

func BenchTwoHashTable_OrAdd_Shift361(words []string, repeat int) int {
	var count int
	for i := 0; i < repeat; i++ {
		for _, word := range words {
			if TwoHashTable_OrAdd_Shift361(word) {
				count++
			}
		}
	}
	return count
}
