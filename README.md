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

## Installation (without go)

```bash
git clone git@github.com:shivamMg/komi.git
cd komi
git checkout temp-br
./install.sh
source ~/.bashrc
source /etc/bash_completion.d/komi
```

`install.sh` does the following:

1. Copies one of the following binaries to `/usr/local/bin` according to the platform. Other platforms are not supported.
   - [linux/amd64](_linux_amd64/komi)
   - [darwin/amd64](_darwin_amd64/komi)

   **Note**: This requires sudo permissions.
2. Creates data directory for komi.
3. Copies example data file from `_data` directory.
4. Appends the following line to your `~/.bashrc`:
   ```bash
   export KOMI_DATA_DIR="/home/$USER/.komi"
   ```
5. Copies the bash completion script to `/etc/bash_completion.d/`.

You can then source your `.bashrc` and komi bash completion script to update your current shell.
```bash
source ~/.bashrc
source /etc/bash_completion.d/komi
```

## Installation (with go)

```bash
go get github.com/shivammg/komi
mkdir /home/$USER/.komi
echo export KOMI_DATA_DIR="/home/$USER/.komi" >> ~/.bashrc
sudo cp _data/bash_autocomplete /etc/bash_completion.d/komi
source ~/.bashrc
source /etc/bash_completion.d/komi
```
