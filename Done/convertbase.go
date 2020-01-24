package student

func ConvertBase(s, t, p string) string {
	ln := 0
	ln2 := 0
	ln3 := 0
	ans := ""
	st_t := map[rune]int{}
	for c := range s {
		ln = c + 1
	}
	for i, c := range t {
		st_t[c] = i
		ln2 = i + 1
	}
	for c := range p {
		ln3 = c + 1
	}
	pw := 1
	cnt := 0
	for i := ln - 1; i >= 0; i-- {
		cnt = cnt + st_t[[]rune(s)[i]]*pw
		pw *= ln2
	}
	for cnt != 0 {
		ans = string(p[cnt%ln3]) + ans
		cnt /= ln3
	}
	return ans
}

// package student

// func tostring(nbr int) string {
// 	n := nbr
// 	dig := 0
// 	for n != 0 {
// 		dig++
// 	}
// 	ans := ""
// 	for j := dig - 1; j >= 0; j-- {
// 		ans += string((nbr % 10) + 48)
// 		nbr /= 10
// 	}
// 	return ans
// }
// func ConvertBase(nbr, baseFrom, baseTo string) string {
// 	firstnumber := AtoiBase(nbr, baseFrom)
// 	final := tostring(AtoiBase(tostring(firstnumber), baseTo))
// 	return final

// }
