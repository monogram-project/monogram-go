[private]
default:
    just --list

# Initialize decision records
init-decisions:
    python3 scripts/decisions.py --init

# Add a new decision record
add-decision TOPIC:
    python3 scripts/decisions.py --add "{{TOPIC}}"

test:
    just -f golang/Justfile build
    python3 runtests.py --tests tests.yaml
