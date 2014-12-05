package kparser

import "fmt"
import "os"
import "strings"
import "kdata"

const EQUAL = '='
const DECIMAL = '.'
const ADD = '+'
const SUB = '-'
const MUL = '*'
const DIV = '/'
const LFP = '('
const RFP = ')'
const POW = '^'
const ISNUM = 0
const ISSTR = 1
const ISOPR = 2
const ISKEY = 3
const ISERR = -1

type Parser struct {
    check Stack
    char Stack
    number Stack
    operator Stack
    strings Stack
    variable Stack
    status int
}

var kparser Parser = new (Parser)

func CharType (char byte) int {
    if char <= '9' && char >= '0' || char == DECIMAL{
        return ISNUM
    } else if char == ADD || char == SUB || char == MUL || char == DIV || char == POW {
        return ISOPR
    } else {
        return ISERR
    }
}

func (p *Parser) Dissolve (input string) bool {
    for i := 0; i < len(input); i++ {
        if CharType (input[i]) == ISNUM {
            p.status = ISNUM
            p.check 
        }
    }
}

func Parse(input string) bool {

    return true
}

/**
 * Parse inputs
 */
func ParseInput (input string) bool {
    if strings.Contains(input, "=") && strings.Count (input, "=") == 1 {
        evenBracket := ParseBracket(input)
        if !evenBracket {
            fmt.Println("Unbalanced parenthesis in", input)
            return true
        }
        left := strings.Trim(strings.Split(input, "=")[0], " ")
        right := strings.Trim(strings.Split(input, "=")[1], " ")
        action := strings.Split(right, "(")[0]
        path := right[strings.Index(right, "(") + 1: strings.Index(right, ")")]
        fmt.Println("var =", left)
        fmt.Println("action =", action)
        fmt.Println("path =", path)
        return ParseExp(left, action, path)
    } else if strings.EqualFold(input, "exit") {
        fmt.Println("Bye!")
        return false
    } else {
        fmt.Printf("! \"%s\" is not a valid expression.\n", input)
        return true
    }
}

func ParseExp (left, action, path string) bool {
    var valid = true
    valid = valid && FileExist(path)
    return valid
}

func FileExist (path string) bool {
    _, erropen := os.Open(path)
    _, errcreate := os.Create(path)
    if erropen != nil && errcreate != nil {
        fmt.Printf("File \"%s\" cannot be opened or created.\n", path)
        return false
    } else {
        if erropen == nil {
            fmt.Println("File \"%s\" can be opened!")
        } else {
            os.Remove(path)
            fmt.Printf("File \"%s\" can be created!\n", path)
        }
        return true
    }
}

/**
 * Check if parentheses are balanced
 */
func ParseBracket (input string) bool {
    if strings.Contains(input, "[") || strings.Contains(input, "]") {
        return false
    } else if strings.Contains(input, "{") || strings.Contains(input, "}") {
        return false
    } else {
        left := 0
        right := 0
        for i := 0; i < len (input); i++ {
            if input[i] == '(' {
                left = left + 1
            } else if input[i] == ')' {
                right = right + 1
            }
            if right > left {
                return false
            }
        }
        return true
    }
}
