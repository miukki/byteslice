package circuit

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// Circuit defines a pre-image knowledge proof
// mimc(secret preImage) = public hash
type Circuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	OrderID             frontend.Variable
	ByteSliceRecipient1 frontend.Variable
	ByteSliceRecipient2 frontend.Variable
	ByteSliceRecipient3 frontend.Variable
	Hash                frontend.Variable `gnark:",public"`
}

// Define declares the circuit's constraints
func (circuit *Circuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {

	// hash function, with hardcoded order ID
	mimcLoc, err := mimc.NewMiMC("TestOTC", curveID)
	if err != nil {
		return err
	}

	// specify constraint using the missing complement slice
	hx := mimcLoc.Hash(cs, circuit.OrderID, circuit.ByteSliceRecipient1, circuit.ByteSliceRecipient2, circuit.ByteSliceRecipient3)

	cs.Println("debug circuit.ByteSliceRecipient1", circuit.ByteSliceRecipient1)
	cs.Println("debug circuit.ByteSliceRecipient2", circuit.ByteSliceRecipient2)
	cs.Println("debug circuit.ByteSliceRecipient3", circuit.ByteSliceRecipient3)
	cs.Println("debug hash", hx)
	cs.Println("debug circuit.Hash", circuit.Hash)

	cs.AssertIsEqual(circuit.Hash, hx)

	return nil
}

//go:generate abigen --sol ../../../../registry-contract/contracts/Verifier.sol --pkg circuit --out wrapper.go
