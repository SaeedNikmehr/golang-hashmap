package main

import (
	"fmt"
)

const size = 10

type HashTable struct {
	array [size]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

//Hash table functions

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//nodes functions

func (b *bucket) insert(newKey string) {
	if b.search(newKey) {
		newNode := &bucketNode{key: newKey}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("The key already exists")
	}
}

func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"test",
		"test1",
		"test2",
	}
	for _, v := range list {
		hashTable.Insert(v)

	}
	fmt.Println(hashTable)
}
