package charging

import "errors"

var ErrMitigationGateState = errors.New("invalid mitigationgate state")

type MitigationGateState struct {
	Mitigated, ProtocolOK, CableOK, ThermalOK, PowerOK bool
	Status                                             string
}

func (s MitigationGateState) CanCertify() bool {
	return s.Mitigated && s.ProtocolOK && s.CableOK && s.ThermalOK && s.PowerOK
}
