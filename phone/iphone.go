package phone

type Iphone struct {
	Phone CellPhone
}

func (ip *Iphone) SetCamera() BuildProcess {
	ip.Phone.Camera = false
	return ip
}

func (ip *Iphone) SetDualSim() BuildProcess {
	ip.Phone.DualSim = false
	return ip
}

func (ip *Iphone) SetTorch() BuildProcess {
	ip.Phone.Torch = true
	return ip
}

func (ip *Iphone) SetColorDisplay() BuildProcess {
	ip.Phone.ColorDisplay = false
	return ip
}

func (ip *Iphone) GetCellPhone() CellPhone {
	return ip.Phone
}
