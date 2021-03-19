package main

func BarycentricX(X, Y []float64, Z [][]float64, y, z float64) float64{
	// binary search for y upper bound
	j, k := 0, len(Y)
	for j < k {
		h := int(uint(j+k) >> 1)
		if Y[h] < y {
			j = h + 1
		} else {
			k = h
		}
	}
	if y < 10 || j == len(Y) {
		return 0
	}
	if y == 10 {
		j += 1
	}
	// linear search for x upper bound
	i := 0
	w := (y - Y[j-1]) / (Y[j] - Y[j-1])
	zOld := (1-w) * Z[j-1][0] + w * Z[j][0]
	lowerTriangle := true
	for k := 1; k < len(X); k++ {
		w := (y - Y[j-1]) / (Y[j] - Y[j-1])
		zBound := (1-w) * Z[j-1][k] + w * Z[j][k]
		triangleBound := (1-w)*Z[j-1][k-1] + w*Z[j][k]
		if zOld <= z && z <= zBound || zBound <= z && z <= zOld {
			i = k
		}
		if zOld <= z && z <= zBound && z <= triangleBound {
			lowerTriangle = false
			if z == 0 { break }
		}
		if zBound <= z && z <= zOld && triangleBound <= z {
			lowerTriangle = false
			if z == 0 { break }
		}
		zOld = zBound
	}
	if i == 0 || i == len(X) {
		return 0
	}
	if lowerTriangle {
		// lower triangle /_|
		A := (Z[j-1][i-1] - Z[j-1][i]) * (Y[j] - Y[j-1])
		w2 := ((Z[j-1][i] - Z[j][i]) * (y - Y[j-1]) + (Y[j] - Y[j-1]) * (z - Z[j-1][i])) / A
		return X[i] - w2 * (X[i] - X[i-1])
	} else {
		// upper triangle |/
		A := (Z[j][i] - Z[j][i-1]) * (Y[j] - Y[j-1])
		w1 := ((Z[j-1][i-1] - Z[j][i-1]) * (y - Y[j]) + (Y[j] - Y[j-1]) * (z - Z[j][i-1])) / A
		return X[i-1] + w1 * (X[i] - X[i-1])
	}
}

func BarycentricZ(X, Y []float64, Z [][]float64, x, y float64) float64{
	// binary search for x upper bound
	i, k := 0, len(X)
	for i < k {
		h := int(uint(i+k) >> 1)
		if X[h] < x {
			i = h + 1
		} else {
			k = h
		}
	}
	if x < 0 || i == len(X) {
		return 0
	}
	if x == 0 {
		i += 1
	}
	// binary search for y upper bound
	j, k := 0, len(Y)
	for j < k {
		h := int(uint(j+k) >> 1)
		if Y[h] < y {
			j = h + 1
		} else {
			k = h
		}
	}
	if y < 10 || j == len(Y) {
		return 0
	}
	if y == 10 {
		j += 1
	}
	w := (x - X[i-1]) / (X[i] - X[i-1])
	triangleBound := (1-w)*Y[j-1] + w*Y[j]
	if y < triangleBound {
		// lower triangle /_|
		w1 := (y - Y[j-1]) / (Y[j] - Y[j-1])
		w2 := (X[i] - x) / (X[i] - X[i-1])
		w3 := 1 - w1 - w2
		return Z[j][i] * w1 + Z[j-1][i-1] * w2 + Z[j-1][i] * w3
	} else {
		// upper triangle |/
		w1 := (x - X[i-1]) / (X[i] - X[i-1])
		w2 := (y - Y[j]) / (Y[j-1] - Y[j])
		w3 := 1 - w1 - w2
		return Z[j][i] * w1 + Z[j-1][i-1] * w2 + Z[j][i-1] * w3
	}
}