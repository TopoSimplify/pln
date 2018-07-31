package pln

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/seg"
	"github.com/TopoSimplify/rng"
)

//Polyline Type
type Polyline struct {
	Coordinates []geom.Point
	Geometry    *geom.LineString
}

//construct new polyline
func New(coordinates []geom.Point) *Polyline {
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

//Polyline segments
func (ln *Polyline) Segments() []*seg.Seg {
	var i, j int
	var lst = make([]*seg.Seg, 0)
	for i = 0; i < ln.Len()-1; i++ {
		j = i + 1
		lst = append(lst, seg.NewSeg(&ln.Coordinates[i], &ln.Coordinates[j], i, j))
	}
	return lst
}

//Range of entire polyline
func (ln *Polyline) Range() rng.Rng {
	return rng.Range(0, ln.Len()-1)
}

//Segment given range
func (ln *Polyline) Segment(rng *rng.Rng) *seg.Seg {
	var i, j = rng.I, rng.J
	return seg.NewSeg(&ln.Coordinates[i], &ln.Coordinates[j], i, j)
}

//generates sub polyline from generator indices
func (ln *Polyline) SubPolyline(rng rng.Rng) *Polyline {
	return New(ln.SubCoordinates(rng))
}

//generates sub polyline from generator indices
func (ln *Polyline) SubCoordinates(rng rng.Rng) []geom.Point {
	var i, n = rng.I, rng.J+1
	return ln.Coordinates[i:n:n]
}

//Length of coordinates in polyline
func (ln *Polyline) Len() int {
	return len(ln.Coordinates)
}