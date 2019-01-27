package cmd

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"testing"
)

var Sample = ` Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis lacinia dapibus interdum. Phasellus malesuada quis leo pulvinar suscipit. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Vestibulum euismod diam in nulla eleifend volutpat. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Duis viverra metus ligula, eget mollis lorem lobortis a. Donec euismod ligula vestibulum libero viverra auctor accumsan vel purus. Quisque accumsan ligula nibh, sed tincidunt eros egestas sit amet. Donec iaculis elementum odio in venenatis. Mauris libero ligula, pretium sed nibh et, elementum porttitor mi. Aenean vel odio mauris. Fusce malesuada id urna sed accumsan. Pellentesque varius libero eu dolor aliquam, sit amet ultricies velit semper. Integer fringilla auctor purus eget consectetur. Sed molestie molestie maximus.

Duis my+email@somedomain.com velit eros, imperdiet nec tortor quis, rhoncus aliquam diam. Fusce elementum turpis at diam molestie mollis. Proin id sapien nunc. Sed fringilla, elit vel blandit ornare, neque dolor facilisis orci, sed maximus metus est quis nisl. Etiam diam orci, consequat vel tempus at, iaculis quis ex. In hac habitasse platea dictumst. Maecenas pulvinar feugiat ultrices. Curabitur lobortis feugiat risus, sit amet mattis eros vulputate et. Nunc tortor nisl, luctus at interdum vel, pulvinar sed orci. Nullam in erat tortor.

Nulla nec massa max@test.com dictum, lobortis ex at, ullamcorper felis. Integer finibus nec turpis eu ultricies. Proin malesuada sit amet augue quis commodo. Phasellus libero nisl, consectetur nec mi eu, consequat maximus tellus. Mauris ligula nisl, vestibulum eu lacus vitae, malesuada egestas neque. Donec condimentum molestie velit in bibendum. Quisque iaculis in dolor a commodo. Sed ullamcorper, arcu non ullamcorper laoreet, odio ex mattis nisl, non aliquam libero lectus in lectus. Cras auctor, ex vel faucibus bibendum, velit diam hendrerit arcu, eu dignissim turpis magna in est.

Etiam sed augue orci. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer porta augue velit, at ullamcorper velit elementum eu. Donec sit amet massa nulla. Mauris iaculis libero a libero luctus lobortis. In porta tortor eu risus euismod, sit amet commodo felis dignissim. Duis consequat purus vitae metus euismod congue. Quisque mattis enim et tristique pulvinar. 
test
this will not work
target@specialdomain.com`

func TestParseBadInputFile(t *testing.T) {
	_, err := parseFile("nonexistent.txt")
	if err == nil {
		t.Fatal()
	}
}

func TestParseFile(t *testing.T) {
	_, createErr := os.Create("test.txt")
	if createErr != nil {
		t.Fatal()
	}
	_, openErr := parseFile("test.txt")
	if openErr != nil {
		t.Fatal(openErr)
	}
	os.Remove("test.txt")
}

func TestParseClipboardFindsEmails(t *testing.T) {
	clipboard.WriteAll(Sample)
	found, err := parseClipboard()
	if err != nil {
		t.Fatal(err)
	}
	if len(found["email"]) < 3 {
		t.Fatal()
	}
}

func TestParseFindsEmails(t *testing.T) {
	fp, err := os.Create("file-with-emails.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()
	fp.Write([]byte(Sample))

	found, err := parseFile("file-with-emails.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(found["email"]) < 3 {
		t.Fatal()
	}

	err = os.Remove("file-with-emails.txt")
	if err != nil {
		fmt.Println("Couldn't delete the file", err)
	}
}
