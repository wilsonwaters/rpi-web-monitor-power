package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func main() {
    // API spec is  monitor-power/{version}/{monitor-id}?power={on,off}
    http.HandleFunc("/monitor-power/1/1", handler)
    http.ListenAndServe(":9990", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
    powers, ok := r.URL.Query()["power"]

    if !ok || len(powers[0]) < 1 {
        fmt.Fprintln(w, "Url Param 'power' is missing")
        return
    }

    // Query()["key"] will return an array of items,
    // we only want the single item.
    power := powers[0]

    if power == "off" {
        fmt.Fprintln(w, "Turning monitor off")
 cmd := exec.Command("/bin/sh", "-c", "sudo tvservice -o")
 err := cmd.Run()
 fmt.Fprintf(w,"Command finished with error: %v", err)
    } else if power == "on" {
        fmt.Fprintln(w, "Turning monitor on")
 cmd := exec.Command("/bin/sh", "-c", "sudo tvservice -p && sudo chvt 2 && sudo chvt 7")
 err := cmd.Run()
 fmt.Fprintf(w,"Command finished with error: %v", err)
    }
}
