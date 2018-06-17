package maker_test

import (
	"os"
	"testing"

	ifttt "github.com/charliemaiors/ifttt-golang-maker"
)

const baseResponsePrefix = "Congratulations! You've fired the "
const baseResponseSuffix = " event"

var iftttValid *ifttt.IFTTTClient
var iftttInvalid *ifttt.IFTTTClient
var err error

func init() {

	if key := os.Getenv("IFTTT_KEY"); key == "" {
		panic("In order to run tests please set a valid environment variable called IFTTT_KEY")
	}

	iftttValid, err = ifttt.NewClient(os.Getenv("IFTTT_KEY"))

	if err != nil {
		panic(err)
	}

	iftttInvalid, err = ifttt.NewClient("ahsuihduahuidhasuidhauishdiau")

	if err != nil {
		panic(err)
	}
}

func TestInvalidClient(t *testing.T) {

	_, err = ifttt.NewClient("")

	if err == nil {
		t.Fatal("The test must fail beacuse there isn't any request key")
	}

	t.Logf("Test succeded, error is %v", err)

}

func TestInvalidFire(t *testing.T) {
	_, err = iftttInvalid.Do("hello_world", ifttt.Values{})

	if err != nil {
		t.Fatal("Invalid request key gets accepted")
	}

	t.Logf("Valid test, error is %v", err)
}

func TestFireEvent(t *testing.T) {
	eventName := "hello_world"
	resp, err := iftttValid.Do(eventName, ifttt.Values{
		FirstValue:  "hello",
		SecondValue: "world",
	})

	if err != nil {
		t.Fatalf("Test must not fail with a valid request key %v", err)
	}

	if resp != baseResponsePrefix+eventName+baseResponseSuffix {
		t.Fatalf("Different responses %s", resp)
	}

	t.Log("Test succedeed")
}
