# The corpus: what Dylan's writing actually looks like

This file is the evidence base behind `voice-profile.md`. It maps which text under his name is genuinely his hand, which is machine-drafted, and the fingerprints that separate the two. Built from a full sweep (August 2026) of his blog, travel journals, career notes, READMEs, commit messages, PR/issue text and review comments, Slack, iMessage, and sent email.

## The two eras: authenticating a sample

Most text published under his name after 2025 is machine-drafted or machine-assisted. Before modelling any sample, date it and check it against these markers.

**His hand (verified across everything hand-typed):**
- Zero em dashes, zero en dashes, zero semicolons. Without exception.
- Australian spelling that never slips: realised, utilising, organisation, behaviour, fibre. He even writes "practise" where "practice" the noun would be standard.
- Contractions everywhere. Uncontracted copulas ("It is easy to view", "You are a single worker") do not appear in anything verifiably hand-typed.
- Typos and grammar slips left in, in published posts: "seperate", "it's" for "its" (both directions), "receieved", "remeber", "We now how 2 more emu files", "This works fine bit is a little less than ideal", "the workflow was execute", "was ran". He does not proofread to perfection and does not care.
- Comma splices as a default connective: "I run the job, the runner picks this up and starts to run no issues."
- Digits for small numbers, always: "2 files, 1 in emu_prod.go", "There's 2 ways", "47 line single Go file", "4 physical locations".

**Machine text under his name (do not model):**
- Em dashes, spaced or unspaced. This is now the single most reliable marker that a given text was not hand-typed by him.
- Flawless spelling, or any US spelling ("utilize" appears only in the AI-assisted swarm-gpus post, "behavior" only in the sapphire post).
- Section scaffolding (Introduction / Assumptions / Conclusion), Title Case headers, tidy parallel triads, uncontracted formal copulas.
- Known assisted texts: "Adding GPUs to Docker Swarm" (2024-09), the sapphire post "Building a GBA Emulator in Go" (2026-02), **"Protect Your Shed" (2026-04)**, the asia-2026 travel entries published by his postcard pipeline, and most post-2025 READMEs (coach, emerald, bolt, among others).

**On "Protect Your Shed" specifically:** it went in hand-typed through Pages CMS over several drafts with no bot trailers, and the ideas and the sustained shed/skyscraper metaphor are his. But it has zero typos, repeated uncontracted copulas, tidy triads, a mirrored close, and one "paycheck" Americanism: none of which appear in anything else he hand-wrote. Treat it as AI-drafted from his ideas and then owned by him. Its *structure* (one committed metaphor, concrete stakes, a hard closing line) is worth modelling; its sanded-smooth *sentence surface* is not his hand. The earlier voice profile leaned on it as "his best writing"; weight it accordingly.

## The gold-standard hand-written set

- Blog, 2024–2025: "Go range over funcs" (2024-05), "Self-hosted traffic ingress" (2024-05), "Self-Dependant, Self-Hosted GitHub Runners" (2024-06), "App debugging not in prod" (2024-06), "12 Months On: Self-Hosting" (2025-05). Local source: `dbut2.github.io/content/blog/`.
- Travel journals, 2024–2025: written as GitHub issues on `devhou-se/www-jp` (the on-site pages are machine translations; the issues are the originals). Titles like "Tilted Tokyo", "The Mighty Hills of Otanacho", "Stranded in Takashima", "Leg Amputation".
- Raw notes: a private first-draft notes file in an internal repo (path and quotes withheld from this public file), ~173 lines of unedited prose. The single rawest sample of his voice that exists; consult it locally, never quote it here.
- Pre-2024 READMEs and micro-copy: butla, advent-of-code, SlackGPT, slack-diffusion, fetch, commuter (public repos only; a few private READMEs share the register and stay unquoted here).
- Commit messages, PR bodies, and issue text across his repos, 2016–2026 (excluding anything with Claude trailers).
- Casual chat: his Slack workspace. Ultra-terse, lowercase, unpunctuated.

## Fingerprint catalogue

