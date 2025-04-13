# Secrethor CLI







Secrethor CLI is a Kubernetes secret auditing tool that complements the [Secrethor Operator](https://github.com/miltlima/secrethor).

---

## Features

- Detect orphaned Secrets not used by any workload
- Search for Secrets by name across namespaces
- Clean, structured output
- JSON/YAML export support
- Ess

---

## Installation

### Install with `go install`

```bash
go install github.com/miltlima/secrethor-cli@latest
```

### Or clone and build manually

```bash
git clone https://github.com/miltlima/secrethor-cli.git
cd secrethor-cli
go build -o secrethor-cli main.go
```

---

## Usage

```bash
secrethor-cli secrets orphan --namespace default
secrethor-cli secrets search my-secret-name --namespace all
```

### Flags

| Flag          | Description                                       |
| ------------- | ------------------------------------------------- |
| `--namespace` | Namespace to scan (or `all`)                      |
| `--output`    | Output format: `json`, `yaml`, or default (table) |
| `--verbose`   | Enable detailed scan output                       |

---

## Example Output

```
Secrethor CLI Orphan Report - Version v0.0.1

In-use Secrets
NAMESPACE        NAME                 USED BY
--------------   ------------------   ---------------------------------------------
  default         db-credentials       Deployments/backend-api

Orphaned Secrets
NAMESPACE        NAME
--------------   ------------------
  default         unused-token
  dev             staging-api-secret

Summary
Secrets in total:   3
Secrets in use:     1
Orphaned secrets:   2
```

---

## Project Structure

```
.
├── cmd/                # CLI commands (orphan, search, etc)
├── internal/secrethor/ # Core logic used by CLI
├── main.go             # Entry point
└── README.md
```

---

## Contributing

Pull requests are welcome! Feel free to open issues for bugs or features.

---

Built by [Milton Lima de Jesus](https://github.com/miltlima)

