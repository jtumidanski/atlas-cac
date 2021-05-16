package skill

const (
   EvanFireBreath                   uint32 = 22151001
   EvanIceBreath                    uint32 = 22121000
   FirePoisonArchMagicianBigBang    uint32 = 2121001
   IceLighteningArchMagicianBigBang uint32 = 2221001
   BishopBigBang                    uint32 = 2321001
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