### Punctuation
- Commas do enormous work; comma splices are routine and are one of his strongest fingerprints.
- The question-then-answer transition: "What happened? When the workflow was execute…", "The issue? This also cancels…", "As for what's next? I'm not sure.", "what manages these runners? Another set of runners?"
- Parentheses for asides and caveats: a mid-paragraph bracket carrying a date, a confession, or a deflating qualifier, then the sentence carries on. (Best examples are in the private notes; the habit shows across registers.)
- Exclamation marks: never in technical writing, freely in travel/personal writing ("It was spectacular!", "so that was fun!", "We managed to get two meals with drinks for just $16!").
- Straight quotes by habit; the few curly apostrophes in published posts are editor paste artifacts, not a preference.

### Sentence shapes
- Long chains glued with "and", "that", "in that", "with", "though", past the point a style guide would allow, then a short flat sentence to land: "…it was all gone. What a dumb mistake."
- Fronted participial openers, sometimes dangling: "Wanting to keep this configuration held in GitHub, …", "Exiting the tunnel, cloudflare has been setup to…", "Rerunning the workflow run for good measure, I wanted to…".
- "Now" and "Then" as pivots at least as often as But/And/So: "Now the runners have been shut down…", "Then I get the following errors:".
- Narrative present tense at the dramatic moment of a build log: "I run the job, the runner picks this up… Then I get the following errors:".
- "There's" with plural nouns: "There's two things here", "There's 2 ways".
- Fragments are register-dependent: build logs use them for the punch ("Done." "Crap."), travel prose runs on momentum and saves the fragment for a punchline ("Classic." "Boy was I wrong.").

### Hedging and stance
- A softener cloud around claims: kind of, basically, mostly, probably, a little, somewhat, fairly, at least, just, "for whatever reason". The mid-clause hedge is very him: "resolved issues, mostly, for services that need it".
- Real self-assessment in both directions, unhedged when it counts: pride in simple things stated plainly, and "What a dumb mistake." when it went wrong.
- Honest uncertainty, plainly: "As for what's next? I'm not sure. I'm happy with how everything is looking at the moment."

### Humour
- Deadpan, never signposted: "Stupid simple / Stupid fast", "if you need a listening ear, or added to a channel for a team's worth of fun", "every second counts on that leaderboard".
- Self-deprecating misadventure, reported flat: "It was funny, but not really. I asked if they were okay and they said yes and got back up and on their way."
- Issue/PR titles as jokes: "What the heck", "Leg Amputation", "devhouse 2!!!".
- Current-era irony about AI, dry and self-aware: "slop gallery".

### Vocabulary and set phrases
- Recurring: "up and running", "here and there", "the only caveat being", "point of contention", "first point of order", "collecting dust", "for good measure", "so long as", "with that all said and done".
- Travel enthusiasm vocabulary (never in tech writing): "boy oh boy", "Boy was I wrong", "absolute bangers", "criminally cheap", "nice and early", "smashed down", "slog out the distance", "amazing", "delicious", and "little" as habitual diminutive ("a little peek", "a little adventure", "a little stressed").
- Abbreviations in notes and chat: "bc", "eg.", "etc", "gha", "u", "ur", "deets".

## Quote bank

### Build log / walkthrough (the default technical register)
> Step 1 to install proxmox, format drives and create a fresh installation. Done. Step 2, boot up and create VMs. Done. Step 3, set up the VMs and get all of the previous services back up and running. Oh, I just formatted the drives. Crap.

> After a few months of increasingly larger GCP bills, and a final month of accidentally leaving a Cloud Workstation running for a little too long, I decided to give the old self-hosted world a try again.

> Now the runners have been shut down, and no new runners have been started, this means I'm left with no runners. This is a problem.

> This isn't really an issue as docker compose still completes successfully but is annoying to look at a red cross when you know everything's ok. But that's an issue for another day.

> Though I'm not quite sure how this will fit into the Go ecosystem, and not quite yet sold on the proposal, should this go ahead hopefully we get it before advent of code this year as part of go1.23 release in August, every second counts on that leaderboard.

> A quick snippet of my docker compose config for managing traffic ingress into my self-hosted setup.

### Travel journal
> I had forgotten to download Japanese maps before I had set off, so I had no good way to guide me other than checking my phone, but this would be too finicky so I decided to just vibe it instead.

