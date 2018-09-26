package main

contract Comparable(t T) {
	t > t
}

func Min(type T Comparable)(a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max(type T Comparable)(a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func minMaxExample() {
	player.HP = Max(player.HP - damage, 0)
}

contract Signed(t T) {
	t = -t
}

func Abs(type T Signed)(n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(type T Signed)(n T) T {
	if n >= 0 {
		return (T)(1)
	}
	return (T)(-1)
}

func absExample() {
	var x0, x1 int
	// ...
	distance := Abs(x0 - x1)
	sign := Sign(distance) // 1 or -1

	// this works too...?
	hmm := Sign(uint(1337)) // always 1
}

// startmax
func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func maxf32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
// endmax

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absi(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(i int) int {
	if i >= 0 {
		return 1
	} else {
		return -1
	}
}