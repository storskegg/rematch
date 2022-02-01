# rematch
Just a little tool to print regex matches from piped input.

## Why?

I had a moment of frustration dealing with inconsistencies between `sed` versions, so I decided to write a tiny tool that does just this one thing.

## Usage

```bash
fortune | rematch [--option(s)] 'valid regex'
```

### Example Usage: 
- `--all`
  - Return all matches.
  - Default behavior returns the first match
- `--posix` 
  - Use POSIX regex.
  - Default behavior is PCRE
