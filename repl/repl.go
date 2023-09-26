package repl 
import ()
const PROMPT=">>"
func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)
	for{
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line := scanner.Text()
		fmt.Println(line)
	}
}