When I ran a SHA-256 checksum on this text file, I got this hash:

- 36BBE50ED96841D10443BCB670D6554F0A34B761BE67EC9C4A8AD2C0C44CA42C

`shasum -a 256 /path/to/file`

or using Powershell

`Get-FileHash <filepath> -Algorithm MD5`

- change the file by one character, then run SHA-256 again
  - 29FE5D4CE9CCDEDD53FDFEFB52FAFCB774B9FF9EC10F06CEF6AD9229879E50C0