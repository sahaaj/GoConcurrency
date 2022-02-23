# GoConcurrency
This repository contains two of the problems which are faced during concurrency, the Reader's Writer's problem and the Producer-Consumer problem.
The problems are solved using:
1. Channels in Go
2. Mutex
3. Wait Groups


Channels in Go:
Channels in Go can be called as pipes which allow two different Go Routines to send data back and forth and using channels we can also attain a certain level
of synchronization in go because some operations in channels are blocking. There are 2 types of channels in Go.
1. Unbuffered Channels: These are used for asynchronous communication and are blocking in nature. 
   So, if we send a value into the channel, 
   it has to be received by another go routine to unblock the sender's go routine
   ```
   syntax: a:=make(chan int)
    ```
2. Buffered Channels: These are used for Asynchronous communication and are non-blocking in nature, and we can specify their sizes.
   But, once a buffered channel is full it block any operation after it until a value has been received from that channel.
    ```
    syntax a:=make(chan int, 3)// making a buffered channel of capacity 3
    ```
   
The channels were mainly used in the Producer-Consumer problem. Where a buffered channel is made and the producers keep on producing items into that channel and consumer
consumes the data from that channel.
It should comply with the following rules:
1. Producer should not produce if the channel is full
2. Consumer should not consume if the channel is empty
