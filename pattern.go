package main

import (
	"crypto/sha1"
	"encoding/json"
	"math/rand"
)

type MemoryPatternT [memoryPatternSize]string

func (mp *MemoryPatternT) replaceOne(newWorldHash string) {
	for i := memoryPatternSize - 1; i != 0; i-- {
		mp[i] = mp[i-1]
	}
	mp[0] = newWorldHash
}

func (mp *MemoryPatternT) reset() { memoryPattern = MemoryPatternT{} }

// ---

var memoryPattern = MemoryPatternT{}

func HashWorld(world WorldT) string {
	hash := sha1.New()

	worldJson, _ := json.Marshal(world)
	hash.Write(worldJson)

	return string(hash.Sum(nil))
}

func isRepeatedPattern(newWorld WorldT) bool {
	newWorldHash := HashWorld(newWorld)

	for _, i := range memoryPattern {
		if i == newWorldHash {
			memoryPattern.reset()
			return true
		}
	}

	memoryPattern.replaceOne(newWorldHash)

	return false
}

func randomInt(min, max int) int { return rand.Intn(max-min) + min }

func getRandomLocations() [][2]int {
	allSells := width * height

	locations := make([][2]int, 0, randomInt(0, allSells))

	for i := 0; i < randomInt(0, allSells); i++ {
		locations = append(locations, [2]int{randomInt(0, height-1), randomInt(0, width-1)})
	}

	return locations
}

func margeWorlds(world1, world2 WorldT) WorldT {
	for x, i := range world1 {
		for y, j := range i {
			if j {
				world2[x][y] = true
			}
		}
	}

	return world2
}

func injectSells(repeatedWorlds WorldT) WorldT {
	randomWorld := WorldT{}

	for _, i := range getRandomLocations() {
		randomWorld[i[0]][i[1]] = true
	}

	return margeWorlds(randomWorld, repeatedWorlds)
}
