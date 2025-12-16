package sovereignledger

import (
	"testing"
	"time"
)

func TestNewBleuSovereignLedger(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	if ledger.CivilianStream.BaseYieldPerSec != 13.6 {
		t.Errorf("Expected Civilian yield 13.6, got %.2f", ledger.CivilianStream.BaseYieldPerSec)
	}

	if ledger.MilitaryStream.BaseYieldPerSec != 6.1 {
		t.Errorf("Expected Military yield 6.1, got %.2f", ledger.MilitaryStream.BaseYieldPerSec)
	}

	if ledger.CosmicStream.BaseYieldPerSec != 9.2 {
		t.Errorf("Expected Cosmic yield 9.2, got %.2f", ledger.CosmicStream.BaseYieldPerSec)
	}

	expectedTotal := 13.6 + 6.1 + 9.2
	if ledger.GetTotalYieldPerSecond() != expectedTotal {
		t.Errorf("Expected total yield %.1f, got %.1f", expectedTotal, ledger.GetTotalYieldPerSecond())
	}
}

func TestPi4Constant(t *testing.T) {
	if Pi4Constant != 4.0 {
		t.Errorf("Expected π₄ constant to be 4.0, got %.1f", Pi4Constant)
	}
}

func TestCalculateYieldWithPi4(t *testing.T) {
	stream := YieldStream{
		Name:            "Test Stream",
		BaseYieldPerSec: 10.0,
		AccelFactor:     Pi4Constant,
	}

	// Test with 1 second
	yield := stream.CalculateYieldWithPi4(1.0)
	if yield <= 0 {
		t.Errorf("Expected positive yield, got %.2f", yield)
	}

	// Test that longer duration produces higher yield
	yield1 := stream.CalculateYieldWithPi4(10.0)
	yield2 := stream.CalculateYieldWithPi4(20.0)
	if yield2 <= yield1 {
		t.Errorf("Expected yield to increase with duration: %.2f vs %.2f", yield1, yield2)
	}
}

func TestCompoundingSafeguard(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	// Test with safe yield
	safeYield := 1000.0
	result, valid := ledger.CompoundingSafeguard(safeYield)
	if !valid {
		t.Error("Expected safe yield to pass validation")
	}
	if result != safeYield {
		t.Errorf("Expected safe yield to remain unchanged: %.2f vs %.2f", safeYield, result)
	}

	// Test with excessive yield
	excessiveYield := 1e15
	result, valid = ledger.CompoundingSafeguard(excessiveYield)
	if valid {
		t.Error("Expected excessive yield to fail validation")
	}
	if result >= excessiveYield {
		t.Error("Expected excessive yield to be capped")
	}
}

func TestBluVaultMirror(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	// Test valid mirroring
	if !ledger.BluVaultMirror("digital123", "physical456") {
		t.Error("Expected valid entries to pass mirroring")
	}

	// Test invalid mirroring (empty entries)
	if ledger.BluVaultMirror("", "physical456") {
		t.Error("Expected empty digital entry to fail mirroring")
	}
	if ledger.BluVaultMirror("digital123", "") {
		t.Error("Expected empty physical entry to fail mirroring")
	}
}

func TestGenerateIncomeTick(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	// Wait a moment for some yield accumulation
	time.Sleep(100 * time.Millisecond)

	tick := ledger.GenerateIncomeTick()

	if tick.CivilianYield <= 0 {
		t.Error("Expected positive civilian yield")
	}
	if tick.MilitaryYield <= 0 {
		t.Error("Expected positive military yield")
	}
	if tick.CosmicYield <= 0 {
		t.Error("Expected positive cosmic yield")
	}
	if tick.TotalYield <= 0 {
		t.Error("Expected positive total yield")
	}

	if !tick.Pi4AccelApplied {
		t.Error("Expected π₄ acceleration to be applied")
	}

	if tick.DivineTimestamp == "" {
		t.Error("Expected divine timestamp to be generated")
	}
	if tick.ReadableTimestamp == "" {
		t.Error("Expected readable timestamp to be generated")
	}
}

func TestFormatDivineTimestamp(t *testing.T) {
	now := time.Date(2025, 12, 16, 15, 30, 45, 0, time.UTC)
	timestamp := FormatDivineTimestamp(now)

	if timestamp == "" {
		t.Error("Expected non-empty timestamp")
	}

	// Check format contains expected components
	if len(timestamp) < 20 {
		t.Errorf("Expected longer timestamp format, got: %s", timestamp)
	}

	// Should contain BLEU prefix
	if timestamp[:4] != "BLEU" {
		t.Errorf("Expected timestamp to start with BLEU, got: %s", timestamp)
	}
}

func TestIncomeTick_ToDeclarativeScroll(t *testing.T) {
	tick := IncomeTick{
		Timestamp:         time.Now(),
		CivilianYield:     100.0,
		MilitaryYield:     50.0,
		CosmicYield:       75.0,
		TotalYield:        225.0,
		Pi4AccelApplied:   true,
		SafeguardValid:    true,
		DivineTimestamp:   "BLEU-2025-12-16-15:30:45-UNIX1765897845",
		ReadableTimestamp: "2025-12-16T15:30:45Z",
	}

	scroll := tick.ToDeclarativeScroll()

	if scroll == "" {
		t.Error("Expected non-empty scroll")
	}

	// Check scroll contains key information
	if len(scroll) < 100 {
		t.Error("Expected longer scroll format")
	}
}

func TestYieldStreamLayers(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	// Test Civilian layers
	expectedCivilianLayers := []string{
		"Retail",
		"Education",
		"EV0L Wearables",
		"Real Estate (ES0IL)",
		"Hospitality",
	}
	if len(ledger.CivilianStream.Layers) != len(expectedCivilianLayers) {
		t.Errorf("Expected %d civilian layers, got %d",
			len(expectedCivilianLayers), len(ledger.CivilianStream.Layers))
	}

	// Test Military layers
	expectedMilitaryLayers := []string{
		"Weapons Targeting",
		"Orbital Defense Grids",
		"AI Maritime Logistics",
	}
	if len(ledger.MilitaryStream.Layers) != len(expectedMilitaryLayers) {
		t.Errorf("Expected %d military layers, got %d",
			len(expectedMilitaryLayers), len(ledger.MilitaryStream.Layers))
	}

	// Test Cosmic layers
	expectedCosmicLayers := []string{
		"Quantum Portal Technology",
		"Inter-Realm Logistics",
		"Dimensional Treasury Protocols",
	}
	if len(ledger.CosmicStream.Layers) != len(expectedCosmicLayers) {
		t.Errorf("Expected %d cosmic layers, got %d",
			len(expectedCosmicLayers), len(ledger.CosmicStream.Layers))
	}
}

func TestYieldStreamAccelFactor(t *testing.T) {
	ledger := NewBleuSovereignLedger()

	if ledger.CivilianStream.AccelFactor != Pi4Constant {
		t.Errorf("Expected civilian accel factor %.1f, got %.1f",
			Pi4Constant, ledger.CivilianStream.AccelFactor)
	}

	if ledger.MilitaryStream.AccelFactor != Pi4Constant {
		t.Errorf("Expected military accel factor %.1f, got %.1f",
			Pi4Constant, ledger.MilitaryStream.AccelFactor)
	}

	if ledger.CosmicStream.AccelFactor != Pi4Constant {
		t.Errorf("Expected cosmic accel factor %.1f, got %.1f",
			Pi4Constant, ledger.CosmicStream.AccelFactor)
	}
}
