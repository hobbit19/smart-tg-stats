package main

import (
	"fmt"

	"github.com/Arman92/go-tdlib"
	"path"
	"errors"
)

var (
	ErrTgNoInteractive = errors.New("No interactive mode is allowed while it was the requirement.")
)

type tg_agent struct {
	client * tdlib.Client
}

func NewTgAgent(session_dir string) * tg_agent {
	tdlib.SetLogVerbosityLevel(1)
	tdlib.SetFilePath(path.Join(session_dir, "errors.txt"))

	// Create new instance of self.client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "187786",
		APIHash:             "e782045df67ba48e441ccb105da8fc85",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   path.Join(session_dir, "tdlib-db"),
		FileDirectory:       path.Join(session_dir, "tdlib-files"),
		IgnoreFileNames:     false,
	})

	return &tg_agent{client: client}
}

func (self * tg_agent) setup(allowsPrompts bool) error {

	var err error

	for {
		currentState, _ := self.client.Authorize()
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			fmt.Print("Enter phone: ")

			if !allowsPrompts {
				return ErrTgNoInteractive
			}

			var number string
			fmt.Scanln(&number)
			_, err = self.client.SendPhoneNumber(number)

			if err != nil {
				fmt.Printf("Error sending phone number: %v", err)
				break
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			fmt.Print("Enter code: ")

			if !allowsPrompts {
				return ErrTgNoInteractive
			}

			var code string
			fmt.Scanln(&code)
			_, err = self.client.SendAuthCode(code)

			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
				break
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
			fmt.Print("Enter Password: ")

			if !allowsPrompts {
				return ErrTgNoInteractive
			}

			var password string
			fmt.Scanln(&password)
			_, err = self.client.SendAuthPassword(password)

			if err != nil {
				fmt.Printf("Error sending auth password: %v", err)
				break
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
			fmt.Println("[Telegram] Authorization Ready! Let's rock")
			break
		}
	}

	return err
}

func (self * tg_agent) update_loop() {
	// Main loop
	for update := range self.client.GetRawUpdatesChannel(10) {
		// Show all updates
		update = update
		// fmt.Println(update.Data)
		// fmt.Print("\n\n")
	}
}

func (self * tg_agent) ProcessRequest(request contract_request, ceo contract_ceo, lock mutex) error {
	defer lock.Release()
	return errors.New("Not implemented (TODO)")
}
