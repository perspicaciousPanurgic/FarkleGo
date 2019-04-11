package main

import (
	"fmt"
)

//Global constants
const MaxDice = 6
const PointsToWin = 2000

// Method to solicit Y or N from user
func GetBinaryChoice() string {
	userInput := "Y"

	fmt.Print("Enter either Y or N : ")
	fmt.Scan(&userInput)

	//check for valid choice
	if !isValidBinaryChoice(userInput) {
		fmt.Println("Invalid choice.")
		userInput = GetBinaryChoice()
	}

	return userInput
}

// Check if user input is valid Y or N
func isValidBinaryChoice(input string) bool {
	valid := false
	if input == "Y" || input == "N" {
		valid = true
	}
	return valid
}

// Convert user input of Y or N to boolean
func ConvertBinaryChoice(userInput string) bool {
	choice := false

	if userInput == "Y" {
		choice = true
	}

	return choice
}

// Method to track whose turn it is
func NextActivePlayer(currentPlayer int, numPlayers int) int {
	// Declare return variable equal to current player
	next := currentPlayer

	// Increment to next player
	next += 1

	// If player is the last player return to first player
	if next > (numPlayers - 1) {
		next = 0
	}

	return next
}

// Method to determine number of dice to roll next round
func DiceLeft(startDice *int, usedDice int) int {
	// Remaining dice to roll is the dice that the round started with minus dice held for points
	dice := *startDice - usedDice
	// If all the dice from round start are held then the player gets to roll 6 dice again
	if dice == 0 {
		dice = 6
	}
	return dice
}

// Method to evaluate dice rolled for points. Returns points kept by player and numDice used
// in a rollResult struct
func FindPoints(dice *[6]int) rollResult {

	// Initialize return values
	points := 0
	numDice := 0

	// Count number of times each value is rolled
	diceCount := CountDice(dice)

	// Check for 1's
	if diceCount[0] != 0 {
		pts := 0
		d := 0
		num := diceCount[0]
		for num != 0 {
			if num > 2 {
				pts += 1000
				num -= 3
				d += 3
				continue
			}
			pts += num * 100
			d += num
			num = 0
		}
		// Ask user if they want to keep points
		fmt.Printf("\nYou rolled %d 1's worth %d points. Do you want to hold them?\n", diceCount[0], pts)
		choice := ConvertBinaryChoice(GetBinaryChoice())

		// add to active player score
		if choice {
			points += pts
			numDice += d
		}
	}

	// Check for 5's
	if diceCount[4] != 0 {
		pts := 0
		d := 0
		num := diceCount[4]
		for num != 0 {
			if num > 2 {
				pts += 500
				num -= 3
				d += 3
				continue
			}
			pts += num * 50
			d += num
			num = 0
		}
		// Ask user if they want to keep points
		fmt.Printf("\nYou rolled %d 5's worth %d points. Do you want to hold them?\n", diceCount[4], pts)
		choice := ConvertBinaryChoice(GetBinaryChoice())

		// add to active player score
		if choice {
			points += pts
			numDice += d
		}
	}

	// Check for remaining 3 of a kind
	values := [...]int{1, 2, 3, 5}
	for _, value := range values {
		if diceCount[value] > 2 {
			pts := 0
			d := 0
			num := diceCount[value]
			for num != 0 {
				if num > 2 {
					pts += 100 * (value + 1)
					num -= 3
					d += 3
					continue
				}
				num = 0
			}
			// Ask user if they want to keep points
			fmt.Printf("\nYou rolled 3 %d's worth %d points. Do you want to hold them? \n", value+1, pts)
			choice := ConvertBinaryChoice(GetBinaryChoice())

			// add to active player score
			if choice {
				points += pts
				numDice += d
			}
		}
	}
	result := rollResult{points, numDice}
	return result
}

// Method to solicit user input on how many players there will be
func GetNumPlayers() int {
	//Get number of players from user
	numPlayers := 0

	fmt.Println("\nHow many players are in this game? Choose a number between 2 & 4 : ")
	fmt.Scan(&numPlayers)

	//check for valid choice
	if !IsValidNumPlayers(numPlayers) {
		fmt.Println("\nInvalid choice.")
		numPlayers = GetNumPlayers()
	}

	return numPlayers
}

// Method to validate number of players is between 2 and 4
func IsValidNumPlayers(num int) bool {
	isValid := true

	if num < 2 {
		isValid = false
	}

	if num > 4 {
		isValid = false
	}

	return isValid
}

// Method to name each active player
func NamePlayer(index int, name string, playerList []Player) {
	playerList[index].name = name
}

