// Package sjtsk2gps provides utility function which converts
// czech S-JTSK coordinates (used in RUIAN)to GPS representation.
package sjtsk2gps

import "math"

// Convert accept czech S-JTSK coordinates and convert them to GPS
// it is rewritten javascript code from http://martin.hinner.info/geo/
func Convert(X float64, Y float64, H float64) (float64, float64, float64) {

	if X < 0 && Y < 0 {
		X = -X
		Y = -Y
	}

	if Y > X {
		X, Y = Y, X
	}

	a := float64(6377397.15508)
	e := float64(0.081696831215303)
	n := float64(0.97992470462083)
	konstURo := float64(12310230.12797036)
	sinUQ := float64(0.863499969506341)
	cosUQ := float64(0.504348889819882)
	sinVQ := float64(0.420215144586493)
	cosVQ := float64(0.907424504992097)
	alfa := float64(1.000597498371542)
	k := float64(1.003419163966575)

	ro := math.Sqrt(X*X + Y*Y)
	epsilon := 2 * math.Atan(Y/(ro+X))

	D := epsilon / n

	S := 2*math.Atan(math.Exp(1/n*math.Log(konstURo/ro))) - math.Pi/2

	sinS := math.Sin(S)
	cosS := math.Cos(S)
	sinU := sinUQ*sinS - cosUQ*cosS*math.Cos(D)
	cosU := math.Sqrt(1 - sinU*sinU)
	sinDV := math.Sin(D) * cosS / cosU
	cosDV := math.Sqrt(1 - sinDV*sinDV)
	sinV := sinVQ*cosDV - cosVQ*sinDV
	cosV := cosVQ*cosDV + sinVQ*sinDV
	Ljtsk := 2 * math.Atan(sinV/(1+cosV)) / alfa
	t := math.Exp(2 / alfa * math.Log((1+sinU)/cosU/k))
	pom := (t - float64(1)) / (t + float64(1))

	for {
		sinB := pom

		pom = t * math.Exp(e*math.Log((1+e*sinB)/(1-e*sinB)))
		pom = (pom - 1) / (pom + 1)

		if !(math.Abs(pom-sinB) > 1e-15) {
			break
		}
	}

	Bjtsk := math.Atan(pom / math.Sqrt(1-pom*pom))

	a = float64(6377397.15508)
	f1 := float64(299.152812853)
	e2 := float64(1) - (float64(1)-float64(1)/f1)*(float64(1)-float64(1)/f1)
	ro = a / math.Sqrt(1-e2*math.Sin(Bjtsk)*math.Sin(Bjtsk))
	x := (ro + H) * math.Cos(Bjtsk) * math.Cos(Ljtsk)
	y := (ro + H) * math.Cos(Bjtsk) * math.Sin(Ljtsk)
	z := ((float64(1)-e2)*ro + H) * math.Sin(Bjtsk)

	dx := 570.69
	dy := 85.69
	dz := 462.84
	wz := -5.2611 / 3600 * math.Pi / 180
	wy := -1.58676 / 3600 * math.Pi / 180
	wx := -4.99821 / 3600 * math.Pi / 180
	m := 3.543e-6
	xn := dx + (float64(1)+m)*(x+wz*y-wy*z)
	yn := dy + (float64(1)+m)*(-wz*x+y+wx*z)
	zn := dz + (float64(1)+m)*(wy*x-wx*y+z)

	a = float64(6378137.0)
	f1 = float64(298.257223563)
	aB := f1 / (f1 - 1)
	p := math.Sqrt(xn*xn + yn*yn)
	e2 = float64(1) - (float64(1)-float64(1)/f1)*(float64(1)-float64(1)/f1)
	theta := math.Atan(zn * aB / p)
	st := math.Sin(theta)
	ct := math.Cos(theta)
	t = (zn + e2*aB*a*st*st*st) / (p - e2*a*ct*ct*ct)
	B := math.Atan(t)
	L := 2 * math.Atan(yn/(p+xn))
	hOut := math.Sqrt(float64(1)+t*t) * (p - a/math.Sqrt(float64(1)+(float64(1)-e2)*t*t))

	lat := B * 180 / math.Pi
	long := L * 180 / math.Pi

	height := math.Floor(hOut*100) / 100

	return lat, long, height
}
