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
		_, _, err = savaContents(ctx, memoText, tagTitle)
		if err != nil {
			return err
		}
	}

	return nil
}

func savaContents(ctx context.Context, memoText, tagTitle string) (string, string, error) {
	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()
	repoTx := getTransactionRepositoryForTest()

	ctx, err := repoTx.Begin(ctx)
	if err != nil {
		_, _ = repoTx.Rollback(ctx)
		return memoText, tagTitle, err
	}

	memo, err := repoM.Save(ctx, memoText)
	if err != nil {
		_, _ = repoTx.Rollback(ctx)
		return memoText, tagTitle, err
	}

	tag, err := repoT.Save(ctx, tagTitle)
	if err != nil {
		_, _ = repoTx.Rollback(ctx)
		return memoText, tagTitle, err
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		_, _ = repoTx.Rollback(ctx)
		return memoText, tagTitle, err
	}

	ctx, err = repoTx.Commit(ctx)
	if err != nil {
		_, _ = repoTx.Rollback(ctx)
		return memoText, tagTitle, err
	}

	m, err := repoM.Get(ctx, memo.ID)
	if err != nil {
		return memoText, tagTitle, err
	}

	t, err := repoT.Get(ctx, tag.ID)
	if err != nil {
		return memoText, tagTitle, err
	}

	return m.Text, t.Title, err
}
