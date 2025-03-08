package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/plugins/log_stash"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.Initgorm()

	log := log_stash.New("127.0.0.1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJxcSIsIm5pY2tfbmFtZSI6Iumhvuaso-aAoSIsInJvbGUiOjEsInVzZXJfaWQiOjMsImV4cCI6MTc0MTMyMzE1NS4zNzc4MTcyLCJpc3MiOiLmiYDmgp_nmobmmK_nqboifQ.H3U523CnpEe4BiHqP10UU264NtKq0WSTjfeZVRKcXFI")
	log.Info("视频观看之141条,6:55")
}
