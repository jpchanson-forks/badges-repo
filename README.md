# human-input-badges

![human input: collaborated](./badge-meta-this-repo.svg)
![license: CC0-1.0](./badge-license-cc0.svg)

SVG provenance badges for indicating the level of human authorship in a
project — from pure AI output to fully hand-crafted work.

Inspired by the GitLab pipeline badge convention, and by Frank Herbert's
*Destination: Void* (1966).

→ [**Read the full concept and rationale**](./docs/concept.md)

---

## The Scale

| Badge | Level | Meaning |
|-------|-------|---------|
| ![untouched](./badge-untouched.svg) | **Untouched** | Pure AI output, unreviewed by any human |
| ![guided](./badge-guided.svg) | **Guided** | Human wrote prompts; AI wrote the content |
| ![collaborated](./badge-collaborated.svg) | **Collaborated** | Roughly equal human/AI back-and-forth |
| ![assisted](./badge-assisted.svg) | **Assisted** | Primarily human; AI filled gaps |
| ![hand-crafted](./badge-handcrafted.svg) | **Hand-Crafted** | Pure human authorship; no AI used |

---

## Usage

Pick the badge that most honestly describes your process and add it to your
README. Like a license badge, this operates on the honor system.

### Markdown — GitHub

```markdown
![human input: collaborated](https://raw.githubusercontent.com/YOUR_USER/human-input-badges/main/badge-collaborated.svg)
```

### Markdown — GitLab

```markdown
![human input: collaborated](https://gitlab.com/YOUR_USER/human-input-badges/-/raw/main/badge-collaborated.svg)
```

### Relative path (badge in the same repo)

```markdown
![human input: collaborated](./badge-collaborated.svg)
```

### HTML

```html
<img src="https://raw.githubusercontent.com/YOUR_USER/human-input-badges/main/badge-collaborated.svg"
     alt="human input: collaborated">
```

---

## Files

```
badge-untouched.svg       ✦  Level 0 — Untouched
badge-guided.svg          ◈  Level 1 — Guided
badge-collaborated.svg    ⬡  Level 2 — Collaborated
badge-assisted.svg        ◐  Level 3 — Assisted
badge-handcrafted.svg     ♦  Level 4 — Hand-Crafted

badge-meta-this-repo.svg  meta: badge for this repo (collaborated)
badge-license-cc0.svg     companion license badge

docs/concept.md           full concept, rationale, and design notes
```

---

## License

[CC0-1.0](./LICENSE) — No rights reserved. Use freely, no attribution required.
