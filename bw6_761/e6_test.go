// Code generated by internal/tower DO NOT EDIT
package bw6_761

import (
	"reflect"
	"testing"
)

type E6TestPoint struct {
	in  [2]E6
	out [8]E6
}

var E6TestPoints []E6TestPoint

// TODO this method is the same everywhere. move it someplace central and call it "compare"
func E6compare(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatal("\nexpect:\t", want, "\ngot:\t", got)
	}
}

func E6check(t *testing.T, f func(*E6, *E6, *E6) *E6, m int) {

	if len(E6TestPoints) < 1 {
		t.Log("no tests to run")
	}

	for i := range E6TestPoints {
		var receiver E6
		var out *E6
		var inCopies [len(E6TestPoints[i].in)]E6

		for j := range inCopies {
			inCopies[j].Set(&E6TestPoints[i].in[j])
		}

		// receiver, return value both set to result
		out = f(&receiver, &inCopies[0], &inCopies[1])

		E6compare(t, receiver, E6TestPoints[i].out[m]) // receiver correct
		E6compare(t, *out, E6TestPoints[i].out[m])     // return value correct
		for j := range inCopies {
			E6compare(t, inCopies[j], E6TestPoints[i].in[j]) // inputs unchanged
		}

		// receiver == one of the inputs
		for j := range inCopies {
			out = f(&inCopies[j], &inCopies[0], &inCopies[1])

			E6compare(t, inCopies[j], E6TestPoints[i].out[m]) // receiver correct
			E6compare(t, *out, E6TestPoints[i].out[m])        // return value correct
			for k := range inCopies {
				if k == j {
					continue
				}
				E6compare(t, inCopies[k], E6TestPoints[i].in[k]) // other inputs unchanged
			}
			inCopies[j].Set(&E6TestPoints[i].in[j]) // reset input for next tests
		}
	}
}

//--------------------//
//     tests		  //
//--------------------//

func TestE6Add(t *testing.T) {
	E6check(t, (*E6).Add, 0)
}

func TestE6Sub(t *testing.T) {
	E6check(t, (*E6).Sub, 1)
}

func TestE6Mul(t *testing.T) {
	E6check(t, (*E6).Mul, 2)
}

func TestE6Square(t *testing.T) {
	E6check(t, (*E6).SquareBinary, 3)
}

func TestE6Inverse(t *testing.T) {
	E6check(t, (*E6).InverseBinary, 4)
}

func TestE6Frobenius(t *testing.T) {
	E6check(t, (*E6).FrobeniusBinary, 5)
}

func TestE6FrobeniusSquare(t *testing.T) {
	E6check(t, (*E6).FrobeniusSquareBinary, 6)
}

func TestE6FrobeniusCube(t *testing.T) {
	E6check(t, (*E6).FrobeniusCubeBinary, 7)
}

//--------------------//
//     benches		  //
//--------------------//

var E6BenchIn1, E6BenchIn2, E6BenchOut E6

func BenchmarkE6Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.Add(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6Sub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.Sub(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.Mul(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6Square(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.SquareBinary(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6Inverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.InverseBinary(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6Frobenius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.FrobeniusBinary(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6FrobeniusSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.FrobeniusSquareBinary(&E6BenchIn1, &E6BenchIn2)
	}
}

func BenchmarkE6FrobeniusCube(b *testing.B) {
	for i := 0; i < b.N; i++ {
		E6BenchOut.FrobeniusCubeBinary(&E6BenchIn1, &E6BenchIn2)
	}
}

//-------------------------------------//
// unary helpers for E6 methods
//-------------------------------------//

// SquareBinary a binary wrapper for Square
func (z *E6) SquareBinary(x, y *E6) *E6 {
	return z.Square(x)
}

// InverseBinary a binary wrapper for Inverse
func (z *E6) InverseBinary(x, y *E6) *E6 {
	return z.Inverse(x)
}

// FrobeniusBinary a binary wrapper for Frobenius
func (z *E6) FrobeniusBinary(x, y *E6) *E6 {
	return z.Frobenius(x)
}

// FrobeniusSquareBinary a binary wrapper for FrobeniusSquare
func (z *E6) FrobeniusSquareBinary(x, y *E6) *E6 {
	return z.FrobeniusSquare(x)
}

// FrobeniusCubeBinary a binary wrapper for FrobeniusCube
func (z *E6) FrobeniusCubeBinary(x, y *E6) *E6 {
	return z.FrobeniusCube(x)
}

//-------------------------------------//
// custom helpers for E6 methods
//-------------------------------------//

// MulByE2Binary a binary wrapper for MulByE2
func (z *E6) MulByE2Binary(x, y *E6) *E6 {
	return z.MulByE2(x, &y.B0)
}
