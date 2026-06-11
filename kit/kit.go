// Package kit is a generic library of daisyUI-styled templ elements.
// Apps compose these elements with data; the HTML structure, classes, and
// behaviour all live here so call sites contain no front-end code.
package kit

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

// Tone is the daisyUI semantic colour of an element.
type Tone string

const (
	Primary Tone = "primary"
	Neutral Tone = "neutral"
	Success Tone = "success"
	Warning Tone = "warning"
	Error   Tone = "error"
	Ghost   Tone = "ghost"
)

// Size is the daisyUI size step of an element.
type Size string

const (
	XS Size = "xs"
	SM Size = "sm"
	LG Size = "lg"
)

// Page content widths.
const (
	WidthNarrow = "max-w-2xl"
	WidthMid    = "max-w-3xl"
	WidthWide   = "max-w-5xl"
)

func classes(parts ...string) string {
	out := parts[:0:0]
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, " ")
}

func iff(cond bool, s string) string {
	if cond {
		return s
	}
	return ""
}

func sized(prefix string, s Size) string {
	return iff(s != "", prefix+"-"+string(s))
}

func toned(prefix string, t Tone) string {
	return iff(t != "", prefix+"-"+string(t))
}

func itoa(n int) string { return strconv.Itoa(n) }

func jsEscape(s string) string {
	return strings.NewReplacer(`\`, `\\`, `'`, `\'`).Replace(s)
}

// ── shell ──────────────────────────────────────────────────────

// NavLink is a top-level app section; Icon names one of the built-in icons,
// drawn in the mobile bottom nav.
type NavLink struct {
	Label, Href, Icon string
}

// MenuEntry is one row of the avatar dropdown: a link, or a POST action
// rendered as a single-button form.
type MenuEntry struct {
	Label, Href, PostAction string
}

// Nav configures the AppShell chrome.
type Nav struct {
	Brand     string
	BrandHref string
	Links     []NavLink
	Active    string // Label of the active link
	Initials  string // avatar placeholder text
	Menu      []MenuEntry
}

// ── buttons ────────────────────────────────────────────────────

type BtnOpts struct {
	Tone     Tone
	Size     Size
	Outline  bool
	Ghost    bool
	Square   bool
	Join     bool
	Disabled bool
	Danger   bool // destructive label colouring
	Fit      bool // shrink to content width inside a column
	NewTab   bool // links only

	// form wiring
	Name, Value, Form, FormAction, Confirm string
	NoSubmit                               bool // type="button"
}

func (o BtnOpts) class() string {
	return classes(
		iff(o.Join, "join-item"),
		"btn",
		toned("btn", o.Tone),
		iff(o.Ghost, "btn-ghost"),
		iff(o.Outline, "btn-outline"),
		sized("btn", o.Size),
		iff(o.Square, "btn-square"),
		iff(o.Disabled, "btn-disabled"),
		iff(o.Danger, "text-error/70"),
		iff(o.Fit, "w-fit"),
	)
}

func (o BtnOpts) attrs() templ.Attributes {
	a := templ.Attributes{}
	if o.Name != "" {
		a["name"] = o.Name
	}
	if o.Value != "" {
		a["value"] = o.Value
	}
	if o.Form != "" {
		a["form"] = o.Form
	}
	if o.FormAction != "" {
		a["formaction"] = o.FormAction
		a["formnovalidate"] = true
	}
	if o.Confirm != "" {
		a["onclick"] = "return confirm('" + jsEscape(o.Confirm) + "')"
	}
	if o.NoSubmit {
		a["type"] = "button"
	}
	if o.NewTab {
		a["target"] = "_blank"
	}
	return a
}

func onclickRemove(selector string) templ.Attributes {
	return templ.Attributes{"onclick": "this.closest('" + jsEscape(selector) + "').remove()"}
}

// ── forms ──────────────────────────────────────────────────────

type FormOpts struct {
	ID     string
	Action string
	Join   bool // render buttons as a joined group
	Row    bool // lay children out in a horizontal row
}

func (o FormOpts) class() string {
	return classes(iff(o.Join, "join"), iff(o.Row, "flex gap-2 items-center"))
}

func (o FormOpts) attrs() templ.Attributes {
	a := templ.Attributes{}
	if o.ID != "" {
		a["id"] = o.ID
	}
	return a
}

type InputOpts struct {
	Name, Value, Placeholder   string
	Size                       Size
	W                          string // width: "full" or a width step like "24"
	Mono                       bool
	TokenTarget                bool // receives tokens inserted by TokenLine clicks
	Form, Pattern, Title, List string
	Required                   bool
}

func (o InputOpts) class() string {
	return classes(
		"input", "input-bordered",
		sized("input", o.Size),
		iff(o.Mono, "font-mono"),
		iff(o.Mono && o.Size == SM, "text-xs"),
		iff(o.W != "", "w-"+o.W),
		iff(o.TokenTarget, "template-input"),
	)
}

func (o InputOpts) attrs() templ.Attributes {
	a := templ.Attributes{"name": o.Name, "value": o.Value}
	if o.Placeholder != "" {
		a["placeholder"] = o.Placeholder
	}
	if o.Form != "" {
		a["form"] = o.Form
	}
	if o.Pattern != "" {
		a["pattern"] = o.Pattern
	}
	if o.Title != "" {
		a["title"] = o.Title
	}
	if o.List != "" {
		a["list"] = o.List
	}
	if o.Required {
		a["required"] = true
	}
	return a
}

// Option is one select choice; Value doubles as the label when Label is empty.
type Option struct {
	Value string
	Label string
}

type OptGroup struct {
	Label   string
	Options []Option
}

// Refresh re-renders the enclosing row from the server when the value changes.
type Refresh struct {
	URL string
	Row string // row marker selector, e.g. ".cond-row"
}

type SelectOpts struct {
	Name, Form, Selected string
	Size                 Size
	Mono, Fit            bool
	MobileWide           bool // span the full row on small screens
	Refresh              *Refresh
}

func (o SelectOpts) class() string {
	return classes(
		"select", "select-bordered",
		sized("select", o.Size),
		iff(o.Mono, "font-mono text-xs"),
		iff(o.Fit, "w-fit"),
		iff(o.MobileWide, "col-span-2 sm:col-span-1"),
	)
}

func (o SelectOpts) attrs() templ.Attributes {
	a := templ.Attributes{"name": o.Name}
	if o.Form != "" {
		a["form"] = o.Form
	}
	if r := o.Refresh; r != nil {
		a["hx-get"] = r.URL
		a["hx-include"] = "closest " + r.Row
		a["hx-target"] = "closest " + r.Row
		a["hx-swap"] = "outerHTML"
	}
	return a
}

type ToggleOpts struct {
	Name    string
	Checked bool
	Post    string // hx-post URL; fires on change without submitting a form
	Tone    Tone
	Size    Size
}

func (o ToggleOpts) class() string {
	return classes("toggle", sized("toggle", o.Size), toned("toggle", o.Tone))
}

func (o ToggleOpts) attrs() templ.Attributes {
	a := templ.Attributes{}
	if o.Name != "" {
		a["name"] = o.Name
	}
	if o.Post != "" {
		a["hx-post"] = o.Post
		a["hx-swap"] = "none"
	}
	return a
}

// ── text ───────────────────────────────────────────────────────

// TextOpts styles a one-off text run.
type TextOpts struct {
	XS, SM                bool
	Strong, Mono          bool
	Faint, Dim            bool // 60% / 40% opacity
	Block, Grow, Truncate bool
	Tone                  Tone
}

func (o TextOpts) class() string {
	return classes(
		iff(o.XS, "text-xs"),
		iff(o.SM, "text-sm"),
		iff(o.Strong, "font-semibold"),
		iff(o.Mono, "font-mono"),
		toned("text", o.Tone),
		iff(o.Faint && o.Tone == "", "opacity-60"),
		iff(o.Dim, "opacity-40"),
		iff(o.Grow, "flex-1"),
		iff(o.Truncate, "truncate"),
	)
}

func toneText(t Tone) string {
	if t == "" {
		return "opacity-40"
	}
	return "text-" + string(t)
}

// ── badges / cards / layout ────────────────────────────────────

type BadgeOpts struct {
	Tone    Tone
	Size    Size
	Outline bool
	Mono    bool
}

func (o BadgeOpts) class() string {
	return classes(
		"badge",
		toned("badge", o.Tone),
		iff(o.Outline, "badge-outline"),
		sized("badge", o.Size),
		iff(o.Mono, "font-mono"),
	)
}

type Crumb struct {
	Label, Href string
}

type BodyOpts struct {
	Pad, Gap int
	Center   bool
}

func (o BodyOpts) class() string {
	return classes(
		"card-body",
		iff(o.Pad > 0, fmt.Sprintf("p-%d", o.Pad)),
		iff(o.Gap > 0, fmt.Sprintf("gap-%d", o.Gap)),
		iff(o.Center, "items-center text-center"),
	)
}

type RowOpts struct {
	Gap                        int
	AlignEnd, AlignStart, Wrap bool
}

func (o RowOpts) class() string {
	align := "items-center"
	if o.AlignEnd {
		align = "items-end"
	}
	if o.AlignStart {
		align = "items-start"
	}
	return classes("flex", align, gapClass(o.Gap), iff(o.Wrap, "flex-wrap"))
}

func gapClass(n int) string {
	return iff(n > 0, fmt.Sprintf("gap-%d", n))
}

func growClass(minW int) string {
	return classes("flex-1", iff(minW > 0, fmt.Sprintf("min-w-%d", minW)))
}

func grid2Class(gapX, gapY int) string {
	return classes(
		"grid grid-cols-1 sm:grid-cols-2 items-start",
		iff(gapX > 0, fmt.Sprintf("gap-x-%d", gapX)),
		iff(gapY > 0, fmt.Sprintf("gap-y-%d", gapY)),
	)
}

// RowGrid is the column template of an EditRow.
type RowGrid string

const (
	// GridFieldOpValue lays out field / operator / value / remove.
	GridFieldOpValue RowGrid = "grid-cols-[1fr_1fr_2rem] sm:grid-cols-[1fr_9rem_1fr_2rem]"
	// GridKeyValue lays out key / value / remove.
	GridKeyValue RowGrid = "grid-cols-[10rem_1fr_2rem]"
)

// ── tables ─────────────────────────────────────────────────────

// Col is one table column; a non-zero W fixes its width to that step.
type Col struct {
	Label string
	W     int
}

func widthClass(w int) string {
	return iff(w > 0, fmt.Sprintf("w-%d", w))
}

func tableClass(s Size) string {
	return classes("table", sized("table", s))
}

type RowStyle struct {
	Hover bool
	Tint  bool
}

func (o RowStyle) class() string {
	return classes(iff(o.Hover, "hover"), iff(o.Tint, "bg-base-200/40"))
}

type CellOpts struct {
	XS, SM, Mono, Strong, Faint bool
	Right, NoWrap, Break        bool
	MaxW, Pad, Colspan          int
}

func (o CellOpts) class() string {
	return classes(
		iff(o.XS, "text-xs"),
		iff(o.SM, "text-sm"),
		iff(o.Mono, "font-mono"),
		iff(o.Strong, "font-semibold"),
		iff(o.Faint, "opacity-60"),
		iff(o.Right, "text-right"),
		iff(o.NoWrap, "whitespace-nowrap"),
		iff(o.Break, "break-all"),
		iff(o.MaxW > 0, fmt.Sprintf("max-w-%d", o.MaxW)),
		iff(o.Pad > 0, fmt.Sprintf("p-%d", o.Pad)),
	)
}

func (o CellOpts) attrs() templ.Attributes {
	a := templ.Attributes{}
	if o.Colspan > 0 {
		a["colspan"] = itoa(o.Colspan)
	}
	return a
}

// ── icons ──────────────────────────────────────────────────────

var iconPaths = map[string]string{
	"activity": "M3 12h4l3-8 4 16 3-8h4",
	"list":     "M4 6h16M4 12h10M4 18h7",
	"layers":   "M4 7l8-4 8 4-8 4-8-4zM4 12l8 4 8-4M4 17l8 4 8-4",
}

func iconPath(name string) string { return iconPaths[name] }
