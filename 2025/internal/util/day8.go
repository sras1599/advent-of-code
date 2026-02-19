package util

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

type junctionConnection struct {
	box1     junctionBox
	box2     junctionBox
	distance float64
}

type junctionConnections []junctionConnection

func newJunctionConnection(box1, box2 junctionBox) junctionConnection {
	distance := getStraightLineDistance(box1, box2)

	return junctionConnection{box1: box1, box2: box2, distance: distance}
}

func sortConnections(jc junctionConnections) {
	sort.Slice(jc, func(i, j int) bool {
		return jc[i].distance < jc[j].distance
	})
}

type junctionBox struct {
	x int
	y int
	z int
}

type junctionBoxes []junctionBox

type circuits []junctionBoxes

type junctiionCircuits struct {
	circuits
}

func newJunctionCircuits() junctiionCircuits {
	circuits := make(circuits, 0)

	return junctiionCircuits{circuits: circuits}
}

func (jc *junctiionCircuits) getCircuitForBox(box junctionBox) int {
	atIndex := -1

	for i, circuit := range jc.circuits {
		if slices.Contains(circuit, box) {
			return i
		}
	}

	return atIndex
}

func (jc *junctiionCircuits) mergeCircuits(i1, i2 int) {
	c1 := jc.circuits[i1]
	c2 := jc.circuits[i2]

	merged := append(c1, c2...)
	jc.circuits[i1] = merged
	jc.circuits = slices.Delete(jc.circuits, i2, i2+1)
}

func (jc *junctiionCircuits) addBoxToCircuit(cIndex int, box junctionBox) {
	jc.circuits[cIndex] = append(jc.circuits[cIndex], box)
}

func (jc *junctiionCircuits) createCircuit(box1 junctionBox, box2 junctionBox) {
	circuit := make(junctionBoxes, 0)
	circuit = append(circuit, box1, box2)

	jc.circuits = append(jc.circuits, circuit)
}

// a private function to calculate the euclidean straight line distance between 2 junction boxes
func getStraightLineDistance(box1, box2 junctionBox) float64 {
	dx := float64(box2.x - box1.x)
	dy := float64(box2.y - box1.y)
	dz := float64(box2.z - box1.z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// read the lines from the input file, and process each line into a junctionBox, returning a slice of junctionBoxes
func parseInputDay8() junctionBoxes {
	lines := ReadInputFile()
	boxes := make([]junctionBox, 0, len(lines))

	for _, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			continue // skip malformed lines
		}
		boxes = append(boxes, junctionBox{x: x, y: y, z: z})
	}

	return boxes
}

func getJunctionConnections() junctionConnections {
	connections := make([]junctionConnection, 0)
	boxes := parseInputDay8()

	for i, box1 := range boxes {
		for _, box2 := range boxes[i+1:] {
			connection := newJunctionConnection(box1, box2)
			connections = append(connections, connection)
		}
	}

	sortConnections(connections)
	return connections
}

func (jc *junctiionCircuits) printCircuits() {
	for _, c := range jc.circuits {
		fmt.Println(c)
	}
}

func (jc *junctiionCircuits) sort() {
	sort.Slice(jc.circuits, func(i, j int) bool {
		return len(jc.circuits[i]) > len(jc.circuits[j])
	})
}

func (jc *junctiionCircuits) getProduct(largest int) int {
	jc.sort()
	product := 1

	for i := range largest {
		if i >= len(jc.circuits) {
			return product
		}

		product *= len(jc.circuits[i])
	}

	return product
}

func (jc *junctiionCircuits) areAllConnectionsMerged(numBoxes int) bool {
	if len(jc.circuits[0]) == numBoxes {
		return true
	}

	return false
}

func processJunctionConnection(circuits *junctiionCircuits, conn *junctionConnection) {
	b1 := conn.box1
	b2 := conn.box2
	c1 := circuits.getCircuitForBox(b1)
	c2 := circuits.getCircuitForBox(b2)

	if c1 == -1 && c2 == -1 {
		circuits.createCircuit(b1, b2)
	} else if c1 == c2 {
		return
	} else if c1 != -1 && c2 == -1 {
		circuits.addBoxToCircuit(c1, b2)
	} else if c2 != -1 && c1 == -1 {
		circuits.addBoxToCircuit(c2, b1)
	} else if c1 != c2 {
		circuits.mergeCircuits(c1, c2)
	}
}

func ProcessJunctionConnections(numConns int, numLargest int) {
	connections := getJunctionConnections()
	circuits := newJunctionCircuits()

	for _, conn := range connections[:numConns] {
		processJunctionConnection(&circuits, &conn)
	}

	fmt.Printf("Product of %d largest connections: %d\n", numLargest, circuits.getProduct(numLargest))
}

func ProcessJunctionConnectionsTillMerge() {
	numBoxes := len(parseInputDay8())
	connections := getJunctionConnections()
	circuits := newJunctionCircuits()

	for _, conn := range connections {
		processJunctionConnection(&circuits, &conn)

		if circuits.areAllConnectionsMerged(numBoxes) {
			fmt.Printf("Product of X coordinates of last 2 junction boxes: %d\n", conn.box1.x*conn.box2.x)
			return
		}
	}
}