> It was not okay. The path that I had been following eventually just ended and I was dumped on the local streets to figure out myself.

> We spend the next few hours blowing out our vocal cords and belting out some absolute bangers. … I came clean the next day that I was actually using my maps.

> we stopped and had some lunch and a small place and we had some sushi, which I followed with a lemon sour. Classic.

> Today was probably the hardest of them all but that probably just makes it all the more worth it, right.

> Until next time! See you in Bali!

> See you then :)

### Raw notes (first-draft voice)
Quotes withheld: the source is private. The register in one line: long comma-chained sentences that follow the thought wherever it goes, "bc" and digit-numbers throughout, hedges mid-clause, zero proofreading, opinions stated flat. It reads like his chat register stretched into paragraphs.

### Commit messages
`moar rules` · `packageify` · `dedebug` (after six commits titled `debug`) · `the big refactor of 2023` · `stupid simple, stupid fast` · `sweepin` · `thwop` · `nap` · `rem deets` · `anti-vpn` · `2025 day 4: sped` · `fix 2: submit pipeline` · `hide controls on mobile`

Default: all-lowercase, 1–4 words, imperative when a verb is present, no trailing full stop, happily repeated verbatim while iterating. From ~2026 he also uses conventional prefixes (`fix:`, `feat:`, `ci:`, plus a personal `tech:`).

### PR / issue text
> Introduce the builtin `testing` package and replace current "examples" with tests, swapping panics for formatted errors

> rename package and fix compile errors to resolve examples not showing up in godocs. tested locally

> Add datepicker to form, defaults to current date else if nothing supplied backend will use time.Now() as was previous / Resolves #4

> self-register hijacks strava webhook target, which is conflicting with my current setup, maybe ill add as optional config later

> As titled, user should be able to submit presence for prior days missed

Pattern: leads with what changed, no greeting, no headings, terse hyphen bullets without trailing stops, a bare "Resolves #4"/"Requires #3", and habitual testing evidence ("tested locally", "Tested in fork"). Single-sentence bodies frequently end without a full stop.

### PR review comments / technical argument
Public-repo samples only; his sharper reviews live on private repos and follow the same shape.

> Have a look into using reflection here, I think it may be more reliable.

> is this attachable for external services?

> What's the server impact of a single text generation? The mental work required for users here is much worth the generation IMO

> Thank you for the notice, this does appear to be an issue with the current hosting provider. Will look into investigating some other options 👍

Pattern: opens on the objection or question, no softener, no "nit:". "Why not just X" is his question shape. A `suggestion` block or one blunt line beats a paragraph, and "No." on its own is a complete reply when he means it. Disagreement is argued straight, no hedging into "have we considered". Warmth is reserved for outside contributors ("As always, appreciate the open source contributions :)").

### Chat (Slack / texts)
Ultra-terse, lowercase, minimal punctuation, no full stops, question marks optional, single-word replies common ("ok", "not yet", "lol"). "u"/"ur" abbreviations. Longer thoughts arrive comma-spliced rather than structured. Dry one-liners over emoji; when emoji appear they're sparse (👍, ":((").

iMessage (verified from sent messages, quotes withheld: private correspondence) is the same voice with phone casing: auto-capitalised, apostrophes intact via autocorrect, still no question marks on quick questions, still comma-spliced. Laughter is "Hahah"/"hahaha"/"lol", emoji sparse and sincere (a single 😭 or 👍, never decoration), banter blunt but affectionate. Casing follows the device, not the register. One thought per message, split across several short messages rather than composed into a paragraph. When answering a numbered list of questions he answers with a numbered list, compact, no restating.

### Email
Verified from his sent mail (a thin corpus, he barely uses email; quotes withheld: private correspondence). The skeleton never varies: "Hi {first name}," then straight into the matter in one or two sentences, then "Thanks" and "Dylan" on their own lines. No "hope you're well" even when the other side pads with it. Dropped articles and subjects, comma splices, the occasional slightly formal stock phrase but never pleasantry scaffolding.

## What was not captured

His 2023-era travel posts survive only as fragments quoted in an internal repo.
