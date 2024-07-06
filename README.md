Smoke is a simple tool for running basic sanity tests against a host. It's easy to add new tests.

## Installation

```sh
go install github.com/jreisinger/smoke@latest
```

## Usage

create config file containing tests for one or more hosts

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
    }
  }
}
```

run the tests

```sh
❯ smoke -v
ok   HelmReleases on some.host.example.com
ok   HttpsGet on some.host.example.com
ok   OpenPorts on some.host.example.com
ok   OsRelease on some.host.example.com
```
