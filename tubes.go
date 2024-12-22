package main

import (
	"fmt"
	"time"
)

// Struktur InventoryItem
type InventoryItem struct {
	ID   int
	Name string
}

// Fungsi utama
func main() {
	var targetID int
	var opsi int

	inventory := make([]InventoryItem, 1000000)
	for i := 0; i < 1000000; i++ {
		inventory[i] = InventoryItem{ID: i + 1, Name: fmt.Sprintf("Item %d", i+1)}
	}

	fmt.Printf("Masukkan ID item yang ingin Anda cari: ")
	fmt.Scan(&targetID)
	fmt.Println("1. Memakai fungsi iterative")
	fmt.Println("2. Memakai fungsi rekursif")
	fmt.Printf("Pilih opsi pencarian (1/2): ")
	fmt.Scan(&opsi)

	start := time.Now()

	var hasil bool
	if opsi == 1 {
		hasil = iterativeSearch(inventory, targetID)
	} else if opsi == 2 {
		hasil = linearSearchByID(inventory, len(inventory), targetID)
	} else {
		fmt.Println("Opsi tidak valid")
		return
	}

	
	duration := time.Since(start).Nanoseconds()

	
	if hasil {
		fmt.Println("Ketemu")
	} else {
		fmt.Println("Tidak ketemu")
	}
	fmt.Printf("Waktu eksekusi: %d ns\n", duration)
}

func iterativeSearch(inventory []InventoryItem, targetID int) bool {
	for i := 0; i < len(inventory); i++ {
		if inventory[i].ID == targetID {
			return true
		}
	}
	return false
}


func linearSearchByID(arr []InventoryItem, size int, key int) bool {
	if size == 0 {
		return false
	} else if arr[size-1].ID == key {
		return true
	}
	return linearSearchByID(arr, size-1, key)
}
