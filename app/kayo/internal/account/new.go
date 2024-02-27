package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	kayoCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/rlp"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"github.com/manifoldco/promptui"
)

func NewAccount() {
	if !kayoCommon.IsInitDir() {
		color.Red("The kaiyunchain hasn't been initialized yet!")
		fmt.Println()
		return
	}
	// Get and Close DbObj
	mptDbObj := db.GetDbObj(db.MptDataPath)
	defer db.CloseDbObj(mptDbObj)
	// New and store an account
	w := wallet.NewWallet()
	passwd := ScanPasswordPrompt()
	p := db.KeystoreDataPath + "/" + w.Address.Hex()
	w.StoreKey(p, passwd)
	// Save sate to mpt
	mptBytes := db.Get([]byte("latest"), mptDbObj)
	var e []interface{}
	err := rlp.DecodeBytes(mptBytes, &e)
	if err != nil {
		log.Panic("Failed to DecodeBytes:", err)
	}
	trie := mpt.NewTrieWithDecodeData(e)
	s := core.NewState()
	err = trie.Put(w.Address.Bytes(), s.Serialize())
	if err != nil {
		log.Panic("Failed to Put:", err)
	}
	db.Set([]byte("latest"), mpt.Serialize(trie.Root), mptDbObj)
	// Prompt
	time := strings.Split(common.GetCurrentTime(), " ")
	color.Green("INFO [%s|%s] Account creation succeeded! address: %s", time[0], time[1], w.Address.Hex())
	fmt.Println()
}

func ScanPasswordPrompt() string {
	fmt.Println("Your new account is locked with a password. Please give a password. Do not forget this password.")
	pass1 := ScanPassword("Password")
	pass2 := ScanPassword("Repeat password")
	if pass1 != pass2 {
		color.Yellow("Passwords do not match!")
		fmt.Println()
		return ""
	}
	return pass1
}

func ScanPassword(s string) string {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		if !kayoCommon.IsContainsDigitAndLetter(input) {
			return errors.New("password must contain letters and digits")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    s,
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	return result
}
