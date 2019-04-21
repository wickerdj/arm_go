package main

import (
	"fmt"
	"log"
	"strings"

	"periph.io/x/periph/conn/onewire"
	"periph.io/x/periph/conn/onewire/onewirereg"
	"periph.io/x/periph/host"
)

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Enumerate all 1-wire buses available and the corresponding pins.
	fmt.Print("1-wire buses available:\n")
	for _, ref := range onewirereg.All() {
		fmt.Printf("- %s\n", ref.Name)
		if ref.Number != -1 {
			fmt.Printf("  %d\n", ref.Number)
		}
		if len(ref.Aliases) != 0 {
			fmt.Printf("  %s\n", strings.Join(ref.Aliases, " "))
		}

		b, err := ref.Open()
		if err != nil {
			fmt.Printf("  Failed to open: %v", err)
		}
		if p, ok := b.(onewire.Pins); ok {
			fmt.Printf("  Q: %s", p.Q())
		}
		if err := b.Close(); err != nil {
			fmt.Printf("  Failed to close: %v", err)
		}
	}

}
