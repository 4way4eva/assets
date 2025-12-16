package sovereignledger

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestIncomeTick_ToENFT(t *testing.T) {
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

	enft := tick.ToENFT("BLEU-TEST-001")

	if enft.TokenID != "BLEU-TEST-001" {
		t.Errorf("Expected token ID BLEU-TEST-001, got %s", enft.TokenID)
	}

	if enft.SovereignData.CivilianYield != 100.0 {
		t.Errorf("Expected civilian yield 100.0, got %.2f", enft.SovereignData.CivilianYield)
	}

	if enft.SovereignData.MilitaryYield != 50.0 {
		t.Errorf("Expected military yield 50.0, got %.2f", enft.SovereignData.MilitaryYield)
	}

	if enft.SovereignData.CosmicYield != 75.0 {
		t.Errorf("Expected cosmic yield 75.0, got %.2f", enft.SovereignData.CosmicYield)
	}

	if enft.SovereignData.TotalYield != 225.0 {
		t.Errorf("Expected total yield 225.0, got %.2f", enft.SovereignData.TotalYield)
	}

	if !enft.SovereignData.PreAuthorized {
		t.Error("Expected ENFT to be pre-authorized")
	}

	if !enft.SovereignData.Pi4Acceleration {
		t.Error("Expected π₄ acceleration to be true")
	}

	if enft.SovereignData.LedgerHash == "" {
		t.Error("Expected ledger hash to be generated")
	}

	if len(enft.SovereignData.PhysicalAssetTags) == 0 {
		t.Error("Expected physical asset tags to be present")
	}
}

func TestENFT_ValidateENFT(t *testing.T) {
	// Test valid ENFT
	validENFT := &ENFTMetadata{
		TokenID: "BLEU-001",
		SovereignData: SovereignENFTData{
			TotalYield:        100.0,
			PreAuthorized:     true,
			LedgerHash:        "abc123",
			PhysicalAssetTags: []string{"TAG-001"},
		},
	}

	valid, errors := validENFT.ValidateENFT()
	if !valid {
		t.Errorf("Expected valid ENFT to pass validation, got errors: %v", errors)
	}

	// Test missing token ID
	invalidENFT := &ENFTMetadata{
		TokenID: "",
		SovereignData: SovereignENFTData{
			TotalYield:        100.0,
			PreAuthorized:     true,
			LedgerHash:        "abc123",
			PhysicalAssetTags: []string{"TAG-001"},
		},
	}

	valid, errors = invalidENFT.ValidateENFT()
	if valid {
		t.Error("Expected ENFT with missing token ID to fail validation")
	}
	if len(errors) == 0 {
		t.Error("Expected validation errors for invalid ENFT")
	}

	// Test negative yield
	invalidYieldENFT := &ENFTMetadata{
		TokenID: "BLEU-002",
		SovereignData: SovereignENFTData{
			TotalYield:        -10.0,
			PreAuthorized:     true,
			LedgerHash:        "abc123",
			PhysicalAssetTags: []string{"TAG-001"},
		},
	}

	valid, _ = invalidYieldENFT.ValidateENFT()
	if valid {
		t.Error("Expected ENFT with negative yield to fail validation")
	}

	// Test not pre-authorized
	notAuthorizedENFT := &ENFTMetadata{
		TokenID: "BLEU-003",
		SovereignData: SovereignENFTData{
			TotalYield:        100.0,
			PreAuthorized:     false,
			LedgerHash:        "abc123",
			PhysicalAssetTags: []string{"TAG-001"},
		},
	}

	valid, _ = notAuthorizedENFT.ValidateENFT()
	if valid {
		t.Error("Expected ENFT without pre-authorization to fail validation")
	}

	// Test missing physical asset tags
	noTagsENFT := &ENFTMetadata{
		TokenID: "BLEU-004",
		SovereignData: SovereignENFTData{
			TotalYield:        100.0,
			PreAuthorized:     true,
			LedgerHash:        "abc123",
			PhysicalAssetTags: []string{},
		},
	}

	valid, _ = noTagsENFT.ValidateENFT()
	if valid {
		t.Error("Expected ENFT without physical asset tags to fail validation")
	}
}

func TestENFT_ToJSON(t *testing.T) {
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

	enft := tick.ToENFT("BLEU-JSON-001")
	jsonStr, err := enft.ToJSON()

	if err != nil {
		t.Errorf("Error converting ENFT to JSON: %v", err)
	}

	if jsonStr == "" {
		t.Error("Expected non-empty JSON string")
	}

	// Verify it's valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		t.Errorf("Generated JSON is not valid: %v", err)
	}

	// Check key fields are present
	if _, ok := parsed["token_id"]; !ok {
		t.Error("Expected token_id in JSON")
	}
	if _, ok := parsed["sovereign_data"]; !ok {
		t.Error("Expected sovereign_data in JSON")
	}
}

