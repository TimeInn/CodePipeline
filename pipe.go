package main

import (
	"fmt"
	"net/http"
	"strings"

	sh "github.com/codeskyblue/go-sh"
)

var tokenStr string

func main() {
	if !CheckFileIsExist(tokenFile) {
		tokenStr = SaveToken()
		fmt.Println("Pipeline .token file not found, aoto generator a token")
	}
	tokenStr = GetToken()
	fmt.Println("Pipeline token: " + tokenStr)

	fmt.Println("Sirasagi Pipeline listening :59000")
	http.HandleFunc("/pipe/", PipeHandler)
	err := http.ListenAndServe("127.0.0.1:59000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func PipeHandler(w http.ResponseWriter, r *http.Request) {
	t := r.PostFormValue("token")
	fmt.Println("deploy token: ", t)
	project := r.PostFormValue("project")
	fmt.Println("Deploy project: ", project)
	version := r.PostFormValue("ver")
	fmt.Println("Deploy version: ", version)
	if t == "" || project == "" || t != tokenStr {
		w.WriteHeader(403)
		fmt.Fprintln(w, "")
		return
	}

	go RunDockerPipeline(project, version)

	fmt.Fprintln(w, "{ message: \"async deploy message\"}")
}

/**
 * Run async deploy
 */
func RunDockerPipeline(project string, version string) {
	projectDir := "/data/docker/" + project
	projectCmd := project
	dockerServiceName := strings.Replace(project, "/", "-", -1)
	if version != "" {
		projectCmd += ":" + version
	}
	// registry host
	commandStr := "registry.docker.io/" + projectCmd
	fmt.Println("Deploy docker service name: ", dockerServiceName)
	fmt.Println("Deploy project: ", projectCmd)
	fmt.Println("Deploy command: ", commandStr)
	fmt.Println("Deploy dir: ", projectDir)
	session := sh.NewSession()
	session.ShowCMD = true
	session.SetDir(projectDir)
	// stop old image
	session.Command("docker", "stop", dockerServiceName).Run()
	// remove old runtime image
	session.Command("docker", "rm", dockerServiceName).Run()
	// pull new image
	session.Command("docker", "pull", commandStr).Run()
	// docker-compose
	session.Command("docker-compose", "up", "-d").Run()
	// docker rm image
	session.Command("docker rmi $(docker images | grep none | awk '{print $3}')").Run()
}
