package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/atotto/clipboard"
	"github.com/kikukafandi/donat/internal/provider"
	"github.com/kikukafandi/donat/internal/session"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "donat",
		Short: "Donat CLI - Do Not Accept Trash",
		Long: `
  .d8888b.  
 d88P  Y88b 
 888    888   DONAT CLI
 888    888   "Do Not Accept Trash"
 888    888 
 Y88b  d88P 
  "Y8888P"  

Usage: donat bake (generate) | donat eat (read)`,
	}

	// bakeCmd generates a new disposable identity.
	var bakeCmd = &cobra.Command{
		Use:   "bake",
		Short: "Generate new disposable email",
		Run: func(cmd *cobra.Command, args []string) {
			client := provider.NewClient()
			fmt.Print("[*] Baking new identity... ")

			email, token, err := client.GenerateEmail()
			if err != nil {
				fmt.Printf("\n[!] Error: %v\n", err)
				return
			}

			// Persist session and copy to clipboard.
			session.Save(email, token)
			clipboard.WriteAll(email)

			fmt.Println("Done.")
			fmt.Println("\n------------------------------------------------")
			fmt.Printf("[+] Identity: \033[32m%s\033[0m\n", email)
			fmt.Println("------------------------------------------------")
			fmt.Println("[+] Copied to clipboard.")
			fmt.Println("[+] Run 'donat eat' to check inbox.")
		},
	}

	// eatCmd handles inbox listing and message reading.
	var eatCmd = &cobra.Command{
		Use:   "eat [id]",
		Short: "Check inbox or read message",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sess, err := session.Load()
			if err != nil {
				fmt.Println("[!] No active session. Run 'donat bake' first.")
				return
			}

			client := provider.NewClient()
			client.SetToken(sess.Token)

			// Case 1: List inbox.
			if len(args) == 0 {
				fmt.Printf("[*] Checking inbox for %s ...\n", sess.Email)
				msgs, err := client.GetMessages()
				if err != nil {
					fmt.Printf("[!] Failed to fetch messages: %v\n", err)
					return
				}

				if len(msgs) == 0 {
					fmt.Println("\n[-] Inbox is empty.")
					return
				}

				fmt.Println("\nINBOX:")

				// Init tabwriter for alignment.
				w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
				fmt.Fprintln(w, "ID\tSENDER\tSUBJECT")
				fmt.Fprintln(w, "--\t------\t-------")

				for _, m := range msgs {
					sub := m.Subject
					if len(sub) > 40 {
						sub = sub[:37] + "..."
					}
					fmt.Fprintf(w, "%d\t%s\t%s\n", m.ID, m.From, sub)
				}
				w.Flush()
				fmt.Println("\n[i] Type 'donat eat <ID>' to read content.")

				// Case 2: Read specific message.
			} else {
				id, _ := strconv.Atoi(args[0])
				fmt.Printf("[*] Fetching message ID %d...\n", id)

				msg, err := client.ReadMessage(id)
				if err != nil {
					fmt.Printf("[!] Failed to read message: %v\n", err)
					return
				}

				fmt.Println("\n========================================")
				fmt.Printf("FROM   : %s\n", msg.From)
				fmt.Printf("SUBJECT: %s\n", msg.Subject)
				fmt.Println("========================================")

				content := msg.TextBody
				if content == "" {
					content = msg.Body
				}

				// Remove upstream footer garbage.
				if idx := strings.Index(content, "Click on the back"); idx != -1 {
					content = content[:idx]
				}
				fmt.Println(strings.TrimSpace(content))
				fmt.Println("\n========================================")
			}
		},
	}

	// crumbsCmd clears the session.
	var crumbsCmd = &cobra.Command{
		Use:   "crumbs",
		Short: "Clear session (logout)",
		Run: func(cmd *cobra.Command, args []string) {
			session.Clear()
			fmt.Println("[+] Session cleared.")
		},
	}

	rootCmd.AddCommand(bakeCmd)
	rootCmd.AddCommand(eatCmd)
	rootCmd.AddCommand(crumbsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
