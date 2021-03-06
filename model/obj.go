package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iCiaran/ray-tracing/maths"
)

type Model struct {
	V []*maths.Point3
	N []*maths.Vec3
	T []*maths.Point3
}

func NewModel() *Model {
	return &Model{make([]*maths.Point3, 0), make([]*maths.Vec3, 0), make([]*maths.Point3, 0)}
}

func check(err error, info interface{}) {
	if err != nil {
		fmt.Fprint(os.Stderr, err, info)
		os.Exit(1)
	}
}

func (m *Model) LoadObj(filepath string) *maths.HittableList {
	file, err := os.Open(filepath)
	check(err, filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l := maths.NewHittableList()
	texture := maths.NewTextureImage("../resources/capsule0.jpg")
	mat := maths.NewMetal(texture, 0.01)
	zv := maths.NewVec3(0.0, 0.0, 0.0)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		switch splitLine[0] {
		case "v":
			x, err := strconv.ParseFloat(splitLine[1], 64)
			check(err, splitLine)
			y, err := strconv.ParseFloat(splitLine[2], 64)
			check(err, splitLine)
			z, err := strconv.ParseFloat(splitLine[3], 64)
			check(err, splitLine)
			m.V = append(m.V, maths.NewVec3(x, y, z))
		case "vt":
			x, err := strconv.ParseFloat(splitLine[1], 64)
			check(err, splitLine)
			y, err := strconv.ParseFloat(splitLine[2], 64)
			check(err, splitLine)
			//z, err := strconv.ParseFloat(splitLine[3], 64)
			//check(err, splitLine)
			m.T = append(m.T, maths.NewVec3(x, y, 0.0))
		case "f":
			iSplit := strings.Split(splitLine[1], "/")
			jSplit := strings.Split(splitLine[2], "/")
			kSplit := strings.Split(splitLine[3], "/")
			i, err := strconv.Atoi(iSplit[0])
			check(err, iSplit)
			j, err := strconv.Atoi(jSplit[0])
			check(err, iSplit)
			k, err := strconv.Atoi(kSplit[0])
			check(err, iSplit)
			if len(iSplit) == 1 {
				e1 := maths.Sub(m.V[j-1], m.V[i-1])
				e2 := maths.Sub(m.V[k-1], m.V[i-1])
				n := maths.Cross(e1, e2).Normalise()
				l.Add(maths.NewTriangle(m.V[i-1], m.V[j-1], m.V[k-1], n, n, n, zv, zv, zv, mat, false))
			} else if len(iSplit) == 3 {
				in, err := strconv.Atoi(iSplit[2])
				check(err, iSplit)
				jn, err := strconv.Atoi(jSplit[2])
				check(err, iSplit)
				kn, err := strconv.Atoi(kSplit[2])
				check(err, iSplit)
				if len(iSplit[1]) > 0 {
					it, err := strconv.Atoi(iSplit[1])
					check(err, iSplit)
					jt, err := strconv.Atoi(jSplit[1])
					check(err, iSplit)
					kt, err := strconv.Atoi(kSplit[1])
					check(err, iSplit)
					l.Add(maths.NewTriangle(m.V[i-1], m.V[j-1], m.V[k-1], m.N[in-1], m.N[jn-1], m.N[kn-1], m.T[it-1], m.T[jt-1], m.T[kt-1], mat, true))
				} else {
					l.Add(maths.NewTriangle(m.V[i-1], m.V[j-1], m.V[k-1], m.N[in-1], m.N[jn-1], m.N[kn-1], zv, zv, zv, mat, true))
				}
			}
		case "vn":
			x, err := strconv.ParseFloat(splitLine[1], 64)
			check(err, splitLine)
			y, err := strconv.ParseFloat(splitLine[2], 64)
			check(err, splitLine)
			z, err := strconv.ParseFloat(splitLine[3], 64)
			check(err, splitLine)
			m.N = append(m.N, maths.NewVec3(x, y, z))
		}
	}

	check(scanner.Err(), nil)

	return l
}
