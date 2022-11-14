---
name: getmemoryinfo
btcversion: 24.0.0
btcgroup: control
permalink: en/doc/24.0.0/rpc/control/getmemoryinfo/
---

getmemoryinfo ( "mode" )

Returns an object containing information about memory usage.

Arguments:
1. mode    (string, optional, default="stats") determines what kind of information is returned.
           - "stats" returns general statistics about memory usage in the daemon.
           - "mallocinfo" returns an XML string describing low-level heap state (only available if compiled with glibc 2.10+).

Result (mode "stats"):
{                         (json object)
  "locked" : {            (json object) Information about locked memory manager
    "used" : n,           (numeric) Number of bytes used
    "free" : n,           (numeric) Number of bytes available in current arenas
    "total" : n,          (numeric) Total number of bytes managed
    "locked" : n,         (numeric) Amount of bytes that succeeded locking. If this number is smaller than total, locking pages failed at some point and key data could be swapped to disk.
    "chunks_used" : n,    (numeric) Number allocated chunks
    "chunks_free" : n     (numeric) Number unused chunks
  }
}

Result (mode "mallocinfo"):
"str"    (string) "<malloc version="1">..."

Examples:
> bitcoin-cli getmemoryinfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getmemoryinfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


