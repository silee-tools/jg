# jg

A frecency-based CLI for quickly jumping to Git repositories.

## 프로젝트 구조

```
cmd/jg/main.go          — CLI 진입점, flag 파싱, 서브커맨드 라우팅
internal/entry/          — Entry 타입, ~/.jg 파일 I/O, file locking
internal/frecency/       — frecency 스코어링 알고리즘
internal/fzf/            — fzf 프로세스 실행, preview 구성
internal/shell/          — jg init zsh/bash 코드 생성
```

## 개발

- Language: Go 1.23
- Task Runner: mise
- Build: `mise run build`
- Test: `mise run test`
- Lint: `mise run lint`
- Format: `mise run fmt`
- Install: `mise run install`

## 릴리스

- `v*` 태그 push → GitHub Actions → GoReleaser → Homebrew tap 자동 업데이트
- `HOMEBREW_TAP_TOKEN` secret 필요 (homebrew-tap 레포 push 권한)
