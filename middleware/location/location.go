package location

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//ExtractLocation function is to to extract the names of the file along the path
//from the string provided ass argument it returns the file paths along with an error if one occurs
func ExtractLocation(content string) ([]string, error) {
	var filepath []string
	dirSep := string(os.PathSeparator)
	expression := dirSep + "([A-Za-z1-9]+" + dirSep + "{0,1})+" + "[A-Za-z1-9]+\\.go:[0-9]+"
	regex, err := regexp.Compile(expression)
	if err == nil {
		filepath = regex.FindAllString(content, -1)
	}
	return filepath, err
}

//ExtractPath would fetch the line number from the the file location detail
//it would return return path, line number and an error if one occurs
func ExtractPath(location string) (string, int, error) {
	strSlice := strings.Split(location, ":")
	if len(strSlice) < 2 || len(strSlice) > 2 {
		return "", 0, errors.New("invalid argument")
	}
	val, err := strconv.Atoi(strSlice[1])
	return strSlice[0], val, err
}
