<h2 align="center"><b>Roblox <a href="https://go.dev/">Golang</a> API Wrapper.</b></h2>
<hr />
<br />
<p align="center">
<h1 align="center"><b>Action Runner Statuses</b></h1>
    <div align="center">
        <h3><u><b>Security Analysis Runners</b></u></h3>
        <p></p>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-windows.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-windows.yml/badge.svg?branch=master" alt="Production Security Analysis Windows"/></a>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-linux.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-linux.yml/badge.svg?branch=master" alt="Production Security Analysis Linux"/></a>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-mac-os.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/security-analysis-mac-os.yml/badge.svg?branch=master" alt="Production Security Analysis Mac OS"/></a>
    </div>
    <hr />
    <div align="center">
        <h3><u><b>Build Runners</b></u></h3>
        <p></p>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/build-windows.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/build-windows.yml/badge.svg?branch=master" alt="Production Build Windows"/></a>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/build-linux.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/build-linux.yml/badge.svg?branch=master" alt="Production Build Linux"/></a>
        <a style="display: block;" href="https://github.com/nkpetko/goblox/actions/workflows/build-mac-os.yml"><img src="https://github.com/nkpetko/goblox/actions/workflows/build-mac-os.yml/badge.svg?branch=master" alt="Production Build Mac OS"/></a>
    </div>
    <hr />
</p>

# Installing

```
$ go get gitub.com/nkpetko/goblox
```

Contributing
============
Run tests:

```
make test
```

Local dev: use `go mod replace` in client code to point to local checkout of
this repository.

Consider running `./script/install-git-hooks` to install local git hooks for this
project.

# LICENSE

Apache2 - see the included LICENSE file for more information
