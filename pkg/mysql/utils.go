package mysql

import "fmt"

type Uri struct {
	ServerUri    string
	Params       string
	DatabaseName string
}
func (uri *Uri) GetFullUri() string {
	return fmt.Sprintf(`%s/%s?%s`, uri.ServerUri, uri.DatabaseName, uri.Params)
}
func (uri *Uri) GetServerUri()string {
	return fmt.Sprintf(`%s/?%s`, uri.ServerUri, uri.Params)
}