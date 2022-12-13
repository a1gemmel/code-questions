# func hasAlternatingBits(n int) bool {
# 	lastBit := n % 2
# 	n /= 2
# 	for n > 0 {
# 		bit := n % 2
# 		n /= 2
# 		if bit == lastBit {
# 			return false
# 		}
# 		lastBit = bit
# 	}
# 	return true
# }


class Solution:
    def hasAlternatingBits(self, n: int) -> bool:
        lastBit = n % 2
        n //= 2
        while n > 0:
            bit = n % 2
            n //= 2
            if bit == lastBit:
                return False
            lastBit = bit
            
        return True