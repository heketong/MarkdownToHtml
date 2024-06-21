// Package MarkdownToHtml is a markdown processor.
//
// It translates plain text with simple formatting rules into an AST, which can
// then be further processed to HTML (provided by Blackfriday itself) or other
// formats (provided by the community).
//
// The simplest way to invoke Blackfriday is to call the Run function. It will
// take a text input and produce a text output in HTML (or other format).
//
// A slightly more sophisticated way to use Blackfriday is to create a Markdown
// processor and to call Parse, which returns a syntax tree for the input
// document. You can leverage Blackfriday's parsing for content extraction from
// markdown documents. You can assign a custom renderer and set various options
// to the Markdown processor.
//
// If you're interested in calling Blackfriday from command line, see
// https://github.com/heketong/MarkdownToHtml-tool.
//
// Sanitized Anchor Names
//
// Blackfriday includes an algorithm for creating sanitized anchor names
// corresponding to a given input text. This algorithm is used to create
// anchors for headings when AutoHeadingIDs extension is enabled. The
// algorithm is specified below, so that other packages can create
// compatible anchor names and links to those anchors.
//
// The algorithm iterates over the input text, interpreted as UTF-8,
// one Unicode code point (rune) at a time. All runes that are letters (category L)
// or numbers (category N) are considered valid characters. They are mapped to
// lower case, and included in the output. All other runes are considered
// invalid characters. Invalid characters that precede the first valid character,
// as well as invalid character that follow the last valid character
// are dropped completely. All other sequences of invalid characters
// between two valid characters are replaced with a single dash character '-'.
//
// SanitizedAnchorName exposes this functionality, and can be used to
// create compatible links to the anchor names generated by MarkdownToHtml.
// This algorithm is also implemented in a small standalone package at
// github.com/shurcooL/sanitized_anchor_name. It can be useful for clients
// that want a small package and don't need full functionality of MarkdownToHtml.
package MarkdownToHtml

// NOTE: Keep Sanitized Anchor Name algorithm in sync with package
//       github.com/shurcooL/sanitized_anchor_name.
//       Otherwise, users of sanitized_anchor_name will get anchor names
//       that are incompatible with those generated by MarkdownToHtml.
