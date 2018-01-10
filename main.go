package main

import (
	"os"
	"compress/zlib"
	"io"
	"io/ioutil"
	"flag"
)



func decode(input io.Reader, output io.Writer) error {
	r, err := zlib.NewReader(input)
	if err != nil {
		return err
	}

	_, err = io.Copy(output, r)
	if err != nil {
		return err
	}

	r.Close()

	return nil
}

func encode(input io.Reader, output io.Writer) error {
	w := zlib.NewWriter(output)

	content, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}

	_, err = w.Write(content)

	w.Close()
	return nil
}

func main() {
	var isDecode bool
	flag.BoolVar(&isDecode, "d",false, "Uncompress the data from stdin using zlib")

	flag.Parse()

	var err error
	if isDecode {
		err = decode(os.Stdin, os.Stdout)
	} else {
		err = encode(os.Stdin, os.Stdout)
	}

	if err!= nil {
		panic(err)
	}
}