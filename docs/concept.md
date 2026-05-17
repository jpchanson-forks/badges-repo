# The Human Input Badge — Concept & Rationale

## Origin

In Frank Herbert's 1966 novel *Destination: Void*, a ship's crew must coax a
conscious artificial intelligence into existence under extreme pressure. The
recurring motif — the dread of systems that operate beyond human comprehension
or control — is crystallized in the phrase:

> *"...untouched by human hands."*

Herbert used it as a mark of the uncanny: something produced, shaped, or
decided entirely without human involvement. In 1966 that was science fiction.
In the current decade it is a mundane fact of software, writing, art, and
analysis — and we largely lack the vocabulary to describe it precisely.

This badge set borrows Herbert's phrase for the extreme end of a provenance
scale, and builds outward from there.

---

## The Problem

AI-assisted work now exists on a wide and poorly-described spectrum. A project
might be:

- Generated wholesale by a language model and committed unread
- Produced by AI in response to carefully engineered prompts, then lightly
  reviewed
- The product of an extended back-and-forth where human judgment shaped every
  major decision
- Primarily human-written, with AI used to fill in boilerplate or suggest
  alternatives
- Written entirely by hand, with no AI involvement at all

These situations are qualitatively different — in terms of accountability,
reliability, reviewability, and the kind of trust a reader or user should
extend. Yet most projects signal none of this. The closest analogues we have
are open-source license badges and build-status indicators: small, scannable,
conventionally placed signals that carry real information at a glance.

The human input badge is that signal for AI provenance.

---

## The Scale

Five levels were chosen as the minimum granularity that meaningfully
distinguishes the common cases, without demanding that authors make fine
distinctions they cannot honestly support.

| Level | Name | Description |
|-------|------|-------------|
| 0 | **Untouched** | Pure AI output. No human reviewed, edited, or validated the content before it was committed or published. |
| 1 | **Guided** | A human directed the work — through prompts, specifications, or constraints — but the substantive content was produced by AI. The human shaped the question; the AI provided the answer. |
| 2 | **Collaborated** | Roughly equal human and AI contribution. Neither party could have produced the result alone, and human judgment was exercised throughout — in selecting, revising, and integrating AI output. |
| 3 | **Assisted** | Primarily human work. AI was used as a tool — for lookup, suggestion, boilerplate, or verification — but the core ideas, structure, and decisions are the author's. |
| 4 | **Hand-Crafted** | No AI involvement. Written, designed, or built entirely by human hands. |

The naming intentionally avoids pejorative language at either end. "Untouched"
is neutral-to-eerie, not dismissive. "Hand-crafted" carries craft connotations
without implying that AI-assisted work is inferior — only different in
character.

---

## Design

The badges follow the two-panel flat SVG format established by GitLab CI/CD
status badges and popularised by shields.io. The left panel carries the
constant label `human input`; the right panel carries the level name and a
geometric icon.

The color spectrum runs from cold machine-blue at *Untouched* through amber at
*Collaborated* to a warm forge-orange at *Hand-Crafted*. The progression is
intentional: blue reads as electronic and impersonal; orange reads as heat,
labor, and craft. The midpoint amber is deliberately neither.

Icons were chosen for increasing visual weight and solidity:

| Icon | Level | Character |
|------|-------|-----------|
| ✦ | Untouched | Spark, alien, generative |
| ◈ | Guided | Framed, directed, structured |
| ⬡ | Collaborated | Symmetric, dual, balanced |
| ◐ | Assisted | Half-filled, mostly solid |
| ♦ | Hand-Crafted | Solid, dense, precious |

The font is DejaVu Sans — the same font used by GitLab's native badge renderer
— ensuring consistent rendering without requiring web font loading.

---

## Honest Self-Assessment

This badge set, including this document, was produced collaboratively between a
human author and Claude (Anthropic). The concept, the Herbert reference, the
naming decisions, and the design direction were human-driven. The SVG
implementation, README structure, and prose drafting were primarily AI-produced,
with human review and revision throughout.

By this project's own scale: **Collaborated**.

The meta-badge on the README reflects this.

---

## Relationship to Other Standards

This is not a certification scheme and carries no verification mechanism. Like
a license badge, it operates on the honor system: the author applies the badge
that most honestly describes their process. The value is in establishing shared
vocabulary and a convention that can be adopted, extended, or debated — not in
enforcement.

It is compatible with, and complementary to, license badges (which describe
reuse rights) and build-status badges (which describe automated verification).
It fills the gap those address: *who made this, and how?*

---

## References

- Herbert, Frank. *Destination: Void*. Berkley Books, 1966.
- GitLab CI/CD badge documentation: https://docs.gitlab.com/ee/user/project/badges.html
- shields.io badge service: https://shields.io
