## find-bytes

This tool shows you the offset of some bytes in a file

```bash
kmille@linbox:find-bytes go run find-bytes.go -h
Usage of find-bytes:
  -b string
        bytes to look for in hex format
  -c    continue after first match
  -f string
        use bytes to look for from file
  -r    find recursive
```

For example if you have a disk image:

```bash
kmille@linbox:find-bytes file disk.img 
disk.img: DOS/MBR boot sector
kmille@linbox:find-bytes dd if=disk.img bs=512 count=1 | xxd
00000000: eb63 9010 8ed0 bc00 b0b8 0000 8ed8 8ec0  .c..............
00000010: fbbe 007c bf00 06b9 0002 f3a4 ea21 0600  ...|.........!..
00000020: 00be be07 3804 750b 83c6 1081 fefe 0775  ....8.u........u
00000030: f3eb 16b4 02b0 01bb 007c b280 8a74 018b  .........|...t..
00000040: 4c02 cd13 ea00 7c00 00eb fe00 0000 0000  L.....|.........
00000050: 0000 0000 0000 0000 0000 0080 0100 0000  ................
00000060: 0000 0000 fffa 9090 f6c2 8074 05f6 c270  ...........t...p
00000070: 7402 b280 ea79 7c00 0031 c08e d88e d0bc  t....y|..1......
00000080: 0020 fba0 647c 3cff 7402 88c2 52be 807d  . ..d|<.t...R..}
00000090: e817 01be 057c b441 bbaa 55cd 135a 5272  .....|.A..U..ZRr
000000a0: 3d81 fb55 aa75 3783 e101 7432 31c0 8944  =..U.u7...t21..D
000000b0: 0440 8844 ff89 4402 c704 1000 668b 1e5c  .@.D..D.....f..\
000000c0: 7c66 895c 0866 8b1e 607c 6689 5c0c c744  |f.\.f..`|f.\..D
000000d0: 0600 70b4 42cd 1372 05bb 0070 eb76 b408  ..p.B..r...p.v..
000000e0: cd13 730d 5a84 d20f 83d8 00be 8b7d e982  ..s.Z........}..
000000f0: 0066 0fb6 c688 64ff 4066 8944 040f b6d1  .f....d.@f.D....
00000100: c1e2 0288 e888 f440 8944 080f b6c2 c0e8  .......@.D......
00000110: 0266 8904 66a1 607c 6609 c075 4e66 a15c  .f..f.`|f..uNf.\
00000120: 7c66 31d2 66f7 3488 d131 d266 f774 043b  |f1.f.4..1.f.t.;
00000130: 4408 7d37 fec1 88c5 30c0 c1e8 0208 c188  D.}7....0.......
00000140: d05a 88c6 bb00 708e c331 dbb8 0102 cd13  .Z....p..1......
00000150: 721e 8cc3 601e b900 018e db31 f6bf 0080  r...`......1....
00000160: 8ec6 fcf3 a51f 61ff 265a 7cbe 867d eb03  ......a.&Z|..}..
00000170: be95 7de8 3400 be9a 7de8 2e00 cd18 ebfe  ..}.4...}.......
00000180: 4752 5542 2000 4765 6f6d 0048 6172 6420  GRUB .Geom.Hard 
00000190: 4469 736b 0052 6561 6400 2045 7272 6f72  Disk.Read. Error
000001a0: 0d0a 00bb 0100 b40e cd10 ac3c 0075 f4c3  ...........<.u..
000001b0: 0000 0000 0000 0000 c03d 65eb 0000 8004  .........=e.....
000001c0: 0104 83fe c2ff 0008 0000 0038 0f00 00fe  ...........8....
000001d0: c2ff 05fe c2ff fe47 0f00 02b0 9000 0000  .......G........
000001e0: 0000 0000 0000 0000 0000 0000 0000 0000  ................
000001f0: 0000 0000 0000 0000 0000 0000 0000 55aa  ..............U.
```

You can find the offset of the magic bytes (`55aa`)

```bash
kmille@linbox:find-bytes find-bytes -b 55aa disk.img 
Match at offset 163 (0xa3) until 165 (0xa5) in file "disk.img"
kmille@linbox:find-bytes find-bytes -c -b 55aa disk.img
Match at offset 163 (0xa3) until 165 (0xa5) in file "disk.img"
Match at offset 510 (0x1fe) until 512 (0x200) in file "disk.img"
kmille@linbox:find-bytes 
```

You can read the search bytes from a file

```bash
kmille@linbox:find-bytes echo -ne '\x55\xaa' > some-bytes
kmille@linbox:find-bytes xxd some-bytes                  
00000000: 55aa                                     U.
kmille@linbox:find-bytes find-bytes -f some-bytes disk.img 
Match at offset 163 (0xa3) until 165 (0xa5) in file "disk.img"
kmille@linbox:find-bytes 
```

Recursive grep is also possible

```bash
kmille@linbox:find-bytes find-bytes -r -b 55aa .    
Match at offset 163 (0xa3) until 165 (0xa5) in file "disk.img"
Match at offset 1359563 (0x14becb) until 1359565 (0x14becd) in file "find-bytes"
Match at offset 163 (0xa3) until 165 (0xa5) in file "tests/1/2/3/disk.img"
Match at offset 182429 (0x2c89d) until 182431 (0x2c89f) in file "tmp/find-bytes-0.1-1-x86_64.pkg.tar.zst"
Match at offset 1359563 (0x14becb) until 1359565 (0x14becd) in file "tmp/pkg/find-bytes/usr/local/bin/find-bytes"
Match at offset 1359563 (0x14becb) until 1359565 (0x14becd) in file "tmp/src/find-bytes/find-bytes"
kmille@linbox:find-bytes 

