package database

import (
	"encoding/binary"

	"github.com/mszalewicz/ardentdb/internal/assert"
)

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

type BNode struct {
	data []byte
}

type BTree struct {
	root uint16

	get func(uint16) BNode
	new func(BNode) uint16
	del func(uint16)
}

const HEADER               = 4
const BTREE_PAGE_SIZE      = 4096
const BTREE_MAX_KEY_SIZE   = 1000
const BTREE_MAX_VALUE_SIZE = 3000
const POINTERS             = 8
const OFFSET_SIZE              = 2

func init() {
	maxNode := HEADER + POINTERS + OFFSET_SIZE + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VALUE_SIZE
	assert.Assert(maxNode <= BTREE_PAGE_SIZE, "node size exceeds page size")
}

// -------- Header

func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

// -------- Pointers

func (node BNode) getPtr(index uint16) uint64 {
	assert.Assert(index < node.nkeys(), "setting pointer: node key index out of bounds")
	position := HEADER + 8 * index
	return uint64(binary.LittleEndian.Uint16(node.data[position:]))
}

func (node BNode) setPtr(index uint16, value uint64) {
	assert.Assert(index < node.nkeys(), "getting pointer: node key index out of bounds")
	position := HEADER + 8 * index
	binary.LittleEndian.PutUint64(node.data[position:], value)
}

