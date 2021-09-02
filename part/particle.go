package part

import (
	"math"

	"github.com/bit101/blgo"
)

type point struct {
	x     float64
	y     float64
	dist  float64
	force float64
}

type Particle struct {
	X            float64
	Y            float64
	heading      float64
	speed        float64
	vx           float64
	vy           float64
	damp         float64
	gravity      float64
	radius       float64
	avoidPoints  []*point
	springPoints []*point
	gravPoints   []*point
}

func NewParticle(x, y float64) *Particle {
	return &Particle{
		X:       x,
		Y:       y,
		heading: 0,
		speed:   0,
		vx:      0,
		vy:      0,
		damp:    1.0,
		gravity: 0.0,
		radius:  0.5,
	}
}

func (p *Particle) AddAvoidPoint(x, y, dist, force float64) {
	p.avoidPoints = append(p.avoidPoints, &point{x, y, dist, force})
}

func (p *Particle) AddSpringPoint(x, y, force float64) {
	p.springPoints = append(p.springPoints, &point{x, y, 0, force})
}

func (p *Particle) AddGravPoint(x, y, force float64) {
	p.gravPoints = append(p.gravPoints, &point{x, y, 0, force})
}

func (p *Particle) SetRadius(r float64) {
	p.radius = r
}

func (p *Particle) Radius() float64 {
	return p.radius
}

func (p *Particle) SetDamp(damp float64) {
	p.damp = damp
}

func (p *Particle) Damp() float64 {
	return p.damp
}

func (p *Particle) SetGravity(gravity float64) {
	p.gravity = gravity
}

func (p *Particle) Gravity() float64 {
	return p.gravity
}

func (p *Particle) SetHeading(heading float64) {
	p.heading = heading
	p.updateVelocity()
}

func (p *Particle) Heading() float64 {
	return p.heading
}

func (p *Particle) AddHeading(amount float64) {
	p.heading += amount
	p.updateVelocity()
}

func (p *Particle) MultHeading(amount float64) {
	p.heading *= amount
	p.updateVelocity()
}

func (p *Particle) SetSpeed(speed float64) {
	p.speed = speed
	p.updateVelocity()
}

func (p *Particle) Speed() float64 {
	return p.speed
}

func (p *Particle) AddSpeed(amount float64) {
	p.speed += amount
	p.updateVelocity()
}

func (p *Particle) MultSpeed(amount float64) {
	p.speed *= amount
	p.updateVelocity()
}

func (p *Particle) SetVx(vx float64) {
	p.vx = vx
	p.updateSpeedAndHeading()
}

func (p *Particle) Vx() float64 {
	return p.vx
}

func (p *Particle) SetVy(vy float64) {
	p.vy = vy
	p.updateSpeedAndHeading()
}

func (p *Particle) Vy() float64 {
	return p.vy
}

func (p *Particle) SetVelocity(vx, vy float64) {
	p.vx = vx
	p.vy = vy
	p.updateSpeedAndHeading()
}

func (p *Particle) AddVelocity(amountX, amountY float64) {
	p.vx += amountX
	p.vy += amountY
	p.updateSpeedAndHeading()
}

func (p *Particle) updateSpeedAndHeading() {
	p.speed = math.Sqrt(p.vx*p.vx + p.vy*p.vy)
	p.heading = math.Atan2(p.vy, p.vx)
}

func (p *Particle) updateVelocity() {
	if p.speed == 0.0 {
		p.vx = 0.0
		p.vy = 0.0
	} else {
		p.vx = math.Cos(p.heading) * p.speed
		p.vy = math.Sin(p.heading) * p.speed
	}
}

func (p *Particle) Update() {
	p.step()
	p.AddVelocity(0, p.gravity)
	if p.damp != 1.0 {
		p.MultSpeed(p.damp)
	}
	for _, ap := range p.avoidPoints {
		p.avoidPoint(ap)
	}
	for _, sp := range p.springPoints {
		p.springTo(sp.x, sp.y, sp.force)
	}
	for _, gp := range p.gravPoints {
		p.gravTo(gp.x, gp.y, gp.force)
	}
}

func (p *Particle) avoidPoint(ap *point) {
	dx := p.X - ap.x
	dy := p.Y - ap.y
	dist := math.Hypot(dx, dy)
	if dist < ap.dist {
		tx := ap.x + dx/dist*ap.dist
		ty := ap.y + dy/dist*ap.dist
		p.springTo(tx, ty, ap.force)
	}
}

func (p *Particle) springTo(x, y, force float64) {
	p.AddVelocity((x-p.X)*force, (y-p.Y)*force)
}

func (p *Particle) gravTo(x, y, force float64) {
	dx := x - p.X
	dy := y - p.Y
	distSQ := dx*dx + dy*dy
	dist := math.Sqrt(distSQ)
	p.AddVelocity(dx/dist*force/distSQ, dy/dist*force/distSQ)
}

func (p *Particle) step() {
	p.X += p.vx
	p.Y += p.vy
}

func (p *Particle) Render(surface *blgo.Surface) {
	surface.FillCircle(p.X, p.Y, p.radius)
}
