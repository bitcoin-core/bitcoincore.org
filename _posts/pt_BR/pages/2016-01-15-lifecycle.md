---
title: Ciclo de vida do software
name: software-life-cycle
id: pt_BR-software-life-cycle
permalink: /pt_BR/lifecycle/
layout: page
type: pages
lang: pt_BR
version: 2
---
{% include toc.html %}

Este documento descreve o ciclo de vida do pacote de software Bitcoin Core lançado pelo projeto Bitcoin Core. Ele está alinhado com a política de manutenção padrão para softwares comerciais.

## Versionamento

Os lançamentos do Bitcoin Core são versionados da seguinte forma: MAJOR.MINOR, e os release candidates recebem os sufixos rc1, rc2 etc.

Nosso objetivo é fazer um lançamento principal (major) a cada 6 meses. Eles serão numerados como 29.0, 30.0 etc.

Forneceremos lançamentos menores (minor/"de manutenção") que corrigem bugs (de segurança ou não) para cada lançamento principal. Eles serão numerados como 29.3, 30.1 etc. Não introduziremos novos recursos importantes em lançamentos de manutenção (exceto mudanças nas regras de consenso, veja abaixo).

## Regras de consenso

As propostas de mudança nas regras de consenso são sempre lançadas primeiro em versões de manutenção, como 22.2, 23.1 etc. Isso facilita a avaliação e o teste da proposta por usuários corporativos, devido ao seu conjunto de mudanças menor em comparação com um lançamento principal. Também permite que usuários que seguem um caminho de atualização mais conservador adotem as mudanças nas regras de consenso em tempo mais adequado.

## Período de manutenção

Sempre mantemos as três últimas versões principais. Quando uma nova versão principal é lançada, a mais antiga sai da janela de manutenção e se torna "End of Life". Por exemplo, se o último lançamento principal for o 30.0, então as versões 29.x e 28.x também são consideradas mantidas. Quando o 31.0 for lançado, a 28.x se torna "End of Life". O critério para fazer backport de uma mudança para uma versão principal mais antiga se torna mais rigoroso à medida que ela envelhece.

As versões principais que estão em "End of Life" geralmente não recebem correções de segurança. Para saber mais sobre nossa política de correções de segurança, consulte nossa página de [avisos de segurança][security advisories]. Recomendamos utilizar a versão de manutenção mais recente da versão principal mais nova para a qual você conseguir atualizar.

## Cronograma

Quando o EOL é atingido, você precisará atualizar para uma versão mais nova.

| Versão | Data de lançamento | End of Life |
|---------|--------------|-------------|
{% include posts/maintenance-table.md %}

\* _Nosso objetivo é fazer um lançamento principal a cada 6-7 meses_

_TBA: a ser anunciado_

## Versionamento do protocolo

A descrição acima trata apenas dos lançamentos do software Bitcoin Core. Muitas outras partes do sistema Bitcoin contêm suas próprias versões. Alguns exemplos:

- Cada **transação** contém um número de versão.
- O **protocolo da rede P2P** usa números de versão para permitir que os nós anunciem quais recursos eles suportam.
- A **carteira integrada** do Bitcoin Core tem seu próprio número de versão interno.

Esses números de versão são deliberadamente desacoplados do número de versão do Bitcoin Core, pois o projeto Bitcoin Core não tem controle direto sobre eles (como é o caso dos blocos e transações), ou tenta manter compatibilidade com outros projetos (como é o caso do protocolo de rede), ou admite a possibilidade de que nenhuma mudança importante seja feita em alguns lançamentos (como às vezes é o caso da carteira integrada).

O protocolo de consenso em si não tem número de versão.

## Relação com o SemVer

O versionamento do software Bitcoin Core não segue o padrão opcional de versionamento [SemVer][], mas o versionamento dos seus lançamentos é superficialmente semelhante. O SemVer foi projetado para uso em bibliotecas de software comuns, nas quais cada pessoa pode escolher atualizar a biblioteca no seu próprio ritmo, ou até permanecer em uma versão mais antiga se não gostar das mudanças.

Partes do Bitcoin, principalmente as regras de consenso, não funcionam assim. Para que uma nova regra de consenso entre em vigor, ela precisa ser aplicada por um certo número de mineradores, nós completos, ou ambos; e, uma vez em vigor, softwares que não conhecem a nova regra podem gerar ou aceitar transações inválidas (embora as atualizações sejam projetadas para evitar que isso aconteça sempre que possível).

Por esse motivo, o Bitcoin Core não segue o SemVer em mudanças nas regras de consenso e em outras atualizações nas quais a adoção em toda a rede é necessária ou desejável. O Bitcoin Core lança essas mudanças como lançamentos de manutenção (`x.y`) em vez de lançamentos principais (`x.0`); isso minimiza o tamanho do patch para facilitar que o maior número possível de pessoas o inspecione, teste e implante. Também torna possível fazer backport do mesmo patch para várias versões principais anteriores, aumentando ainda mais o número de usuários que podem atualizar facilmente, embora nem sempre haja voluntários suficientes para gerenciar isso.

[SemVer]: https://semver.org/
[bitcoin-transifex-link]: https://explore.transifex.com/bitcoin/bitcoin/
[security advisories]: /en/security-advisories
