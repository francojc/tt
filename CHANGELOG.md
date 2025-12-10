# 0.5.0:

Major configuration and output improvements:

## Breaking Changes
- **Config format changed** from JSON to YAML:
  - Config file location: `~/.config/tt/config.yaml` (previously config.json)
  - Existing JSON config files will be ignored
  - YAML format offers better readability and commenting support

- **Installation method changed**:
  - Pre-built binaries removed from releases
  - Users must build from source using Go
  - Updated installation instructions provided

## New Features
- **ZenQuotes API integration**:
  - Use `-quotefile zen` to fetch random quotes from ZenQuotes API
  - Fallback mechanisms ensure smooth operation
  - No API key required for basic usage

- **Enhanced CSV logging**:
  - Added 'n' field to track test group size in CSV stats
  - Added file tracking to CSV output
  - CSV format now includes: `timestamp,wpm,cpm,accuracy,file,n`

## Improvements
- **YAML configuration support**:
  - All configuration options now supported in YAML
  - Comments and inline documentation in config file
  - Automatic config file creation with defaults
  - XDG Base Directory compliant

- **Flag behavior improvements**:
  - Better handling of `-words` vs `-wordfile`
  - Improved `-quotes` vs `-quotefile` distinction
  - Clearer error messages for invalid combinations

- **CSV file output**:
  - CSV output now writes to files instead of stdout
  - Stats file: `~/.local/share/tt/results/{mode}-stats.csv`
  - Errors file: `~/.local/share/tt/results/{mode}-errors.csv`
  - Headers automatically added on first write

## Fixes
- Config defaults now properly applied with correct precedence
- CSV file headers correctly managed
- Mode flag handling improved
- Binary file removed from repository (reduced bloat)

**Note**: This release includes breaking changes that may require users to update their configuration and scripts.

## Additional Cleanup Changes
- **Removed Ctrl-C exit handling** - Use Esc to exit for consistent behavior
- **Updated directory structure** - Custom resources now use `~/.config/tt/` (XDG compliant)
- **Removed non-functional features**:
  - Left/Right arrow navigation (never worked)
  - Outdated curl example (now core feature with `-quotefile zen`)
- **Documentation updates** to reflect current functionality

# 0.4.4:

Enhanced keyboard shortcuts and CSV functionality:

- **Changed keyboard shortcuts**:
  - `Esc` now quits the application (previously restarted test)
  - `Tab` restarts test during active test, starts new test on results screen (previously unused)
  - `Ctrl-W` deletes previous word during typing (already existed, now documented)

- **Enhanced CSV output** (Breaking change):
  - `-csv` flag now writes to files instead of stdout
  - Stats file: `~/.local/share/tt/results/{mode}-stats.csv` (timestamp,wpm,cpm,accuracy)
  - Errors file: `~/.local/share/tt/results/{mode}-errors.csv` (timestamp,word,error)
  - Headers automatically added on first write
  - Supports all test modes: words, quotes, file, stdin

- **Config file support**:
  - Configuration file: `~/.config/tt/config.json`
  - Customize CSV output directory with `{"csvdir": "/custom/path"}`
  - Supports tilde expansion (`~/Documents/stats`)
  - XDG Base Directory compliant

**Note**: The `-csv` flag behavior change is a breaking change. Users with scripts that relied on stdout output will need to update them.

# 0.4.3:

Code modernization and build improvements:

- Removed deprecated io/ioutil package usage (deprecated since Go 1.16)
- Updated to modern Go stdlib equivalents:
  - ioutil.ReadAll() → io.ReadAll()
  - ioutil.ReadFile() → os.ReadFile()
  - ioutil.WriteFile() → os.WriteFile()
  - ioutil.Discard → io.Discard
- Removed unnecessary rand.Seed() call (auto-initialized in Go 1.20+)
- Added clean target to Makefile

# 0.4.2:

Added -notheme, -blockcursor and -bold.

# 0.4.0:
  Too numerous to list (see the man page)

  Highlights:

 - Added -quotes.
 - Added support for navigating between tests via right/left.
 - Now store the user's position within a file if one is specified.
 - Improved documentation.

# 0.3.0:
 - Added support for custom word lists (`-words).
 - `-theme` now accepts a path.
 - Added `~/.tt/themes` and `~/.tt/words`.
 - Scrapped ~/.ttrc in favour of aliases/flags.
 - Included more default word lists. (`-list words`)

# 0.2.2:
 - Modified -g to correspond to the number of groups rather than the group size.
 - Added -multi
 - Added -v
 - Changed the default behaviour to restart the currently generated test rather than generating a new one
 - Added a CHANGELOG :P
