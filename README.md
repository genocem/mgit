# mgit - Multi-Repository Git Management Tool


**mgit** is a command-line tool for managing and executing commands across multiple Git repositories simultaneously.


## Download
Download the latest version from the [Releases page](https://github.com/genocem/mgit/releases).

## Install on Linux

### App Binary
```bash
chmod +x mgit
sudo mv mgit /usr/local/bin/mgit
```
### Shell Completion

#### Bash/Zsh/Fish (Linux)

```bash
sudo mv <mgit_completion> /etc/bash_completion.d/<mgit_completion>
source /etc/bash_completion.d/<mgit_completion>
```
## Install on MacOS

### App Binary
```bash
chmod +x mgit
sudo mv mgit /usr/local/bin/mgit
```
### Shell Completion

#### Zsh (macOS)

```bash
mkdir -p ~/.zsh/completions
mv <_mgit> ~/.zsh/completions/
echo 'fpath+=~/.zsh/completions' >> ~/.zshrc
autoload -Uz compinit && compinit
```
## Install on Windows

### App Binary
```ps1
Move-Item mgit.exe "$env:USERPROFILE\bin\mgit.exe"
# Ensure $env:USERPROFILE\bin is in your PATH
```
### Shell Completion
#### PowerShell temporary use

```powershell
go run ./ completion powershell > mgit_completion.ps1

. .\mgit_completion.ps1 # for temporary use 
```


## Build from source (requires Go to be installed)

### Build Binaries

#### Linux/macOS 

```bash
cd mgit
go build -o mgit
chmod +x mgit
sudo mv mgit /usr/local/bin/mgit
```

#### Windows (PowerShell or CMD)

```powershell
cd mgit
go build -o mgit.exe
Move-Item mgit.exe "$env:USERPROFILE\bin\mgit.exe"
# Ensure $env:USERPROFILE\bin is in your PATH
```


### Completion

#### bash
```bash
mgit completion bash > mgit_completion.bash
```

#### zsh

```zsh
mgit completion zsh > mgit_completion.zsh
```

#### fish

```fish
mgit completion fish > mgit_completion.fish
```

#### powershell

```powershell
mgit completion powershell > mgit_completion.ps1
```

---




## Usage

### Add repositories/projects
```bash
# Add a new repository to current project
mgit add repo --path <full-path> #name of the repository is automatically sourced from the end of the path

mgit add repo --path <full-path> --name custom-name

# Add a new project
mgit add project --name <project-name>
```

### Delete repositories/projects
```bash
# Delete a repository from current project
mgit delete repo <repo-name>

# Delete a project
mgit delete project <project-name>
```

### List repositories
```bash
# List repositories in current project
mgit list

# List repositories in a sepecified project
mgit list --project name

# List repositories in all projects
mgit list --all
```
![list preview](https://github.com/genocem/mgit/blob/main/image.png)

### Switch between projects
```bash
mgit switch-project <project-name>
```

### Execute commands on repositories
```bash
mgit exec repo1 repo2 repo3 -- command
mgit exec -- command
```


## Examples

1. **Add repositories to your workspace:**
```bash
mgit add repo --path /home/username/projects/project1
mgit add repo -p /home/username/projects/project2
```

2. **Create and switch to a new project:**
```bash
mgit add project work-projects
mgit switch-project work-projects
```

3. **Check status of specific repositories:**
```bash
mgit exec project1 project2 -- git status
```

4. **Pull latest changes from all repositories in current project:**
```bash
mgit -- git pull
```

5. **List all tracked repositories across projects:**
```bash
mgit list --all
```

## Configuration

mgit uses a SQLite database to store your repository information and projects. The database is automatically created at `~/.mgit` as `db.sqlite` on first run. 

It also keeps track of the current project in `~/.mgit/config.json`



## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the Apache-2.0 License - see the [LICENSE](LICENSE) file for details.