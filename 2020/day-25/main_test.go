package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []byte(`5764801
17807724`)

func TestPart1(t *testing.T) {
	cardKeys := keyHolder{
		public: 5764801,
	}

	doorKeys := keyHolder{
		public: 17807724,
	}

	require.Equal(t, 8, cardKeys.calculateLoopSize())
	require.Equal(t, 11, doorKeys.calculateLoopSize())

	require.Equal(t, 14897079, cardKeys.calculateEncryptionKey(doorKeys.public))
	require.Equal(t, 14897079, doorKeys.calculateEncryptionKey(cardKeys.public))
}

func TestPart1Full(t *testing.T) {
	input := testInput
	encryptionKey := calculatePart1EncryptionKey(input)
	require.Equal(t, 14897079, encryptionKey)
}
