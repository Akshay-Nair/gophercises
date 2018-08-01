package fileHandle

// func TestNegFile(t *testing.T) {
// 	getHomeDir = func() (string, error) {
// 		return "", errors.New("error occured")
// 	}
// 	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 	oldStdout := os.Stdout
// 	os.Stdout = file
// 	temp, _ := GetSecret("abc")
// 	content, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	val, _ := regexp.Match("error occured", content)
// 	assert.Equalf(t, val, true, "they should be equal")
// 	file.Truncate(0)
// 	file.Seek(0, 0)
// 	file.Close()
// 	os.Stdout = oldStdout
// 	fmt.Println(content, temp)
// }
