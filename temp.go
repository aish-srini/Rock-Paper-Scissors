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

func compPlay() {
        var compMovesIntForm = map[int]string {0: "rock", 1: "paper", 2: "scissors"}
        return compMovesIntForm[rand.Intn(3)]
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
                fmt.Println("Enter if your species and opponent type")
        }

}

func client(ipAddress string, port int) {
        clientConn, err := net.Dial("tcp", *port)
         if err != nil { fmt.Println("Client Connection Error:", err)
                return
        }

        reader := bufio.NewReader(clientConn)
        numGames := 3
        myScore := 0
        oppScore := 0

        for i := 0; i < numGames; i++ {
                recvMsg, err := reader.ReadString('\n')
                if err != nil { fmt.Println("Error:", err)
                        return
                        }

                 myMove := askforMove()
                 oppMove := compPlay()

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

                 score(myScore, oppScore)


                if _, err := clientConn.Write([]byte(myMove)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }

                score(myScore, oppScore)

                }

        clientConn.Close()
}


func askforMove() {
        fmt.Println("Choose to play 'rock', 'paper', or 'scissors'")
        move := flag.String("player", "random", "Choice of rock, paper, or scissors")

        if *move != "rock" || *move != "paper" || *move != "scissors" {
                fmt.Println("Choose 'rock', 'paper', or 'scissors' to proceed with game")
        } else {
                fmt.Println(*move)
        }

}

func score(myScore int, oppScore int) {
        switch {
                case oppScore == 0 && myScore == 0:
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
                recvMsgBytes, err := reader.ReadBytes("\n")
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                fmt.Printf("(%d) Recieved: %s", i, string(recvMsgBytes))


                if string(recvMsgBytes) == "" {
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

