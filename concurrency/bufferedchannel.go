package concurrency

/*

A buffer allows the channel to hold a fixed number of values before sending a block.
This means sending to the channel only gets blocked if the buffer is full,
and receiving blocks if the buffer is empty.

We can provide a second argument in make() to create a buffered channel.

```
make(chan int, 100)
```

*/

// Assignment

/*
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered
channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them
out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the
emails it's given. It should then write the emails to the channel in order, and finally return the channel.
*/

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, email := range emails {
		ch <- email
	}
	return ch
}
