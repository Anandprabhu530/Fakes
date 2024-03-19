package main

import (
	"fmt"
	"math/rand"
	"time"
)

var animalname = []string{"Ant", "Bear", "Cat", "Dog", "Elephant", "Fish", "Frog", "Giraffe",
	"Horse", "Kangaroo", "Lion", "Monkey", "Mouse", "Owl", "Panda", "Penguin",
	"Pig", "Rabbit", "Snake", "Tiger", "Turtle", "Whale", "Zebra", "Cheetah",
	"Chimpanzee", "Crocodile", "Dolphin", "Eagle", "Fox", "Gorilla", "Hippopotamus",
	"Hyena", "Jaguar", "Jellyfish", "Koala", "Leopard", "Lionfish", "Manatee", "Narwhal",
	"Octopus", "Orangutan", "Parrot", "Peacock", "Python", "Quail", "Raccoon", "Rhinoceros",
	"Salamander", "Seahorse", "Shark", "Sheep", "Sloth", "Squid", "Starfish", "Stingray",
	"Tapir", "Tasmanian Devil", "Toucan", "Turkey", "Vulture", "Walrus", "Wolverine",
	"Aardvark", "Albatross", "Armadillo", "Alpaca", "Antelope", "Baboon", "Badger", "Bat",
	"Beagle", "Beetle", "Bison", "Boar", "Budgie", "Camel", "Capybara", "Caribou", "Cassowary",
	"Chameleon", "Cheetah", "Chinchilla", "Cobra", "Condor", "Cormorant", "Coyote", "Crab",
	"Crane", "Cricket", "Crow", "Deer", "Dingo", "Donkey", "Duck", "Echidna", "Eel", "Emu",
	"Ferret", "Finch", "Flamingo", "Flea", "Fly", "Fox", "Gazelle", "Gecko", "Gnu", "Goose",
	"Gorilla", "Grasshopper", "Grebe", "Greyhound", "Guinea Pig", "Hamster", "Hare", "Hawk",
	"Hedgehog", "Heron", "Honeybee", "Hummingbird", "Hyena", "Ibex", "Iguana", "Impala",
	"Jackal", "Jaguar", "Jay", "Jellyfish", "Kangaroo", "Koala", "Komodo Dragon", "Kookaburra"}

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

var countrycode = []string{
	"Afghanistan",
	"Armenia",
	"Azerbaijan",
	"Bahrain",
	"Bangladesh",
	"Bhutan",
	"Brunei",
	"Cambodia",
	"China",
	"Cyprus",
	"East Timor",
	"Egypt",
	"Georgia",
	"India",
	"Indonesia",
	"Iran",
	"Iraq",
	"Israel",
	"Japan",
	"Jordan",
	"Kazakhstan",
	"North Korea",
	"South Korea",
	"Kuwait",
	"Kyrgyzstan",
	"Laos",
	"Lebanon",
	"Malaysia",
	"Maldives",
	"Mongolia",
	"Myanmar",
	"Nepal",
	"Oman",
	"Pakistan",
	"Palestine",
	"Philippines",
	"Qatar",
	"Russia",
	"Saudi Arabia",
	"Singapore",
	"Sri Lanka",
	"Syria",
	"Tajikistan",
	"Thailand",
	"Turkey",
	"Turkmenistan",
	"United Arab Emirates",
	"Uzbekistan",
	"Vietnam",
	"Yemen",
}

func Name(num int) []string {
	return namesdata[:num]
}

func Number() uint {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	randomNumber := 1000000000 + generator.Intn(9999999999)
	return uint(randomNumber)
}

func AnimalName(num int) []string {
	return animalname[:num]
}

func Country() string {
	sources := rand.NewSource(time.Now().Unix())
	generator := rand.New(sources)
	randomnumber := generator.Intn(len(countrycode)) + 1
	return countrycode[randomnumber]
}

func RandomColor() string {
	source := rand.NewSource(time.Now().Unix())
	generator := rand.New(source)
	r := generator.Intn(256)
	g := generator.Intn(256)
	b := generator.Intn(256)
	return fmt.Sprintf("#%02x%02x%02x\n", r, g, b)
}

func main() {
	fmt.Println(Country())
}
