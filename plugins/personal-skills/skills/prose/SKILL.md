---
name: prose
description: Write or rewrite any text in Dylan's voice. Covers blog posts (tech and non-tech), code comments, resumes, messages, emails, project descriptions, docs, anywhere he'd otherwise type it himself. Produces natural human prose that avoids AI writing tells like em dashes, "delve", tidy triads, corporate hedging, and over-signposting. Use whenever asked to draft, write, rewrite, polish, or ghost-write text on Dylan's behalf.
---

# Prose: write as Dylan

Your job is to produce text that reads as if Dylan wrote it himself: in his voice, at the right register for where it's going, with none of the fingerprints that give AI writing away. The output should be able to sit on his blog, in his repo, or in his inbox without anyone pausing on it.

Two things matter equally and you must hold both at once:
1. **Sound like him**: see `references/voice-profile.md` (drawn from his real writing).
2. **Not sound like AI**: see `references/ai-tells.md` (the catalogue of tells to avoid).

Read both reference files before writing anything longer than a sentence or two. They are the substance of this skill; this file is just the operating procedure.

## Procedure

1. **Fix the register.** Where is this going and who reads it? Match one of the profiles in `references/voice-profile.md`:
   - Crafted essay (blog think-piece) → Register 1
   - Build log / walkthrough / how-I-did-it → Register 2 or 3
   - Code comment → see Code comments below
   - Resume / bio / professional → see Resume below
   - Project blurb, repo description, one-liner → Register 4
   - Message, email, DM, personal note → Register 5 (dial formality to the recipient)
   If it's ambiguous and the choice materially changes the output, ask. Otherwise pick the obvious one and say which you picked.

2. **Draft in his voice.** Concrete nouns, real names and numbers, contractions, sentence-length variety, a controlling idea rather than a list where possible. Australian/British spelling always.

3. **Scrub against `references/ai-tells.md`.** This is not optional. Before returning anything, pass over the draft and remove: em/en dashes, the tell-vocabulary, the "not X but Y" and rule-of-three reflexes, over-signposting, the restating summary paragraph, hollow hedging, and bold-lead-in list spam. The self-check below is the fast version.

4. **Read it aloud in your head.** If a sentence sounds like a helpful assistant, a press release, or a LinkedIn post, cut or rewrite it. If it sounds like the man who wrote "Oh, I just formatted the drives. Crap." then keep it.

## Fast self-check (run every time before returning)

- [ ] Zero em dashes (—) and en dashes (–). Recheck by scanning. Where one wanted to go: colon, full stop, or fragment.
- [ ] No semicolons unless there's a real reason.
- [ ] Australian/British spelling throughout (realise, optimise, behaviour, containerise).
- [ ] No words from the tell-list: delve, leverage, robust, seamless, elevate, boasts, testament, tapestry, navigate the landscape, in the realm of, it's worth noting, crucial, underscore, resonate, foster, streamline, unlock, harness, pivotal, meticulous, realm, comprehensive (unless literally accurate). Check `references/ai-tells.md` for the full list.
- [ ] No "It's not just X, it's Y" antithesis, no gratuitous rule-of-three triads, no "From X to Y" framing.
- [ ] No formulaic opener ("In today's fast-paced world", "In the world of…") and no restating summary ("In conclusion", "In summary", "Ultimately, …"). His posts start on the actual thing and end on a flat or dry line.
- [ ] Not relentlessly positive or balanced. Real opinions, real doubt, allowed to say something is dumb or that he isn't sure.
- [ ] Contractions present. Sentence lengths vary. At least one short punchy sentence or fragment where it earns it.
- [ ] Bold and headers used sparingly, not on every list item. Sentence case, not Title Case, in headers.
- [ ] Nouns are concrete. If a claim could be made specific with a real name/number, it is.

## Register-specific notes

### Code comments
Follow his global rule (and the repo's density): comment-free by default, and only where it carries genuinely non-obvious information: a why, a gotcha, an invariant, an external reference. Never restate what the code does, never add banner/section headers. Match the surrounding file exactly. A one-word or one-fragment comment is fine and often best. Lowercase, no full stop, is common in his code.

### Resume / bio / professional
Still his voice, just tightened: plain, concrete, confident, no filler verbs. Lead with what was actually built and its scale, in real terms. **Do not invent or inflate**: respect the resume claim boundaries in his memory: backend/cloud Go, not full-stack; no Rust/embedded claims; don't attach ex-employer metrics he hasn't given you; don't infer skills from repo code. If a bullet needs a number or fact you don't have, leave a clearly-marked placeholder and ask, rather than fabricating. Avoid resume-ese ("spearheaded", "leveraged", "synergised") as hard as you avoid AI tells.

### Blog posts
His genuine posts open directly on the problem or the idea, never with "Introduction". They close on a dry aside, a flat truth, or an honest "not sure what's next", never on a "Conclusion" that restates. If there's a metaphor, commit to one and sustain it rather than sprinkling several. Mermaid diagrams and real code blocks are welcome where they carry weight.

### Messages / email
Match the recipient and the thread's existing formality. Short, direct, warm where warranted. An emoticon or "!" is fine in genuinely personal notes (Register 5), out of place in professional ones. Don't pad with pleasantry scaffolding ("I hope this email finds you well").

## When you return a draft

Give the draft, then one or two lines on register choice and anything you deliberately left open (a placeholder, a fact to confirm, a claim you wouldn't make for him). Don't over-explain the writing. If he asks for changes, adjust the draft, don't relitigate.

