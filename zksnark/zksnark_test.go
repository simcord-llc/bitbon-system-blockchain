package zksnark

import (
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
}

func assertEqual(t *testing.T, a, b interface{}, message string) { // nolint:unparam
	if a == b {
		return
	}
	if message == "" {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestNew(t *testing.T) {
	Init()
	var hashInput = []byte("gkdnsrtlmgudkvmtjsjvbgufmyjgnror")
	var hashInputWrong = []byte("gkdnsrtlmgudkvmtjsjvbgufmyjgnroe")
	start := time.Now()
	pk, vk := GenerateKeys(1000)
	generateTime := time.Since(start)

	startP := time.Now()
	proof, err := Prove(hashInput, pk, 1000)
	if err != nil {
		t.Fatal(err)
	}
	proofTime := time.Since(startP)
	startV := time.Now()

	ok, err := Verify(proof, vk, hashInput)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, ok, true, "proof was not valid as expected")
	verifyTime := time.Since(startV)

	ok, err = Verify(proof, vk, hashInputWrong)
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, ok, false, "proof was not valid as expected")
	verifyTime2 := time.Since(startV)

	fmt.Printf("* Create keys(PK,VK): %s\n", generateTime)
	fmt.Printf("* Create proof: %s\n", proofTime)
	fmt.Printf("* Verify valid data: %s\n", verifyTime)
	fmt.Printf("* Verify wrong data: %s\n", verifyTime2)
}

func TestTransfer(t *testing.T) {
	Init()
	var code [64]byte
	var wrongCode [64]byte

	copy(code[:], "qwerty1234adsfasdfasdfasd")
	copy(wrongCode[:], "qwerfdsgty123asdfasdfasdfasdf")
	start := time.Now()

	fmt.Print("Geterate prof & pk/vk\n")
	proof, vk := ProveTransfer(code, 20000)
	proofTime := time.Since(start)
	startV := time.Now()

	fmt.Printf("proof: % x\n", proof)

	fmt.Print("Check valid data\n")
	assertEqual(t, VerifyTransfer(proof, vk, code), true, "proof was not valid as expected")
	verifyTime := time.Since(startV)
	startV = time.Now()

	fmt.Print("Check invalid data\n")
	assertEqual(t, VerifyTransfer(proof, vk, wrongCode), false, "proof was not valid as expected")
	verifyTime2 := time.Since(startV)

	fmt.Printf("* Create keys(PK,VK) and proof: %s\n", proofTime)
	fmt.Printf("* Verify valid data: %s\n", verifyTime)
	fmt.Printf("* Verify wrong data: %s\n", verifyTime2)
}

func TestVerifyTransferByHash(t *testing.T) {
	Init()
	var code [64]byte
	copy(code[:], "qwerty1234adsfasdfasdfasd")

	start := time.Now()
	fmt.Print("Geterate prof & pk/vk\n")
	proof, vk := ProveTransfer(code, 12000)
	proofTime := time.Since(start)
	startV := time.Now()

	fmt.Printf("proof: % x\n", proof)

	fmt.Print("Check valid data\n")
	codeHash := Circuit(code)
	fmt.Print("Check valid data\n")

	assertEqual(t, VerifyTransferByHash(proof, vk, codeHash), true, "proof was not valid as expected")
	verifyTime := time.Since(startV)

	fmt.Printf("* Create keys(PK,VK) and proof: %s\n", proofTime)
	fmt.Printf("* Verify valid data: %s\n", verifyTime)
}
