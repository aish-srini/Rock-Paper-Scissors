package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "flag"
        "math/rand"
//        "io"
)

// command line prompt  -> ./rps -iora=interactive -player=client -ipAddress=localhost -port=5867

var compMovesIntForm = map[int]string {0: "rock", 1: "paper", 2: "scissors"}


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

        var move string
        fmt.Scanln(&move)
        return move
}


func score(myScore int, oppScore int) string {  //current total score
        switch {
                case oppScore == 2:
                        fmt.Printf("Opponent wins, with a final score of (%d):(%d)!", oppScore, myScore)
                        os.Exit(1)
                case myScore == 2:
                        fmt.Printf("You win, with a final score of (%d):(%d)!", myScore, oppScore)
                        os.Exit(1)
                default:
                        fmt.Printf("Continue playing. Current score, you versus opponent, is (%d):(%d)", myScore, oppScore)
                        return "Game continues"
        }
        return ""
}


func rules(myMove string, oppMove string, game int, myScore int, oppScore int) string {
        switch {
        case myMove == "quit":
                os.Exit(1)
        case myMove == "":
                fmt.Println("No move given")
        case oppMove == myMove:
                game -= 1
                return "tie"
        case myMove == "rock"
                if oppMove == "paper" {
                        oppScore += 1
                        return oppMove
                } else if oppMove == "scissors" {
                        myScore += 1
                        return myMove
                }
        case myMove == "paper":
                if oppMove == "scissors" {
                        oppScore += 1
                        return oppMove
                } else if oppMove == "rock" {
                        myScore += 1
                        return myMove
                }
        case myMove == "scissors":
                if oppMove == "rock" {
                        oppScore += 1
                        return oppMove
                } else if oppMove == "paper" {
                        myScore += 1
                        return myMove
                }
        default:
                return myMove

        }

        return ""
}



func main() {
        interorauto := flag.String("iora", "filler", "interactive or automatic play?")
        player := flag.String("player", "filler", "Are you a server or a client?")
        opponent := flag.String("opponent", "filler", "Are you playing against a server or client?")


        ipAddress := flag.String("ipAddress", "169.229.50.188", "or 169.229.50.175")
        port := flag.Int("port", 5867, "8333")


        flag.Parse()

        if *player != "" && *opponent != "" && *interorauto != "" {
                if *player == "server" && *interorauto == "interactive" {
                        fmt.Println("Beginning game as an interactive server...")
                        serverint(*port)
                } else if *player == "server" && *interorauto == "automatic" {
                        fmt.Println("Beginning game as an automatic server...")
                        serverauto(*port)
                } else if *player == "client" && *interorauto == "interactive" {
                        fmt.Println("Beginning game as an interactive client...")
                        clientint(*ipAddress, *port)
                } else if *player == "client" && *interorauto == "automatic" {
                        fmt.Println("Beginning game as an automatic client...")
                        clientauto(*ipAddress, *port)
                } else {
                        fmt.Println("Please enter your player type, and how you want to play this game")
                }      

        } else {
                fmt.Println("Please select who you and your opponent are, and if you want to play interactively or automatically")
        }

}



func clientint(ipAddress string, port int) {
        ipAddressPort := fmt.Sprintf("%s:%d", ipAddress, port)
        clientConn, err := net.Dial("tcp", ipAddressPort)
        if err != nil { fmt.Println("Client Connection Error:", err)
                return
        }
        fmt.Println("Passed net.dial checkpoint")

        numGames := 3
        myScore := 0
        oppScore := 0

        for game := 0; game < numGames; game++ {
                fmt.Println("Beginning loop now!")

                myMove := askforMove()

                if _, err := clientConn.Write([]byte(myMove + "\n")); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }

                reader := bufio.NewReader(clientConn)
                recvMsgBytes, err := reader.ReadString('\n')
                if err != nil { fmt.Println("Error in reading opponent's play:", err)
                        return
                }

                fmt.Println("recvMsg checkpoint passed")

                oppMove := string(recvMsgBytes)

                fmt.Println("Rock...Paper...Scissors...GO!")

                fmt.Printf("(%d) You played (%s) and Opponent played (%s).", game, myMove, oppMove)

                fmt.Println("Who played what, just passed")
                gamewinner := rules(myMove, oppMove, game, myScore, oppScore)

                if gamewinner == "tie" { game -= 1
                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round!")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, your opponent won this round!")
                        oppScore += 1

                } else {
                        fmt.Println("Please enter a move")
                }

                score(myScore, oppScore)

        }

        clientConn.Close()
}



func clientauto(ipAddress string, port int) {
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

        for game := 0; game < numGames; game++ {
                fmt.Println("Rock...Paper...Scissors...GO!")

                recvMsg, err := reader.ReadString('\n')
                if err != nil { fmt.Println("Error in reading opponent's play:", err)
                        return
                        }

                oppMove := recvMsg
                myMove := compPlay(oppMove)

                if _, err := clientConn.Write([]byte(myMove)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }

                fmt.Printf("(%d) Player 1 played (%s) and Player 2 played (%s).", game, myMove, oppMove)

                rules(myMove, oppMove, game, myScore, oppScore)
                score(myScore, oppScore)

        }

        clientConn.Close()
}

        

func serverint(port int) {
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
        myScore := 0
        oppScore := 0

        for game:= 0; game < numGames; game++ {
                fmt.Println("Beginning loop now!")
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                oppMove := string(recvMsgBytes)

                fmt.Println("Rock...Paper...Scissors...GO!")
                fmt.Println("Choose to play either rock, paper, or scissors")

                myMove :=  askforMove()

                if _, err := serverConn.Write([]byte(myMove + "\n")); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }


                fmt.Printf("(%d) You played (%s) and Opponent played (%s).", game, myMove, oppMove)

                fmt.Println("Who played what just passed!")

                gamewinner := rules(myMove, oppMove, game, myScore, oppScore)

                if gamewinner == "tie" { game -= 1
                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round!")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, your opponent won this round!")
                        oppScore += 1

                } else {
                        fmt.Println("Technical difficulty. Conside replaying round?")
                }

                score(myScore, oppScore)
                
        }
        serverConn.Close()
}


func serverauto(port int) {
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
        myScore := 0
        oppScore := 0

        for game:= 0; game < numGames; game++ {
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                oppMove := string(recvMsgBytes)
                myMove := compPlay(oppMove)

                fmt.Printf("(%d) You played (%s) and Opponent played (%s).", game, myMove, oppMove)
                gamewinner := rules(myMove, oppMove, game, myScore, oppScore)

                if gamewinner == "tie" { game -= 1
                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round!")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, your opponent won this round!")
                        oppScore += 1

                } else {
                        fmt.Println("Technical difficulty. Consider replaying round?")
                }

                score(myScore, oppScore)


                if _, err := serverConn.Write([]byte(myMove + "\n")); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                } else {
                        fmt.Println("Message sent!")
                }

//                io.Copy(os.Stdout, serverConn)
        }
        serverConn.Close()
}

