package gdrv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/drive/v3"
	"github.com/pkg/errors"
	"log"
)

var (
	ErrParseClientSecretFile = errors.New("Unable to parse client secret file.")
	ErrUnableReadAuthCode = errors.New("Unable to read authorization code.")
	ErrUnableRetrieveToken = errors.New("Unable to retrieve token from web.")
	ErrUnableRetrieveSheetsClient = errors.New("Unable to retrieve Sheets Client.")
	ErrUnableRetrieveDataFromSheet = errors.New("Unable to retrieve data from sheet.")
	ErrUnableCacheOaAuthToken = errors.New("Unable to cache oauth token.")
	ErrUnableGetPathCredentialFile = errors.New("Unable to get path to cached credential file.")
)

var (
	Scopes = []string { drive.DriveMetadataReadonlyScope, sheets.SpreadsheetsReadonlyScope }
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) (*http.Client, error) {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		fmt.Printf("Unable to get path to cached credential file. %v", err)
		return nil, ErrUnableGetPathCredentialFile
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}

		err := saveToken(cacheFile, tok)
		if err != nil {
			return nil, err
		}
	}
	return config.Client(ctx, tok), nil
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, ErrUnableReadAuthCode
		//log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web %v", err)
		return nil, ErrUnableRetrieveToken
	}
	return tok, nil
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {

	workingDir := "."
	for _, name := range []string{"HOME", "USERPROFILE"} { // *nix, windows
		if dir := os.Getenv(name); dir != "" {
			workingDir = dir
		}
	}
	tokenCacheDir := filepath.Join(workingDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,  url.QueryEscape("sheets.googleapis.com-go-quickstart.json")), nil
	//usr, err := user.Current()
	//if err != nil {
	//	return "", err
	//}
	//tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	//os.MkdirAll(tokenCacheDir, 0700)
	//return filepath.Join(tokenCacheDir,  url.QueryEscape("sheets.googleapis.com-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v", err)
		return ErrUnableCacheOaAuthToken
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func RequestToken() error {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return ErrParseClientSecretFile
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, Scopes...)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return ErrParseClientSecretFile
	}

	_, err = getClient(ctx, config)
	if err != nil {
		return err
	}

	return nil
}

func ReadSpreadsheet(spreadsheetId string, readRange string) ([][]interface{}, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return nil, ErrParseClientSecretFile
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, Scopes...)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return nil, ErrParseClientSecretFile
	}

	client, err := getClient(ctx, config)
	if err != nil {
		return nil, err
	}

	srv, err := sheets.New(client)
	if err != nil {
		fmt.Printf("Unable to retrieve Sheets Client %v", err)
		return nil, ErrUnableRetrieveSheetsClient
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	//spreadsheetId := "1hy6wW1B0PkAUBu9f8bhKizGnFOXm_UUA3mDLf39iIhk"
	//readRange := "Лист1!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		fmt.Printf("Unable to retrieve data from sheet. %v", err)
		return nil, ErrUnableRetrieveDataFromSheet
	}

	return resp.Values, nil
}

func GetAllSheets(spreadsheetId string) ([]*sheets.Sheet, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return nil, ErrParseClientSecretFile
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/sheets.googleapis.com-go-quickstart.json
	config, err := google.ConfigFromJSON(b, Scopes...)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return nil, ErrParseClientSecretFile
	}

	client, err := getClient(ctx, config)
	if err != nil {
		return nil, err
	}

	srv, err := sheets.New(client)
	if err != nil {
		fmt.Printf("Unable to retrieve Sheets Client %v", err)
		return nil, ErrUnableRetrieveSheetsClient
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	//spreadsheetId := "1hy6wW1B0PkAUBu9f8bhKizGnFOXm_UUA3mDLf39iIhk"
	//readRange := "Лист1!A2:E"
	resp, err := srv.Spreadsheets.Get(spreadsheetId).Do()

	if err != nil {
		fmt.Printf("Unable to retrieve data from sheet. %v", err)
		return nil, ErrUnableRetrieveDataFromSheet
	}

	return resp.Sheets, nil
}

func Walk(fileId string) ([]*drive.File, error) {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return nil, ErrParseClientSecretFile
	}

	config, err := google.ConfigFromJSON(b, Scopes...)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return nil, ErrParseClientSecretFile
	}

	client, err := getClient(ctx, config)
	if err != nil {
		return nil, err
	}

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	pageToken := ""
	var files []*drive.File

	for {
		r, err := srv.Files.List().
			Q(fmt.Sprintf("'%s' in parents", fileId)).PageToken(pageToken).
			PageSize(20).Fields("nextPageToken, files(id, name, mimeType, size)").Do()

		if err != nil {
			return nil, err
		}

		pageToken = r.NextPageToken
		files = append(files, r.Files...)

		if err != nil {
			return nil, err
		}

		if len(pageToken) == 0 {
			break
		}
	}

	return files, nil
}