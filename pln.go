package pln

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/rng"
)

//Polyline Type
type Polyline struct {
	*geom.LineString
}

//construct new polyline
func New(coordinates geom.Coords) *Polyline {
	return &Polyline{geom.NewLineString(coordinates)}
}

//Polyline segments
func (ln *Polyline) Segments() []*geom.Segment {
	var i int
	var n = ln.Len() - 1
	var lst = make([]*geom.Segment, 0, n)
	for i = 0; i < n; i++ {
		lst = append(lst, geom.NewSegment(ln.Coordinates, i, i+1))
	}
	return lst
}

//Range of entire polyline
func (ln *Polyline) Range() rng.Rng {
	return rng.Range(0, ln.Len()-1)
}

//Segment given range
func (ln *Polyline) Segment(rng *rng.Rng) *geom.Segment {
	return geom.NewSegment(ln.Coordinates, rng.I, rng.J)
}

//generates sub polyline from generator indices
func (ln *Polyline) SubPolyline(rng rng.Rng) *Polyline {
	return New(ln.SubCoordinates(rng))
}

//generates sub polyline from generator indices
func (ln *Polyline) SubCoordinates(rng rng.Rng) geom.Coords {
	var coords = ln.Coordinates
	coords.Idxs = make([]int, 0, rng.J-rng.I+1)
	for i := rng.I; i <= rng.J; i++ {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

//Length of coordinates in polyline
func (ln *Polyline) Len() int {
	return ln.Coordinates.Len()
}
