# mgit - Multi-Repository Git Management Tool


**mgit** is a powerful command-line tool for managing and executing commands across multiple Git repositories simultaneously. Built in Go, it simplifies workflows when working with multiple repositories by allowing you to perform operations across all your projects with a single command.


## Installation

### Prerequisites
- Go 1.23 or higher
- Git installed on your system

### Build and Install
```bash
go build -o mgit
chmod +x mgit
sudo mv mgit /usr/local/bin/

# Enable bash completion
go run ./ completion > mgit_completion
sudo mv mgit_completion /etc/bash_completion.d/
source /etc/bash_completion.d/mgit_completion
```

## Usage

### Basic Structure
```bash
mgit [command] [subcommand] [flags] -- [git command]
```

### Core Commands

#### Add repositories/namespaces
```bash
# Add a new repository to current namespace
mgit add repo <repository-url> [flags]

# Add a new namespace
mgit add namespace <namespace-name>
```

#### Delete repositories/namespaces
```bash
# Delete a repository from current namespace
mgit delete repo <repo-name>

# Delete a namespace
mgit delete namespace <namespace-name>
```

#### List repositories
```bash
# List repositories in current namespace
mgit list

# List repositories in all namespaces
mgit list --all
```
![list preview](https://github.com/genocem/mgit/blob/main/image.png)

#### Switch namespaces
```bash
# Change to a different namespace
mgit switch-namespace <namespace-name>
```

#### Run commands on repositories
```bash
# Run 'git status' on selected repositories
mgit --repos repo1,repo2 -- git status

# Run 'git pull' on all repositories in current namespace
mgit -- git pull
```

### Global Flags
```bash
-n, --namespace string   Namespace for the resource (default "default")
-r, --repos strings      List of repositories to run commands on
```

## Examples

1. **Add repositories to your workspace:**
```bash
mgit add repo https://github.com/user/project1
mgit add repo https://github.com/user/project2
```

2. **Create and switch to a new namespace:**
```bash
mgit add namespace work-projects
mgit switch-namespace work-projects
```

3. **Check status of specific repositories:**
```bash
mgit --repos project1,project2 -- git status
```

4. **Pull latest changes from all repositories in current namespace:**
```bash
mgit -- git pull
```

5. **List all tracked repositories across namespaces:**
```bash
mgit list --all
```

## Configuration

mgit uses a SQLite database to store your repository information and namespaces. The database is automatically created at `~/.mgit/db.sqlite` on first run.

## Building from Source

1. Clone the repository:
```bash
git clone https://github.com/genocem/mgit.git
cd mgit
```

2. Build and install:
```bash
go build -o mgit
sudo mv mgit /usr/local/bin/
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.