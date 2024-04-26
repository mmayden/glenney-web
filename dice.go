package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Dice object
type Dice struct {
	Value	int //Current value of the dice
	Size	int // number of sides of the dice
}

// Creates new Dice object
func NewDice(size int) *Dice {
	return &Dice {
		Size: size,
	}
}

func (d *Dice) GetValue() int {
	return d.Value
}

//Set number of sides of dice
func (d *Dice) SetValue(size int) {
	d.Size = size
}

// Array of Dice objects
type DiceList struct {
	Dices []*Dice
}

// Create DiceList object
func NewDiceList(numDice int) *DiceList{
	dices := make([]*Dice, numDice)

	fmt.Println("\nPlease indicate a whole integer representing the size, or number of sides for each dice")
	//set_size := 0
	for i := 0; i < numDice; i++ {
		for {
			fmt.Println("\nSize of 1 of the dice: ")
				dices[i] = NewDice(ScanInteger())
				break
		}
	}
	return &DiceList {
		Dices: dices,
	}
}

func GetDiceList(d *DiceList) {
	for i := 0; i < len(d.Dices); i++ {
		fmt.Println(d.Dices[i].Value)
	}
}

func (d *DiceList) Roll() {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(d.Dices); i++ {
		d.Dices[i].Value = rand.Intn(d.Dices[i].Size) + 1
	}
	GetDiceList(d)
}