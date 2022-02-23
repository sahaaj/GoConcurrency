# GoConcurrency
This repository contains two of the problems which are faced druing concurrency, the Reader's Writer's problem and the Producer-Consumer problem.
The problems are solved using:
1. Channels in Go
2. Mutex
3. Wait Groups


Channels in Go:
Channels in Go can be called as pipes which allow two different Go Routines to send data back and forth and using channels we can also attain a certain level
of synchronization in go because some operations in channels are blocking. There are 2 types of channels in Go.
1. Unbuffered Channels
2. Buffered Channels
