package chessboard

type File []bool
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0

	for _, value := range cb[file] {
		if value {
			occupiedSquares++
		}
	}

	return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedSquares := 0

	for _, file := range cb {
		for i, occupied := range file {
			if rank-1 == i && occupied {
				occupiedSquares++
			}
		}
	}

	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0

	for _, file := range cb {
		for range file {
			squares++
		}
	}

	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0

	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				occupiedSquares++
			}
		}
	}

	return occupiedSquares
}
