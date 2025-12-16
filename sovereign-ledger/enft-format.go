package sovereignledger

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// ENFTMetadata represents the Enhanced Non-Fungible Token metadata structure
type ENFTMetadata struct {
	TokenID           string                 `json:"token_id"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Image             string                 `json:"image"`
	ExternalURL       string                 `json:"external_url"`
	Attributes        []ENFTAttribute        `json:"attributes"`
	SovereignData     SovereignENFTData      `json:"sovereign_data"`
	BluVaultMirrorTag string                 `json:"blu_vault_mirror_tag"`
	CreatedAt         time.Time              `json:"created_at"`
	Properties        map[string]interface{} `json:"properties,omitempty"`
}

// ENFTAttribute represents metadata attributes for the ENFT
type ENFTAttribute struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type,omitempty"`
}

// SovereignENFTData contains BLEU Sovereign Ledger specific data
type SovereignENFTData struct {
	YieldStream       string    `json:"yield_stream"`
	CivilianYield     float64   `json:"civilian_yield"`
	MilitaryYield     float64   `json:"military_yield"`
	CosmicYield       float64   `json:"cosmic_yield"`
	TotalYield        float64   `json:"total_yield"`
	Pi4Acceleration   bool      `json:"pi4_acceleration"`
	QuadLockApplied   bool      `json:"quad_lock_applied"`
	DivineTimestamp   string    `json:"divine_timestamp"`
	PhysicalAssetTags []string  `json:"physical_asset_tags"`
	PreAuthorized     bool      `json:"pre_authorized"`
	LedgerHash        string    `json:"ledger_hash"`
	MintedAt          time.Time `json:"minted_at"`
}

// ToENFT converts an IncomeTick to ENFT format
func (it *IncomeTick) ToENFT(tokenID string) *ENFTMetadata {
	now := time.Now()
	ledgerHash := generateLedgerHash(it)

	enft := &ENFTMetadata{
		TokenID:     tokenID,
		Name:        fmt.Sprintf("BLEU Sovereign Income Tick #%s", tokenID),
		Description: "Pre-authorized ledger compoundment representing triple-sphere yield streams with π₄ acceleration",
		Image:       fmt.Sprintf("ipfs://bleu-sovereign-ledger/%s.png", tokenID),
		ExternalURL: fmt.Sprintf("https://bleu-sovereign.io/enft/%s", tokenID),
		Attributes: []ENFTAttribute{
			{
				TraitType:   "Yield Stream Type",
				Value:       "Triple-Sphere (Civilian + Military + Cosmic)",
				DisplayType: "string",
			},
			{
				TraitType:   "Civilian Yield (Million $)",
				Value:       fmt.Sprintf("%.2f", it.CivilianYield),
				DisplayType: "number",
			},
			{
				TraitType:   "Military Yield (Million $)",
				Value:       fmt.Sprintf("%.2f", it.MilitaryYield),
				DisplayType: "number",
			},
			{
				TraitType:   "Cosmic Yield (Million $)",
				Value:       fmt.Sprintf("%.2f", it.CosmicYield),
				DisplayType: "number",
			},
			{
				TraitType:   "Total Yield (Million $)",
				Value:       fmt.Sprintf("%.2f", it.TotalYield),
				DisplayType: "number",
			},
			{
				TraitType: "π₄ Acceleration Applied",
				Value:     it.Pi4AccelApplied,
			},
			{
				TraitType: "Safeguard Validation",
				Value:     it.SafeguardValid,
			},
		},
		SovereignData: SovereignENFTData{
			YieldStream:       "Triple-Sphere",
			CivilianYield:     it.CivilianYield,
			MilitaryYield:     it.MilitaryYield,
			CosmicYield:       it.CosmicYield,
			TotalYield:        it.TotalYield,
			Pi4Acceleration:   it.Pi4AccelApplied,
			QuadLockApplied:   it.SafeguardValid,
			DivineTimestamp:   it.DivineTimestamp,
			PhysicalAssetTags: []string{"BLU-VAULT-ASSET-001", "ES0IL-TAG-001"},
			PreAuthorized:     true,
			LedgerHash:        ledgerHash,
			MintedAt:          now,
		},
		BluVaultMirrorTag: fmt.Sprintf("BLU-VAULT-%s", ledgerHash[:16]),
		CreatedAt:         now,
		Properties: map[string]interface{}{
			"irreversible_economic_ops": true,
			"compounding_safeguards":    true,
			"blu_vault_mirrored":        true,
		},
	}

	return enft
}

// generateLedgerHash creates a unique hash for the income tick ledger entry
func generateLedgerHash(it *IncomeTick) string {
	data := fmt.Sprintf("%v-%f-%f-%f-%f",
		it.Timestamp.Unix(),
		it.CivilianYield,
		it.MilitaryYield,
		it.CosmicYield,
		it.TotalYield,
	)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// ToJSON converts ENFT metadata to JSON format
func (enft *ENFTMetadata) ToJSON() (string, error) {
	jsonData, err := json.MarshalIndent(enft, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// ValidateENFT validates the ENFT structure and sovereign data
func (enft *ENFTMetadata) ValidateENFT() (bool, []string) {
	var errors []string

	if enft.TokenID == "" {
		errors = append(errors, "Token ID is required")
	}

	if enft.SovereignData.TotalYield <= 0 {
		errors = append(errors, "Total yield must be positive")
	}

	if !enft.SovereignData.PreAuthorized {
		errors = append(errors, "ENFT must be pre-authorized for sovereign ledger")
	}

	if enft.SovereignData.LedgerHash == "" {
		errors = append(errors, "Ledger hash is required for validation")
	}

	if len(enft.SovereignData.PhysicalAssetTags) == 0 {
		errors = append(errors, "Physical asset tags required for Blu-Vault mirroring")
	}

	return len(errors) == 0, errors
}

// ENFTCollection represents a collection of ENFTs
type ENFTCollection struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ENFTs       []*ENFTMetadata `json:"enfts"`
	TotalYield  float64         `json:"total_yield"`
	CreatedAt   time.Time       `json:"created_at"`
}

// NewENFTCollection creates a new ENFT collection
func NewENFTCollection(name, description string) *ENFTCollection {
	return &ENFTCollection{
		Name:        name,
		Description: description,
		ENFTs:       make([]*ENFTMetadata, 0),
		TotalYield:  0,
		CreatedAt:   time.Now(),
	}
}

// AddENFT adds an ENFT to the collection
func (ec *ENFTCollection) AddENFT(enft *ENFTMetadata) {
	ec.ENFTs = append(ec.ENFTs, enft)
	ec.TotalYield += enft.SovereignData.TotalYield
}

// ToJSON converts ENFT collection to JSON
func (ec *ENFTCollection) ToJSON() (string, error) {
	jsonData, err := json.MarshalIndent(ec, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
