package p3d

type Point3DList struct {
	list []*Point3D
}

func NewPoint3DList() *Point3DList {
	var list []*Point3D
	return &Point3DList{list}
}

func (l *Point3DList) List() []*Point3D {
	return l.list
}

func (l *Point3DList) Get(index int) *Point3D {
	return l.list[index]
}

func (l *Point3DList) AddPoint(p *Point3D) {
	l.list = append(l.list, p)
}

func (l *Point3DList) RotateX(angle float64) {
	for _, p := range l.list {
		p.RotateX(angle)
	}
}

func (l *Point3DList) RotateY(angle float64) {
	for _, p := range l.list {
		p.RotateY(angle)
	}
}

func (l *Point3DList) RotateZ(angle float64) {
	for _, p := range l.list {
		p.RotateZ(angle)
	}
}

func (l *Point3DList) Rotate(x, y, z float64) {
	for _, p := range l.list {
		p.Rotate(x, y, z)
	}
}

func (l *Point3DList) Translate(x, y, z float64) {
	for _, p := range l.list {
		p.Translate(x, y, z)
	}
}

func (l *Point3DList) TranslateX(x float64) {
	for _, p := range l.list {
		p.TranslateX(x)
	}
}

func (l *Point3DList) TranslateY(y float64) {
	for _, p := range l.list {
		p.TranslateY(y)
	}
}

func (l *Point3DList) TranslateZ(z float64) {
	for _, p := range l.list {
		p.TranslateZ(z)
	}
}

func (l *Point3DList) Scale(s float64) {
	for _, p := range l.list {
		p.Scale(s)
	}
}

func (l *Point3DList) ScaleX(s float64) {
	for _, p := range l.list {
		p.ScaleX(s)
	}
}

func (l *Point3DList) ScaleY(s float64) {
	for _, p := range l.list {
		p.ScaleY(s)
	}
}

func (l *Point3DList) ScaleZ(s float64) {
	for _, p := range l.list {
		p.ScaleZ(s)
	}
}
