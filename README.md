# reader
osu database file format flexible reader

#Example usage

```go
package main

import (
	"fmt"

	"github.com/l3lackShark/reader"
	types "github.com/l3lackShark/reader/types"
)

type agent struct {
	OsuDB    types.OsuDB
	ScoresDB types.ScoresDB
}

func (a *agent) parseDB() error {

	rdr := reader.New()
	err := rdr.Read("osu.db", &a.OsuDB)
	if err != nil {
		return err
	}
	err = rdr.Read("scores.db", &a.ScoresDB)
	if err != nil {
		return err
	}

	return nil
}



func main() {
	a := &agent{}
	err := a.parseDB()
	if err != nil {
		//handle error
	}
  //a.OsuDB and a.ScoresDB should be populated with data at this point
}
