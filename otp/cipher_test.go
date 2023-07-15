package otp

import (
	"reflect"
	"testing"
)

func TestRandomDistribution(t *testing.T) {
	runs := 256 * 100

	for i := 0; i < runs; i++ {
		v := sampleRandomDistribution(0, 256)

		if v > 255 {
			t.Errorf("Generated value is out of range. got=%b", v)
		}
	}
}

func TestGenerateOTPKey(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{
			10,
			10,
		},
		{
			100,
			100,
		},
		{
			1000,
			1000,
		},
	}

	for _, tt := range tests {
		k := GenerateOTPKey(tt.input)

		if len(k) != tt.output {
			t.Errorf("Key is not expected length. got=%d, expected=%d", len(k), tt.output)
		}

		for i := 0; i < len(k); i++ {
			if k[i] > 255 {
				t.Errorf("Generated value is out of at key[%d]. got=%b", i, k[i])
			}
		}
	}
}

func TestEncryptDecrypt(t *testing.T) {

	ptString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	pt := []byte(ptString)
	k := GenerateOTPKey(len(pt))

	if len(k) != len(pt) {
		t.Errorf("Key is not the same length as the plaintext. expected=%d, got=%d", len(pt), len(k))
	}

	c := NewOTPCipher(k)
	ct := c.Encrypt(pt)

	if len(ct) != len(pt) {
		t.Errorf("Ciphertext is not the same length as the plaintext. expected=%d, got=%d", len(pt), len(ct))
	}

	pt2 := c.Decrypt(ct)

	if len(ct) != len(pt2) {
		t.Errorf("Output plaintext is not the same length as the ciphertext. expected=%d, got=%d", len(ct), len(pt2))
	}

	if !reflect.DeepEqual(pt2, pt) {
		t.Errorf("Output plaintext is not equal to input plaintext.")

	}

}
