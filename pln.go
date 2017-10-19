package pln

import (
	"simplex/seg"
	"simplex/rng"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
)

//Polyline Type
type Polyline struct {
	Coordinates []*geom.Point
	Geometry    *geom.LineString
}

//construct new polyline
func New(coordinates []*geom.Point) *Polyline {
	var n = len(coordinates)
	return &Polyline{
		Coordinates: coordinates[:n:n],
		Geometry:    geom.NewLineString(coordinates),
	}
}

//Bounding box of polyline
func (ln *Polyline) BBox() *mbr.MBR {
	return ln.Geometry.BBox()
}

//polyline
func (ln *Polyline) Polyline() *Polyline {
	return ln
}

//Coordinates at index i
func (ln *Polyline) Coordinate(i int) *geom.Point {
	return ln.Coordinates[i]
}

//Polyline segments
func (ln *Polyline) Segments() []*seg.Seg {
	var i, j int
	var lst = make([]*seg.Seg, 0)
	for i = 0; i < ln.Len()-1; i++ {
		j = i + 1
		lst = append(lst, seg.NewSeg(ln.Coordinates[i], ln.Coordinates[j], i, j))
	}
	return lst
}

//Range of entire polyline
func (ln *Polyline) Range() *rng.Range {
	return rng.NewRange(0, ln.Len()-1)
}

//Segment given range
func (ln *Polyline) Segment(rng *rng.Range) *seg.Seg {
	var i, j = rng.I(), rng.J()
	return seg.NewSeg(ln.Coordinates[i], ln.Coordinates[j], i, j)
}

//generates sub polyline from generator indices
func (ln *Polyline) SubPolyline(rng *rng.Range) *Polyline {
	return New(ln.Coordinates[rng.I():rng.J()+1])
}

//Length of coordinates in polyline
func (ln *Polyline) Len() int {
	return len(ln.Coordinates)
}
