# To add autocompletion to vmount, please copy this script to
# /etc/bash_completion.d and log off / log in again

OPTIONS="--help -m -d -e -l"
complete -W "${OPTIONS}" 'vmount'
