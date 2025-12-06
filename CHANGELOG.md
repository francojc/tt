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
