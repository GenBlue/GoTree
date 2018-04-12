package main

import (
	"fmt"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func main() {
	// address: GB6S3XHQVL6ZBAF6FIK62OCK3XTUI4L5Z5YUVYNBZUXZ4AZMVBQZNSAU
	from := "SCRUYGFG76UPX3EIUWGPIQPQDPD24XPR3RII5BD53DYPKZJGG43FL5HI"

	// seed: SDLJZXOSOMKPWAK4OCWNNVOYUEYEESPGCWK53PT7QMG4J4KGDAUIL5LG
	to := "GA3A7AD7ZR4PIYW6A52SP6IK7UISESICPMMZVJGNUTVIZ5OUYOPBTK6X"

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: from},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: to},
			b.NativeAmount{Amount: "0.1"},
		),
	)
	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(from)
	if err != nil {
		panic(err)
	}



	txeB64, err := txe.Base64()

	if err != nil {
		panic(err)
	}
	fmt.Printf("tx base64: %s", txeB64)


	// Second
	blob := "AAAAAH0t3PCq/ZCAvioV7ThK3edEcX3PcUrhoc0vngMsqGGWAAAAZAA1fpcAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAANg+Af8x49GLeB3Un+Qr9ESJJAnsZmqTNpOqM9dTDnhkAAAAAAAAAAAAPQkAAAAAAAAAAASyoYZYAAABA/L7Du9hlYzCtR9F89Mp9/alkCXsq9CWuJ1Mpql+Q16fHE4P2+H62p4cx+b2YUp/fUX73ucW+RPxOgSXmeV6uBQ=="

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)





}
