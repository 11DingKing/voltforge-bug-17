package charging

func ApplyMitigationGate(state MitigationGateState) MitigationGateState {
	if state.CanCertify() {
		state.Status = "certified"
		return state
	}
	state.Status = "retest_required"
	return state
}
func IsMitigationGateReady(state MitigationGateState) bool { return state.Status == "certified" }
