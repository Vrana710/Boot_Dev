package main

import "fmt"

const issueList = `[ 
    { 
        "id": 0, 
        "name": "Fix the thing", 
        "estimate": 0.5, 
        "completed": false 
    }, 
    { 
        "id": 1, 
        "name": "Unstick the widget", 
        "estimate": 30, 
        "completed": false 
    } 
]`

const userObject = `{
    "name": "Wayne Lagner",
    "role": "Developer",
    "remote": true
}`

func main() {
    fmt.Println("JSON Data:")
    fmt.Println(issueList)
    fmt.Println(userObject)
}
