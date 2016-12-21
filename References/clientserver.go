# References to be used for the RPS Assignment

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
        ednamode := flag.String("flagname", "Activating EDNAMode", "clientserver")

        ipAddress := flag.String("ipAddress", "169.229.50.178", "input ip address")
        port := flag.Int("port", 5888, "input port number")
        flag.Parse()

        if *ednamode != "" {
                if *ednamode == "client" {
                        fmt.Println("Initiating client mode")
                        client(*ipAddress, *port)
                } else if *ednamode == "server" {

fmt.Println("Initiating server mode")
                        server(*port)
                }
        } else {
                fmt.Println("Please enter either 'client' or 'server'!")
        }
}

func client(ipAddress string, port int) {
        IPAddressPort := fmt.Sprintf("[%s]:%d", ipAddress, port)
        clientConn, err := net.Dial("tcp", IPAddressPort)
        if err != nil {
                fmt.Println("Error:”, err)
                return
        }
        fmt.Println("Completed")
        reader := bufio.NewReader(clientConn)
        sendMsg := "pingpong client\n"
        numIters := 100
        for i := 0; i < numIters; i++ {
            fmt.Printf("(%d) Sending: %s", i, sendMsg)
                if _, err := clientConn.Write([]byte(sendMsg)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }
                recvMsg, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println("Error:”, err)
                        return
                }
                fmt.Printf("(%d) Recieved: %s", i, recvMsg)

                time.Sleep(5 * time.Second)

        }

        clientConn.Close()

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

        numIters := 100
        sendMsg := "pingpong server"

        for i:= 0; i < numIters; i++ {
                recvMsgBytes, err := reader.ReadBytes(‘\n’)
                if err != nil {
                        fmt.Println("Receive failed", err)
                        os.Exit(1)
                }
                fmt.Printf("(%d) Recieved: %s", i, string(recvMsgBytes))

                fmt.Printf("(%d) Sending: %s\n", i, sendMsg)
                if _, err := serverConn.Write([]byte(sendMsg)); err != nil {
                        fmt.Println("Send failed:", err)
                        os.Exit(1)
                }
        }

        serverConn.Close()
}
