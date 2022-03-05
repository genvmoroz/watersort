package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlaskLifeCycle(t *testing.T) {
	t.Parallel()

	flask := NewFlask()

	require.True(t, flask.Empty()) // flask: [Non, Non, Non, Non], empty
	require.False(t, flask.Full())
	require.False(t, flask.Completed())
	require.True(t, flask.Put(Gray)) // flask: [Gray, Non, Non, Non]¬
	require.False(t, flask.Empty())
	require.False(t, flask.Full())
	require.True(t, flask.Put(Red)) // flask: [Gray, Red, Non, Non]
	require.False(t, flask.Empty())
	require.False(t, flask.Full())
	require.True(t, flask.Put(LightGreen)) // flask: [Gray, Red, LightGreen, Non]
	require.False(t, flask.Empty())
	require.False(t, flask.Full())
	require.True(t, flask.Put(DarkGreen)) // flask: [Gray, Red, LightGreen, DarkGreen]
	require.False(t, flask.Empty())
	require.True(t, flask.Full())            // flask: [Gray, Red, LightGreen, DarkGreen], full
	require.False(t, flask.Put(Yellow))      // flask: [Gray, Red, LightGreen, DarkGreen], unable to put more
	require.Equal(t, flask.Get(), DarkGreen) // flask: [Gray, Red, LightGreen, Non]
	require.False(t, flask.Full())
	require.False(t, flask.Empty())
	require.Equal(t, flask.Get(), LightGreen) // flask: [Gray, Red, Non, Non]
	require.False(t, flask.Full())
	require.False(t, flask.Empty())
	require.False(t, flask.Completed())
	require.Equal(t, flask.Get(), Red) // flask: [Gray, Non, Non, Non]
	require.False(t, flask.Full())
	require.False(t, flask.Empty())
	require.Equal(t, flask.Get(), Gray) // flask: [Non, Non, Non, Non]
	require.False(t, flask.Completed())
	require.False(t, flask.Full())
	require.True(t, flask.Empty())
	require.Equal(t, flask.Get(), Non) // flask: [Non, Non, Non, Non], unable to get more
	require.True(t, flask.Put(Red))    // flask: [Red, Non, Non, Non]
	require.True(t, flask.Put(Red))    // flask: [Red, Red, Non, Non]
	require.True(t, flask.Put(Red))    // flask: [Red, Red, Red, Non]
	require.True(t, flask.Put(Red))    // flask: [Red, Red, Red, Red]
	require.True(t, flask.Completed())
	require.True(t, flask.Full())
}

func BenchmarkFlaskPut(b *testing.B) {
	flash := NewFlask()

	for i := 0; i < b.N; i++ {
		_ = flash.Put(Purple)
	}
}

func BenchmarkFlaskGet(b *testing.B) {
	flash := NewFlask()
	flash.Put(Green)
	flash.Put(Blue)

	for i := 0; i < b.N; i++ {
		_ = flash.Get()
	}
}

func BenchmarkFlaskEmpty(b *testing.B) {
	flash := NewFlask()
	flash.Put(Pink)

	for i := 0; i < b.N; i++ {
		_ = flash.Empty()
	}
}

func BenchmarkFlaskFull(b *testing.B) {
	flash := NewFlask()
	flash.Put(Purple)
	flash.Put(Green)

	for i := 0; i < b.N; i++ {
		_ = flash.Full()
	}
}

func BenchmarkFlaskCompleted(b *testing.B) {
	flash := NewFlask()
	flash.Put(Purple)
	flash.Put(Purple)
	flash.Put(Purple)
	flash.Put(Purple)

	for i := 0; i < b.N; i++ {
		_ = flash.Completed()
	}
}
