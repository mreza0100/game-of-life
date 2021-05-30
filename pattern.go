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
	randomWorld := genesis(initialPopulation * 4)

	return margeWorlds(randomWorld, repeatedWorlds)
}
