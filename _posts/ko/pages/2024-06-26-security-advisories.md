---
title: 보안 권고
name: Security Advisories
id: ko-security-advisories
permalink: /ko/security-advisories/
layout: page
type: pages
lang: ko
version: 1
---
{% include toc.html %}

이 페이지는 Bitcoin Core 취약점 공개 정책을 요약하고, 과거 보안 권고(Security Advisories) 목록을 제공합니다.

## 정책

모든 취약점은 security@bitcoincore.org로 제보해 주세요([SECURITY.md](https://github.com/bitcoin/bitcoin/blob/master/SECURITY.md) 참고).
제보된 취약점은 심각도 기준으로 분류되며, 프로젝트는 4가지 등급을 사용합니다.

* **Critical**: 비트코인 네트워크 전체의 근본적인 보안/무결성을 위협하는 버그입니다.
  프로토콜 수준에서 코인 탈취를 가능하게 하거나, 발행 스케줄과 다른 코인 생성,
  또는 네트워크 전체의 영구적인 체인 분할을 유발할 수 있는 유형입니다.
  <details markdown="1">

  <summary>
  예시
  </summary>

  * 동일 트랜잭션 출력(output)을 블록 내에서 두 번 소비해 통화량을 부풀릴 수 있는 버그
    ([CVE-2018-17144](/en/2018/09/20/notice/)).
  * 기반 데이터베이스 제한 때문에 구버전 노드가 신버전 노드가 수용한 블록을 거부해
    네트워크 전반의 체인 분할을 유발한 합의 실패
    ([BIP 50](https://github.com/bitcoin/bips/blob/master/bip-0050.mediawiki)).

  </details>
* **High**: 영향을 받는 노드나 네트워크에 중대한 영향을 주는 버그입니다.
  일반적으로 기본 설정에서도 원격 악용이 가능하며 광범위한 장애를 일으킬 수 있습니다.
  <details markdown="1">

  <summary>
  예시
  </summary>

  * 원격 트리거로 다수 노드를 오프라인 상태로 만들 수 있는 크래시
    ([CVE-2024-35202](/en/2024/10/08/disclose-blocktxn-crash/)).
  * 장시간 노드를 정지 상태로 만들어 신규 트랜잭션과 블록 처리를 방해하는
    서비스 거부(DoS) 공격
    ([CVE-2024-52914](/en/2024/07/03/disclose-orphan-dos/)).
  * 과도한 블록 헤더 저장을 유도해 원격으로 노드를 크래시시킬 수 있는
    메모리 고갈 취약점
    ([CVE-2019-25220](/en/2024/09/18/disclose-headers-oom/)).

  </details>
* **Medium**: 네트워크 또는 노드 성능/기능을 눈에 띄게 저하시킬 수 있지만,
  영향 범위 또는 악용 가능성은 제한적인 버그입니다. 비기본 설정 등 특정 조건에서만
  트리거되거나, 완전 장애 대신 서비스 품질 저하를 유발하는 경우가 많습니다.
  <details markdown="1">

  <summary>
  예시
  </summary>

  * UPnP 같은 비기본 기능이 활성화된 경우에만 악용 가능한
    로컬 네트워크 기반 잠재적 원격 코드 실행(RCE) 취약점
    ([CVE-2015-20111](/en/2024/07/03/disclose_upnp_rce/)).
  * 변형된 블록(mutated block) 전송으로 블록 전파를 저해해
    노드의 신규 블록 수신을 지연시키는 공격
    ([CVE-2024-52921](/en/2024/10/08/disclose-mutated-blocks-hindering-propagation/)).
  * 공격자가 블록을 광고한 뒤 실제 데이터를 주지 않아,
    피해 노드가 다른 피어에서 받기까지 최대 10분 대기하게 만드는 공격
    ([CVE-2024-52922](/en/2024/11/05/cb-stall-hindering-propagation/)).

  </details>
* **Low**: 악용이 어렵거나 노드 동작에 미치는 영향이 제한적인 버그입니다.
  비기본 설정 또는 로컬 네트워크에서만 트리거되는 경우가 많고, 즉각적이거나
  광범위한 위협으로 이어지지는 않습니다.
  <details markdown="1">

  <summary>
  예시
  </summary>

  * 변조된 `getdata` 메시지로 피어 연결이 무한 루프에 들어가 CPU를 소모할 수 있지만,
    노드의 블록 처리나 다른 피어 처리 능력에는 직접적 영향이 없는 경우
    ([CVE-2024-52920](/en/2024/07/03/disclose-getdata-cpu/)).
  * 의존성 버그로 노드가 크래시할 수 있으나, UPnP 같은 비기본 기능이 활성화된
    경우에만 해당되는 취약점
    ([CVE-2024-52917](/en/2024/07/31/disclose-upnp-oom/)).
  * 노드 크래시를 유발할 수 있지만 악용 난이도가 매우 높은 버그
    ([CVE-2024-52919](/en/2025/04/28/disclose-cve-2024-52919/)).

  </details>

일반 정책상 **Low**는 수정이 포함된 메이저 버전 배포 후 2주 뒤 공개됩니다.
**Medium**/**High**는 마지막 영향 버전이 [EOL](/ko/lifecycle/)이 된 뒤 2주 후 공개됩니다
(대략 수정 포함 메이저 버전의 첫 배포로부터 약 1년 후).

취약점 상세 공개 2주 전에는 사전 공지가 게시되며, 보통 신규 메이저 버전 릴리스와 함께
수정된 취약점 수와 심각도 수준이 안내됩니다.

**Critical** 버그는 표준 정책 범위를 벗어날 수 있으며, 별도 절차가 적용될 수 있습니다.
또한 제보된 이슈가 취약점으로 분류되지 않을 수도 있고, 심각하더라도 엠바고가 필요하지 않을 수 있습니다.

## 과거 보안 권고

{% assign advisories=site.posts | where:"lang", 'en' | where:"type", 'advisory' | sort: "date" | reverse %}
{% for advisory in advisories %}
{% assign post=advisory %}
  <article>
    <h2><a href="{{ post.url }}" title="{{ post.title | xml_escape }}">{{ post.title }}</a></h2>
    <p>{{ post.excerpt | markdownify | strip_html | truncate: 200 }}</p>
  </article>
{% endfor %}
