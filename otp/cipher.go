package otp

import "math/rand"

type OTPCipher struct {
	key []byte
}

func NewOTPCipher(k []byte) OTPCipher {
	return OTPCipher{
		key: k,
	}
}

func GenerateOTPKey(l int) []byte {
	key := make([]byte, l)

	i := 0
	for i < len(key) {
		key[i] = sampleRandomDistribution(0, 256)
		i++
	}
	return key
}

func sampleRandomDistribution(min, max int) byte {
	value := rand.Intn(max-min) + max

	return byte(value)
}

func (c *OTPCipher) Encrypt(pt []byte) []byte {
	return c.applyKey(pt)
}

func (c *OTPCipher) Decrypt(ct []byte) []byte {
	return c.applyKey(ct)
}

func (c *OTPCipher) applyKey(input []byte) []byte {
	output := make([]byte, len(input))

	for i, b := range input {
		output[i] = b ^ c.key[i]
	}

	return output
}