```

There is a PKGBUILD file for building a package on Arch Linux.

## KNOWN ISSUES

- the code does not contain concurrency right now

- files are read completely to memory before searching. This can lead to 

  ```bash
  kmille@linbox:grub-luks find-bytes -r -f mods-debian11/luks.mod .
  Could not read file "grub/include/grub/cpu": read grub/include/grub/cpu: is a directory
  Could not read file "grub/include/grub/machine": read grub/include/grub/machine: is a directory
  Match at offset 0 (0x0) until 6944 (0x1b20) in file "mods-debian11/luks.mod"
  Match at offset 188603392 (0xb3ddc00) until 188610336 (0xb3df720) in file "test-vm/_disk.img.extracted/100000"
  zsh: killed     find-bytes -r -f mods-debian11/luks.mod .
  
  kmille@linbox:tmp journalctl -f                                       
  -- Journal begins at Thu 2021-12-16 14:30:16 CET. --                                                                                                                     
  Dec 21 13:28:12 linbox kernel: [ 100350]  1000 100350     2732      341    61440      306             0 zsh                                                              
  Dec 21 13:28:12 linbox kernel: [ 100564]  1000 100564     2739        2    65536      641             0 zsh                                                              
  Dec 21 13:28:12 linbox kernel: [ 100844]  1000 100844      593        0    45056       25             0 xxd                                                              Dec 21 13:28:12 linbox kernel: [ 100845]  1000 100845     1685        0    53248       68             0 less                                                             
  Dec 21 13:28:12 linbox kernel: [ 109921]  1000 109921  5842410  3645393 43708416  1791476             0 find-bytes                                                       Dec 21 13:28:12 linbox kernel: oom-kill:constraint=CONSTRAINT_NONE,nodemask=(null),cpuset=/,mems_allowed=0,global_oom,task_memcg=/user.slice/user-1000.slice/session-1.sc
  ope,task=find-bytes,pid=109921,uid=1000                                                                                                                                  
  Dec 21 13:28:12 linbox kernel: Out of memory: Killed process 109921 (find-bytes) total-vm:23369640kB, anon-rss:14581572kB, file-rss:0kB, shmem-rss:0kB, UID:1000 pgtables
  :42684kB oom_score_adj:0                                                                                                                                                 
  Dec 21 13:28:12 linbox systemd[1]: session-1.scope: A process of this unit has been killed by the OOM killer.
  ```

  