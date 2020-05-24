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
            return 'a'
    }
    return 'a'
}
