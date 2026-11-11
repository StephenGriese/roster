#!/usr/bin/env bash
# Guard: ensure built binaries match the project source. `make verify` (and
# any pipeline that trusts a prebuilt binary) refuses to proceed when a binary
# is missing, was built from a different revision, or is older than the
# tracked Go sources. Each probe tolerates a failing subcommand instead of
# tripping errexit and exiting silently.
#
# Usage: scripts/verify-build.sh BINARY...
set -euo pipefail
cd "$(dirname "$0")/.."

if [[ $# -eq 0 ]]; then
  echo "usage: $0 BINARY..." >&2
  exit 2
fi

# The working tree's revision must be resolvable (fail closed otherwise).
if ! EXPECTED_REV=$(git rev-parse HEAD 2>/dev/null); then
  echo "Error: cannot resolve the source revision (git HEAD). Run from the project's Git checkout and rebuild with \`make build\`." >&2
  exit 1
fi
EXPECTED_REV="${EXPECTED_REV:0:7}"

for BINARY in "$@"; do
  if [[ ! -x "$BINARY" ]]; then
    echo "Error: $BINARY is missing or not executable. Run \`make build\` to build it." >&2
    exit 1
  fi

  # Revision check only when the binary advertises a --version flag. Probing
  # the binary rather than invoking it unconditionally keeps servers that
  # ignore flags from starting their main loop and hanging the check.
  if grep -aq -- '--version' "$BINARY" 2>/dev/null; then
    if command -v timeout >/dev/null 2>&1; then
      VERSION_OUT=$(timeout 5 "$BINARY" --version 2>/dev/null) || VERSION_OUT=""
    elif command -v gtimeout >/dev/null 2>&1; then
      VERSION_OUT=$(gtimeout 5 "$BINARY" --version 2>/dev/null) || VERSION_OUT=""
    else
      VERSION_OUT=$("$BINARY" --version 2>/dev/null) || VERSION_OUT=""
    fi
    BIN_REV="${VERSION_OUT%% *}"
    # Only compare when the first token actually looks like a git revision;
    # a human-readable version (e.g. "v1.2.3") is not a revision to match.
    if [[ "$BIN_REV" =~ ^[0-9a-f]{7,40}$ && "$BIN_REV" != "$EXPECTED_REV" ]]; then
      echo "Error: $BINARY was built from a different revision ($BIN_REV vs $EXPECTED_REV). Run \`make build\` to rebuild it." >&2
      exit 1
    fi
  fi

  # No tracked Go source may be newer than the binary (uncommitted edits).
  NEWER_FILES=""
  while IFS= read -r f; do
    if [[ -f "$f" && "$f" -nt "$BINARY" ]]; then
      NEWER_FILES+="$f"$'\n'
    fi
  done < <(git ls-files -- '*.go' 'go.mod' 'go.sum' 2>/dev/null)
  if [[ -n "$NEWER_FILES" ]]; then
    echo "Error: $BINARY is older than tracked source files that affect the build:" >&2
    echo "${NEWER_FILES%$'\n'}" >&2
    echo "Run \`make build\` to rebuild it." >&2
    exit 1
  fi

  echo "✓ $BINARY matches the current source ($EXPECTED_REV)"
done
