
import (
	"fmt"
	"sort"
)

// Job represents a job with weight and length
type Job struct {
	Weight, Length int
}

// ByDifference sorts jobs by (weight - length) difference, and by weight if the difference is the same
type ByDifference []Job

func (a ByDifference) Len() int { return len(a) }
func (a ByDifference) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDifference) Less(i, j int) bool {
	diff1 := a[i].Weight - a[i].Length
	diff2 := a[j].Weight - a[j].Length
	if diff1 == diff2 {
		return a[i].Weight > a[j].Weight // Higher weight first if differences are the same
	}
	return diff1 > diff2 // Higher difference first
}

func main() {
	// Example list of jobs
	jobs := []Job{
		{Weight: 74, Length: 59},
		{Weight: 1, Length: 2},
		{Weight: 100, Length: 50},
		// Add more jobs as needed
	}

	// Sort jobs by difference (weight - length) in descending order
	sort.Sort(ByDifference(jobs))

	// Calculate the weighted sum of completion times
	totalCompletionTime := 0
	weightedCompletionTime := 0

	for _, job := range jobs {
		totalCompletionTime += job.Length
		weightedCompletionTime += job.Weight * totalCompletionTime
	}

	// Output the result
	fmt.Println("Total Weighted Completion Time:", weightedCompletionTime)
}
