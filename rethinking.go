package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "flag"
        "math/rand"
        "time"
)

// command line prompt  -> ./rps -iora=interactive -player=client -ipAddress=localhost -port=5867

var compMovesIntForm = map[int]string {0: "rock", 1: "paper", 2: "scissors"}


func compPlay(recvMsg string) string {  //opponent's move
//        var compMovesIntForm = map[int]string {0: "rock", 1: "paper", 2: "scissors"}

        if recvMsg == "rock\n" {
                return "paper"
        } else if recvMsg == "paper\n" {
                return "scissors"
        } else if recvMsg == "scissors\n" {
                return "rock"
        } else {
                return compMovesIntForm[rand.Intn(3)]
        }
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


func rules(myMove string, oppMove string) string {
//        fmt.Println("My move is: ", myMove)
//        fmt.Println("Opponent move is: ", oppMove)     

        switch {
        case myMove == "quit":
                fmt.Println("quit")
                return "quit"
        case myMove == "":
                fmt.Println("No move given")
                return "empty"
        case (myMove + "\n") == oppMove:
                fmt.Println("tie")
                return "tie"
        case myMove == "rock":
                if oppMove == "paper\n" {
                        fmt.Println("I played rock and opp played paper")
                        return oppMove
                } else if oppMove == "scissors\n" {
                        fmt.Println("I played rock and opp played scissors")
                        return myMove
                }
                return "d"
        case myMove == "paper":
                if oppMove == "scissors\n" {
                        fmt.Println("i played paper and opp played scissors")
                        return oppMove
                } else if oppMove == "rock\n" {
                        fmt.Println("i played paper and opp played rock")
                        return myMove
                }
                return "e"
        case myMove == "scissors":
                if oppMove == "rock\n" {
                        fmt.Println("i played scissors and opp played rock")
                        return oppMove
                } else if oppMove == "paper\n" {
                        fmt.Println("i played scissors and opp played paper")
                        return myMove
                }
                return "f"
        default:
                fmt.Println("default")
                return myMove

        }

        return "h"

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
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil { fmt.Println("Error in reading opponent's play:", err)
                        return
                }


                oppMove := string(recvMsgBytes)

                fmt.Println("------------------------------")
                fmt.Println("Rock...Paper...Scissors...GO!")

                fmt.Printf("(%d) You played (%s) and Opponent played (%s)", game, myMove, oppMove)

                fmt.Println("You played: ", myMove)
                fmt.Println("Opponent played: ", oppMove)

                fmt.Println("Gamewinner about to be printed")
                gamewinner := rules(myMove, oppMove)
                fmt.Println("Winner is: ", gamewinner)

                if gamewinner == "tie" {
                        game -= 1
                        fmt.Println("It's a tie. Play round again!")

                } else if gamewinner == "quit" {
                        fmt.Println("Quitting game")
                        os.Exit(1)

                } else if gamewinner == "empty" {
                        fmt.Println("No move entered, quitting game!")
                        os.Exit(1)

                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round")
                        myScore += 1
                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, you lose this round")
                        oppScore += 1

                } else {
                        fmt.Printf("Technical difficulty. Gamewinner: %s", gamewinner)
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


        numGames := 3
        myScore := 0
        oppScore := 0

        for game := 0; game < numGames; game++ {
                fmt.Println("Beginning loop now!")

                myMove := compMovesIntForm[rand.Intn(3)]

                if _, err := clientConn.Write([]byte(myMove + "\n")); err != nil {
                fmt.Println("Send failed:", err)
                os.Exit(1)
                }

                reader := bufio.NewReader(clientConn)
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil { fmt.Println("Error in reading opponent's play:", err)
                        return
                } else {
                        fmt.Println("No error in reading opponent's play")
                }


                oppMove := string(recvMsgBytes)
                fmt.Println(myMove)
                fmt.Println(oppMove)
                fmt.Println("------------------------------")

                fmt.Println("Rock...Paper...Scissors...GO!")

                fmt.Printf("(%d) You played (%s) and Opponent played (%s)", game, myMove, oppMove)

                fmt.Println("You played: ", myMove)
                fmt.Println("Opponent played: ", oppMove)

                fmt.Println("Gamewinner about to be printed.")
                gamewinner := rules(myMove, oppMove)
                fmt.Println("Winner is: ", gamewinner)


                if gamewinner == "tie" {
                        game -= 1
                        fmt.Println("It's a tie. Play round again!")

                } else if gamewinner == "quit" {
                        fmt.Println("Quitting game")
                        os.Exit(1)

                } else if gamewinner == "empty" {
                        fmt.Println("No move entered, quitting game!")
                        os.Exit(1)

                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, you lose this round")
                        oppScore += 1

                } else {
                        fmt.Printf("Technical difficulty. Gamewinner: %s", gamewinner)
                }

                score(myScore, oppScore)

                time.Sleep(3 * time.Second)

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

//              scanner := bufio.NewScanner(os.Stdin)
//                userInput := scanner.Text()
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                oppMove := string(recvMsgBytes)

                fmt.Println("------------------------------")
                fmt.Println("Rock...Paper...Scissors...GO!")

                myMove :=  askforMove()

                if _, err := serverConn.Write([]byte(myMove + "\n")); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }


                fmt.Printf("(%d) You played (%s) and Opponent played (%s).", game, myMove, oppMove)

                fmt.Println("Gamewinner about to be printed.")
                gamewinner := rules(myMove, oppMove)
                fmt.Println("Winner is: ", gamewinner)

                if gamewinner == "tie" {
                        game -= 1
                        fmt.Println("It's a tie. Play round again!")


                } else if gamewinner == "quit" {
                        fmt.Println("Quitting game")
                        os.Exit(1)

                } else if gamewinner == "empty" {
                        fmt.Println("No move entered, quitting game!")
                        os.Exit(1)

                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, you lose this round")
                        oppScore += 1

                } else {
                        fmt.Printf("Technical difficulty. Gamewinner: %s", gamewinner)
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
                fmt.Println("Beginning loop now")

//              myMove := compMovesIntForm[rand.Intn(3)]
                recvMsgBytes, err := reader.ReadBytes('\n')
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }

                oppMove := string(recvMsgBytes)

                fmt.Println("------------------------------")
                fmt.Println("Rock...Paper...Scissors...GO!")

//              myMove := compMovesIntForm[rand.Intn(3)]
                myMove := compPlay(oppMove)

                fmt.Println(myMove)
                fmt.Println(oppMove)

                if _, err := serverConn.Write([]byte(myMove + "\n")); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                } else {
                        fmt.Println("No error in sending message")
                }

                fmt.Printf("(%d) You played (%s) and Opponent played (%s).", game, myMove, oppMove)

                fmt.Println("You played: ", myMove)
                fmt.Println("Opponent played: ", oppMove)

               fmt.Println("Gamewinner about to be printed.")
                gamewinner := rules(myMove, oppMove)
                fmt.Println("Winner is: ", gamewinner)

                if gamewinner == "tie" {
                        game -= 1
                        fmt.Println("It's a tie. Play round again!")

                } else if gamewinner == "quit" {
                        fmt.Println("Quitting game")
                        os.Exit(1)

                } else if gamewinner == "empty" {
                        fmt.Println("No move entered, quitting game!")
                        os.Exit(1)

                } else if gamewinner == myMove {
                        fmt.Println("Congratulations, you win this round")
                        myScore += 1

                } else if gamewinner == oppMove {
                        fmt.Println("Sorry, you lose this round")
                        oppScore += 1

                } else {
                        fmt.Printf("Technical difficulty. Gamewinner: %s", gamewinner)
                }

                score(myScore, oppScore)
//              time.Sleep(3 * time.Second)
        }
        serverConn.Close()
}
