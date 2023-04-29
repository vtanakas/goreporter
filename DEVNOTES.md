# Development Notes
Notes for various development topics

## Linter Utilization
Which linters included in the code base are actually used?

| Linter Name    | Has Vendored Linter | Has Strategy      | Loaded | Source path         | Remote Repo Location                                         | Repo Status (Active, Dormant, Deprecated, Can't Find)        |
| -------------- | ------------------- | ----------------- | ------ | ------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| aligncheck     | Yes                 | No                |        | linters/aligncheck  | https://gitlab.com/opennota/check                            |                                                              |
| copycheck      | Yes                 | Yes               |        | linters/copycheck   |                                                              | Can't Find?                                                  |
| countcode      | Yes                 | Yes               | Yes    | linters/countcode   |                                                              | Can't Find?                                                  |
| cyclo          | Yes                 | Yes               | Yes    | linters/cyclo       | https://pkg.go.dev/github.com/fzipp/gocyclo                  | Active                                                       |
| deadcode       | Yes                 | Yes               | Yes    | linters/deadcode    | https://github.com/tsenart/deadcode                          | Active                                                       |
| depend         | Yes                 | Yes (dependgraph) | Yes    | linters/depend      |                                                              | Can't Find, but may be based on: https://github.com/kisielk/godepgraph |
| depth          | Yes                 | Yes               | Yes    | linters/depth       | https://github.com/arthurgustin/godepth                      | Dormant                                                      |
| errorcheck     | Yes                 | No                |        | linters/errorcheck  |                                                              | Can't Find, but may be based on: https://github.com/kisielk/errcheck |
| flen           | Yes                 | No                |        | linters/flen        | https://github.com/lafolle/flen                              | Dormant                                                      |
| gofmt          | Yes                 | Yes               | Yes    | linters/gofmt       | https://pkg.go.dev/cmd/gofmt                                 | active                                                       |
| golint         | Yes                 | Yes (golint)      | Yes    | linters/golint      | https://github.com/golang/lint                               | Deprecated. Use `staticcheck` and `govet` instead.           |
| govet          | Yes                 | Yes               | Yes    | linters/govet       | https://pkg.go.dev/cmd/vet                                   | active                                                       |
| importpackages | No                  | Yes               | Yes    |                     | Local                                                        |                                                              |
| interfacer     | Yes                 | Yes               | Yes    | linters/interfacer  | https://github.com/mvdan/interfacer                          | dormant (author concerned about false positives)             |
| simplecode     | Yes                 | Yes               | Yes    | linters/simplecode  |                                                              | Can't Find.  May have been rolled into "staticcheck"<br />(https://github.com/dominikh/go-tools) |
| simpler        | Yes                 | No                |        | linters/simpler     | https://pkg.go.dev/honnef.co/go/tools@v0.4.3/simple<br />https://github.com/dominikh/go-tools | This seems to be an extension of simple check<br />(https://github.com/dominikh/go-tools) |
| spellcheck     | Yes                 | Yes               | Yes    | linters/spellcheck  | https://pkg.go.dev/github.com/client9/misspell               | dormant since 2018                                           |
| staticcheck    | Yes                 | No                |        | linters/staticcheck | https://staticcheck.io/                                      | active                                                       |
| structcheck    | Yes                 | No                |        | linters/structcheck | https://gitlab.com/opennota/check                            | dormant                                                      |
| unittest       | Yes                 | Yes               | Yes    | linters/unittest    | Local                                                        |                                                              |
| varcheck       | Yes                 | No                |        | linters/varcheck    | https://gitlab.com/opennota/check                            | dormant                                                      |
|                |                     |                   |        |                     |                                                              |                                                              |

## Linter Resources and Lists

https://github.com/golangci/awesome-go-linters

https://freshman.tech/linting-golang/

https://analysis-tools.dev/tag/go
