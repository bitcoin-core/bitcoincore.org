---
title: Ciclo de Vida do Software
name: software-life-cycle
id: pt_br-software-life-cycle
permalink: /pt_br/lifecycle/
layout: page
type: pages
lang: pt_br
version: 2
---
{% include toc.html %}

Este documento descreve o ciclo de vida do pacote de software
Bitcoin Core lançado pelo projeto Bitcoin Core. Ele está de acordo com
a política padrão de manutenção dos softwares comerciais.


## Versionamento

Os lançamentos do Bitcoin Core são versionados seguindo: MAJOR.MINOR e
lançamentos candidatos são atribuidos com sufixo rc1, rc2 e etc.

## Lançamentos Major

Nosso objetivo é fazer um grande lançamento a cada 6-7 meses.

Essas serão numeradas 22.0, 23.0 etc.

## Lançamentos de manutenção

Nós forneceremos "lançamentos minor" para consertar bugs que
vieram nos lançamentos major. Por regra geral, nós não introduzimos grandes
mudanças em um lançamento de manutenção (exceto para regras de consenso).
No entanto, nós podemos adicionar pequenas mudanças quando for necessário e
faremos o back-port de alterações de regras de consenso, como soft forks.

Lançamentos minor serão numerados 22.1, 22.2, 23.1, 23.2 etc.

## Regra de consenso

Propostas para mudar regras de consenso são sempre enviadas primeiro para a
versão de manutenção como 22.2, 23.1 e etc. Isso torna mais fácil para usuários
empresariais acessarem e testarem a proposta devido ao seu menor conjunto de
alterações em comparação com um lançamento major.

## Período de manutenção

Nós mantemos a versão major até o seu "Fim de Manutenção". Nós geralmente
mantemos os lançamentos das versões major corrente e a anterior.
Uma vez lançada a versão 24.0, então a 22.0 será considerada como "Fim de
Manutenção". Com o envelhecimento da versão major, problemas tem que ser cada
vez mais críticos para serem reportados para ela e uma quantidade crescente de
problemas severos é requerido  para justificar uma nova versão minor.
Uma vez que o software chegou ao período de "Fim de Manutenção", ela só
receberá manutenção criticas de segurança até a data do Fim do Ciclo de vida
(FCV). Após o FCV, usuários devem atualizar para a última versão para receber
atualizações de segurança, mesmo que a comunidade possa fornecer correções para
problemas criticos baseado nos melhores esforços.
Geralmente é recomendo executar a última versão de manutenção lançada (lançamento
pontual) da versão major atual ou anterior.

Por favor, verifique se a versão minor tem correções, atualizações de tradução
e soft forks. Tradução no [Transifex][bitcoin-transifex-link] só está
disponível para os últimos dois lançamentos major.

Por exemplo, a versão 22.0 foi lançada em 13-09-2021 e nós fornecemos
manutenção de correção (lançamento atual) até 15-11-2022.
Problemas criticos de segurança devem ser consertados até o Fim Ciclo de vida
"FCV" na data de 01-04-2024.
Portanto, para ter as vantagens dos bugs consertados, você deveria atualizar
para a última versão major.


## Agenda

Uma vez que o FCV é atingido, você irá precisar atualizar para a versão mais
nova.

| Versão | Data de Lançamento | Fim manutenção | Fim do Ciclo de vida |
|--------|--------------------|----------------|----------------------|
{% include posts/maintenance-table.md %}

\* _Nosso objetivo é fazer um grande lançamento a cada 6-7 meses._

_TBA: a ser anunciado_

## Versões do protocolo

A descrição abaixo somente descreve os lançamentos do software Bitcoin Core.
Outras partes do projeto do sistema Bitcoin contém suas próprias versões.
Alguns exemplos são:
- Toda **transação** contem um número de versão.
- O **protocolo da rede P2P** usa números de versões para permitir os nós
informarem quais funcionalidades eles suportam.
- **Carteira integrada** Bitcoin Core tem seu próprio número de versionamento
interno.

Estes números de versões são deliberadamente desvinculados do número de versão
do projeto Bitcoin Core, uma vez que o projeto Bitcoin Core ou não tem controle
direto sobre eles (como é o caso dos blocos e transações), ou tenta manter a
compatibilidade com outros projetos (como é o caso com o protocolo de rede), ou
permite a possibilidade que as modificações sejam feitas não em grandes
lançamentos (como acontece, as vezes, na construção da carteira).

O protocolo de consenso, por si só, não tem um número de versão.

## Relacionamento com SemVer

O versionamento do software Bitcoin Core não segue o padrão de versionamento
opcional [SemVer][], más é lançado um versionamento superficialmente similar.
SemVer foi designado para uso em bibliotecas de softwares normais
onde indivíduos podem escolher atualizar a biblioteca no seu próprio local, ou
podem permanecer em uma versão anterior caso não goste das modificações.

Partes do Bitcoin, mais notadamente a regra de consenso, não funciona dessa
maneira.  A fim que uma nova regra de consenso entre em atividade, ela
precisa ser aceita por um número de mineradores, full nodes ou ambos; e uma vez
que ela tenha efeito, softwares que não conhecem aquelas novas regras podem
gerar ou aceitar transações inválidas (embora as atualizações sejam desenhadas
para prevenir que esse tipo de comportamento aconteça).

Por essa razão, o Bitcoin Core desvia da regra de consenso do SemVer e outras
atualizações onde a adoção em toda a rede é necessária ou desejável. Bitcoin
Core lança essas mudanças como lançamentos de manutenção (`x.y`) ao invés de
um lançamento major (`x.0`); isso minimiza o tamanho da modificação para tornar
mais fácil possível para a maioria das pessoas inspecionar, testar e implanta-
la.  Isso também permite enviar a modificação para vários outros lançamentos
majors, crescendo ainda mais o número de usuários que podem atualizar
facilmente, embora não haja voluntários suficientes para gerenciar isso.

[SemVer]: https://semver.org/
[bitcoin-transifex-link]: https://www.transifex.com/bitcoin/bitcoin/
