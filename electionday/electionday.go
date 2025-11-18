package electionday

import "fmt"


func NewVoteCounter(numberOfInitialVotes int) *int {
	return &numberOfInitialVotes
}

func VoteCount(counter *int) int{
	if counter == nil{
		return 0
	}
	return *counter
}

func IncrementVoteCount(counter *int, numberOfVotes int) {
	*counter += numberOfVotes
}

func NewElectionResult(candidateName string, votes int) *ElectionResult {

	result := ElectionResult{
		name: candidateName,
		votes: votes,
	}

	return &result

}

func DisplayResult(result *ElectionResult) string{
	newResult := ElectionResult{
		name: result.name,
		votes: result.votes,
	}
	display := fmt.Sprintf("%v (%v)", newResult.name, newResult.votes)
	return display 
}

func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}