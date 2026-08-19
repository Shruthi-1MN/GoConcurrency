

### Structure your responses
* Situation: Set the scene by sharing relevant details about the context and challenges faced
* Task: Describe your specific role in the situation and what was expected of you
* Action: Outline the precise steps you took to address the situation or archive the intended result
* Result: Show off the outcomes achieved explaining the impact your actions had

### Core concepts

1. Unbuffered Channels (make(chan int))
Synchronous: sender waits until receiver is ready

Both sender and receiver must be ready simultaneously

Blocking behavior:

Send blocks until another goroutine receives

Receive blocks until another goroutine sends

2. Buffered Channels (make(chan int, n))
Asynchronous up to buffer capacity

Sender blocks only when buffer is full

Receiver blocks only when buffer is empty