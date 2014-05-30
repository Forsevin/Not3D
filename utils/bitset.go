package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

// Word size of a bit set
const wordSize = uint(64)

// for laster arith.
const log2WordSize = uint(6)

// The zero value of a BitSet is an empty set of length 0.
type BitSet struct {
	length uint
	set    []uint64
}

type BitSetError string

// fixup b.set to be non-nil and return the field value
func (b *BitSet) safeSet() []uint64 {
	if b.set == nil {
		b.set = make([]uint64, wordsNeeded(0))
	}
	return b.set
}

func wordsNeeded(i uint) int {
	if i > ((^uint(0)) - wordSize + 1) {
		return int((^uint(0)) >> log2WordSize)
	}
	return int((i + (wordSize - 1)) >> log2WordSize)
}

func New(length uint) *BitSet {
	return &BitSet{length, make([]uint64, wordsNeeded(length))}
}

func Cap() uint {
	return ^uint(0)
}

func (b *BitSet) Len() uint {
	return b.length
}

//
func (b *BitSet) extendSetMaybe(i uint) {
	if i >= b.length { // if we need more bits, make 'em
		nsize := wordsNeeded(i + 1)
		if b.set == nil {
			b.set = make([]uint64, nsize)
		} else if len(b.set) < nsize {
			newset := make([]uint64, nsize)
			copy(newset, b.set)
			b.set = newset
		}
		b.length = i + 1
	}
}

/// Test whether bit i is set.
func (b *BitSet) Test(i uint) bool {
	if i >= b.length {
		return false
	}
	return b.set[i>>log2WordSize]&(1<<(i&(wordSize-1))) != 0
}

// Set bit i to 1
func (b *BitSet) Set(i uint) *BitSet {
	b.extendSetMaybe(i)
	b.set[i>>log2WordSize] |= 1 << (i & (wordSize - 1))
	return b
}

// Clear bit i to 0
func (b *BitSet) Clear(i uint) *BitSet {
	if i >= b.length {
		return b
	}
	b.set[i>>log2WordSize] &^= 1 << (i & (wordSize - 1))
	return b
}

// Set bit i to value
func (b *BitSet) SetTo(i uint, value bool) *BitSet {
	if value {
		return b.Set(i)
	}
	return b.Clear(i)
}

// Flip bit at i
func (b *BitSet) Flip(i uint) *BitSet {
	if i >= b.length {
		return b.Set(i)
	}
	b.set[i>>log2WordSize] ^= 1 << (i & (wordSize - 1))
	return b
}

// return the next bit set from the specified index, including possibly the current index
// along with an error code (true = valid, false = no set bit found)
// for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
func (b *BitSet) NextSet(i uint) (uint, bool) {
	x := int(i >> log2WordSize)
	if x >= len(b.set) {
		return 0, false
	}
	w := b.set[x]
	w = w >> (i & (wordSize - 1))
	if w != 0 {
		return i + trailingZeroes64(w), true
	}
	x = x + 1
	for x < len(b.set) {
		if b.set[x] != 0 {
			return uint(x)*wordSize + trailingZeroes64(b.set[x]), true
		}
		x = x + 1

	}
	return 0, false
}

// Clear entire BitSet
func (b *BitSet) ClearAll() *BitSet {
	if b != nil && b.set != nil {
		for i := range b.set {
			b.set[i] = 0
		}
	}
	return b
}

// Query words used in a bit set
func (b *BitSet) wordCount() int {
	return wordsNeeded(b.length)
}

// Clone this BitSet
func (b *BitSet) Clone() *BitSet {
	c := New(b.length)
	if b.set != nil { // Clone should not modify current object
		copy(c.set, b.set)
	}
	return c
}

// Copy this BitSet into a destination BitSet
// Returning the size of the destination BitSet
// like array copy
func (b *BitSet) Copy(c *BitSet) (count uint) {
	if c == nil {
		return
	}
	if b.set != nil { // Copy should not modify current object
		copy(c.set, b.set)
	}
	count = c.length
	if b.length < c.length {
		count = b.length
	}
	return
}

// From Wikipedia: http://en.wikipedia.org/wiki/Hamming_weight
const m1 uint64 = 0x5555555555555555  //binary: 0101...
const m2 uint64 = 0x3333333333333333  //binary: 00110011..
const m4 uint64 = 0x0f0f0f0f0f0f0f0f  //binary:  4 zeros,  4 ones ...
const m8 uint64 = 0x00ff00ff00ff00ff  //binary:  8 zeros,  8 ones ...
const m16 uint64 = 0x0000ffff0000ffff //binary: 16 zeros, 16 ones ...
const m32 uint64 = 0x00000000ffffffff //binary: 32 zeros, 32 ones
const hff uint64 = 0xffffffffffffffff //binary: all ones
const h01 uint64 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3...

