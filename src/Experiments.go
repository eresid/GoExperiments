package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"math/rand"
	"math"
	"time"
	"math/cmplx"
	"crypto/md5"
	"hash"
	"io"
	"encoding/hex"
	"os/exec"
	"runtime"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// global
var c, python, java bool

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

// Типи int, uint та uintptr зазвичай займають 32 біти на 32-бітній системі та 64 біти на 64-бітній. Ви маєте
// використовувати int для цілих значень, за винятком коли є певні причини для використання розмірних або беззнакових типів.
func main() {
	//testVars()
	//testInputParameters()
	//testStringsEqualFold()
	//testRandom()
	//testThread()
	// f01fc92b23faa973f3492a23d5a705c5
	// f01fc92b23faa973f3492a23d5a705c5
	//testMd5()
	//testExec()
	//showVersion()
	//testStruct()
	testSH1()
}

type Person struct {
	Name string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("Im at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func testStruct() {
	p := Person{Name: "Steve"}
	p.Address = Address{ Number: "13", Street: "Main" }
	p.Address.City = "Gotham"
	p.Address.State = "NY"
	p.Address.Zip = "01313"
	p.Talk()
	p.Location()
}

func showVersion() {
	version := runtime.Version()

	fmt.Println(strings.SplitAfter(version, "go")[1])
}

func testExec() {
	exeCmdSh("uname -a")
	exeCmdBash("uname -a")
}

func exeCmdSh(cmd string) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

func exeCmdBash(cmd string) []byte {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic("some error found")
	}
	fmt.Printf("%s", out)
	return out
}

func testSH1() {
	s := "Ukraine"
	fmt.Println(s)

	fmt.Println(getSha1(s))

	key := "f01fc92b23faa973f3492a23d5a705c5" + "918e862585716e5f6be3899347d4ae4c" + "23a9686dc6e4cd60";
	fmt.Println("sha1", getSha1(key))

	fmt.Println("sha256", getSha256(key))

	fmt.Println("sha512", getSha512(key))
}

func getSha1(value string) (string) {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(value))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	// fmt.Printf("%x\n", bs)
	return hex.EncodeToString(bs);
}

func getSha256(value string) (string) {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha256.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(value))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	// fmt.Printf("%x\n", bs)
	return hex.EncodeToString(bs);
}

func getSha512(value string) (string) {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha512.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(value))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	// fmt.Printf("%x\n", bs)
	return hex.EncodeToString(bs);
}

func testMd5() {
	timeStart := time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		testMd5One("Ukraine")
	}
	// Time 5.905044112 sec
	fmt.Println("testMd5One = ", time.Now().UnixNano() - timeStart)

	timeStart = time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		testMd5Two("Ukraine")
	}
	// Time 3.870982354 sec
	fmt.Println("testMd5Two = ", time.Now().UnixNano() - timeStart)

	timeStart = time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		GetMD5Hash("Ukraine")
	}
	// Time 5.482504281 sec
	fmt.Println("GetMD5Hash = ", time.Now().UnixNano() - timeStart)

	fmt.Println(testMd5One("Ukraine"))
	fmt.Println(testMd5Two("Ukraine"))
	fmt.Println(GetMD5Hash("Ukraine"))
}

func testMd5One(value string) string {
	var hasher hash.Hash = md5.New()
	io.WriteString(hasher, value)

	return hex.EncodeToString(hasher.Sum(nil))
}

// TODO fix
func testMd5Two(value string) string {
	data := []byte(value)
	resultBytes := md5.Sum(data)

	return hex.EncodeToString(resultBytes[:]) // slice the array from beginning to end
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func testVars() {
	fmt.Println("The time is", time.Now().Unix())
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(3, 50))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("ololo", "trololo")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Println("Pi = ", Pi + 5)

	var ii int
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)

	var xxx, yyy int = 3, 4
	var fff float64 = math.Sqrt(float64(xxx * xxx + yyy * yyy))
	var zzz int = int(fff)
	fmt.Println(xxx, yyy, zzz)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// or add(x, y int)
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func testInputParameters() {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}

func testStringsEqualFold() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var result string = scanner.Text()
		fmt.Println(result)

		if strings.EqualFold(result, "go") {
			break
		} else {
			fmt.Println("Enter text: " + result)
		}
	}
}

func testRandom() {
	// Try changing this number!
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		fmt.Println(random(1, 10))
	}


	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
}

func random(min, max int) int {
	// rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func testThread() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}