func TestNewENFTCollection(t *testing.T) {
	collection := NewENFTCollection(
		"Test Collection",
		"Test Description",
	)

	if collection.Name != "Test Collection" {
		t.Errorf("Expected name 'Test Collection', got %s", collection.Name)
	}

	if collection.Description != "Test Description" {
		t.Errorf("Expected description 'Test Description', got %s", collection.Description)
	}

	if len(collection.ENFTs) != 0 {
		t.Errorf("Expected empty collection, got %d ENFTs", len(collection.ENFTs))
	}

	if collection.TotalYield != 0 {
		t.Errorf("Expected zero total yield, got %.2f", collection.TotalYield)
	}
}

func TestENFTCollection_AddENFT(t *testing.T) {
	collection := NewENFTCollection("Test", "Test")

	tick1 := IncomeTick{
		Timestamp:   time.Now(),
		TotalYield:  100.0,
		Pi4AccelApplied: true,
		SafeguardValid: true,
	}
	enft1 := tick1.ToENFT("BLEU-001")

	collection.AddENFT(enft1)

	if len(collection.ENFTs) != 1 {
		t.Errorf("Expected 1 ENFT in collection, got %d", len(collection.ENFTs))
	}

	if collection.TotalYield != 100.0 {
		t.Errorf("Expected total yield 100.0, got %.2f", collection.TotalYield)
	}

	// Add another ENFT
	tick2 := IncomeTick{
		Timestamp:   time.Now(),
		TotalYield:  200.0,
		Pi4AccelApplied: true,
		SafeguardValid: true,
	}
	enft2 := tick2.ToENFT("BLEU-002")

	collection.AddENFT(enft2)

	if len(collection.ENFTs) != 2 {
		t.Errorf("Expected 2 ENFTs in collection, got %d", len(collection.ENFTs))
	}

	if collection.TotalYield != 300.0 {
		t.Errorf("Expected total yield 300.0, got %.2f", collection.TotalYield)
	}
}

func TestENFTCollection_ToJSON(t *testing.T) {
	collection := NewENFTCollection("Test Collection", "Description")

	tick := IncomeTick{
		Timestamp:   time.Now(),
		TotalYield:  100.0,
		Pi4AccelApplied: true,
		SafeguardValid: true,
	}
	enft := tick.ToENFT("BLEU-001")
	collection.AddENFT(enft)

	jsonStr, err := collection.ToJSON()
	if err != nil {
		t.Errorf("Error converting collection to JSON: %v", err)
	}

	if jsonStr == "" {
		t.Error("Expected non-empty JSON string")
	}

	// Verify it's valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		t.Errorf("Generated JSON is not valid: %v", err)
	}

	// Check key fields
	if _, ok := parsed["name"]; !ok {
		t.Error("Expected name in JSON")
	}
	if _, ok := parsed["enfts"]; !ok {
		t.Error("Expected enfts in JSON")
	}
}

func TestENFTAttributes(t *testing.T) {
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

	enft := tick.ToENFT("BLEU-ATTR-001")

	if len(enft.Attributes) == 0 {
		t.Error("Expected ENFT to have attributes")
	}

	// Check for specific attributes
	hasYieldStreamType := false
	hasTotalYield := false

	for _, attr := range enft.Attributes {
		if attr.TraitType == "Yield Stream Type" {
			hasYieldStreamType = true
		}
		if attr.TraitType == "Total Yield (Million $)" {
			hasTotalYield = true
		}
	}

	if !hasYieldStreamType {
		t.Error("Expected Yield Stream Type attribute")
	}
	if !hasTotalYield {
		t.Error("Expected Total Yield attribute")
	}
}

func TestBluVaultMirrorTag(t *testing.T) {
	tick := IncomeTick{
		Timestamp:   time.Now(),
		TotalYield:  100.0,
		Pi4AccelApplied: true,
		SafeguardValid: true,
	}

	enft := tick.ToENFT("BLEU-MIRROR-001")

	if enft.BluVaultMirrorTag == "" {
		t.Error("Expected Blu-Vault mirror tag to be generated")
	}

	if !strings.HasPrefix(enft.BluVaultMirrorTag, "BLU-VAULT-") {
		t.Errorf("Expected mirror tag to start with BLU-VAULT-, got %s", enft.BluVaultMirrorTag)
	}
}

func TestENFTProperties(t *testing.T) {
	tick := IncomeTick{
		Timestamp:   time.Now(),
		TotalYield:  100.0,
		Pi4AccelApplied: true,
		SafeguardValid: true,
	}

	enft := tick.ToENFT("BLEU-PROPS-001")

	if enft.Properties == nil {
		t.Error("Expected properties to be initialized")
	}

	if val, ok := enft.Properties["irreversible_economic_ops"]; !ok || val != true {
		t.Error("Expected irreversible_economic_ops property to be true")
	}

	if val, ok := enft.Properties["compounding_safeguards"]; !ok || val != true {
		t.Error("Expected compounding_safeguards property to be true")
	}

	if val, ok := enft.Properties["blu_vault_mirrored"]; !ok || val != true {
		t.Error("Expected blu_vault_mirrored property to be true")
	}
}
