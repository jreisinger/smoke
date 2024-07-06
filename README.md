Smoke is a simple tool for running basic sanity tests against a host. It's easy to add new tests.

## Installation

```sh
go install github.com/jreisinger/smoke@latest
```

## Usage

Create config file containing tests for one or more hosts

```json
{
  "some.host.example.com": {
    "HelmReleases": {
      "Count": 2
    },
    "HttpsGet": {
      "StatusCode": 200
    },
    "OpenPorts": ["22", "443"],
    "OsRelease": {
      "ID": "ubuntu",
      "VERSION_ID": "\"20.04\""
    },
    "PodsNotRunning": {
      "Count": 0
    }
  }
}
```

Run the tests (exit code is the number of failed tests)

```sh
❯ smoke -v
fail HelmReleases on some.host.example.com: ssh "helm ls -A": failed to run: Process exited with status 127
ok   HttpsGet on some.host.example.com
ok   OpenPorts on some.host.example.com
fail OsRelease on some.host.example.com: want VERSION_ID="22.04", got VERSION_ID="20.04"
ok   PodsNotRunning on some.host.example.com
✗ echo $?
2
```
