## komi

`komi` is a simple command saver with the ability to group commands in categories so they can be retrieved easily. Commands and their uses can be added, modified and deleted. They can quickly be copied to the system clipboard. All the data is saved in a JSON file which can be exported.

Features:

- Bash completion for categories
- Command copying to system clipboard
- Searching for string (includes case-insensitive search)
- Exporting data file

## Dependency

If you're on Linux, you would need the `xclip` package for copying commands. On Debian/Ubuntu you can install it with:

```bash
sudo apt-get install xclip
```

Also note, if you're using `komi copy` command via ssh on a remote linux machine, you would need to enable X11 forwarding. You can do by adding the `-X` flag:

```bash
ssh <user>@<ip> -X
```

## Installation

The default data directory is kept as:

```
/home/$USER/.komi
```

If you want a different data directory, you can export `KOMI_DATA_DIR` env variable to that directory.

```bash
export KOMI_DATA_DIR="/home/$USER/diff_komi"
```

**Note**: If specifying data dir through env var, make sure to include the export statement in your `.bashrc`.


### Installation with Go

```bash
go get github.com/shivammg/komi
sudo cp data/bash_autocomplete /etc/bash_completion.d/komi
source /etc/bash_completion.d/komi

# If you want to see the example data file (Optional)
mkdir /home/$USER/.komi
cp data/komi.json /home/$USER/.komi/komi.json
```

### Installation without Go

```bash
git clone git@github.com:shivamMg/komi.git
cd komi
./install.sh
source /etc/bash_completion.d/komi
```

`install.sh` does the following:

1. Copies one of the following binaries to `/usr/local/bin` according to the platform. Binaries for other platforms are not included.
   - [linux/amd64](data/bin/linux_amd64)
   - [darwin/amd64](data/bin/darwin_amd64)

   **Note**: This requires sudo permissions.
2. Creates data directory at `/home/$USER/.komi` and copies example data file from `data/komi.json`.
3. Copies the bash completion script to `/etc/bash_completion.d/`.
   **Note**: This also requires sudo permissions.

You can then source your komi bash completion script to update your current shell.

```bash
source /etc/bash_completion.d/komi
```

