package charging

import "testing"

func TestVoltForge17(t *testing.T) {
	state := ApplyMitigationGate(MitigationGateState{Mitigated: true, ProtocolOK: true, CableOK: true, ThermalOK: false, PowerOK: true})
	if IsMitigationGateReady(state) || state.Status != "retest_required" {
		t.Fatalf("unsafe state was certified: %+v", state)
	}
}
