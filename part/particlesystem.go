package part

import "github.com/bit101/blgo"

type ParticleSystem struct {
	particles []*Particle
}

func NewParticleSystem() *ParticleSystem {
	return &ParticleSystem{}
}

func (s *ParticleSystem) AddParticle(x, y float64) *Particle {
	p := NewParticle(x, y)
	s.particles = append(s.particles, p)
	return p
}

func (s *ParticleSystem) ForEach(callback func(*Particle)) {
	for _, p := range s.particles {
		callback(p)
	}
}

func (s *ParticleSystem) Update() {
	for _, p := range s.particles {
		p.Update()
	}
}

func (s *ParticleSystem) Render(surface *blgo.Surface) {
	for _, p := range s.particles {
		p.Render(surface)
	}
}
