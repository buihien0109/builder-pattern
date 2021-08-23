package phone

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (d *Director) BuildPhone() CellPhone {
	d.builder.SetCamera().SetDualSim().SetTorch().SetColorDisplay()
	return d.builder.GetCellPhone()
}
