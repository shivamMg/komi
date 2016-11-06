## komi

`komi` is a simple command saver with the ability to group commands in categories so they can be retrieved easily. Commands and their uses can be added, modified and deleted. They can quickly be copied to the system clipboard. All the data is saved in a JSON file which can be exported.

Features:

- Bash completion for categories
- Command copying to system clipboard
- Searching for string (includes case-insensitive search)
- Exporting data file

## Installation (without go)

```bash
git clone git@github.com:shivamMg/komi.git
git checkout temp-br
./install.sh
source ~/.bashrc
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

You can then source your `.bashrc` to update your current shell.
```bash
source ~/.bashrc
```

## Installation (with go)

```bash
go get github.com/shivammg/komi
mkdir /home/$USER/.komi
echo export KOMI_DATA_DIR="/home/$USER/.komi" >> ~/.bashrc
source ~/.bashrc
```
