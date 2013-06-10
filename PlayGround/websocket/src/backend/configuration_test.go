
package backend

import "testing"

func TestQueryServerName(t *testing.T) {
    var serverId = 2
    if serverName := QueryServerName(serverId); serverName != "Server II" {
        t.Errorf("QueryServerName failed")
    }

    serverId = 3
    if serverName := QueryServerName(serverId); serverName == "Unknow Server" {
        t.Errorf("QueryServerName sucess but we should provide an valid name.")
    }
}