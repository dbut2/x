# AI writing tells: the catalogue to avoid

The goal is not to pass a detector. It's that the text reads like a person wrote it. These are the patterns that make text read as machine-generated. Avoid the whole family, not just the listed instances. The point is the underlying reflex, and the examples are illustrative.

**Why AI writes this way (the one thing to internalise):** models are tuned to *look* thoughtful, balanced, and polished to someone skimming. They use real rhetorical devices, but with no taste about when to stop, so the devices saturate and go mechanical. The forensic signature underneath every tell below is *low variance*: uniform sentence length, uniform structure, uniform positivity, uniform hedging. Everything human is the opposite of uniform. (Sources at the end; the vocabulary in particular shifts by model era, so treat the word-lists as live, not fixed.)

## 1. Punctuation

- **Em dash (—).** The loudest single tell. AI reaches for it constantly to join clauses, insert asides, and build the "not X — but Y" swing. Dylan never uses it. Replacements: a colon (setup: payoff), a full stop and a new sentence, a comma if the join is light, a fragment, or a plain "and"/"but"/"so".
  - AI: "The project was small — but it taught me everything."
  - Human: "The project was small. It taught me everything anyway."
- **En dash (–)** in prose (outside genuine number ranges). Same treatment: avoid it.
- **Semicolons.** Overused by AI to sound literate. Dylan uses none. Split into two sentences.
- **The "spaced hyphen" dodge ( - ).** People substitute this for an em dash and it reads just as artificial when overused. Don't swap one tell for another; restructure the sentence instead.
- **Curly/smart quotes** appearing in otherwise plain source, or a suspiciously perfect mix of them. Type straight quotes.
- **Perfectly balanced parenthetical asides** on every other sentence.

## 2. Vocabulary (the tell-words)

If any of these appear, stop and rewrite. They are the surface of AI's register.

delve, leverage (as a verb), robust, seamless, seamlessly, elevate, boasts, testament (a testament to), tapestry, navigate/navigating the landscape, in the realm of, realm, it's worth noting (that), it's important to note, crucial, pivotal, underscore, underscores, resonate, foster, fostering, streamline, unlock, harness, meticulous, meticulously, comprehensive, holistic, nuanced, intricate, myriad, plethora, embark, journey (metaphorical), landscape (metaphorical), ecosystem (metaphorical), game-changer, cutting-edge, state-of-the-art, unparalleled, unwavering, bustling, vibrant, rich (metaphorical), treasure trove, whether you're a X or a Y, at the end of the day, when it comes to, dive into / deep dive, first and foremost, notably, arguably, ultimately, moreover, furthermore, additionally, that being said, needless to say.

Also the hype adjectives that do sales work: powerful, blazing-fast, lightning-fast, effortless, intuitive, elegant (as filler), stunning, revolutionary, transformative. Say the concrete thing instead.

- AI: "This robust, seamless solution leverages cutting-edge tooling to streamline your workflow."
- Human: "It uses k3s and Traefik. Deploys are a git push."

**Copula avoidance** is the deeper version of this tell: AI dodges plain "is" and "are" for inflated stand-ins like "serves as", "stands as", "represents", "marks", "boasts", "features", "offers". Restore the plain verb.
- AI: "The library serves as a testament to the city's rich cultural landscape."
- Human: "The library is one of the oldest buildings in the city."

## 3. Structure and sentence shapes

