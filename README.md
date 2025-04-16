<p align="center">
  <img src="assets/images/secrethor-cli-logo.png" alt="Secrethor Logo" width="300"/>
</p>

<p align="center">
  <a href="https://opensource.org/licenses/MIT"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg"/></a>
</p>

<p align="center">
  <!-- License -->
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg"/>
  </a>
  <!-- Release -->
  <a href="https://github.com/miltlima/secrethor-cli/releases">
    <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/miltlima/secrethor-cli"/>
  </a>
  <!-- Go Version -->
  <a href="https://golang.org">
    <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/miltlima/secrethor-cli"/>
  </a>
  <!-- Build Status -->
  <a href="https://github.com/miltlima/secrethor-cli/actions/workflows/release.yml">
    <img alt="Build Status" src="https://github.com/miltlima/secrethor-cli/actions/workflows/release.yml/badge.svg"/>
  </a>
  <!-- Go Report Card -->
  <a href="https://goreportcard.com/report/github.com/miltlima/secrethor-cli">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/miltlima/secrethor-cli"/>
  </a>
  <!-- Downloads -->
  <a href="https://github.com/miltlima/secrethor-cli/releases">
    <img alt="GitHub all releases" src="https://img.shields.io/github/downloads/miltlima/secrethor-cli/total"/>
  </a>
  <!-- Last Commit -->
  <a href="https://github.com/miltlima/secrethor-cli/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/miltlima/secrethor-cli"/>
  </a>
  <!-- Open Issues -->
  <a href="https://github.com/miltlima/secrethor-cli/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/miltlima/secrethor-cli"/>
  </a>
</p>

# Secrethor CLI


Secrethor CLI is a Kubernetes secret auditing tool that complements the [Secrethor Operator](https://github.com/miltlima/secrethor).

---
## Core Functionalities

### 1. Secret Scanning
- **Orphaned Secret Detection**: 
  - Automatically identifies secrets not used by any workload
  - Scans across all Kubernetes workload types:
    - Deployments
    - StatefulSets
    - DaemonSets
    - ReplicaSets
    - CronJobs
    - Jobs
    - Pods
  - Provides clear visual indicators (🔒 for used, ❗ for orphaned)

### 2. Secret Search
- **Cross-Namespace Search**:
  - Find secrets by name across all namespaces
  - Detailed information display:
    - ✅ Secret location (namespace/name)
    - 🔐 Secret type
    - 📦 Available data keys
  - Supports single namespace or all-namespace search

## Features

- Detect orphaned Secrets not used by any workload
- Search for Secrets by name across namespaces
- Clean, structured output
- JSON/YAML export support


## Workload Coverage

Secrethor CLI scans for secrets used in:

| Workload Type | Secret References Checked |
|---------------|-------------------------|
| Deployments   | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| StatefulSets  | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| DaemonSets    | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| ReplicaSets   | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| CronJobs      | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| Jobs          | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
| Pods          | - Environment variables<br>- Volume mounts<br>- Image pull secrets |
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
### Flags

| Flag          | Description                                       |
| ------------- | ------------------------------------------------- |
| `--namespace` | Namespace to scan (or `all`)                      |
| `--output`    | Output format: `json`, `yaml`, or default (table) |
| `--verbose`   | Enable detailed scan output                       |

---

## Usage Examples

### Scan for Orphaned Secrets
```bash
# Scan all namespaces
secrethor-cli secrets orphan --namespace all

# Scan specific namespace
secrethor-cli secrets orphan --namespace default

# With verbose output
secrethor-cli secrets orphan --namespace all --verbose

# With different output format
secrethor-cli secrets orphan --namespace all --output json
```

### Search for Specific Secrets
```bash
# Search across all namespaces
secrethor-cli secrets search my-secret-name --namespace all

# Search in specific namespace
secrethor-cli secrets search my-secret-name --namespace default
```

### Output Format Examples
```bash
# Table output (default)
secrethor-cli secrets orphan --output table

# JSON output
secrethor-cli secrets orphan --output json

# YAML output
secrethor-cli secrets orphan --output yaml
```


## Example Output

```
 _______                            __    __
|     __|.-----..----..----..-----.|  |_ |  |--..-----..----.
|__     ||  -__||  __||   _||  -__||   _||     ||  _  ||   _|
|_______||_____||____||__|  |_____||____||__|__||_____||__|


In-use Secrets
  NAMESPACE              NAME                  USED BY                                                                                                                                                      
                                                                                                                                                                                                            
  🔒  mongo              my-mongo-db-config    StatefulSets/my-mongo-db, StatefulSets/my-mongo-db-arb, Pods/my-mongo-db-0, Pods/my-mongo-db-1, Pods/my-mongo-db-2                                           
  🔒  secrethor-system   webhook-server-cert   Deployments/controller-manager, ReplicaSets/controller-manager-54486fd57, ReplicaSets/controller-manager-9d64b7c76, Pods/controller-manager-54486fd57-tcwd9  

Orphaned Secrets
  NAMESPACE          NAME                                      
                                                               
  ❗  cert-manager   cert-manager-webhook-ca                   
  ❗  default        allowed-secret                            
  ❗  default        both-ok                                   
  ❗  default        only-password                             
  ❗  default        only-username                             
  ❗  default        totally-unused                            
  ❗  dev            test-secret                               
  ❗  dev            test-secret1                              
  ❗  mongo          my-mongo-db-admin-my-user                 
  ❗  mongo          my-mongo-db-agent-password                
  ❗  mongo          my-mongo-db-keyfile                       
  ❗  mongo          my-user-password                          
  ❗  mongo          my-user-scram-scram-credentials           
  ❗  mongo          sh.helm.release.v1.community-operator.v1  

Summary
🔑  Secrets in total:   16
🔒  Secrets in use:     2
❗  Orphaned secrets:   14
```

---

## Project Structure

```
.
├── CHANGELOG.md
├── README.md
├── cmd
│   ├── expired_
│   ├── orphan.go
│   ├── root.go
│   ├── search.go
│   ├── secrets.go
│   └── version.go
├── go.mod
├── go.sum
├── internal
│   └── secrethor
│       ├── banner.go
│       ├── expired_
│       ├── orphan.go
│       ├── search.go
│       └── utils.go
├── main.go
└── makefile

4 directories, 17 files

```

---

## Contributing

We welcome contributions! Here's how you can help:

- Report bugs by opening issues
- Suggest new features
- Submit pull requests
- Improve documentation

Please ensure your commits follow conventional commit format for automatic versioning.

---

## License

MIT License - see LICENSE file for details.

---
Built with ❤️ by [Milton Lima de Jesus](https://github.com/miltlima)
