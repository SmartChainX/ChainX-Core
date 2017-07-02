# ChainX-Core
The ChainX Protocol enables  issue, exchange, and controll of blockchain network on as ChainX Sidechain . It permits a solitary substance or a gathering of associations to work a system, underpins the concurrence of various sorts of advantages, and is interoperable with other free systems.

Each network maintains a cryptographically-secured transaction log, known as a blockchain, which allows partipicants to define, issue, and transfer digital assets on a multi-asset shared ledger. Digital assets share a common, interoperable format and can represent any units of value that are guaranteed by a trusted issuer — such as currencies, bonds, securities, IOUs, or loyalty points. Each Chain Core holds a copy of the ledger and independently validates each update, or “block,” while a federation of block signers ensures global consistency of the ledger.

### Environment

Set the `CHAINX` environment variable, in `.profile` in your home
directory, to point to the root of the Chain source code repo:

```sh
export CHAINX=$(go env GOPATH)/src/chainX
```

You should also add `$CHAINX/bin` to your path (as well as
`$(go env GOPATH)/bin`, if it isn’t already):

```sh
PATH=$(go env GOPATH)/bin:$CHAINX/bin:$PATH
```

You might want to open a new terminal window to pick up the change.
Set up the database:

```sh
$ createdb core
```

Start Chain Core:

```sh
$ ./cored
```

Access the dashboard:

```sh
$ open http://localhost:1999/
```

Run tests:

```sh
$ go test $(go list ./... | grep -v vendor)
```

### Building from source

There are four build tags that change the behavior of the resulting binary:
  - `reset`: allows the core database to be reset through the api
  - `localhost_auth`: allows unauthenticated requests on the loopback device (localhost)
  - `no_mockhsm`: disables the MockHSM provided for development
  - `http_ok`: allows plain HTTP requests
  - `init_cluster`: automatically creates a single process cluster

The default build process creates a binary with three build tags enabled for a
friendlier experience. To build from source with build tags, use the following
command:

> NOTE: when building from source, make sure to check out a specific
tag to build. The `main` branch is __not considered__ stable, and may
contain in progress features or an inconsistent experience.

```sh
$ go build -tags 'http_ok localhost_auth init_cluster' chain/cmd/cored
$ go build chain/cmd/corectl
```

## Developing Chain Core

### Updating the schema with migrations

```sh
$ go run cmd/dumpschema/main.go
```

### Dependencies

To add or update a Go dependency at import path `x`, do the following:

Copy the code from the package's directory
to `$CHAINX/vendor/x`. For example, to vendor the package
`github.com/kr/pretty`, run

```sh
$ mkdir -p $CHAINX/vendor/github.com/kr
$ rm -r $CHAIN/vendor/github.com/kr/pretty
$ cp -r $(go list -f {{.Dir}} github.com/kr/pretty) 
$ rm -rf $CHAINX/vendor/github.com
```

(Note: don’t put a trailing slash (`/`) on these paths.
It can change the behavior of cp and put the files
in the wrong place.)

In your commit message, include the commit hash of the upstream repo
for the dependency. (You can find this with `git rev-parse HEAD` in
the upstream repo.) Also, make sure the upstream working tree is clean.
(Check with `git status`.)

## License

ChainX Core Developer Edition is licensed under the terms of the [GNU
Affero General Public License Version 3 (AGPL)](LICENSE).
