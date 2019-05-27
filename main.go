package main

import	(
  "github.com/mattn/go-isatty"
  "os"
  "fmt"
	"io/ioutil"
  "strings"
)

var (
  allowsPrompts     = false
  ceo_ks_file       string
  ceo_pwd_file      string

  agent * tg_agent
  contract * contract_interface
)

func getEnv(key, defaultValue string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return defaultValue
}

func init() {

  if isatty.IsTerminal(os.Stdout.Fd()) {
    allowsPrompts = true
  } else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
    allowsPrompts = true
  }

  dir := getEnv("AGENT_SESSION_DIR", "./session")
  agent = NewTgAgent(dir)

  node := getEnv("ETH_NODE_ADDRESS", "")
  addr := getEnv("ETH_CONTRACT_ADDRESS", "")
  contract = NewContractInterface(node, addr, agent)

  ceo_ks_file = getEnv("ETH_CONTRACT_CEO_KEYSTORE", "./ceo.json")
  if len(ceo_ks_file) == 0 {
    fmt.Println("No CEO keystore filename specified.")
    os.Exit(1)
  }

  ceo_pwd_file = getEnv("ETH_CONTRACT_CEO_PASSWORD", "./ceo_password.txt")
  if len(ceo_pwd_file) == 0 {
    fmt.Println("No CEO keystore [password] filename specified.")
    os.Exit(1)
  }
}

func main() {

  // ============================= CONTRACT & NODE =============================

  if err := contract.setup(); err != nil {
    fmt.Printf("An error occured during contract [setup] %s\n", err.Error())
    os.Exit(1)
  }

  // ============================ CONTRACT CEO =================================

  var err error
	var dat, ceo_keystore []byte

	ceo_keystore, err = ioutil.ReadFile(ceo_ks_file)
  if err != nil {
    fmt.Printf("Error reading ceo_keystore file %s: %s\n", ceo_ks_file, err.Error())
    os.Exit(1)
  }

  dat, err = ioutil.ReadFile(ceo_pwd_file)
  if err != nil {
    fmt.Printf("Error reading ceo_password file %s: %s\n", ceo_pwd_file, err.Error())
    os.Exit(1)
  }

  ceo_password := string(dat)
  ceo_password = strings.Trim(ceo_password, "\n ")
  if len(ceo_password) == 0 {
    fmt.Println("Empty keystore password read")
    os.Exit(1)
  }

  if err := contract.setup_ceo(ceo_keystore, ceo_password); err != nil {
    fmt.Printf("An error occured during contract [setup_ceo] %s\n", err.Error())
    os.Exit(1)
  }

  // ============================== TG AGENT ===================================

  if err := agent.setup(allowsPrompts); err != nil {
    fmt.Printf("An error occured during agent [setup] %s\n", err.Error())
    os.Exit(1)
  }

  // ========================= BACKGROUND ROUTINES =============================

  go contract.poll_loop()
  go agent.update_loop()

  /* run forever */ for  {}
}
