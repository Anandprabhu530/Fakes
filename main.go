package main

import (
	"fmt"
	"math/rand"
	"time"
)

var namesdata = []string{"Liam", "Olivia", "Noah", "Emma", "William", "Ava", "James", "Sophia",
	"Oliver", "Isabella", "Benjamin", "Charlotte", "Elijah", "Mia", "Lucas",
	"Amelia", "Mason", "Evelyn", "Ethan", "Abigail", "Alexander", "Harper",
	"Aiden", "Sofia", "Daniel", "Evelyn", "Matthew", "Elizabeth", "Michael",
	"Avery", "Aiden", "Ella", "Henry", "Scarlett", "Joseph", "Chloe",
	"Jackson", "Brianna", "David", "Victoria", "Luke", "Madison", "Anthony",
	"Ella", "Jack", "Riley", "Dylan", "Aubrey", "Grayson", "Zoe", "Andrew",
	"Layla", "Isaac", "Aaliyah", "Gabriel", "Camila", "Logan", "Penelope",
	"William", "Luna", "James", "Nova", "Benjamin", "Eliana", "Mason", "Violet",
	"Ethan", "Mila", "Alexander", "Claire", "Aiden", "Everly", "Daniel", "Stella",
	"Matthew", "Natalie", "Michael", "Oaklyn", "Joseph", "Aurora", "Jackson",
	"Eleanor", "David", "Maya", "Luke", "Valentina", "Anthony", "Riley",
	"John", "Sarah", "Owen", "Amelia", "Wyatt", "Evelyn", "Robert", "Isabella",
	"Carter", "Charlotte", "Gabriel", "Mia", "Daniel", "Harper", "Elijah",
	"Sofia", "Lucas", "Evelyn", "Mason", "Abigail", "Ethan", "Elizabeth",
	"Alexander", "Avery", "Aiden", "Ella", "Henry", "Scarlett", "Joseph", "Chloe"}

func Name(num int) []string {
	return namesdata[:num]
}

func Number() uint {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	randomNumber := 1000000000 + generator.Intn(9999999999)
	return uint(randomNumber)
}

func main() {
	fmt.Println(Number())
}
