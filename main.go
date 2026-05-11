package main
import ("encoding/json";"fmt";"io";"os";"strconv";"strings")
func main() {
	data, _ := io.ReadAll(os.Stdin)
	var v any
	if err := json.Unmarshal(data, &v); err != nil { fmt.Fprintln(os.Stderr,"JSON error:",err); os.Exit(1) }
	result := make(map[string]any); flatten("", v, result)
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
func flatten(prefix string, v any, result map[string]any) {
	switch tv := v.(type) {
	case map[string]any:
		for k, v2 := range tv {
			p := k
			if prefix != "" { p = prefix + "." + k }
			flatten(p, v2, result)
		}
	case []any:
		for i, v2 := range tv {
			p := prefix + "[" + strconv.Itoa(i) + "]"
			flatten(p, v2, result)
		}
	default:
		result[prefix] = v
	}
}
