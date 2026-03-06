# jg - frecency-based git repo jumper (lazy loaded)

# chpwd hook: runs immediately for cd tracking
_jg_chpwd() { command jg --add "$PWD" &! }
autoload -Uz add-zsh-hook
add-zsh-hook chpwd _jg_chpwd

# jg function: lazy stub replaced on first invocation
jg() {
  unfunction "$0"
  jg() {
    local result
    result=$(command jg "$@")
    local ret=$?
    if [[ $ret -eq 0 && -d "$result" ]]; then
      builtin cd "$result"
    elif [[ -n "$result" ]]; then
      echo "$result"
    fi
    return $ret
  }
  jg "$@"
}
