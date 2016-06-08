package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	adjs  = []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer", "icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient", "twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing", "broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering", "bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small", "sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black", "young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral", "restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}
	nouns = []string{"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning", "snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest", "hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly", "bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass", "haze", "mountain", "night", "pond", "darkness", "snowflake", "silence", "sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower", "wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog", "frost", "voice", "paper", "frog", "smoke", "star"}
)

func codename() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%s-%s-%d", adjs[r.Intn(len(adjs))], nouns[r.Intn(len(nouns))], r.Intn(9999-1000)+1000)
}

func createProject(args []string, gitUrl string) string {
	project := ""
	if len(args) < 2 {
		// Generating random name
		project = codename()
	} else {
		project = args[1]
	}
	fmt.Println("Generating new pen:", project)

	cloneCmd := "git"
	cloneArgs := []string{"clone", gitUrl, project}
	if cloneErr := exec.Command(cloneCmd, cloneArgs...).Run(); cloneErr != nil {
		fmt.Println("Oops! Sorry I can't do it right now :(")
		fmt.Fprintln(os.Stderr, cloneErr)
		os.Exit(1)
	}

	gitRmCmd := "rm"
	gitRmArgs := []string{"-rf", project + "/.git"}
	if gitRmErr := exec.Command(gitRmCmd, gitRmArgs...).Run(); gitRmErr != nil {
		fmt.Println("Oops! You have to remove git in your pen manually, sorry dude!")
		fmt.Fprintln(os.Stderr, gitRmErr)
		os.Exit(1)
	}
	return project
}

func main() {
	if _, giterr := exec.LookPath("git"); giterr != nil {
		println("Sorry dude! You're cool but you don't have Git!")
		println(giterr)
		os.Exit(0)
	}

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "new" {
		// Create new protopen project
		projectName := createProject(args, "http://github.com/huytd/protopen")
		fmt.Println("OK, we're done!\nDon't forget to run:\n\n   npm install\n\nIn " + projectName + " folder before you start!")
	} else if len(args) > 0 && args[0] == "iron" {
		// Create new irongo project
		projectName := createProject(args, "http://github.com/huytd/irongo")
		fmt.Println("OK, we're done!\nDon't forget to run:\n\n   make\n\nIn " + projectName + " folder before you start!")
	} else {
		fmt.Println("What's up? Don't know what to do?\nTry this:\n\n   pen new/iron something-cool-123")
	}
}
