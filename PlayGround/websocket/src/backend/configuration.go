
package backend

func QueryServerName(id int) string {

    var serverName string
    if id == 0 {
        serverName = "Server I"

    } else if id == 2 {
        serverName = "Server II"
    } else {
        serverName = "Unknow Server"
    }

    return serverName
}