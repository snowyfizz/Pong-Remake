// main.go - A comprehensive Go starter script
package main
import "fmt"
type Project struct { Name, Version string }
func (p Project) displayInfo() { fmt.Printf("Project: %s, v%s\n", p.Name, p.Version) }
func main() {
    myProject := Project{Name: "GitHub Auto-Repo Project", Version: "1.0.0"}
    myProject.displayInfo()
    fmt.Println("\nFeatures: Structs, Methods, Loops")
    for i := 0; i < 5; i++ { fmt.Printf("  - Loop %d\n", i+1) }
}