- **The antithesis swing: "It's not just X, it's Y" / "This isn't about X. It's about Y."** AI's favourite rhetorical move. Kill on sight. State the thing directly.
- **Rule of three.** Triads everywhere: "fast, reliable, and scalable." "It's clean, it's simple, and it works." One or two is human; the relentless three-beat is a tell. Break the pattern: use one, or four, or an uneven list.
- **"From X to Y" framing.** "From startups to enterprises, from hobbyists to professionals…" Fakes comprehensive range. Name the specific cases you actually mean, or drop it.
- **"Not only X but also Y" / "more than just X, it's Y."** Same negative-parallelism family as the antithesis swing, equally overused. Just say Y.
- **Inverted parallelism / chiasmus.** "Stacks I pick up quickly; domain, I bring." "You don't choose the tool, the tool chooses you." Mirrored constructions that exist to sound writerly. Say it straight.
- **The constructed fragment punchline.** "Unglamorous work — the kind X lives or dies by." A deliberately crafted fragment deployed as a mic-drop. Different from Dylan's genuine plain fragments ("Done." "Crap."), which are short, flat, and unshowy. If the fragment feels written for applause, it's a tell.
- **The mirrored close.** Ending by echoing the opening line or image back at the reader for a tidy bow. His writing ends flat, dry, or uncertain, not with a callback ribbon.
- **One-off metaphor flourishes.** "That's the water I've been swimming in." A decorative metaphor dropped in once and abandoned. Either commit to one sustained metaphor for the whole piece (his *Protect Your Shed* move) or use none.
- **Formulaic openers.** "In today's fast-paced world", "In the world of software", "In an era of", "Imagine a scenario where", "Have you ever wondered". His posts open on the actual subject.
- **The restating conclusion.** "In conclusion", "In summary", "To sum up", "Ultimately", a final paragraph that repeats what was just said in tidier words. His posts end on a dry line, a flat truth, or honest uncertainty, and then stop.
- **Over-signposting.** "Let's dive in", "Now, let's explore", "First, we'll look at…, then we'll cover…", "As we've seen". Just move.
- **Every list item a bolded lead-in.** "**Performance:** it's fast. **Reliability:** it's stable." Occasionally fine; as a default format it's a tell. Prefer prose. If a list is genuinely right, let items be plain.
- **Section scaffolding: Introduction / Overview / Assumptions / Conclusion / Key Takeaways / TL;DR** stapled onto everything. His genuine posts don't do this.
- **Uniform paragraph length.** AI emits even, 3-4 sentence blocks. Humans write a six-line paragraph then a one-line one.
- **Topic-sentence-then-elaborate on every paragraph**, forever. Predictable and flat.

## 4. Tone

- **Relentless positivity / enthusiasm.** Everything is exciting, powerful, a great choice. Real writing has flat stretches, annoyance, doubt, and jokes. Dylan calls his own mistake "a dumb mistake" and says he's "not quite sold" on a proposal.
- **Hollow balance / both-sidesing.** "While X has its benefits, it's important to consider Y." "There are pros and cons to each approach." Take a position.
- **Over-hedging.** "It might be worth considering that perhaps…", "generally speaking, in many cases…". Say it. Where there's real uncertainty, name it plainly ("I'm not sure what's next") rather than wrapping it in qualifiers.
- **Over-explaining the obvious.** Defining terms the reader knows, spelling out every implication. Trust the reader.
- **The helpful-assistant register.** "Great question!", "Certainly!", "I'd be happy to…", "Here's a breakdown of…". None of it.
- **Unearned profundity.** Grave, weighty beats with nothing under them: "Something shifted." "Everything changed." "But here's the thing." "Here's the part most people miss." Cut them or earn them with a concrete detail.
- **Corporate neutrality.** No personality, no stakes, could have been written by any company about any product.

## 5. Formatting

- **Bold everywhere** for emphasis. A little goes far; AI over-bolds. Match the surrounding document.
- **Emoji as section markers** (🚀 ✨ 🔥 📌) or sprinkled through prose. Out, unless the venue genuinely calls for it and he does it (an emoticon `:)` in a personal note is fine; rocket-ship headers are not).
- **Title Case Headers.** He uses sentence case. "Network challenges and solutions", not "Network Challenges And Solutions". (Note some of his own older posts slip into Title Case; sentence case is the better default.)
- **Everything forced into lists and tables** when prose would read better. Lists are for genuinely enumerable things.
- **A "Key Takeaways" / "TL;DR" box** bolted on by reflex.
- **Decorative Unicode and stray structure.** Inline arrows/operators (→, ×, ✓) used as prose glue, a `---` horizontal rule dropped in before a heading, heading levels skipped (H2 straight to H4). All read as machine formatting.
- **Markdown leakage into plain-text media.** `**bold**`, `###`, or `---` showing up in a plain email, a Slack message, a commit message, a form field, or anywhere markdown doesn't render. Dead giveaway of an LLM source. Strip formatting to suit the actual destination.

