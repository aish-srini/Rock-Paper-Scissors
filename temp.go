// temp file to place entire rps file (straight from ubuntu)

package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "flag"
        "math/rand"
)


func compPlay(recvMsg string) string {  //opponent's move
        var compMovesIntForm = map[int]string {0: "rock", 1: "paper", 2: "scissors"}

        if recvMsg == "rock" {
                return "paper"
        } else if recvMsg == "paper" {
                return "scissors"
        } else {
                return "rock"
        }
        return compMovesIntForm[rand.Intn(3)]
}


func askforMove() string { //your move
        fmt.Println("Choose to play 'rock', 'paper', or 'scissors'")
        move := flag.String("move", "random", "Choice of rock, paper, or scissors")

        if *move != "rock" || *move != "paper" || *move != "scissors" {
                fmt.Println("Choose 'rock', 'paper', or 'scissors' to proceed with game")
        } else {
                fmt.Println(*move)
        }

        return ""
}


func score(myScore int, oppScore int) string {  //current total score
        switch {
                case oppScore == 2:
                        fmt.Printf("Opponent wins, with a final score of (%d):(%s)!", oppScore, myScore)
                        return "Game over"
                case myScore == 2:
                        fmt.Printf("You win, with a final score of (%d):(%s)!", myScore, oppScore)
                        return "Game over"
                default:
                        fmt.Printf("Continue playing. Current score, you versus opponent, is (%d):(%s)", myScore, oppScore)
                        return "Game continues"
        }
}


func rules(myMove string, oppMove string, game int, myScore int, oppScore int) string {
        switch {
        case oppMove == myMove:
              fmt.Println("Tie! Replay round.")
              game -= 1
        case myMove == "rock" && oppMove == "paper", myMove == "paper" && oppMove == "scissors", myMove == "scissors" && oppMove == "rock":
              fmt.Println("Opponent wins this round! Play another round.")
              oppScore += 1
        case oppMove == "rock" && myMove == "paper", oppMove == "paper" && myMove == "scissors", oppMove == "scissors" && myMove == "rock":
              fmt.Println("You win this round! Play another round.")
              myScore += 1
        default:
              return myMove
        }

        return ""
}

func main() {
        playerType := flag.String("player", "Beginning game now...", "Are you a computer or a human?")
        chooseOpponent := flag.String("opponent", "Beginning game...", "Are you playing a computer or a human?")


        myipAddress := flag.String("ipAddress", "169.229.50.188", "INPUT IP ADDRESS")
        myport := flag.Int("port", 5867, "INPUT PORT NUMBER")

        johnipAddress := flag.String("johnip", "169.229.50.175", "INPUT JOHN's IP")
        johnport := flag.Int("johnport", 8333, "INPUT JOHNS PORT")

        flag.Parse()


        if *chooseOpponent != "" && *playerType != ""   {
                if *playerType == "human" {
                        if *chooseOpponent == "computer" {
                                fmt.Println("Beginning game...")
                                client(*myipAddress, *myport)
                        } else if *chooseOpponent == "human" {
                                client(*johnipAddress, *johnport) 
                        }

                } else if *playerType == "computer" {
                        if *chooseOpponent == "human" {
                                fmt.Println("Waiting for human player...")
                                server(*myport)
                        }
                } else {
                        fmt.Println("Enter who you are, so that the game can begin.")
                }
        } else {
                fmt.Println("Enter whether you're a human or computer, and what species your opponent is as well")
        }

}


func client(ipAddress string, port int) {
        ipAddressPort := fmt.Sprintf("%s:%d", ipAddress, port)
        clientConn, err := net.Dial("tcp", ipAddressPort)
        if err != nil { fmt.Println("Client Connection Error:", err)
                return
        } else {
                fmt.Println("No error with Client Connection")
        }

        reader := bufio.NewReader(clientConn)
        numGames := 3
        myScore := 0
        oppScore := 0

        for i := 0; i < numGames; i++ {
                recvMsg, err := reader.ReadString('\n')
                if err != nil { fmt.Println("Error in reading opponent's play:", err)
                        return
                        }

                myMove := askforMove()
                oppMove := compPlay(recvMsg)

                fmt.Printf("(%d) Player 1 played (%s) and Player 2 played (%s).", i, myMove, oppMove)

                rules(myMove, oppMove, i, myScore, oppScore)
                score(myScore, oppScore)
 
                if _, err := clientConn.Write([]byte(myMove)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }

                score(myScore, oppScore)

        }

        clientConn.Close()
}



func server(port int) {
        portString := fmt.Sprintf(":%d", port)
        ln, err := net.Listen("tcp", portString)
        if err != nil {
                fmt.Println("Listen failed:", err)
                os.Exit(1)
        } else {
                fmt.Println("No error found in listening")
        }
        
        serverConn, err := ln.Accept()
        if err != nil {
                fmt.Println("Accept failed:", err)
                os.Exit(1)
        } else {
                fmt.Println("Message accepted, no error found")
        }

        reader := bufio.NewReader(serverConn)

        numGames := 3

        for game:= 0; game < numGames; game++ {
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                fmt.Printf("(%d) Recieved: %s", game, string(recvMsgBytes))

                recvMsg := string(recvMsgBytes)
                sendMsg := compPlay(recvMsg)
                
                fmt.Printf("(%d) Sending: %s\n", game, sendMsg)
                if _, err := serverConn.Write([]byte(sendMsg)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }
        }
        serverConn.Close()
}

