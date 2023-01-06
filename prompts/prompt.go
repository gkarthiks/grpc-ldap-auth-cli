package prompts

import (
	"encoding/base64"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"grpc-ldap-auth-cli/utils"
	"grpc-ldap-auth-cli/validators"
	"os"
	"strings"
)

func PromptUser(label interface{}, validate func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func PromptPassword(label interface{}, validate func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validate,
		Mask:      '⎈',
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// PromptYesNo prompts Yes/No question
func PromptYesNo(label interface{}) (promptResult string) {
	items := []string{utils.StringYes, utils.StringNo}
	index := -1
	var err error
	for index < 0 {
		prompt := promptui.Select{
			Label: label,
			Items: items,
		}
		index, promptResult, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	}
	return
}

func ConfirmSubmitAndGenerateAuth() (encodedAuthBytes string) {
	confirmSubmit := PromptYesNo("Confirm submit the request?")
	if strings.Compare(confirmSubmit, "Yes") == 0 {
		userName := PromptUser("Enter your username:", validators.UserNameValidate)
		passwd := PromptPassword("Enter your password:", validators.PasswdLenValidate)
		encodedAuthBytes = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", userName, passwd)))
	} else {
		logrus.Info("Please start over!!! Thanks!!!")
		os.Exit(0)
	}
	return
}
