import random

print("\nDEFAULT VALUE:")
print("Max Martingale = 13")
print("Sample= 3000")
print("Batch= 1000 \n")
print("Write 0 if u want to use default value \n")

maxMartingale = int(input('Max Martingale: '))
startingSample = int(input('How many sample: '))
startingBatch = int(input('How many batch: '))

martingale = 1
maxSample = startingSample
batch,sample,number = 1,0,0
number = random.randint(0,1)
prevNumber = number
status = ""
losedoll = 0
lose = 0

if startingSample == 0:
	maxSample = 3000
if startingBatch == 0:
	startingBatch = 1000
if maxMartingale == 0:
	maxMartingale = 13

def printing():
	global status, martingale, sample, batch, number
	if number == 1:
		stat = "Banker"
	else:
		stat = "Player"
	print(str(batch), str(sample) + ', ' + str(stat) + ', Bets no. ' + str(int(martingale)) , status)

def baccarat():
	global maxSample, startingSample, startingBatch, batch, sample, number, prevNumber, martingale, status
	if batch <= startingBatch:
		if sample < maxSample:
			number = random.randint(0,1)
			if prevNumber != number:
				status = "LOSE"
			else:
				status = "WIN"
			sample += 1
			printing()
			prevNumber = number
			if status == "LOSE":
				martingale += 1
			else:
				martingale = 1
		else:
			sample = 0
			batch += 1

while(1):
	if batch <= startingBatch:
		baccarat()
		if martingale > maxMartingale:
			lose += 1
			martingale = 1
			sample = 0
			batch += 1
	else:
		winrate = ((batch - lose)/batch)*100
		print("LOSEDOLL = ", lose, "WINRATE = ", winrate, "%")
		break