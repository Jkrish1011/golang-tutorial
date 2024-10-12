package concurrency

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func test2(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func TestSelectInChannels(t *testing.T) {
	rand.Seed(0)
	test2(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test2(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	// Spin of a goroutine
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			// create a new token channel
			done := make(chan struct{})
			// get the current sms and email
			s := sms[i]
			e := emails[i]
			// calculate random wait times
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			// Spin of another gorountine
			go func() {
				time.Sleep(t1)
				// send data to sms channel
				chSms <- s
				// Send to token channel
				done <- struct{}{}
			}()
			// Spin of another gorountine
			go func() {
				time.Sleep(t2)
				// send data to email channel
				chEmails <- e
				// Send to token channel
				done <- struct{}{}
			}()
			// Receive the tokens
			// So this will block the execution until the current sms and emails are being sent to the channel
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		// close both the channels
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

/*
Tickers
-------

time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
time.After() sends a value once after the duration has passed.
time.Sleep() blocks the current goroutine for the specified amount of time.
*/
func TestSaveBackups(t *testing.T) {
	expectedLogs := []string{
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"All backups saved!",
	}

	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println(actualLog)
		actualLogs = append(actualLogs, actualLog)
	}

	if !slices.Equal(expectedLogs, actualLogs) {
		t.Errorf(`
---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets2(expectedLogs), sliceWithBullets2(actualLogs))
	} else {
		fmt.Printf(`
---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets2(expectedLogs), sliceWithBullets2(actualLogs))
	}
}

func sliceWithBullets2[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}