// From Wikipedia: count number of set bits.
// This is algorithm popcount_2 in the article retrieved May 9, 2011

func popcount_2(x uint64) uint64 {
	x -= (x >> 1) & m1             //put count of each 2 bits into those 2 bits
	x = (x & m2) + ((x >> 2) & m2) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & m4        //put count of each 8 bits into those 8 bits
	x += x >> 8                    //put count of each 16 bits into their lowest 8 bits
	x += x >> 16                   //put count of each 32 bits into their lowest 8 bits
	x += x >> 32                   //put count of each 64 bits into their lowest 8 bits
	return x & 0x7f
}

// Count (number of set bits)
func (b *BitSet) Count() uint {
	if b != nil && b.set != nil {
		cnt := uint64(0)
		for _, word := range b.set {
			cnt += popcount_2(word)
		}
		return uint(cnt)
	}
	return 0
}

var deBruijn = [...]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

func trailingZeroes64(v uint64) uint {
	return uint(deBruijn[((v&-v)*0x03f79d71b4ca8b09)>>58])
}

// Test the equvalence of two BitSets.
// False if they are of different sizes, otherwise true
// only if all the same bits are set
func (b *BitSet) Equal(c *BitSet) bool {
	if c == nil {
		return false
	}
	if b.length != c.length {
		return false
	}
	if b.length == 0 { // if they have both length == 0, then could have nil set
		return true
	}
	// testing for equality shoud not transform the bitset (no call to safeSet)

	for p, v := range b.set {
		if c.set[p] != v {
			return false
		}
	}
	return true
}

func panicIfNull(b *BitSet) {
	if b == nil {
		panic(BitSetError("BitSet must not be null"))
	}
}

// Difference of base set and other set
// This is the BitSet equivalent of &^ (and not)
func (b *BitSet) Difference(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	result = b.Clone() // clone b (in case b is bigger than compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	for i := 0; i < l; i++ {
		result.set[i] = b.set[i] &^ compare.set[i]
	}
	return
}

