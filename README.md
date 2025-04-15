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

