package encryption

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryption(t *testing.T) {
	os.Chdir("..")
	os.Chdir("..")

	encrypt := New()
	t.Run("EncryptFunc", func(t *testing.T) {
		inputRightAlgorithm := []string{"MD5", "SHA-256"}
		inputWrongAlgorithm := []string{"md5", "SHA256", "sha256"}
		for _, alg := range inputRightAlgorithm {
			_, err := encrypt.Encrypt("chipi chipi chapa chapa", alg)
			require.NoError(t, err)
		}
		for _, alg := range inputWrongAlgorithm {
			_, err := encrypt.Encrypt("chipi chipi chapa chapa", alg)
			require.ErrorIs(t, err, ErrUndefinedAlgorithm)
		}

		_, err := encrypt.Encrypt("                              ", inputRightAlgorithm[0])
		require.NoError(t, err)

		_, err = encrypt.Encrypt("", inputRightAlgorithm[1])
		require.ErrorIs(t, err, ErrEmptyString)
	})

	t.Run("EncryptionFuncsSHA256", func(t *testing.T) {
		inputStrings := []string{"Hello World", "There is a string", "Wingardium Leviosa"}
		expectedCyphers := []string{"a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e",
			"7b0d00e115025f52d2d94d4b13767f027668cc81cc97ba599509484d80c417cb",
			"90717b0e2f60025596a11fc62796abc9b0e2e2a0687b466d909b99b7a33583ce",
		}
		for i, str := range inputStrings {
			cypher, err := encrypt.EncryptSHA256(str)
			require.NoError(t, err)
			require.Equal(t, expectedCyphers[i], cypher)
		}
	})

	t.Run("EncryptionFuncsMD5", func(t *testing.T) {
		inputStrings := []string{"Hello World", "There is a string", "Wingardium Leviosa"}
		expectedCyphers := []string{"b10a8db164e0754105b7a99be72e3fe5",
			"b5e82ae5cde0fafb66cfc6fd7af9c163",
			"7dc6d4550f8421416d55eeceb1e46d2a",
		}
		for i, str := range inputStrings {
			cypher, err := encrypt.EncryptMD5(str)
			require.NoError(t, err)
			require.Equal(t, expectedCyphers[i], cypher)
		}
	})
}
