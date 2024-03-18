package enemy

func (e *Enemy) Update() {
	e.Y += e.Speed
}
