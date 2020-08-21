package proj2

import (
	"testing"
	"reflect"
	"github.com/cs161-staff/userlib"
	_ "encoding/json"
	_ "encoding/hex"
	_ "github.com/google/uuid"
	_ "strings"
	_ "errors"
	_ "strconv"
)

func clear() {
	// Wipes the storage so one test does not affect another
	userlib.DatastoreClear()
	userlib.KeystoreClear()
}

func TestAllFunctionalities(t *testing.T) {
	clear()
	//File sharing functionality test
	userlib.SetDebugStatus(true)

  //Initializing user 'gen'
  g, err := InitUser("gen", "I love boba")

	//Creating and storing file
  file := []byte("Thanks for checking out my file sharing project!")
	g.StoreFile("test_file", file)

	//Downloading the stored file from the client
  test, err1 := g.LoadFile("test_file")

	//Editing file to add a new line
  extra := []byte(" If you have any questions, email me and let me know!")
  g.AppendFile("test_file", extra)

	//Downloading the edited stored file from the client
  test1, err2 := g.LoadFile("test_file")

  //Initializing new user 'sachi' to share files with
  s, err3 := InitUser("sachi", "fubar")

  //Sharing user gen's test file with user 'sachi'
  magic_string, err4 := g.ShareFile("test_file", "sachi")
	err5 := s.ReceiveFile("file_shared", "gen", magic_string)

	//Downloading sachi's shared version of the file from the client
  file_sachi, err6 := s.LoadFile("file_shared")

  //Revoking user sachi's access to user gen's file
  err7 := g.RevokeFile("test_file", "sachi")

	//Attempting to download sachi's version of the file after she was revoked (should fail)
  _, err8 := s.LoadFile("file_shared")

  t.Log("End of file sharing client test")
}

func TestAllFunctionalities(t *testing.T) {
	clear()
	t.Log("File sharing functionality test")
	userlib.SetDebugStatus(true)

  t.Log("Initializing user 'gen'")
  g, err := InitUser("gen", "I love boba")
	if err != nil {
		t.Error("Failed to create user 'gen'", err)
		return
	} else {
		t.Log("Successfully created user 'gen'")
	}

  file := []byte("Thanks for checking out my file sharing project!")
  t.Log("Created and storing test file with contents: ", string(file))
	g.StoreFile("test_file", file)

  test, err1 := g.LoadFile("test_file")
	if err1 != nil {
		t.Error("Failed to upload and download file", err1)
		return
	} else {
    t.Log("Uploaded and downloaded test file successfully")
  }

  t.Log("Checking file contents for correctness...")
	if !reflect.DeepEqual(test, file) {
		t.Error("Downloaded file is not the same", test, file)
		return
	} else {
    t.Log("Downloaded file contents is correct:", string(test))
  }

  extra := []byte(" If you have any questions, email me and let me know!")
  t.Log("Editing test file to add new line:", string(extra))
  g.AppendFile("test_file", extra)

  test1, err2 := g.LoadFile("test_file")
	if err2 != nil {
		t.Error("Failed to upload and download", err2)
		return
	} else {
    t.Log("Appended and downloaded test file successfully")
  }

  new_file := []byte("Thanks for checking out my file sharing project! If you have any questions, email me and let me know!")
	if !reflect.DeepEqual(test1, new_file) {
		t.Error("Downloaded file is not the same", string(test1), string(new_file))
		return
	} else {
		t.Log("Downloaded file is correct:", string(test1))
	}

  t.Log("Initializing new user 'sachi' to share files with")
  s, err3 := InitUser("sachi", "fubar")
	if err3 != nil {
		t.Error("Failed to initialize user 'sachi'", err3)
		return
	} else {
    t.Log("Successfully created user 'sachi'")
  }

  t.Log("Sharing user gen's test file with user 'sachi'")
  magic_string, err4 := g.ShareFile("test_file", "sachi")
	if err4 != nil {
		t.Error("Failed to share the file", err4)
		return
	}

	err5 := s.ReceiveFile("file_shared", "gen", magic_string)
	if err5 != nil {
		t.Error("Failed to receive the share message", err5)
		return
	} else {
    t.Log("Successfully shared test file from user 'gen' with user 'sachi'")
  }

  file_sachi, err6 := s.LoadFile("file_shared")
	if err6 != nil {
		t.Error("Failed to download the file after sharing", err6)
		return
	} else {
    t.Log("Successfully downloaded shared file")
  }
	if !reflect.DeepEqual(new_file, file_sachi) {
		t.Error("Shared file is not the same", new_file, file_sachi)
		return
	} else {
    t.Log("Shared file contents match original contents:", string(new_file))
  }

  t.Log("Revoking user sachi's access to user gen's test file")
  err7 := g.RevokeFile("test_file", "sachi")
	if err7 != nil {
		t.Error("Failed to revoke 'sachi'", err7)
		return
	} else {
    t.Log("Revoke successful")
  }

  _, err8 := s.LoadFile("file_shared")
	if err8 != nil {
		t.Log("User 'sachi' cannot access user gen's test file after being revoked. Error message:", err8)
		return
	} else {
    t.Error("User 'sachi' could still access file after revocation")
  }

  t.Log("User 'sachi' successfully revoked.")
  t.Log("End of file sharing client test")
}
