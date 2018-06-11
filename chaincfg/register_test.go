package chaincfg_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	. "github.com/daglabs/btcd/chaincfg"
)

// Define some of the required parameters for a user-registered
// network.  This is necessary to test the registration of and
// lookup of encoding magics from the network.
var mockNetParams = Params{
	Name:            "mocknet",
	Net:             1<<32 - 1,
	Bech32HRPSegwit: "tc",
	HDPrivateKeyID:  [4]byte{0x01, 0x02, 0x03, 0x04},
	HDPublicKeyID:   [4]byte{0x05, 0x06, 0x07, 0x08},
}

func TestRegister(t *testing.T) {
	type registerTest struct {
		name   string
		params *Params
		err    error
	}
	type prefixTest struct {
		prefix string
		valid  bool
	}
	type hdTest struct {
		priv []byte
		want []byte
		err  error
	}

	tests := []struct {
		name           string
		register       []registerTest
		segwitPrefixes []prefixTest
		hdMagics       []hdTest
	}{
		{
			name: "default networks",
			register: []registerTest{
				{
					name:   "duplicate mainnet",
					params: &MainNetParams,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate regtest",
					params: &RegressionNetParams,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate testnet3",
					params: &TestNet3Params,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate simnet",
					params: &SimNetParams,
					err:    ErrDuplicateNet,
				},
			},
			segwitPrefixes: []prefixTest{
				{
					prefix: MainNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: TestNet3Params.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: RegressionNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: SimNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: strings.ToUpper(MainNetParams.Bech32HRPSegwit + "1"),
					valid:  true,
				},
				{
					prefix: mockNetParams.Bech32HRPSegwit + "1",
					valid:  false,
				},
				{
					prefix: "abc1",
					valid:  false,
				},
				{
					prefix: "1",
					valid:  false,
				},
				{
					prefix: MainNetParams.Bech32HRPSegwit,
					valid:  false,
				},
			},
			hdMagics: []hdTest{
				{
					priv: MainNetParams.HDPrivateKeyID[:],
					want: MainNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: TestNet3Params.HDPrivateKeyID[:],
					want: TestNet3Params.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: RegressionNetParams.HDPrivateKeyID[:],
					want: RegressionNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: SimNetParams.HDPrivateKeyID[:],
					want: SimNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: mockNetParams.HDPrivateKeyID[:],
					err:  ErrUnknownHDKeyID,
				},
				{
					priv: []byte{0xff, 0xff, 0xff, 0xff},
					err:  ErrUnknownHDKeyID,
				},
				{
					priv: []byte{0xff},
					err:  ErrUnknownHDKeyID,
				},
			},
		},
		{
			name: "register mocknet",
			register: []registerTest{
				{
					name:   "mocknet",
					params: &mockNetParams,
					err:    nil,
				},
			},
			segwitPrefixes: []prefixTest{
				{
					prefix: MainNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: TestNet3Params.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: RegressionNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: SimNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: strings.ToUpper(MainNetParams.Bech32HRPSegwit + "1"),
					valid:  true,
				},
				{
					prefix: mockNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: "abc1",
					valid:  false,
				},
				{
					prefix: "1",
					valid:  false,
				},
				{
					prefix: MainNetParams.Bech32HRPSegwit,
					valid:  false,
				},
			},
			hdMagics: []hdTest{
				{
					priv: mockNetParams.HDPrivateKeyID[:],
					want: mockNetParams.HDPublicKeyID[:],
					err:  nil,
				},
			},
		},
		{
			name: "more duplicates",
			register: []registerTest{
				{
					name:   "duplicate mainnet",
					params: &MainNetParams,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate regtest",
					params: &RegressionNetParams,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate testnet3",
					params: &TestNet3Params,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate simnet",
					params: &SimNetParams,
					err:    ErrDuplicateNet,
				},
				{
					name:   "duplicate mocknet",
					params: &mockNetParams,
					err:    ErrDuplicateNet,
				},
			},
			segwitPrefixes: []prefixTest{
				{
					prefix: MainNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: TestNet3Params.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: RegressionNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: SimNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: strings.ToUpper(MainNetParams.Bech32HRPSegwit + "1"),
					valid:  true,
				},
				{
					prefix: mockNetParams.Bech32HRPSegwit + "1",
					valid:  true,
				},
				{
					prefix: "abc1",
					valid:  false,
				},
				{
					prefix: "1",
					valid:  false,
				},
				{
					prefix: MainNetParams.Bech32HRPSegwit,
					valid:  false,
				},
			},
			hdMagics: []hdTest{
				{
					priv: MainNetParams.HDPrivateKeyID[:],
					want: MainNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: TestNet3Params.HDPrivateKeyID[:],
					want: TestNet3Params.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: RegressionNetParams.HDPrivateKeyID[:],
					want: RegressionNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: SimNetParams.HDPrivateKeyID[:],
					want: SimNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: mockNetParams.HDPrivateKeyID[:],
					want: mockNetParams.HDPublicKeyID[:],
					err:  nil,
				},
				{
					priv: []byte{0xff, 0xff, 0xff, 0xff},
					err:  ErrUnknownHDKeyID,
				},
				{
					priv: []byte{0xff},
					err:  ErrUnknownHDKeyID,
				},
			},
		},
	}

	for _, test := range tests {
		for _, regTest := range test.register {
			err := Register(regTest.params)
			if err != regTest.err {
				t.Errorf("%s:%s: Registered network with unexpected error: got %v expected %v",
					test.name, regTest.name, err, regTest.err)
			}
		}
		for i, prxTest := range test.segwitPrefixes {
			valid := IsBech32SegwitPrefix(prxTest.prefix)
			if valid != prxTest.valid {
				t.Errorf("%s: segwit prefix %s (%d) valid mismatch: got %v expected %v",
					test.name, prxTest.prefix, i, valid, prxTest.valid)
			}
		}
		for i, magTest := range test.hdMagics {
			pubKey, err := HDPrivateKeyToPublicKeyID(magTest.priv[:])
			if !reflect.DeepEqual(err, magTest.err) {
				t.Errorf("%s: HD magic %d mismatched error: got %v expected %v ",
					test.name, i, err, magTest.err)
				continue
			}
			if magTest.err == nil && !bytes.Equal(pubKey, magTest.want[:]) {
				t.Errorf("%s: HD magic %d private and public mismatch: got %v expected %v ",
					test.name, i, pubKey, magTest.want[:])
			}
		}
	}
}
