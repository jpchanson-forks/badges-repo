BADGES = "Assisted"
THEME = src/theme.css
TEMPLATE = src/badge.svg.tmpl
LABEL = Human Influence:

CONTAINER := $(shell \
	if command -v docker >/dev/null 2>&1; then echo docker; \
	elif command -v podman >/dev/null 2>&1; then echo podman; \
	else echo ""; \
	fi)

help: ## Show help
	@echo "== Core =="
	@grep -E '^[a-zA-Z_-]+:.*?##' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo ""
	@echo "== Badges =="
	@for b in $(BADGES); do \
		echo "  $$b -> make badge-$$b"; \
	done

.PHONY: all $(BADGES)
all: $(BADGES) ## make all badges

image: ## Build the container image, uses either docker or podman whichever is present
	@if [ -z "$(CONTAINER)" ]; then \
		echo "Error: neither docker nor podman found"; \
		exit 1; \
	fi
	$(CONTAINER) build --progress=plain -t badgegen src

$(BADGES):
	$(MAKE) badge-$@

badge-%:
	$(CONTAINER) run --rm \
		-v "$(CURDIR):/work" \
		-w /work \
		badgegen \
		"$(LABEL)" "$*" "$(THEME)" "$(TEMPLATE)" \
		> test-badge-$*.svg

