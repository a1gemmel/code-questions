package lib

type BigInt struct {
	a []int
}

func (a BigInt) Length() int {
	return len(a.a)
}

func (a BigInt) Add(b BigInt) BigInt {
	if b.Length() > a.Length() {
		return b.Add(a)
	}

	var c []int
	carry := 0
	for i := 0; i < len(a.a); i++ {
		bDig := 0
		if b.Length() > i {
			bDig = b.a[i]
		}
		sum := a.a[i] + bDig + carry
		carry = sum / 10
		sum = sum % 10
		c = append(c, sum)
	}
	if carry != 0 {
		c = append(c, carry)
	}

	return BigInt{a: c}
}

func (bi BigInt) Multiply(x int) BigInt {
	length := len(bi.a)
	for i := length - 1; i >= 0; i-- {
		bi.a[i] *= x
		if bi.a[i] > 9 {
			carry := bi.a[i] / 10
			bi.a[i] %= 10
			for j := i + 1; carry != 0; j++ {
				if j >= len(bi.a) {
					bi.a = append(bi.a, 0)
				}
				bi.a[j] += carry
				if bi.a[j] > 9 {
					carry = bi.a[j] / 10
					bi.a[j] %= 10
				} else {
					carry = 0
				}
			}
		}
	}
	return bi
}

func (bi BigInt) Repr() []int {
	return bi.a
}

func NewBigInt(n int) BigInt {
	a := []int{}
	for n != 0 {
		a = append(a, n%10)
		n = n / 10
	}
	return BigInt{a: a}
}

func BigIntFactorial(n int) BigInt {
	b := NewBigInt(1)
	for i := 1; i <= n; i++ {
		b = b.Multiply(i)
	}
	return b
}
