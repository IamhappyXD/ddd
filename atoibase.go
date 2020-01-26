package student

func AtoiBase(s string, base string) int {
 out := 0
 zxc := 0
 q := map[rune]bool{}
 for _, t := range base {
  if q[t] || t == '-' || t == '+' {
   return out
  }
  q[t] = true
  zxc++
 }
 if zxc > 1 {
  for _, t := range s {
   d := 0
   if q[t] {
    for _, j := range base {
     if j == t {
      break
     }
     d++
    }
    out = out*zxc + d
   }
  }
 }
 return out