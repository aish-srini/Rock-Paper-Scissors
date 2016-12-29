///func main, func client, func server
/// questions: should the strategy be automatically carried out by the program, from the beginning move throughout the rest of the game, or should the user actually input each move? Especially when playing against a server?
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
Rules: 
        1. best 2 out of 3
        2. Your code should be able to act as either a client or server.
        3. your rock-paper-scissors program should be able to accept input from the command line, so that a human can play 
           interactively, and it should also be able to play automatically, submitting moves according to some strategy you
           have coded up. 
        4. efficiently handle errors 

Tips:
        1. use your knowledge of socket programming to complete the task
        2. work out a common protocol (set of rules for the format and order of the messages that are exchanged between the two sides)
                - must agree upon things such as which of the three rounds you are on, who made which moves, who won, etc.
 */


func main() {
        playerType := flag.String("player", "Beginning game now...", "Are you a computer or a human?")
        chooseOpponent := flag.String("opponent", "Beginning game...", "Are you playing a computer or a human?")
        
        
//         ipAddress := flag.String("ipAddress", "169.229.50.178", "INPUT IP ADDRESS")
//         port := flag.Int("port", 2003, "INPUT PORT NUMBER")

//        
//         flag.Parse()
        

        if *chooseOpponent != "" and *playerType != "" {
                if *playerType == "human" {
                        if *chooseOpponent == "computer" {
                                fmt.Println("Beginning game...")
                                client(*ipAddress, *port)   //human vs computer, using default port and ipaddress
                        } else if *chooseOpponent == "human" {
                                client(/* INSERT IP ADDRESS */, /* INSERT PORT */)  //human vs other computer, using alternate ipaddress and port!
                 
                } else if *playerType == "computer" {
                        if *chooseOpponent == "human" {
                                fmt.Println("Waiting for human player...")
                                server(*port)   //computer will act as a server, and respond to the client human's (will change based on port # flag provided) moves
                        }
                } else { 
                        fmt.Println("Please enter who you are, so that the game can begin.")
        } else {
                fmt.Println("Please enter your player type (human, computer) and opponent type (human, computer)!!")
        }
}

func client(ipAddress string, port int) {
        clientConn, err := net.Dial("tcp", IPAddressPort)
        if err != nil {
                fmt.Println("Client Connection Error:”, err)
                return
        }
                            
        reader := bufio.NewReader(clientConn)
        numGames := 3
        myScore := 0
        oppScore := 0
                            
        for i := 0; i < numGames; i++ {
                recvMsg, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println("Error:”, err)
                        return
                        }
                                    
                // insert function to allow player to select choose move!!          
                                            
//                 if oppMove == nil {
//                         myMove := "scissors\n"
//                 } else if oppMove == "scissors" {
//                         myMove := "paper\n"
//                 } else if oppMove == "paper" {
//                         myMove := "rock\n"
//                 } else if oppMove == "rock" {
//                         myMove := "scissors\n"
//                 }    
                   myMove := askforMove()
                                    
                   switch {
                   case oppMove == nil:
                           return myMove
                   case oppMove == myMove:
                           fmt.Println("Tie! Replay round.")
                           i -= 1
                   case myMove == "rock" && oppMove == "paper", myMove == "paper" && oppMove == "scissors", myMove == "scissors" && oppMove == "rock":
                           oppScore += 1
                   default:
                           myScore += 1
                   }
                                    
                   fmt.Printf("(%d) Player 1 played (%s) and Player 2 played (%s).", i, myMove, oppMove)
                                    
                   finalStance(myScore, oppScore)

                                   
                if _, err := clientConn.Write([]byte(myMove)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }
                
                finalStance(myScore, oppScore)

                /*
                 insert strategy for playing, and depending on whatever the opponent plays, initialize "sendMsg" to what could beat that!!
                if blah blah:
                   sendMsg := blah blah rps
                elif blah blah:
                   sendMsg := blah blah rps
                then, continue on by printing and sending that message!!!
                */
                }                    
                            
        clientConn.Close()

}

                                    
                                    
func askforMove() {
        fmt.Println("Please choose whether to play 'rock', 'paper', or 'scissors'.")
        move := flag.String("player", "random", "Choice of rock, paper, or scissors")
        
        if *move != "rock" or *move != "paper" or *move !- "scissors {
                fmt.Println("Please select a move ['rock', 'paper', or 'scissors']")
        } else {   //regardless of the move chosen, you want to return it as "whichMove", mentioned above in the client function
                return *move
        }
}
                                    }
                                   
func finalstance(myScore, oppScore string) string {
        switch {
                case oppscore == nil and myMove == nil:
                        fmt.Println("Game has not begun! No score reported (0:0)")
                case oppScore == myScore:
                        fmt.Printf("It's a tie! Play one more round for final score. Player 1 has (%s) points and Player 2 has (%s) points", myScore, oppScore)
                case oppScore == 2 && myScore == 0, oppScore == 2 && myScore == 1, oppScore == 3 && myScore == 0:
                        fmt.Println("Player 2 wins!")
                case myScore == 2 && oppScore == 0, oppScore == 1 && myScore == 2, oppScore == 0 && myScore == 3:
                        fmt.Println("Player 1 wins!")
                default:
                        fmt.Println("Continue playing, till three rounds have been completed")
        }
}
                
func server(port int) {
        //message received is (myMove) in byte form, message then sent through server is the (oppMove) sent in String form!!
        portString := fmt.Sprintf(":%d", port)
        ln, err := net.Listen("tcp", portString)
        if err != nil {
                fmt.Println("Listen failed:", err)
                os.Exit(1)
        }

        serverConn, err := ln.Accept()
        if err != nil {
                fmt.Println("Accept failed:", err)
                os.Exit(1)
        }

        reader := bufio.NewReader(serverConn)
        
        numGames := 3
        
        for i:= 0; i < numGames; i++ {
                recvMsgBytes, err := reader.ReadBytes(‘\n’)
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }
                
                fmt.Printf("(%d) Recieved: %s", i, string(recvMsgBytes))
               
                if string(recvMsgBytes) == nil {
                        sendMsg := "scissors\n"
                } else if string(recvMsgBytes) == "scissors" {
                        sendMsg := "paper\n"
                } else if string(recvMsgBytes) == "paper" {
                        sendMsg := "rock\n"
                } else if string(recvMsgBytes) == "rock" {
                        sendMsg := "scissors\n"
                }
                           
                fmt.Printf("(%d) Sending: %s\n", i, sendMsg)
                if _, err := serverConn.Write([]byte(sendMsg)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }
        }
        serverConn.Close()
}

                           

