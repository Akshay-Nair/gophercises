package content

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//FindAndReplacePath will take a stackTrace as argument and replace the path in it
//with the hyperlink to access the source code by clicking on it
func FindAndReplacePath(stackTrace string) (string, error) {
	data := bytes.NewBufferString("<html><body>")
	reader := bufio.NewScanner(strings.NewReader(stackTrace))
	dirSep := string(os.PathSeparator)
	hyperLinkTemplate := "&nbsp&nbsp&nbsp&nbsp&nbsp<a href=\"/debug?path=%s&line=%d\">%s</a> : %d<br>"
	expression := dirSep + "([A-Za-z1-9]+" + dirSep + "{0,1})+" + "[A-Za-z1-9]+\\.go:[0-9]+"
	re, err := regexp.Compile(expression)
	if err == nil {
		i := 0
		for reader.Scan() {
			i++
			if (i%2 == 1) && (i > 1) {
				locationDetail := re.FindString(reader.Text())
				path, line, err := extractPath(locationDetail)
				if err == nil {
					fmt.Fprintln(data, fmt.Sprintf(hyperLinkTemplate, path, line, path, line))
				}
			} else {
				fmt.Fprintln(data, reader.Text()+"<br>")
			}
		}
	}
	fmt.Fprintln(data, "</body></html>")
	return data.String(), err
}

func extractPath(location string) (string, int, error) {
	strSlice := strings.Split(location, ":")
	if len(strSlice) < 2 || len(strSlice) > 2 {
		return "", 0, errors.New("invalid argument")
	}
	val, err := strconv.Atoi(strSlice[1])
	return strSlice[0], val, err
}
