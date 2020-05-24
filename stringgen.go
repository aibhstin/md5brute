package md5brute

import "errors"

func MutateCharArray(array []rune, cursor int) ([]rune, error) {
    maxIndex := len(array) - 1 // the maximum index of the array
    newArray := array // copies the array, or rather a pointer to it
    if maxIndex < cursor { // if the cursor exceeds the maximum index possible
        return nil, errors.New("Cursor must not exceed slice length")
    }

    currentChar := array[cursor] // save the current character at cursor position
    nextChar := MutateRune(currentChar) // creates a new character based on the old one
    if currentChar == 'f' && nextChar == 'a' {
        if cursor != maxIndex {
            newArray[cursor] = nextChar
            return MutateCharArray(newArray, cursor + 1)
        } else {
            newArray[cursor] = nextChar
            newArray = append(newArray, 'a')
            return newArray, nil
        }
    } else {
        newArray[cursor] = nextChar
        return newArray, nil
    }
}

func MutateRune(r rune) rune {
    switch r {
        case 'a':
            return 'b'
        case 'b':
            return 'c'
        case 'c':
            return 'd'
        case 'd':
            return 'e'
        case 'e':
            return 'f'
        case 'f':
            return 'g'
        case 'g':
            return 'h'
	case 'h':
            return 'i'
	case 'i':
            return 'j'
	case 'j':
            return 'k'
	case 'k':
            return 'l'
	case 'l':
            return 'm'
	case 'm':
            return 'n'
	case 'n':
            return 'o'
	case 'o':
            return 'p'
	case 'p':
            return 'q'
	case 'q':
            return 'r'
	case 'r':
            return 's'
	case 's':
            return 't'
	case 't':
            return 'u'
	case 'u':
            return 'v'
	case 'v':
            return 'w'
	case 'w':
            return 'x'
	case 'x':
            return 'y'
	case 'y':
            return 'z'
	case 'z':
            return 'A'
	case 'A':
            return 'B'
	case 'B':
            return 'C'
	case 'C':
            return 'D'
	case 'D':
            return 'E'
	case 'E':
            return 'F'
	case 'F':
            return 'G'
	case 'G':
            return 'H'
	case 'H':
            return 'I'
	case 'I':
            return 'J'
	case 'J':
            return 'K'
	case 'K':
            return 'L'
	case 'L':
            return 'M'
	case 'M':
            return 'N'
	case 'N':
            return 'O'
	case 'O':
            return 'P'
	case 'P':
            return 'Q'
	case 'Q':
            return 'R'
	case 'R':
            return 'S'
	case 'S':
            return 'T'
	case 'T':
            return 'U'
	case 'U':
            return 'V'
	case 'V':
            return 'W'
	case 'W':
            return 'X'
	case 'X':
            return 'Y'
	case 'Y':
            return 'Z'
	case 'Z':
            return 'a'
    }
    return 'a'
}
