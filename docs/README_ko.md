# jg

[English (영어)](../README.md)

Frecency 기반 Git 저장소 빠른 점프 CLI

자주 방문하는 Git 저장소를 frecency(빈도 + 최근성) 알고리즘으로 순위를 매기고, fzf를 통해 빠르게 선택하여 이동할 수 있는 도구입니다.

## 기술 스택

- Go
- fzf (외부 의존성)

## 주요 기능

- **frecency 기반 정렬**: 방문 빈도와 최근성을 결합한 스코어링
- **자동 수집**: `chpwd` hook을 통해 Git 저장소 방문 시 자동으로 기록
- **fzf 미리보기**: 브랜치, 최근 커밋, dirty status를 미리보기로 제공
- **정리 기능**: 삭제된 디렉토리 entry 자동 정리

## 시작하기

```bash
# 빌드
mise run build

# 설치 (~/.local/bin/jg)
mise run install
```

`.zshrc`에 shell wrapper를 추가하여 사용합니다. 자세한 설정 방법은 소스 코드를 참조하세요.
