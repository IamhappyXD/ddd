package student

func tonumber(s string) int {
	val := []rune(s)
	numb := 0
	for _, v := range val {
		numb *= 10
		numb += int(v) - 48
	}
	return numb
}
func slen(s string) int {
	check := []rune(s)
	ans := 0
	for range check {
		ans++
	}
	return ans
}
func AtoiBase(s string, base string) int {
	if base == "0123456789" {
		numb := tonumber(s)
		return numb
	} else if base == "01" {
		numb := tonumber(s)
		dec := 0
		i := 0
		remainder := 0
		for numb != 0 {
			remainder = numb % 10
			numb /= 10
			dec += remainder * RecursivePower(2, i)
			i++
		}
		return dec
	} else if base == "0123456789ABCDEF" {
		dec := 0
		numb := []rune(s)
		i := 1
		ln := 0
		for range numb {
			ln++
		}
		for j := ln - 1; j >= 0; j-- {
			if numb[j] >= 'A' && numb[j] <= 'F' {
				dec += (int(numb[j]) - 55) * i
			}
			if numb[j] >= '0' && numb[j] <= '9' {
				dec += (int(numb[j]) - 48) * i
			}
			i *= 16
		}
		return dec
	} else {
		strln := slen(s)
		blen := slen(base)
		bs := []rune(base)
		st := []rune(s)
		ans := 0
		if strln > blen {
			return 0
		}
		for i, x := range st {
			ind := 0
			for j := 0; j < blen; j++ {
				if x == bs[j] {
					ind = j
					break
				}
			}
			ans += ind * RecursivePower(6, strln-i-1)
		}
		return ans
	}
	return 0

}
