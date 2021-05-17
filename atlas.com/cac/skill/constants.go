package skill

const (
	EvanFireBreath                           uint32 = 22151001
	EvanIceBreath                            uint32 = 22121000
	FirePoisonWizardMPEater                  uint32 = 2100000
	FirePoisonArchMagicianBigBang            uint32 = 2121001
	IceLightningWizardMPEater                uint32 = 2200000
	IceLighteningArchMagicianBigBang         uint32 = 2221001
	ClericMPEater                            uint32 = 2300000
	BishopBigBang                            uint32 = 2321001
	FirePoisonMagicianElementAmplification   uint32 = 2110001
	IceLigthningMagicianElementAmplification uint32 = 2210001
	BlazeWizardElementAmplification          uint32 = 12110001
	EvanMagicAmplification                   uint32 = 22150000
)

// Is determines if the reference skill matches with any of the choices provided. Returns true if so.
func Is(reference uint32, choices ...uint32) bool {
	for _, c := range choices {
		if reference == c {
			return true
		}
	}
	return false
}
