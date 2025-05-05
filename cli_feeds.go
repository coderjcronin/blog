package main

import (
	"context"
	"fmt"
)

func commandFeeds(programState *state, args ...string) error {

	returnData, err := programState.db.ListFeedsWithCreators(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Feeds:")
	for _, record := range returnData {
		fmt.Printf("\tName:\t\t %s\n\tURL:\t\t %s\n\tCreated By:\t %s\n\n", record.Name, record.Url, record.Name_2)
	}

	return nil
}
