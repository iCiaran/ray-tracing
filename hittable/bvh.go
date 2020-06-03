package hittable

import (
	"fmt"
	"os"
	"sort"

	"github.com/iCiaran/ray-tracing/hittable/hitmatrecord"
	"github.com/iCiaran/ray-tracing/maths"
)

type BVHNode struct {
	left  Hittable
	right Hittable
	box   *AABB
}

func NewBVHNode(list *List, t0, t1 float64) *BVHNode {
	a := newBVHNode(list.objects, 0, len(list.objects), t0, t1)
	fmt.Fprintf(os.Stderr, "%v %v\n", a.box.min, a.box.max)
	return a
}

func newBVHNode(objects []Hittable, start, end int, t0, t1 float64) *BVHNode {
	axis := maths.RandomIntInRange(0, 2)
	objectSpan := end - start

	n := &BVHNode{nil, nil, NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))}
	if objectSpan == 1 {
		n.left = objects[start]
		n.right = objects[start]
	} else if objectSpan == 2 {
		if boxCompare(objects[start], objects[start+1], axis) {
			n.left = objects[start]
			n.right = objects[start+1]
		} else {
			n.left = objects[start+1]
			n.right = objects[start]
		}
	} else {
		// Sort objects[start:end] here
		sort.Slice(objects[start:end], func(i, j int) bool { return boxCompare(objects[i], objects[j], axis) })

		mid := start + objectSpan/2
		n.left = newBVHNode(objects, start, mid, t0, t1)
		n.right = newBVHNode(objects, mid, end, t0, t1)
	}

	boxLeft := NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))
	boxRight := NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))

	if !n.left.BoundingBox(t0, t1, boxLeft) || !n.right.BoundingBox(t0, t1, boxRight) {
		fmt.Fprint(os.Stderr, "No bounding box in bvh_node constructor.\n")
	}
	*n.box = *surroundingBox(boxLeft, boxRight)

	return n
}

func (n *BVHNode) Hit(r *maths.Ray, tMin, tMax float64, rec *hitmatrecord.HitMatRecord) bool {
	if !n.box.Hit(r, tMin, tMax, rec) {
		return false
	}

	hitLeft := n.left.Hit(r, tMin, tMax, rec)
	var hitRight bool

	if hitLeft {
		hitRight = n.right.Hit(r, tMin, rec.Rec.T, rec)
	} else {
		hitRight = n.right.Hit(r, tMin, tMax, rec)
	}

	return hitLeft || hitRight
}

func (n *BVHNode) BoundingBox(t0, t1 float64, outputBox *AABB) bool {
	*outputBox = *n.box
	return true
}

func boxCompare(a, b Hittable, axis int) bool {
	boxA := NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))
	boxB := NewAABB(maths.NewVec3(0.0, 0.0, 0.0), maths.NewVec3(0.0, 0.0, 0.0))

	if !a.BoundingBox(0.0, 0.0, boxA) || !b.BoundingBox(0.0, 0.0, boxB) {
		fmt.Fprint(os.Stderr, "No bounding box in bvh_node constructor.\n")
	}

	return boxA.min[axis] < boxB.min[axis]
}
