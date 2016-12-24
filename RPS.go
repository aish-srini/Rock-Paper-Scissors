///func main, func client, func server

package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "time"
        "flag"
)

/*
 * Fill in the missing parts of this code to complete the client-server
 * implementation. You will need to complete the following high level steps:
 *   1) Parse the command line arguments, so that running "pinpong client"
 *      executes the client code, while running "pingpong server" executes
 *                      the server code.
 *       2) Complete the client function so that it sends a message to the server
 *      once every second for 100 seconds. You can hardcode the address and port
 *                      of the server.
 *       3)     Complete the server function. All it needs to do is check for messages
 *          from the client and respond with its own message. The server should
 *                      stop listenting after it has received 100 messages.
 *   4) Add the ability to specify custom client and server messages from the
 *      command line.
 */
func main() {
        playerType := flag.String("player", "Beginning game now...", "Are you a computer or a human?")
        chooseOpponent := flag.String("opponent", "Beginning game...", "Are you playing a computer or a human?")
        
//         ipAddress := flag.String("ipAddress", "169.229.50.178", "input ip address")
//         port := flag.Int("port", 2003, "input port number")
//         flag.Parse()

        if *chooseOpponent != "" and *playerType != "" {
                if *playerType == "human" {
                        if *chooseOpponent == "computer" {
                                fmt.Println("Beginning game...")
                                clientcomp(*ipAddress, *port)   //if a human chooses to play the computer, then the client code will run
                        } 
                        // else if opponent is a human!! in this case, you have to be both a client and a server! so have two more functions
                } else if *playerType == "computer" {
                        if *chooseOpponent == "human" {
                                fmt.Println("Waiting for human player...")
                                servercomp(*port)   //computer will act as a server, and respond to the client human's moves
                        }
                } else { 
                        fmt.Println("Please enter who you are, so that the game can begin.")
        } else {
                fmt.Println("Please enter your player type (human, computer) and opponent type (human, computer)!!")
        }
}

func clientcomp(ipAddress string, port int) {
        

        clientConn.Close()

}

func servercomp(port int) {

        serverConn.Close()
}



///will need to create a separate client function to connect with John's client function!!
