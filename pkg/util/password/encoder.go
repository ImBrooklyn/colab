package password

import (
    "crypto/rand"
    "crypto/sha512"
    "encoding/hex"
    "hash"

    "golang.org/x/crypto/pbkdf2"
)

// Options is a struct for custom values of salt length, number of iterations, the encoded key's length,
// and the hash function being used. If set to `nil`, default options are used:
// &Options{ 256, 10000, 512, "sha512" }

type HashFunc struct {
    Name string
    Func func() hash.Hash
}

type Options struct {
    SaltLen      int
    Iterations   int
    KeyLen       int
    HashFunction HashFunc
}

var (
    Sha512 = HashFunc{"sha512", sha512.New}
)

func DefaultOptions() *Options {
    return &Options{
        SaltLen:      16,
        Iterations:   100,
        KeyLen:       32,
        HashFunction: Sha512,
    }
}

func generateSalt(length int) []byte {
    const charPool = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    salt := make([]byte, length)
    _, _ = rand.Read(salt)
    for key, val := range salt {
        salt[key] = charPool[val%byte(len(charPool))]
    }
    return salt
}

// Encode takes two arguments, a raw password, and a pointer to an Options struct.
// In order to use default options, pass `nil` as the second argument.
// It returns the generated salt and encoded key for the user.
func Encode(rawPwd string, options *Options) (string, string) {
    if options == nil {
        options = DefaultOptions()
    }
    salt := generateSalt(options.SaltLen)
    hashFunc := options.HashFunction
    encodedPwd := pbkdf2.Key([]byte(rawPwd), salt, options.Iterations, options.KeyLen, hashFunc.Func)

    return string(salt), hex.EncodeToString(encodedPwd)
}

// Match takes four arguments, the raw password, its generated salt, the encoded password,
// and a pointer to the Options struct, and returns a boolean value determining whether the password is the correct one or not.
// Passing `nil` as the last argument resorts to default options.
func Match(rawPwd string, salt string, encodedPwd string, options *Options) bool {
    if options == nil {
        options = DefaultOptions()
    }
    return encodedPwd == hex.EncodeToString(pbkdf2.Key([]byte(rawPwd), []byte(salt), options.Iterations, options.KeyLen, options.HashFunction.Func))
}
