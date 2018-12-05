package db

import (
	"context"
	"fmt"
	"testing"
)

type Result string

var errChan chan error
var resChan chan Result

func TestMutiThreadSuccess(t *testing.T) {

	connectTestDB()
	defer closeTestDB()

	threadCount := 2
	errChan = make(chan error, threadCount)
	resChan = make(chan Result, threadCount)

	for i := 0; i < threadCount; i++ {
		name := fmt.Sprintf("Thread%d", i)
		go executeTestOnThread(name)
	}

	for i := 0; i < threadCount; i++ {
		select {
		case err := <-errChan:
			t.Error(err)
		case name := <-resChan:
			t.Log("Finished:", name)
		}
	}
}

func executeTestOnThread(word string) {
	ctx := context.Background()

	//ctx = connectTestDB(ctx)
	//defer closeTestDB(ctx)

	err := executeTestRepositoryOnThread(ctx, word)
	if err != nil {
		errChan <- err
		return
	}
	resChan <- Result(word)
}

func executeTestRepositoryOnThread(ctx context.Context, word string) error {

	baseWord := word + ":%d"

	loopCount := 2
	for i := 0; i < loopCount; i++ {
		var err error
		memoText, tagTitle := fmt.Sprintf(baseWord, i), fmt.Sprintf(baseWord, i)
		memoText, tagTitle, err = savaContents(ctx, memoText, tagTitle)
		if err != nil {
			return err
		}
	}

	return nil
}

func savaContents(ctx context.Context, memoText, tagTitle string) (string, string, error) {
	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		return memoText, tagTitle, err
	}

	memo, ctx, err := repoM.Save(ctx, memoText)
	if err != nil {
		repoM.Rollback(ctx)
		return memoText, tagTitle, err
	}

	tag, ctx, err := repoT.Save(ctx, tagTitle)
	if err != nil {
		repoT.Rollback(ctx)
		return memoText, tagTitle, err
	}

	ctx, err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		repoT.Rollback(ctx)
		return memoText, tagTitle, err
	}

	ctx, err = repoM.Commit(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		return memoText, tagTitle, err
	}

	m, ctx, err := repoM.Get(ctx, memo.ID)
	if err != nil {
		return memoText, tagTitle, err
	}

	t, ctx, err := repoT.Get(ctx, tag.ID)
	if err != nil {
		return memoText, tagTitle, err
	}

	return m.Text, t.Title, err
}
