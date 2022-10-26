# timplit

This is as easy as it gets! A minimal CLI for go templates.

Read JSON data from stdin and apply it to the template file. Alternatively,
read the template from stdin and the JSON data from a file!

## Top-level JSON array

If the top level of the incoming JSON is an array, it will get wrapped in an
object (under the `items` property) to ensure that it can be unmarshaled. 

For instance, `[1, 2, 3]` will become `{"items": [1, 2, 3]}`. See
`example/array.tmpl` and `example/array.json` for a practical demonstration.

## Examples

Read JSON from stdin and apply it to a template file:
```bash
cat data.json | timplit template.tmpl

# Same as
timplit template.tmpl <data.json
```

Read JSON from a file and apply it to the template from stdin:
```bash
# Invert the logic with the -j flag
cat template.tmpl | timplit -j data.json

# Same as
timplit -j data.json <template.tmpl
```

