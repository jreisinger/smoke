Smoke is a simple tool for running basic sanity tests against one or more hosts.

## Installation

```sh
go install github.com/jreisinger/smoke@latest
```

## Usage

Create config file containing tests for one or more hosts

```json
{
  "some.host.example.com": {
    "FilesPresent": ["/etc/passwd"],
    "HelmReleases": 2,
    "HttpsGetStatusCode": 200,
    "OpenPorts": ["22", "443"],
    "OsRelease": {
      "ID": "ubuntu",
      "VERSION_ID": "\"20.04\""
    },
    "PodsNotRunning": 0
  }
}
```

Run the tests - exit code is the number of failed tests

```sh
$ smoke
--- some.host.example.com ---
ok   OpenPorts: 22, 443
ok   HttpsGetStatusCode: 200
fail OsRelease: want VERSION_ID="22.04", got VERSION_ID="20.04"
ok   FilesPresent: /etc/passwd
ok   PodsNotRunning: 0
fail HelmReleases: ssh "helm ls -A": failed to run: Process exited with status 127
$ echo $?
2
```

## Development

It's easy to add a new test:

* write function of type `tests.TestFunc`
* add it to `tests.Available`
* test, install and run it: `make && smoke`
* update config file and smoke output in the Usage section above