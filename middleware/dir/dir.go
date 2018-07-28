package dir

import (
	"os"
	"regexp"
)

//ExtractPathDetail function is to to extract the names of the file along the path
//from the string provided as argument it returns the file paths along with an error if one occurs
func ExtractPathDetail(content string) ([]string, error) {
	var filepath []string
	dirSep := string(os.PathSeparator)
	expression := dirSep + "([A-Za-z1-9]+" + dirSep + "{0,1})+" + "[A-Za-z1-9]+\\.go"
	regex, err := regexp.Compile(expression)
	if err == nil {
		filepath = regex.FindAllString(content, 1)
	}
	return filepath, err
}
