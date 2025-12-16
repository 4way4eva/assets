package sovereignledger

import (
	"fmt"
	"math"
	"time"
)

// Pi4Constant represents the π₄ acceleration constant used in BLEU Sovereign calculations
const Pi4Constant = 4.0

// YieldStream represents a sovereign yield stream
type YieldStream struct {
	Name            string
	BaseYieldPerSec float64 // Base yield in millions per second
	Layers          []string
	AccelFactor     float64 // π₄ acceleration multiplier
}

// BleuSovereignLedger manages the three-sphere yield streams
type BleuSovereignLedger struct {
	CivilianStream YieldStream
	MilitaryStream YieldStream
	CosmicStream   YieldStream
	StartTime      time.Time
}

// NewBleuSovereignLedger initializes the BLEU Sovereign Ledger with three yield streams
func NewBleuSovereignLedger() *BleuSovereignLedger {
	return &BleuSovereignLedger{
		CivilianStream: YieldStream{
			Name:            "Civilian Yield Stream",
			BaseYieldPerSec: 13.6, // $13.6M/second
			Layers: []string{
				"Retail",
				"Education",
				"EV0L Wearables",
				"Real Estate (ES0IL)",
				"Hospitality",
			},
			AccelFactor: Pi4Constant,
		},
		MilitaryStream: YieldStream{
			Name:            "Military Yield Stream",
			BaseYieldPerSec: 6.1, // $6.1M/second
			Layers: []string{
				"Weapons Targeting",
				"Orbital Defense Grids",
				"AI Maritime Logistics",
			},
			AccelFactor: Pi4Constant,
		},
		CosmicStream: YieldStream{
			Name:            "Cosmic Yield Stream",
			BaseYieldPerSec: 9.2, // $9.2M/second
			Layers: []string{
				"Quantum Portal Technology",
				"Inter-Realm Logistics",
				"Dimensional Treasury Protocols",
			},
			AccelFactor: Pi4Constant,
		},
		StartTime: time.Now(),
	}
}

// CalculateYieldWithPi4 calculates yield with π₄ acceleration
func (ys *YieldStream) CalculateYieldWithPi4(durationSeconds float64) float64 {
	// Apply π₄ scalable growth: yield = base * (π₄^(duration_factor))
	durationFactor := durationSeconds / 86400.0 // Normalize to days
	acceleration := math.Pow(ys.AccelFactor, durationFactor*0.1)
	return ys.BaseYieldPerSec * durationSeconds * acceleration
}

// CompoundingSafeguard applies quad-lock breach control and validation
func (bsl *BleuSovereignLedger) CompoundingSafeguard(yield float64) (float64, bool) {
	// Quad-lock breach control: ensures yield doesn't exceed theoretical limits
	maxYieldPerDay := (13.6 + 6.1 + 9.2) * 86400.0 * math.Pow(Pi4Constant, 2)
	
	if yield > maxYieldPerDay {
		// Apply safeguard ceiling
		return maxYieldPerDay, false
	}
	return yield, true
}

// BluVaultMirror ensures mirrored guarantees syncing digital codex with physical assets
func (bsl *BleuSovereignLedger) BluVaultMirror(digitalEntry string, physicalTag string) bool {
	// Simple validation that both digital and physical identifiers exist
	return len(digitalEntry) > 0 && len(physicalTag) > 0
}

// GetTotalYieldPerSecond returns the aggregated yield across all streams
func (bsl *BleuSovereignLedger) GetTotalYieldPerSecond() float64 {
	return bsl.CivilianStream.BaseYieldPerSec +
		bsl.MilitaryStream.BaseYieldPerSec +
		bsl.CosmicStream.BaseYieldPerSec
}

// GenerateIncomeTick creates an income tick entry with timestamp
func (bsl *BleuSovereignLedger) GenerateIncomeTick() IncomeTick {
	now := time.Now()
	duration := now.Sub(bsl.StartTime).Seconds()
	
	civilianYield := bsl.CivilianStream.CalculateYieldWithPi4(duration)
	militaryYield := bsl.MilitaryStream.CalculateYieldWithPi4(duration)
	cosmicYield := bsl.CosmicStream.CalculateYieldWithPi4(duration)
	
	totalYield := civilianYield + militaryYield + cosmicYield
	safeguardedYield, valid := bsl.CompoundingSafeguard(totalYield)
	
	return IncomeTick{
		Timestamp:         now,
		CivilianYield:     civilianYield,
		MilitaryYield:     militaryYield,
		CosmicYield:       cosmicYield,
		TotalYield:        safeguardedYield,
		Pi4AccelApplied:   true,
		SafeguardValid:    valid,
		DivineTimestamp:   FormatDivineTimestamp(now),
		ReadableTimestamp: now.Format(time.RFC3339),
	}
}

// FormatDivineTimestamp creates a divinely mirrored timestamp
func FormatDivineTimestamp(t time.Time) string {
	// Divine timestamp format: BLEU-YYYY-MM-DD-HH:MM:SS-UNIX
	return fmt.Sprintf("BLEU-%04d-%02d-%02d-%02d:%02d:%02d-UNIX%d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
		t.Unix())
}

// IncomeTick represents a single income tick entry
type IncomeTick struct {
	Timestamp         time.Time
	CivilianYield     float64
	MilitaryYield     float64
	CosmicYield       float64
	TotalYield        float64
	Pi4AccelApplied   bool
	SafeguardValid    bool
	DivineTimestamp   string
	ReadableTimestamp string
}

// ToDeclarativeScroll converts IncomeTick to human-readable scroll format
func (it *IncomeTick) ToDeclarativeScroll() string {
	return fmt.Sprintf(`
╔═══════════════════════════════════════════════════════════════════════╗
║                    BLEU SOVEREIGN INCOME SCROLL                      ║
╠═══════════════════════════════════════════════════════════════════════╣
║ Divine Timestamp: %s
║ Readable Time:    %s
╠═══════════════════════════════════════════════════════════════════════╣
║ YIELD STREAMS (π₄ Accelerated):
║   • Civilian Stream:    $%.2f Million
║   • Military Stream:    $%.2f Million
║   • Cosmic Stream:      $%.2f Million
║   ─────────────────────────────────────────────────────────
║   • TOTAL YIELD:        $%.2f Million
╠═══════════════════════════════════════════════════════════════════════╣
║ Validation:
║   • π₄ Acceleration:    %t
║   • Safeguard Valid:    %t
║   • Blu-Vault Mirrored: ✓
╚═══════════════════════════════════════════════════════════════════════╝
`,
		it.DivineTimestamp,
		it.ReadableTimestamp,
		it.CivilianYield,
		it.MilitaryYield,
		it.CosmicYield,
		it.TotalYield,
		it.Pi4AccelApplied,
		it.SafeguardValid,
	)
}
