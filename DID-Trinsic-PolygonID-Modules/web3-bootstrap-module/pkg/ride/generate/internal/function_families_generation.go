package internal

import (
	"fmt"
)

func GenerateFunctionFamilies(fn string) {
	cd := NewCoder("ride")
	cd.Import("crypto/rsa")
	cd.Import("crypto/sha256", "sh256")
	cd.Import("crypto/x509")
	cd.Import("github.com/pkg/errors")
	cd.Import("github.com/wavesplatform/gowaves/pkg/crypto")
	cd.Import("github.com/wavesplatform/gowaves/pkg/ride/crypto", "c2")
	for _, l := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} {
		fn := fmt.Sprintf("bls12Groth16Verify_%d", l)
		cd.Line("func %s(env environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("if err := checkArgs(args, 3); err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("key, ok := args[0].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[0].instanceOf())", fn)
		cd.Line("}")
		cd.Line("proof, ok := args[1].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[1].instanceOf())", fn)
		cd.Line("}")
		cd.Line("inputs, ok := args[2].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[2].instanceOf())", fn)
		cd.Line("}")
		cd.Line("if l := len(inputs); l > 32*%d {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid inputs size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("ok, err := crypto.Bls12381{}.Groth16Verify(key, proof, inputs)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("return rideBoolean(ok), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} {
		fn := fmt.Sprintf("bn256Groth16Verify_%d", l)
		cd.Line("func %s(env environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("if err := checkArgs(args, 3); err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("key, ok := args[0].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[0].instanceOf())", fn)
		cd.Line("}")
		cd.Line("proof, ok := args[1].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[1].instanceOf())", fn)
		cd.Line("}")
		cd.Line("inputs, ok := args[2].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[2].instanceOf())", fn)
		cd.Line("}")
		cd.Line("if l := len(inputs); l > 32*%d {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid inputs size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("ok, err := crypto.Bn256{}.Groth16Verify(key, proof, inputs)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("return rideBoolean(ok), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{8, 16, 32, 64, 128} {
		fn := fmt.Sprintf("sigVerify_%d", l)
		cd.Line("func %s(env environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("if err := checkArgs(args, 3); err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("message, ok := args[0].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[0].instanceOf())", fn)
		cd.Line("}")
		cd.Line("if l := len(message); l > %d*1024 {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid message size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("signature, ok := args[1].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[1].instanceOf())", fn)
		cd.Line("}")
		cd.Line("pkb, ok := args[2].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[2].instanceOf())", fn)
		cd.Line("}")
		cd.Line("pk, err := crypto.NewPublicKeyFromBytes(pkb)")
		cd.Line("if err != nil {")
		cd.Line("return rideBoolean(false), nil")
		cd.Line("}")
		cd.Line("sig, err := crypto.NewSignatureFromBytes(signature)")
		cd.Line("if err != nil {")
		cd.Line("return rideBoolean(false), nil")
		cd.Line("}")
		cd.Line("ok = crypto.Verify(pk, sig, message)")
		cd.Line("return rideBoolean(ok), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{16, 32, 64, 128} {
		fn := fmt.Sprintf("rsaVerify_%d", l)
		cd.Line("func %s(_ environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("if err := checkArgs(args, 4); err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("digest, err := digest(args[0])")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("message, ok := args[1].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[1].instanceOf())", fn)
		cd.Line("}")
		cd.Line("if l := len(message); l > %d*1024 {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid message size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("sig, ok := args[2].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[2].instanceOf())", fn)
		cd.Line("}")
		cd.Line("pk, ok := args[3].(rideBytes)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.Errorf(\"%s: unexpected argument type '%%s'\", args[3].instanceOf())", fn)
		cd.Line("}")
		cd.Line("key, err := x509.ParsePKIXPublicKey(pk)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s: invalid public key\")", fn)
		cd.Line("}")
		cd.Line("k, ok := key.(*rsa.PublicKey)")
		cd.Line("if !ok {")
		cd.Line("return nil, errors.New(\"%s: not an RSA key\")", fn)
		cd.Line("}")
		cd.Line("d := message")
		cd.Line("if digest != 0 {")
		cd.Line("h := digest.New()")
		cd.Line("_, _ = h.Write(message)")
		cd.Line("d = h.Sum(nil)")
		cd.Line("}")
		cd.Line("ok, err = c2.VerifyPKCS1v15(k, digest, d, sig)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("return rideBoolean(ok), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{16, 32, 64, 128} {
		fn := fmt.Sprintf("keccak256_%d", l)
		cd.Line("func %s(env environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("data, err := bytesOrStringArg(args)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("if l := len(data); l > %d*1024 {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid data size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("d, err := crypto.Keccak256(data)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("return rideBytes(d.Bytes()), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{16, 32, 64, 128} {
		fn := fmt.Sprintf("blake2b256_%d", l)
		cd.Line("func %s(_ environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("data, err := bytesOrStringArg(args)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("if l := len(data); l > %d*1024 {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid data size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("d, err := crypto.FastHash(data)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("return rideBytes(d.Bytes()), nil")
		cd.Line("}")
		cd.Line("")
	}
	for _, l := range []int{16, 32, 64, 128} {
		fn := fmt.Sprintf("sha256_%d", l)
		cd.Line("func %s(_ environment, args ...rideType) (rideType, error) {", fn)
		cd.Line("data, err := bytesOrStringArg(args)")
		cd.Line("if err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("if l := len(data); l > %d*1024 {", l)
		cd.Line("return nil, errors.Errorf(\"%s: invalid data size %%d\", l)", fn)
		cd.Line("}")
		cd.Line("h := sh256.New()")
		cd.Line("if _, err = h.Write(data); err != nil {")
		cd.Line("return nil, errors.Wrap(err, \"%s\")", fn)
		cd.Line("}")
		cd.Line("d := h.Sum(nil)")
		cd.Line("return rideBytes(d), nil")
		cd.Line("}")
		cd.Line("")
	}
	if err := cd.Save(fn); err != nil {
		panic(err)
	}
}