// computes the cardinality of the differnce
func (b *BitSet) DifferenceCardinality(compare *BitSet) uint {
	panicIfNull(b)
	panicIfNull(compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	cnt := uint64(0)
	for i := 0; i < l; i++ {
		cnt += popcount_2(b.set[i] &^ compare.set[i])
	}
	for i := l; i < len(b.set); i++ {
		cnt += popcount_2(b.set[i])
	}
	return uint(cnt)
}

// Difference of base set and other set
// This is the BitSet equivalent of &^ (and not)
func (b *BitSet) InPlaceDifference(compare *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	for i := 0; i < l; i++ {
		b.set[i] &^= compare.set[i]
	}
}

// Convenience function: return two bitsets ordered by
// increasing length. Note: neither can be nil
func sortByLength(a *BitSet, b *BitSet) (ap *BitSet, bp *BitSet) {
	if a.length <= b.length {
		ap, bp = a, b
	} else {
		ap, bp = b, a
	}
	return
}

// Intersection of base set and other set
// This is the BitSet equivalent of & (and)
func (b *BitSet) Intersection(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	result = New(b.length)
	for i, word := range b.set {
		result.set[i] = word & compare.set[i]
	}
	return
}

// Computes the cardinality of the union
func (b *BitSet) IntersectionCardinality(compare *BitSet) uint {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	cnt := uint64(0)
	for i, word := range b.set {
		cnt += popcount_2(word & compare.set[i])
	}
	return uint(cnt)
}

// Intersection of base set and other set
// This is the BitSet equivalent of & (and)
func (b *BitSet) InPlaceIntersection(compare *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	for i := 0; i < l; i++ {
		b.set[i] &= compare.set[i]
	}
	for i := l; i < len(b.set); i++ {
		b.set[i] = 0
	}
	if compare.length > 0 {
		b.extendSetMaybe(compare.length - 1)
	}
	return
}

// Union of base set and other set
// This is the BitSet equivalent of | (or)
func (b *BitSet) Union(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	result = compare.Clone()
	for i, word := range b.set {
		result.set[i] = word | compare.set[i]
	}
	return
}

func (b *BitSet) UnionCardinality(compare *BitSet) uint {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	cnt := uint64(0)
	for i, word := range b.set {
		cnt += popcount_2(word | compare.set[i])
	}
	for i := len(b.set); i < len(compare.set); i++ {
		cnt += popcount_2(compare.set[i])
	}

	return uint(cnt)
}

// Union of base set and other set
// This is the BitSet equivalent of | (or)
func (b *BitSet) InPlaceUnion(compare *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	if compare.length > 0 {
		b.extendSetMaybe(compare.length - 1)
	}
	for i := 0; i < l; i++ {
		b.set[i] |= compare.set[i]
	}
	if len(compare.set) > l {
		for i := l; i < len(compare.set); i++ {
			b.set[i] = compare.set[i]
		}
	}
}

// SymmetricDifference of base set and other set
// This is the BitSet equivalent of ^ (xor)
func (b *BitSet) SymmetricDifference(compare *BitSet) (result *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	// compare is bigger, so clone it
	result = compare.Clone()
	for i, word := range b.set {
		result.set[i] = word ^ compare.set[i]
	}
	return
}

// computes the cardinality of the symmetric difference
func (b *BitSet) SymmetricDifferenceCardinality(compare *BitSet) uint {
	panicIfNull(b)
	panicIfNull(compare)
	b, compare = sortByLength(b, compare)
	cnt := uint64(0)
	for i, word := range b.set {
		cnt += popcount_2(word ^ compare.set[i])
	}
	for i := len(b.set); i < len(compare.set); i++ {
		cnt += popcount_2(compare.set[i])
	}

	return uint(cnt)
}

// SymmetricDifference of base set and other set
// This is the BitSet equivalent of ^ (xor)
func (b *BitSet) InPlaceSymmetricDifference(compare *BitSet) {
	panicIfNull(b)
	panicIfNull(compare)
	l := int(compare.wordCount())
	if l > int(b.wordCount()) {
		l = int(b.wordCount())
	}
	if compare.length > 0 {
		b.extendSetMaybe(compare.length - 1)
	}
	for i := 0; i < l; i++ {
		b.set[i] ^= compare.set[i]
	}
	if len(compare.set) > l {
		for i := l; i < len(compare.set); i++ {
			b.set[i] = compare.set[i]
		}
	}
}

// Is the length an exact multiple of word sizes?
func (b *BitSet) isEven() bool {
	return b.length%wordSize == 0
}

// Clean last word by setting unused bits to 0
func (b *BitSet) cleanLastWord() {
	if !b.isEven() {
		// Mask for cleaning last word
		const allBits uint64 = 0xffffffffffffffff
		b.set[wordsNeeded(b.length)-1] &= allBits >> (wordSize - b.length%wordSize)
	}
}

// Return the (local) Complement of a biset (up to length bits)
func (b *BitSet) Complement() (result *BitSet) {
	panicIfNull(b)
	result = New(b.length)
	for i, word := range b.set {
		result.set[i] = ^word
	}
	result.cleanLastWord()
	return
}

// Returns true if all bits are set, false otherwise
func (b *BitSet) All() bool {
	panicIfNull(b)
	return b.Count() == b.length
}

// Return true if no bit is set, false otherwise
func (b *BitSet) None() bool {
	panicIfNull(b)
	if b != nil && b.set != nil {
		for _, word := range b.set {
			if word > 0 {
				return false
			}
		}
		return true
	}
	return true
}

// Return true if any bit is set, false otherwise
func (b *BitSet) Any() bool {
	panicIfNull(b)
	return !b.None()
}

// Dump as bits
func (b *BitSet) DumpAsBits() string {
	if b.set == nil {
		return "."
	}
	buffer := bytes.NewBufferString("")
	i := len(b.set) - 1
	for ; i >= 0; i-- {
		fmt.Fprintf(buffer, "%064b.", b.set[i])
	}
	return string(buffer.Bytes())
}

func (b *BitSet) MarshalJSON() ([]byte, error) {
	// Put the bitset length in front of the string
	length := uint64(b.length)
	dataCap := binary.Size(length) + binary.Size(b.set)
	buffer := bytes.NewBuffer(make([]byte, 0, dataCap))

	// Write length
	err := binary.Write(buffer, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	// Write set
	err = binary.Write(buffer, binary.BigEndian, b.set)
	if err != nil {
		return nil, err
	}

	// URLEncode all bytes
	return json.Marshal(base64.URLEncoding.EncodeToString(buffer.Bytes()))
}

func (b *BitSet) UnmarshalJSON(data []byte) error {
	// Unmarshal as string
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	// URLDecode string
	buf, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(buf)
	var length uint64

	// Read length first
	err = binary.Read(reader, binary.BigEndian, &length)
	if err != nil {
		return err
	}
	newset := New(uint(length))

	if uint64(newset.length) != length {
		return errors.New("Unmarshalling error: type mismatch")
	}

	// Read remaining bytes as set
	err = binary.Read(reader, binary.BigEndian, newset.set)

	if err != nil {
		return err
	}

	*b = *newset
	return nil
}
