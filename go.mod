module github.com/charmbracelet/charm

go 1.22

require (
	github.com/auth0/go-jwt-middleware/v2 v2.2.1
	github.com/caarlos0/env/v6 v6.10.1
	github.com/calmh/randomart v1.1.0
	github.com/charmbracelet/bubbles v0.18.0
	github.com/charmbracelet/bubbletea v0.25.0
	github.com/charmbracelet/keygen v0.5.0
	github.com/charmbracelet/lipgloss v0.10.0
	github.com/charmbracelet/log v0.3.1
	github.com/charmbracelet/ssh v0.0.0-20240301204039-e79ff702f5b3
	github.com/charmbracelet/wish v1.3.2
	github.com/dgraph-io/badger/v3 v3.2103.5
	github.com/go-jose/go-jose v2.6.3+incompatible
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/uuid v1.6.0
	github.com/jacobsa/crypto v0.0.0-20190317225127-9f44e2d11115
	github.com/mattn/go-isatty v0.0.20
	github.com/meowgorithm/babylogger v1.2.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/muesli/go-app-paths v0.2.2
	github.com/muesli/mango-cobra v1.2.0
	github.com/muesli/reflow v0.3.0
	github.com/muesli/roff v0.1.0
	github.com/muesli/sasquatch v0.0.0-20220506032543-a98cc9b4d8ec
	github.com/muesli/toktok v0.1.0
	github.com/prometheus/client_golang v1.19.0
	github.com/spf13/cobra v1.8.0
	github.com/tursodatabase/libsql-client-go v0.0.0-20240220085343-4ae0eb9d0898
	goji.io v2.0.2+incompatible
	golang.org/x/crypto v0.21.0
	golang.org/x/sync v0.6.0
	modernc.org/sqlite v1.29.3
)

replace github.com/libsql/sqlite-antlr4-parser => github.com/developing-today-forks/sqlite-antlr4-parser v0.0.0-20240310060246-cdb695e0ba74
replace github.com/tursodatabase/libsql-client-go => github.com/developing-today-forks/libsql-client-go v0.0.0-20240310053440-e5fba2a8eb64

require (
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/charmbracelet/x/errors v0.0.0-20240307183742-ad8dadc14f31 // indirect
	github.com/charmbracelet/x/exp/term v0.0.0-20240307183742-ad8dadc14f31 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/creack/pty v1.1.21 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v24.3.7+incompatible // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jacobsa/oglematchers v0.0.0-20150720000706-141901ea67cd // indirect
	github.com/jacobsa/oglemock v0.0.0-20150831005832-e94d794d06ff // indirect
	github.com/jacobsa/ogletest v0.0.0-20170503003838-80d50a735a11 // indirect
	github.com/jacobsa/reqtrace v0.0.0-20150505043853-245c9e0234cb // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/libsql/sqlite-antlr4-parser v0.0.0-20230802215326-5cb5bb604475 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/mango v0.2.0 // indirect
	github.com/muesli/mango-pflag v0.1.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.50.0 // indirect
	github.com/prometheus/procfs v0.13.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xrash/smetrics v0.0.0-20231213231151-1d8dd44e695e // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/go-jose/go-jose.v2 v2.6.3 // indirect
	modernc.org/gc/v3 v3.0.0-20240304020402-f0dba7c97c2b // indirect
	modernc.org/libc v1.43.1 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
	nhooyr.io/websocket v1.8.10 // indirect
)
