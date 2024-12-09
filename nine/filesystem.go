package nine

import (
	"bufio"
	"cmp"
	"io"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Block struct {
	IsFile bool
	Id     int
}

type DiskMap []Block

func (dm DiskMap) String() string {
	var sb strings.Builder
	for _, b := range dm {
		name := "."
		if b.IsFile {
			name = strconv.Itoa(b.Id)
		}

		sb.WriteString(name)
	}

	return sb.String()
}

func (dm DiskMap) CheckSum() int {
	chk := 0
	for i := range dm {
		if !dm[i].IsFile {
			continue
		}
		chk += i * dm[i].Id
	}

	return chk
}

// FirstFree returns the index of the first free space
func (dm DiskMap) FirstFree() int {
	return slices.IndexFunc(dm, func(n Block) bool {
		return !n.IsFile
	})
}

// LastFile searches for the last file in the map
func (dm DiskMap) LastFile() int {
	for i := len(dm) - 1; i > 0; i-- {
		if dm[i].IsFile {
			return i
		}
	}

	return -1
}

func (dm DiskMap) removeGap() {
	freeIdx := dm.FirstFree()
	if freeIdx < 0 {
		return
	}

	fileIdx := dm.LastFile()
	if fileIdx < 0 {
		return
	}

	// simple swap
	dm[freeIdx], dm[fileIdx] = dm[fileIdx], dm[freeIdx]

}

func (dm DiskMap) IsCompacted() bool {
	// Go though the map until a free piece is found. If a file comes after,
	// we know it's not fully compacted.
	foundFree := false

	for i := range dm {
		if !dm[i].IsFile {
			foundFree = true
		} else if foundFree {
			return false
		}
	}

	return true
}

func (dm DiskMap) Compact() {
	for !dm.IsCompacted() {
		dm.removeGap()
	}
}

// DescendingIDs returns all IDs in decreasing order
func (dm DiskMap) DescendingIDs() []int {
	set := make(map[int]struct{})
	for _, b := range dm {
		set[b.Id] = struct{}{}
	}

	// get a sorted list of unique IDs
	var ids []int

	for k := range set {
		ids = append(ids, k)
	}

	slices.SortFunc(ids, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return ids
}

// IndicesOf returns the start and end indices of the specified ID blocks
func (dm DiskMap) IndicesOf(id int) []int {
	var idx []int
	for k := range dm {
		if dm[k].Id == id {
			idx = append(idx, k)
		}
	}
	return idx
}

// FindGap returns the index of the first appropriately-sized free space.
func (dm DiskMap) FindGap(size int) int {
	end := len(dm) - size

	// Go through each window
	for i := 0; i < end; i++ {
		freeBlocks := 0
		for d := 0; d < size; d++ {
			if !dm[i+d].IsFile {
				freeBlocks++
			} else {
				break
			}
		}
		if freeBlocks == size {
			return i
		}
	}

	return -1
}

func (dm DiskMap) Defrag() {
	order := dm.DescendingIDs()
	for _, id := range order {
		indices := dm.IndicesOf(id)
		size := len(indices)
		addr := dm.FindGap(size)

		// Skip if no gap found
		if addr < 0 || addr > indices[0] {
			continue
		}

		// Swap out
		for i := 0; i < size; i++ {
			dm[addr+i], dm[indices[i]] = dm[indices[i]], dm[addr+i]
		}
	}
}

func ParseMap(r io.Reader) DiskMap {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)

	var blocks []Block

	// Alternate between file or space length:
	// true means we are reading file length,
	// false means we are reading space length.
	file := true
	id := 0

	for scanner.Scan() {
		char := scanner.Text()

		// Skip non-numbers
		if !unicode.IsDigit(rune(char[0])) {
			continue
		}

		len := int(char[0] - '0')

		newBlock := Block{
			IsFile: file,
		}

		if newBlock.IsFile {
			newBlock.Id = id
			id++
		}

		for range len {
			blocks = append(blocks, newBlock)
		}

		file = !file
	}

	return blocks
}
