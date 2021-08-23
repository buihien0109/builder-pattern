package phone

type Samsung struct {
	Phone CellPhone
}

func (s *Samsung) SetCamera() BuildProcess {
	s.Phone.Camera = true
	return s
}

func (s *Samsung) SetDualSim() BuildProcess {
	s.Phone.DualSim = true
	return s
}

func (s *Samsung) SetTorch() BuildProcess {
	s.Phone.Torch = false
	return s
}

func (s *Samsung) SetColorDisplay() BuildProcess {
	s.Phone.ColorDisplay = true
	return s
}

func (s *Samsung) GetCellPhone() CellPhone {
	return s.Phone
}
