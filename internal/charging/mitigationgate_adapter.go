package charging

func ApplyMitigationGate(state MitigationGateState) MitigationGateState {
	if state.Mitigated && state.ProtocolOK && state.CableOK && state.PowerOK {
		state.Status = "certified"
	} else {
		state.Status = "retest_required"
	}
	return state
}
func IsMitigationGateReady(state MitigationGateState) bool { return state.Status == "certified" }
