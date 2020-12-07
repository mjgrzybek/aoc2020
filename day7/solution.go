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
type Package2PackagesMap = map[Package]PackagesSet

type Packages2Num = map[Package]int
type Package2Packages2NumMap = map[Package]Packages2Num

func parsePackage(s string) (int, Package) {
	split := strings.SplitN(s, " ", 2)
	atoi, _ := strconv.Atoi(split[0])
	bag := split[1][0:strings.Index(split[1], " bag")]
	return atoi, bag
}

func loadRelations1(r io.Reader) Package2PackagesMap {
	fileScanner := bufio.NewScanner(r)
	p2c := make(Package2PackagesMap)

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
	p2c := loadRelations1(f)

	root := "shiny gold"

	leaves := outermost(root, p2c)
	fmt.Println(leaves)

	return len(leaves)
}

func outermost(root string, p2c Package2PackagesMap) PackagesSet {
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

func loadRelations2(r io.Reader) Package2Packages2NumMap {
	fileScanner := bufio.NewScanner(r)
	c2p2n := make(Package2Packages2NumMap)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		// [A A] bag contain (<X> [B B] bag[s],no other bag)[, <X> [C C] bag[s]]
		p := strings.SplitN(text, " bags contain ", 2)
		container := p[0]
		packages := strings.Split(p[1], ", ")

		for _, p := range packages {
			num, pkg := parsePackage(p)
			if pkg == "other" {
				continue
			}

			if c2p2n[container] == nil {
				c2p2n[container] = make(Packages2Num)
			}
			c2p2n[container][pkg] = num
		}
	}

	for k, v := range c2p2n {
		fmt.Println(k, "...", v)
	}

	return c2p2n
}

func Solve2() int {
	f, _ := os.Open("day7/input.txt")
	c2p2n := loadRelations2(f)

	topBag := "shiny gold"
	_ = topBag
	fmt.Println(c2p2n)

	return calculateBagsInside(topBag, c2p2n) - 1
}

func calculateBagsInside(bag string, c2p2n Package2Packages2NumMap) int {
	sum := 1
	for pkg, num := range c2p2n[bag] {
		sum += num * calculateBagsInside(pkg, c2p2n)
	}
	return sum
}


