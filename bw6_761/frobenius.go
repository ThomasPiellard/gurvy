// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gurvy/internal/generators DO NOT EDIT

package bw6_761

import "github.com/consensys/gurvy/bw6_761/fp"

// Frobenius set z to Frobenius(x), return z
func (z *PairingResult) Frobenius(x *PairingResult) *PairingResult {
	// Adapted from https://eprint.iacr.org/2010/354.pdf (Section 3.2)

	z.B0.Conjugate(&x.B0)
	z.B1.Conjugate(&x.B1)
	z.B2.Conjugate(&x.B2)

	z.B1.MulByNonResidue1Power1(&z.B1)
	z.B2.MulByNonResidue1Power2(&z.B2)

	return z
}

// FrobeniusSquare set z to Frobenius^2(x), and return z
func (z *PairingResult) FrobeniusSquare(x *PairingResult) *PairingResult {
	// Adapted from https://eprint.iacr.org/2010/354.pdf (Section 3.2)

	z.Set(x)

	z.B1.MulByNonResidue2Power1(&z.B1)
	z.B2.MulByNonResidue2Power2(&z.B2)

	return z
}

// FrobeniusCube set z to Frobenius^3(x), return z
func (z *PairingResult) FrobeniusCube(x *PairingResult) *PairingResult {
	// Adapted from https://eprint.iacr.org/2010/354.pdf (Section 3.2)

	z.B0.Conjugate(&x.B0)
	z.B1.Conjugate(&x.B1)
	z.B2.Conjugate(&x.B2)

	z.B1.MulByNonResidue3Power1(&z.B1)
	z.B2.MulByNonResidue3Power2(&z.B2)

	return z
}

// MulByNonResidue1Power1 set z=x*(1,1)^(1*(p^1-1)/3) and return z
func (z *E2) MulByNonResidue1Power1(x *E2) *E2 {
	// (0,4922464560225523242118178942575080391082002530232324381063048548642823052024664478336818169867474395270858391911405337707247735739826664939444490469542109391530482826728203582549674992333383150446779312029624171857054392282775648)
	b := E2{
		A0: fp.Element{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		A1: fp.Element{
			9193734820520314185,
			15390913228415833887,
			5309822015742495676,
			5431732283202763350,
			17252325881282386417,
			298854800984767943,
			15252629665615712253,
			11476276919959978448,
			6617989123466214626,
			293279592164056124,
			3271178847573361778,
			76563709148138387,
		},
	}
	z.Mul(x, &b)
	return z
}

// MulByNonResidue1Power2 set z=x*(1,1)^(2*(p^1-1)/3) and return z
func (z *E2) MulByNonResidue1Power2(x *E2) *E2 {
	// 4922464560225523242118178942575080391082002530232324381063048548642823052024664478336818169867474395270858391911405337707247735739826664939444490469542109391530482826728203582549674992333383150446779312029624171857054392282775649
	b := fp.Element{
		10159193990637832851,
		5286779382647858051,
		15149190582698529379,
		10172307932521123666,
		7672315572788794062,
		4504265454330324035,
		8586997380578354686,
		5916374020980521403,
		9559933215456904989,
		10407926721244239843,
		3712625600415690514,
		17752318063289862,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResidue2Power1 set z=x*(1,1)^(1*(p^2-1)/3) and return z
func (z *E2) MulByNonResidue2Power1(x *E2) *E2 {
	// 1968985824090209297278610739700577151397666382303825728450741611566800370218827257750865013421937292370006175842381275743914023380727582819905021229583192207421122272650305267822868639090213645505120388400344940985710520836292650
	b := fp.Element{
		7467050525960156664,
		11327349735975181567,
		4886471689715601876,
		825788856423438757,
		532349992164519008,
		5190235139112556877,
		10134108925459365126,
		2188880696701890397,
		14832254987849135908,
		2933451070611009188,
		11385631952165834796,
		64130670718986244,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResidue2Power2 set z=x*(1,1)^(2*(p^2-1)/3) and return z
func (z *E2) MulByNonResidue2Power2(x *E2) *E2 {
	// 4922464560225523242118178942575080391082002530232324381063048548642823052024664478336818169867474395270858391911405337707247735739826664939444490469542109391530482826728203582549674992333383150446779312029624171857054392282775648
	b := fp.Element{
		9193734820520314185,
		15390913228415833887,
		5309822015742495676,
		5431732283202763350,
		17252325881282386417,
		298854800984767943,
		15252629665615712253,
		11476276919959978448,
		6617989123466214626,
		293279592164056124,
		3271178847573361778,
		76563709148138387,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}

// MulByNonResidue3Power1 set z=x*(1,1)^(1*(p^3-1)/3) and return z
func (z *E2) MulByNonResidue3Power1(x *E2) *E2 {
	// (0,1)
	b := E2{
		A0: fp.Element{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		A1: fp.Element{
			144959613005956565,
			6509995272855063783,
			11428286765660613342,
			15738672438262922740,
			17071399330169272331,
			13899911246788437003,
			12055474021000362245,
			2545351818702954755,
			8887388221587179644,
			5009280847225881135,
			15539704305423854047,
			23071597697427581,
		},
	}
	z.Mul(x, &b)
	return z
}

// MulByNonResidue3Power2 set z=x*(1,1)^(2*(p^3-1)/3) and return z
func (z *E2) MulByNonResidue3Power2(x *E2) *E2 {
	// 6891450384315732539396789682275657542479668912536150109513790160209623422243491736087683183289411687640864567753786613451161759120554247759349511699125301598951605099378508850372543631423596795951899700429969112842764913119068298
	b := fp.Element{
		17481284903592032950,
		10104133845767975835,
		8607375506753517913,
		13706168424391191299,
		9580010308493592354,
		14241333420363995524,
		6665632285037357566,
		5559902898979457045,
		15504799981718861253,
		8332096944629367896,
		18005297320867222879,
		58811391084848524,
	}
	z.A0.Mul(&x.A0, &b)
	z.A1.Mul(&x.A1, &b)
	return z
}
