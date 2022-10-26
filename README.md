# timplit

This is as easy as it gets! CLI for go templates.

Read JSON data from stdin and apply it to the template file. Alternatively,
read the template from stdin and the JSON data from a file!

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

