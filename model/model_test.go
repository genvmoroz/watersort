package model

import (
	"testing"
)

func TestFlaskLifeCycle(t *testing.T) {
	t.Parallel()

	flask := NewFlask()

	if got := flask.Empty(); got != true { // flask: [Non, Non, Non, Non], empty
		t.Errorf("Empty() = %v, want %v", got, true)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Completed(); got != false {
		t.Errorf("Completed() = %v, want %v", got, false)
	}

	if got := flask.Put(Gray); got != true { // flask: [Gray, Non, Non, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Put(Red); got != true { // flask: [Gray, Red, Non, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Put(LightGreen); got != true { // flask: [Gray, Red, LightGreen, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Put(DarkGreen); got != true { // flask: [Gray, Red, LightGreen, DarkGreen]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Full(); got != true { // flask: [Gray, Red, LightGreen, DarkGreen], full
		t.Errorf("Full() = %v, want %v", got, true)
	}

	if got := flask.Put(Yellow); got != false { // flask: [Gray, Red, LightGreen, DarkGreen], unable to put more
		t.Errorf("Put() = %v, want %v", got, false)
	}

	if got := flask.Get(); got != DarkGreen { // flask: [Gray, Red, LightGreen, Non]
		t.Errorf("Get() = %v, want %v", got, DarkGreen)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Get(); got != LightGreen { // flask: [Gray, Red, Non, Non]
		t.Errorf("Get() = %v, want %v", got, LightGreen)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Completed(); got != false {
		t.Errorf("Completed() = %v, want %v", got, false)
	}

	if got := flask.Get(); got != Red { // flask: [Gray, Non, Non, Non]
		t.Errorf("Get() = %v, want %v", got, Red)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	if got := flask.Get(); got != Gray { // flask: [Non, Non, Non, Non]
		t.Errorf("Get() = %v, want %v", got, Gray)
	}

	if got := flask.Completed(); got != false {
		t.Errorf("Completed() = %v, want %v", got, false)
	}

	if got := flask.Full(); got != false {
		t.Errorf("Full() = %v, want %v", got, false)
	}

	if got := flask.Empty(); got != true { // flask: [Non, Non, Non, Non], empty
		t.Errorf("Empty() = %v, want %v", got, true)
	}

	if got := flask.Get(); got != Non { // flask: [Non, Non, Non, Non], unable to get more
		t.Errorf("Get() = %v, want %v", got, Non)
	}

	if got := flask.Put(Red); got != true { // flask: [Red, Non, Non, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Put(Red); got != true { // flask: [Red, Red, Non, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Put(Red); got != true { // flask: [Red, Red, Red, Non]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Put(Red); got != true { // flask: [Red, Red, Red, Red]
		t.Errorf("Put() = %v, want %v", got, true)
	}

	if got := flask.Completed(); got != true { // flask: [Red, Red, Red, Red], full and completed
		t.Errorf("Completed() = %v, want %v", got, true)
	}
	if got := flask.Full(); got != true {
		t.Errorf("Full() = %v, want %v", got, true)
	}
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
