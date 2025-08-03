# mgit - Multi-Repository Git Management Tool


**mgit** is a command-line tool for managing and executing commands across multiple Git repositories simultaneously.

## Installation

### Prerequisites
- Go 1.23 or higher
- Git installed on your system



### Build and Install

#### Linux/macOS (Bash/Zsh)

```bash
go build -o mgit
chmod +x mgit
sudo mv mgit /usr/local/bin/mgit
```

#### Windows (PowerShell or CMD)

```powershell
go build -o mgit.exe
Move-Item mgit.exe "$env:USERPROFILE\bin\mgit.exe"
# Ensure $env:USERPROFILE\bin is in your PATH
```

---

### Shell Completion

#### Bash (Linux)

```bash
go run ./ completion > mgit_completion
sudo mv mgit_completion /etc/bash_completion.d/mgit
source /etc/bash_completion.d/mgit
```

#### Zsh (macOS)

```bash
go run ./ completion zsh > _mgit
mkdir -p ~/.zsh/completions
mv _mgit ~/.zsh/completions/
echo 'fpath+=~/.zsh/completions' >> ~/.zshrc
autoload -Uz compinit && compinit
```

#### PowerShell (Windows)

```powershell
go run ./ completion powershell > mgit_completion.ps1

. .\mgit_completion.ps1 # for temporary use 
```


## Usage

### Basic Structure
```bash
mgit [flags] --repos [repos] -- [command to run on repos]
mgit [subcommand] [flags] 
```

### Core Commands

#### Add repositories/namespaces
```bash
# Add a new repository to current namespace
mgit add repo --<path> full-path #name of the repository is automatically sourced from the end of the path

mgit add repo --<path> full-path --name custom-name

# Add a new namespace
mgit add namespace --name <namespace-name>
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

# List repositories in a sepecified namespace
mgit list --namespace name

# List repositories in all namespaces
mgit list --all
```
![list preview](https://github.com/genocem/mgit/blob/main/image.png)

#### Switch between namespaces
```bash
mgit switch-namespace <namespace-name>
```

#### Run commands on repositories
```bash
# Run 'git status' on selected repositories
mgit --repos repo1,repo2,repo3 -- git status
mgit -r repo1 -r repo2 -- git status
# Run 'git pull' on all repositories in current namespace
mgit -- git pull
```


## Examples

1. **Add repositories to your workspace:**
```bash
mgit add repo --path /home/username/projects/project1
mgit add repo -p /home/username/projects/project2
```

2. **Create and switch to a new namespace:**
```bash
mgit add namespace work-projects
mgit switch-namespace work-projects
```

3. **Check status of specific repositories:**
```bash
mgit --repos project1 --repos project2 -- git status
mgit -r project1 -r project2 -- git status
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