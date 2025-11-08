package chessboard
import "slices"

type File []bool

type Chessboard map[string]File
// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    accepted := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    if !slices.Contains(accepted, file) {
        return 0
    }
	sumOccupied := 0
   	for i := 0; i < len(cb[file]); i++ {
        if cb[file][i] {
            sumOccupied++
        }
    }
    return sumOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > 8 || rank < 1 {
        return 0
    }
    sumOccupied := 0
    for _, v := range cb {
        if v[rank-1] {
            sumOccupied++
        }
    }
    return sumOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb["A"]) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sumOccupied := 0
    for k, _ := range cb {
        sumOccupied += CountInFile(cb, k)
    }
    return sumOccupied
}
