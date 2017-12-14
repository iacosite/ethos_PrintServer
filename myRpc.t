MyRpc interface {
	Increment(n uint64) (count uint64)
	Box(buffer Box) (count uint64)
	FileTransfer(buffer Param) (count uint64)
	TestVars(tmp Par) (count uint64)
}
Param struct {
	buff []uint8
	t uint32
	name string
}
Par struct {
	a uint32
	b uint64
	c string
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
