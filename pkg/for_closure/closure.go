package for_closure //nolint:golint,nosnakecase // underscore in package name

func doClosure(values []string) []string {
	resultCh := make(chan string, len(values))
	results := make([]string, 0, len(values))
	for _, v := range values {
		go func() {
			resultCh <- v
		}()
	}

	for range values {
		results = append(results, <-resultCh)
	}

	return results
}
