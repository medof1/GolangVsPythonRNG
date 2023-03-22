package main
import (
	"fmt"
	"math/rand"
)
var startingBatch, startingSample, maxMartingale, sample, number, prevNumber, lose, losedoll, maxSample int
var winrate float32
var martingale, batch = 1, 1
var status, stat string


func init(){
	fmt.Println("\nDEFAULT VALUE:")
	fmt.Println("Max Martingale = 13")
	fmt.Println("Sample= 3000")
	fmt.Println("Batch= 1000 \n")
	fmt.Println("Write 0 if u want to use default value \n")
	fmt.Print("Max Martingale: ")
	fmt.Scanln(&maxMartingale)
	fmt.Print("How many sample: ")
	fmt.Scanln(&startingSample)
	fmt.Print("How many batch: ")
	fmt.Scanln(&startingBatch)
	maxSample = startingSample
	if startingSample == 0{
		maxSample = 3000
	}
	if startingBatch == 0{
		startingBatch = 1000
	}
	if maxMartingale == 0{
		maxMartingale = 13
	}
	number = 0
	prevNumber = number
}

func printing(){
	if number == 0{
		stat = "Banker"
	}else{
		stat = "Player"
	}
	fmt.Println(batch, sample, ",", stat ,", Bets no.", martingale , status)
}

func baccarat(){
	if batch <= startingBatch{
		if sample < maxSample{
			number = rand.Intn(2)
			if prevNumber != number{
				status = "LOSE"
			}else{
				status = "WIN"
			}
			sample += 1
			printing()
			prevNumber = number
			if status == "LOSE"{
				martingale += 1
			}else{
				martingale = 1
			}
		}else{
			sample = 0
			batch += 1
		}
	}
}

func main(){
	for{
		if batch <= startingBatch{
			baccarat()
			if martingale > maxMartingale{
				lose += 1
				martingale = 1
				sample = 0
				batch += 1
			}
		}else{
			winrate = ((float32(batch) - float32(lose))/float32(batch))*100
			fmt.Println("Martingale:", maxMartingale, "LOSEDOLL = ", lose, " WINRATE = ", winrate, "%")
			break
		}
	}
}