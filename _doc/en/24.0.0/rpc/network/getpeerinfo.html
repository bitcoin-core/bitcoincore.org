---
name: getpeerinfo
btcversion: 24.0.0
btcgroup: network
permalink: en/doc/24.0.0/rpc/network/getpeerinfo/
---

getpeerinfo

Returns data about each connected network peer as a json array of objects.

Result:
[                                         (json array)
  {                                       (json object)
    "id" : n,                             (numeric) Peer index
    "addr" : "str",                       (string) (host:port) The IP address and port of the peer
    "addrbind" : "str",                   (string, optional) (ip:port) Bind address of the connection to the peer
    "addrlocal" : "str",                  (string, optional) (ip:port) Local address as reported by the peer
    "network" : "str",                    (string) Network (ipv4, ipv6, onion, i2p, cjdns, not_publicly_routable)
    "mapped_as" : n,                      (numeric, optional) The AS in the BGP route to the peer used for diversifying
                                          peer selection (only available if the asmap config flag is set)
    "services" : "hex",                   (string) The services offered
    "servicesnames" : [                   (json array) the services offered, in human-readable form
      "str",                              (string) the service name if it is recognised
      ...
    ],
    "relaytxes" : true|false,             (boolean, optional) Whether peer has asked us to relay transactions to it
    "lastsend" : xxx,                     (numeric) The UNIX epoch time of the last send
    "lastrecv" : xxx,                     (numeric) The UNIX epoch time of the last receive
    "last_transaction" : xxx,             (numeric) The UNIX epoch time of the last valid transaction received from this peer
    "last_block" : xxx,                   (numeric) The UNIX epoch time of the last block received from this peer
    "bytessent" : n,                      (numeric) The total bytes sent
    "bytesrecv" : n,                      (numeric) The total bytes received
    "conntime" : xxx,                     (numeric) The UNIX epoch time of the connection
    "timeoffset" : n,                     (numeric) The time offset in seconds
    "pingtime" : n,                       (numeric, optional) ping time (if available)
    "minping" : n,                        (numeric, optional) minimum observed ping time (if any at all)
    "pingwait" : n,                       (numeric, optional) ping wait (if non-zero)
    "version" : n,                        (numeric) The peer version, such as 70001
    "subver" : "str",                     (string) The string version
    "inbound" : true|false,               (boolean) Inbound (true) or Outbound (false)
    "bip152_hb_to" : true|false,          (boolean) Whether we selected peer as (compact blocks) high-bandwidth peer
    "bip152_hb_from" : true|false,        (boolean) Whether peer selected us as (compact blocks) high-bandwidth peer
    "startingheight" : n,                 (numeric, optional) The starting height (block) of the peer
    "presynced_headers" : n,              (numeric, optional) The current height of header pre-synchronization with this peer, or -1 if no low-work sync is in progress
    "synced_headers" : n,                 (numeric, optional) The last header we have in common with this peer
    "synced_blocks" : n,                  (numeric, optional) The last block we have in common with this peer
    "inflight" : [                        (json array, optional)
      n,                                  (numeric) The heights of blocks we're currently asking from this peer
      ...
    ],
    "addr_relay_enabled" : true|false,    (boolean, optional) Whether we participate in address relay with this peer
    "addr_processed" : n,                 (numeric, optional) The total number of addresses processed, excluding those dropped due to rate limiting
    "addr_rate_limited" : n,              (numeric, optional) The total number of addresses dropped due to rate limiting
    "permissions" : [                     (json array) Any special permissions that have been granted to this peer
      "str",                              (string) bloomfilter (allow requesting BIP37 filtered blocks and transactions),
                                          noban (do not ban for misbehavior; implies download),
                                          forcerelay (relay transactions that are already in the mempool; implies relay),
                                          relay (relay even in -blocksonly mode, and unlimited transaction announcements),
                                          mempool (allow requesting BIP35 mempool contents),
                                          download (allow getheaders during IBD, no disconnect after maxuploadtarget limit),
                                          addr (responses to GETADDR avoid hitting the cache and contain random records with the most up-to-date info).
                                          
      ...
    ],
    "minfeefilter" : n,                   (numeric, optional) The minimum fee rate for transactions this peer accepts
    "bytessent_per_msg" : {               (json object)
      "msg" : n,                          (numeric) The total bytes sent aggregated by message type
                                          When a message type is not listed in this json object, the bytes sent are 0.
                                          Only known message types can appear as keys in the object.
      ...
    },
    "bytesrecv_per_msg" : {               (json object)
      "msg" : n,                          (numeric) The total bytes received aggregated by message type
                                          When a message type is not listed in this json object, the bytes received are 0.
                                          Only known message types can appear as keys in the object and all bytes received
                                          of unknown message types are listed under '*other*'.
      ...
    },
    "connection_type" : "str"             (string) Type of connection: 
                                          outbound-full-relay (default automatic connections),
                                          block-relay-only (does not relay transactions or addresses),
                                          inbound (initiated by the peer),
                                          manual (added via addnode RPC or -addnode/-connect configuration options),
                                          addr-fetch (short-lived automatic connection for soliciting addresses),
                                          feeler (short-lived automatic connection for testing addresses).
                                          Please note this output is unlikely to be stable in upcoming releases as we iterate to
                                          best capture connection behaviors.
  },
  ...
]

Examples:
> bitcoin-cli getpeerinfo 
> curl --user myusername --data-binary '{"jsonrpc": "1.0", "id": "curltest", "method": "getpeerinfo", "params": []}' -H 'content-type: text/plain;' http://127.0.0.1:8332/


