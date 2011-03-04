# Start with the digits 1 to 9, and alternate between sequences from previous and current iteration
sequences_a = range(1, 10)
sequences_b = []
sequences = [sequences_a, sequences_b]
#
# List prime divisors
primes = [2, 3, 5, 7, 11, 13, 17]
#
# Create new sequences 2 to 10 digits long
for i in range(2, 11):
	#
	# Examine each sequence from the last set (initially 1 to 9)
	for prev_seq in sequences[i % 2]:
		#
		# For each possible next digit...
		for next_digit in range(0, 10):
			#
			# ... check it isn't in the sequence already
			if (str(prev_seq).find(str(next_digit)) == -1):
				#
				# If not, add to end of the sequence
				new_seq = prev_seq * 10 + next_digit
				#
				# Check divisibility by prime if sequence at least 4 digits long
				if (i < 4) or (int(str(new_seq)[(i - 3):]) % primes[i - 4] == 0):
					#
					# Add any valid longer sequences to the next sequence set
					sequences[(i + 1) % 2].append(new_seq)
					#
					#print str(i) + "-digits: " + str(new_seq)
					#
					# Erase last sequence set, ready for next iteration
					sequences[i % 2] = []
#
# Now add up the valid 10-digit pandigitals...
sum = 0
#
for each_seq in sequences[1]:
	#
	sum = sum + each_seq
	#
	# ... and print the result
	print each_seq, sum