## 6. What detectors and forensic readers actually key on

- **Low burstiness**: uniform sentence length and low variance. Fix by mixing long and very short sentences deliberately.
- **Low perplexity**: the most predictable next word every time, no surprising diction or idiosyncratic phrasing. Fix with concrete specifics, a real aside, an unexpected but apt word.
- **Absence of specific, checkable detail**: AI stays general. Names, numbers, dates, model numbers, the actual error message: these read human and AI tends to omit them.
- **No first-person stakes or lived detail**: the writer never got annoyed, never lost data, never made a call they regret. Dylan's writing is full of these.
- **Perfect, sanded-smooth grammar with zero voice.** A stray fragment, a sentence that starts with "But", a dry one-liner: these are features, not flaws.

## Caveats: don't over-correct

- **No single marker is proof, and the aim isn't a clean detector score.** Humans use em dashes, triads, and "delve" naturally. The tell is *saturation and uniformity*, not any one instance. Fixing the voice matters; gaming a detector doesn't.
- **Don't just file the serial numbers off.** Swapping an em dash for a spaced hyphen, or "delve" for "explore", while keeping the uniform rhythm and hollow balance, still reads as AI. Fix the underlying reflex.
- For Dylan specifically the em-dash rule is a hard *no*, because he genuinely never uses them (verified across everything he has hand-typed). His site does now contain machine-published text with em dashes (bot-posted travel entries, assisted posts); that text is not the model. See `corpus.md` for how to authenticate a sample before imitating it.
- Dylan-specific tells beyond the general list: uncontracted copulas ("It is easy to view"), flawless spelling with zero loose grammar, US spellings ("behavior", "utilize"), Title Case headers, and section scaffolding on a short post. Any of these mean the text wasn't his hand, even if it's published under his name.

## What actually reads as human (aim for these, not just away from the tells)

The strongest human signals are the inverse of the list above:
- **Burstiness**: hard variation in sentence and paragraph length. A long clause-heavy sentence, then a three-word one.
- **Specificity**: concrete, checkable, idiosyncratic detail. Real names, numbers, dates, the actual error. AI stays general; this is the single biggest giveaway of a human.
- **A committed stance**: a real opinion, real doubt, or "I don't know", instead of hollow balance.
- **Register consistency**: no helpful-assistant voice bleeding into an essay or a message.
- **Restraint**: sparse bold, sparse hedging, sparse triads. Emphasis that's load-bearing because it's rare.

## The test

Read the draft and ask: would this survive next to "Oh, I just formatted the drives. Crap. This was probably the hardest lesson to learn... What a dumb mistake."? If the draft is smoother, more balanced, more enthusiastic, and more generic than that line, it's wrong. Rough it up toward the real voice.

## Sources

Assembled from web research (2025-2026). Fuller list and quotes worth a look:
- [Wikipedia: Signs of AI writing](https://en.wikipedia.org/wiki/Wikipedia:Signs_of_AI_writing). the most exhaustive single catalogue.
- [Gorrie, "Why ChatGPT writes like that"](https://www.deadlanguagesociety.com/p/rhetorical-analysis-ai) and [Guo, "The Field Guide to AI Slop"](https://www.ignorance.ai/p/the-field-guide-to-ai-slop). rhetorical and tonal tics.
- [GC AI: contrastive negation](https://gc.ai/blog/ai-writing-pattern-to-know-contrastive-negation) and [GPTZero: the rule of three](https://gptzero.me/news/the-rule-of-three/). the structural swings.
- [GPTZero: perplexity & burstiness](https://gptzero.me/news/perplexity-and-burstiness-what-is-it/). the statistics detectors key on.
- [PubMed word-overuse study (PMC12679996)](https://pmc.ncbi.nlm.nih.gov/articles/PMC12679996/) and [FSU on "delve"](https://news.fsu.edu/news/science-technology/2025/02/17/why-does-chatgpt-delve-so-much-fsu-researchers-begin-to-uncover-why-chatgpt-overuses-certain-words/). quantitative vocabulary evidence.