func SetupGame(numPlayers int) []Player {
	// Initialize array of players
	playerList := []Player{
		Player{1, "Adam", 0, 0},
		Player{2, "Bob", 0, 0},
		Player{3, "Chris", 0, 0},
		Player{4, "Dave", 0, 0},
	}

	// name the players
	for i := 0; i < numPlayers; i++ {
		//Get name of player from user
		name := ""
		fmt.Printf("\nWhat is the name of Player %d : ", i+1)
		fmt.Scan(&name)

		// Change player name
		NamePlayer(i, name, playerList)
	}

	return playerList
}

func PlayGame() {

	// Setup the Game
	numPlayers := GetNumPlayers()
	playerList := SetupGame(numPlayers)
	activePlayer := -1

	// Initialize Dice Array & eligible dice to roll
	activeDice := [6]int{}
	numDice := MaxDice

	// Variables to track winner
	scoreToBeat := 0
	winningPlayer := 0
	lastRound := 0

	finalRound := false
	for finalRound == false {
		// State who is the active player
		activePlayer = NextActivePlayer(activePlayer, numPlayers)
		fmt.Printf("\nThe active player is Player %d : %v\n", activePlayer+1, playerList[activePlayer].name)

		// Reset dice
		numDice = MaxDice

		// Round loop
		farkle := false
		for farkle == false {
			farkle = PlayRound(&activeDice, &numDice, &playerList[activePlayer])
		}

		// Check if player has enough points to win
		if playerList[activePlayer].score >= PointsToWin {
			// Trigger final round
			finalRound = true
			// Store score to beat
			scoreToBeat = playerList[activePlayer].score
			// Remember whose turn will end the game
			lastRound = activePlayer
			// Which player is winning
			winningPlayer = activePlayer
			// Let the players know this is the final round
			fmt.Println("\nEntering the final round")
		}
	}

	gameOver := false
	for gameOver == false {
		// State who is the active player
		activePlayer = NextActivePlayer(activePlayer, numPlayers)
		fmt.Printf("\nThe active player is Player %d : %v\n", activePlayer+1, playerList[activePlayer].name)

		// Check if we are back to the player who triggered the final round and end the game if true
		if activePlayer == lastRound {
			gameOver = true
			break
		}

		// Reset dice
		numDice = MaxDice

		// Round loop
		farkle := false
		for farkle == false {
			farkle = PlayRound(&activeDice, &numDice, &playerList[activePlayer])
		}

		// Check if player has more points than the current winning player
		if playerList[activePlayer].score > scoreToBeat {
			// Record new score to beat
			scoreToBeat = playerList[activePlayer].score
			// Record new winning player
			winningPlayer = activePlayer
		}
	}

	// Declare the winner
	fmt.Printf("\n%v has won the game with a score of %d points!\n", playerList[winningPlayer].name, playerList[winningPlayer].score)

	// Farewell
	println("\nThank you for playing!")
}

// Method to determine if a players turn is over
func EndRound(points int, keptDice int, numDice int, player *Player) bool {
	farkle := true

	// If no points were rolled then the round is over
	if points == 0 {
		println("\nFarkle! Your turn is over.")
		// Reset player point count to 0
		player.ResetPoints()
		return farkle
	}

	// If all dice rolled have scored points then player must roll again
	if keptDice == numDice {
		println("\nAll dice have scored points. Roll again!")
		// Store points for next round
		player.AddPoints(points)
		fmt.Printf("\n%v now has %d points.", player.name, player.points)
		farkle = false
		return farkle
	}

	// Store points for next round
	player.AddPoints(points)
	fmt.Printf("\n%v now has %d points.\n", player.name, player.points)

	// If player score is less than 1000 and they have not rolled 1000 points they must roll again
	if player.score < 1000 {
		if player.points < 1000 {
			// Store points for next round
			println("\nYou do not have 1000 points yet. Roll again!")
			farkle = false
			return farkle
		}
	}

	// Ask if player wants to keep points and end turn
	fmt.Printf("\nDo you want to end your turn and keep %d points? There are %d dice left to roll.\n", player.points, numDice-keptDice)
	choice := ConvertBinaryChoice(GetBinaryChoice())

	// If player chooses to roll
	if choice == false {
		farkle = false
		return farkle
	}

	// Add player points to player score
	player.AddScore()
	fmt.Printf("\n%v's score is now %d.\n", player.name, player.score)
	// Reset player points to 0
	player.ResetPoints()

	return farkle
}

func PlayRound(activeDice *[6]int, numDice *int, player *Player) bool {
	// Roll Dice
	activeDice = Roll(activeDice, numDice)

	// Print Dice
	PrintActiveDice(activeDice)

	// Find points
	result := FindPoints(activeDice)

	// End of round evaluation to determine if round is over
	farkle := EndRound(result.points, result.keptDice, *numDice, player)

	// Determine dice to roll
	*numDice = DiceLeft(numDice, result.keptDice)

	return farkle
}
