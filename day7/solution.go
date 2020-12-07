package day7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Package = string
type PackagesSet = map[Package]struct{}
type Package2ContainersMap = map[Package]PackagesSet

func parsePackage(s string) (int, Package) {
	split := strings.SplitN(s, " ", 2)
	atoi, _ := strconv.Atoi(split[0])
	bag := split[1][0:strings.Index(split[1], " bag")]
	return atoi, bag
}

func loadRelations(r io.Reader) Package2ContainersMap {
	fileScanner := bufio.NewScanner(r)
	p2c := make(Package2ContainersMap)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		// [A A] bag contain (<X> [B B] bag[s],no other bag)[, <X> [C C] bag[s]]
		p := strings.SplitN(text, " bags contain ", 2)
		container := p[0]
		packages := strings.Split(p[1], ", ")

		for _, p := range packages {
			_, pkg := parsePackage(p)

			if p2c[pkg] == nil {
				p2c[pkg] = make(PackagesSet)
			}
			p2c[pkg][container] = struct{}{}
		}
	}

	for k, v := range p2c {
		fmt.Println(k, "...", v)
	}

	return p2c
}

func Solve1() int {
	f, _ := os.Open("day7/input.txt")
	p2c := loadRelations(f)

	root := "shiny gold"

	leaves := outermost(root, p2c)
	fmt.Println(leaves)

	return len(leaves)
}

func outermost(root string, p2c Package2ContainersMap) PackagesSet {
	result := make(PackagesSet)

	packagesSet, exists := p2c[root]
	if exists {
		for pkg, _ := range packagesSet {
			result[pkg] = struct{}{}

			// merge sets
			for v, _ := range outermost(pkg, p2c) {
				result[v] = struct{}{}
			}
		}
	}

	return result
}

func Solve2() int {
	return -1
}
