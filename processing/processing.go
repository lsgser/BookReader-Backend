package processing

import(
	//"fmt"
	//"io/ioutil"
	//"log"
	"os"
)

/*
	This package will be responsible for reading or writing 
	data to external files such a to a log file or a text file
*/

/*
	WriteToText will write a tokens or other
	data to a text file
*/
func WriteToTextFile(text string,path string) error{
	file,err := os.OpenFile(path,os.O_APPEND|os.O_WRONLY,0644)

	if err != nil{
		return err
	}

	defer file.Close()

	_,err = file.WriteString(text+"\n\n")

	if err != nil{
		return err
	}

	return err
}
