package extra

import (
	"fmt"
	"os"
	"time"
)

func Log(text string) {
	path := "C:/Games/test.txt"

	const layout = "Jan 2, 2006 at 3:04pm (MST)"

	err := appendStringToFile(path, time.Now().Format(layout)+":"+text+"\r\n")
	if err != nil {
		fmt.Println("Log error: ", err)
	}
}

/**
 * Append string to the end of file
 *
 * path: the path of the file
 * text: the string to be appended. If you want to append text line to file,
 *       put a newline '\n' at the of the string.
 */
func appendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
