MyRpc interface {
	Increment(n uint64) (count uint64)
	Chunk(chunk []byte) (count uint64)
	Box(buffer Box) (count uint64)
}
// Point is the representation of a coordinate in (X,Y)
Point struct {
	x int32
	y int32
}

// Box is the representation of th ebox itself. It only contains two points, since they are essential in order to describe a square.
// 	      ur
//  .________o
//  |	     |
//  |	     |
//  |	     |
//  |	     |
//  o________|
// ll
Box struct {
	ll Point
	ur Point
}
