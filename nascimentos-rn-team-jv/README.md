# Trabalho 1: Estatísticas de Nascidos Vivos no Rio Grande do Norte

<sub>Última atualização: 26/04/2023</sub>

## Sumário

- [Visão geral e objetivos](#visão-geral-e-objetivos)  
- [Tarefas](#tarefas)
  - [Tarefa 1](#tarefa-1)
  - [Tarefa 2](#tarefa-2)
  - [Tarefa 3](#tarefa-3)
- [Requisitos](#requisitos)
- [Orientações gerais](#orientações-gerais)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Visão geral e objetivos

O objetivo deste trabalho é colocar em prática as operações de leitura e escrita em arquivos na linguagem de programação Go. O programa a ser implementado utiliza como fonte de dados um arquivo [CSV (*Comma-Separated Values*)](https://en.wikipedia.org/wiki/Comma-separated_values) contendo o número de nascidos vivos em cada um dos 167 municípios do Estado do Rio Grande do Norte entre os  anos de 1994 e 2020. Os dados são provenientes do Sistema de Informações sobre Nascidos Vivos (SINASC) do Departamento de Informática do Sistema Único de Saúde (DATASUS), vinculado ao Ministério da Saúde do Brasil, e estão disponíveis neste [link](http://tabnet.datasus.gov.br/cgi/deftohtm.exe?sinasc/cnv/nvbr.def).

Este trabalho explora os seguintes elementos da programação em Go, cujos conhecimentos são, portanto, ora necessários:

- Entrada e saída formatada via console
- Entrada e saída via arquivos
- Tipos estruturados (*structs*)
- Funções
- *Arrays* e *slices*
- Mapas

## Tarefas

### Tarefa 1

A primeira tarefa deste trabalho consiste em implementar um programa chamado `nascimentos` que recebe como entrada, via linha de comando, o arquivo de texto [`Nascimentos_RN.csv`](Nascimentos_RN.csv) disponível neste repositório. Esse arquivo, no formato CSV, contém os números de nascidos vivos em cada município para cada ano contabilizado. Cada linha do arquivo refere-se a um município e os números de nascimentos em cada ano são separados por vírgulas. Enquanto o arquivo de entrada for sendo lido, os dados de cada município deverão ser armazenados em memória. Para tal armazenamento, pode-se fazer uso de tipos estruturados (*structs*) e/ou de estruturas de dados como *arrays*, *slices* ou mapas.

Ao abrir o arquivo CSV de entrada em um editor de texto ou editor de planilhas, será possível observar que a primeira linha do arquivo diz respeito a um cabeçalho para a tabela de dados e que a primeira coluna contém o nome do município antecedido por um código numérico que o identifica, de forma aglutinada. Durante a leitura do arquivo, o programa deverá separar o código numérico do nome do município. A última linha e a última coluna do arquivo CSV de entrada apresentam, respectivamente, o total de nascimentos em cada ano e o total de nascimentos em cada município somados os 27 anos da série 1994-2020. No entanto, esses totais deverão ser usados **apenas** para fins de verificação de consistência. Para isso, deverá ser implementada uma função de verificação para verificar se os totais informados no arquivo conferem com a série de dados lida.

Após a leitura do arquivo, o programa deverá computar as seguintes estatísticas que resumem o conjunto de dados em questão na série histórica:

1. o *maior* número de nascimentos em cada ano;
2. o *menor* número de nascimentos em cada ano;
3. a *média* do número de nascimentos em cada ano;
4. o *desvio padrão* do número de nascimentos em cada ano, e;
5. o número *total* de nascimentos em cada ano.

Para calcular o desvio padrão $\sigma$ do número de nascimentos em um determinado ano, pode-se utilizar a seguinte equação:

$$ \sigma = \sqrt{\frac{1}{M} \sum_{i=1}^{M} \(n_i - MD\)^2} $$

em que $M$ é o número de municípios, $n_i$ é o número de nascimentos do $i$-ésimo município no ano em questão e $MD$ é a média do número de nascimentos nesse ano. Um baixo desvio padrão indica que os pontos dos dados tendem a estar próximos da médiado, enquanto que um alto desvio padrão indica que os pontos dos dados estão espalhados por uma ampla gama de valores.

Como saída, o programa deverá gerar automaticamente dois arquivos:

1. um arquivo de texto no formato CSV chamado `estatisticas.csv`, no qual cada linha corresponde a um ano e suas respectivas estatísticas acerca do número de nascimentos, cada valor sendo separado por vírgulas, e;
2. um arquivo de texto chamado `totais.dat` contendo **apenas** o ano e o respectivo número total de nascimentos nesse ano, separados por **espaço**.

Uma vez gerado o arquivo `totais.dat` como saída da execução do programa implementado, deverá ser gerado automaticamente um [**histograma**](https://en.wikipedia.org/wiki/Histogram) que mostrará a evolução do número de nascimentos entre os anos de 1994 e 2020. Para gerar esse gráfico, pode-se utilizar o programa [`gnuplot`](http://www.gnuplot.info/) ou outro programa ou biblioteca que ofereça a mesma funcionalidade, seja em Go ou em outra linguagem de programação. Contudo, a fim de minimizar uma maior curva de aprendizado, sugere-se o uso do `gnuplot`, com o qual o histograma pode ser gerado de forma simples executando o seguinte comando no terminal do sistema operacional:

```bash
$ gnuplot -e  "filename='totais.dat'" histograma.gnuplot
```

sendo `totais.dat` é o arquivo anteriormente gerado e que contém os dados a serem plotados no histograma e o arquivo [`histograma.gnuplot`](histograma.gnuplot) é um *script* de configuração para instrução do quando da geração do gráfico. Esse *script*, disponível neste repositório, possui o seguinte conteúdo:

```
# Inicialização
clear
reset
set key off

# Configurações de saida: inclui formato de exportação, tamanho do gráfico, 
# fontes utilizadas e nome do arquivo de saída

# Exportação para o formato PNG
set terminal png size 640,480 enhanced font 'Helvetica,12'
set output 'histograma.png'

# Exportação para o formato JPG
# set terminal jpeg size 640,480 enhanced font 'Helvetica' 12
# set output 'histograma.jpg'

# Exportação para o formato SVG
# set terminal svg size 640,480 enhanced background rgb 'white' fname 'Helvetica' fsize 14 butt solid
# set output 'histograma.svg'

# Título do gráfico
set title 'Total de nascidos vivos no RN (1994-2020)'

# Configurações do eixo horizontal
set xrange[1994:2020]           # Faixa de valores
set xtics 1                     # Salto entre valores
set xtic rotate by -45 scale 0  # Rotação dos rótulos

# Configurações do eixo vertical
set yrange[0:80000]             # Faixa de valores

# Seleção do tipo de gráfico a ser gerado (histograma)
set style data histogram
set style histogram clustered gap 1
set style fill solid border -1      # Preenchimento e contorno
set linetype 1 lc rgb 'green'       # Cor
set boxwidth 0.6                    # Largura das barras verticais

# Plotagem do gráfico
# Os dados a serem plotados constam no arquivo totais.dat
plot 'totais.dat' using 1:2 title '' smooth freq with boxes
```

Por padrão, esse *script* gerará um gráfico na forma de um arquivo de imagem no formato PNG, a saber, `histograma.png`. Para gerar em outros formatos de arquivos de imagem, o caractere `#` (o qual indica um comentário) deve ser removido das respectivas linhas responsáveis pela geração no formato em questão, iniciadas pelos comandos `set terminal` e `set output`. É importante destacar que **não é possível** gerar mais de um arquivo de imagem em uma única plotagem (comando `plot`), ou seja, é necessário repetir o comando de plotagem de dados para cada arquivo de imagem a ser gerado, caso deseje-se gerar saída em múltiplos formatos.

## Tarefa 2

A segunda tarefa deste trabalho consiste em estender o programa implementado até a realização da Tarefa 1 para gerar automaticamente um [**gráfico de linha**](https://en.wikipedia.org/wiki/Line_chart) mostrando a evolução do número de nascimentos para um determinado conjunto de municípios (em número indefinido) em toda a série histórica. Tais municípios alvo deverão ser informados de forma configurável ao programa através de um arquivo de entrada chamado `alvos.dat` que deverá conter **apenas** uma lista de códigos dos municípios a serem considerados para o gráfico de linha a ser gerado, um por linha. Um exemplo de conteúdo para o arquivo `alvos.dat` considerando a cidade do Natal e seus municípios limítrofes seria:

```
240810
240360
240325
240710
241200
```

tais códigos referindo-se respectivamente aos municípios de Natal, Extremoz, Parnamirim, Macaíba e São Gonçalo do Amarante. Um exemplo de execução do programa seria, portanto:

```bash
$ nascimentos Nascimentos_RN.csv
> Arquivo estatisticas.csv gerado
> Arquivo totais.dat gerado

> Lendo arquivo alvos.dat
Municípios definidos como alvo (5):
NATAL
EXTREMOZ
PARNAMIRIM
MACAIBA
SAO GONCALO DO AMARANTE

> Arquivo nascimentos-alvos.dat gerado
```

A geração automática do gráfico de linha fazendo uso do `gnuplot` ou de outro programa ou biblioteca deverá seguir a mesma ideia desenvolvida na Tarefa 1 em termos de utilizar um arquivo de entrada (no caso o arquivo `nascimentos-alvos.dat`) contendo os dados a serem plotados no gráfico. No caso de ser utilizado o `gnuplot`, para possibilitar a geração do gráfico, deverá ser criado um *script* de configuração próprio, podendo inclusive adaptar o arquivo [`histograma.gnuplot`](histograma.gnuplot) apresentado anteriormente. Além da [documentação oficial](http://gnuplot.info/documentation.html) do `gnuplot`, bons exemplos de uso e demonstração podem ser encontrados neste [link](http://alvinalexander.com/technology/gnuplot-charts-graphs-examples) e neste outro [link](http://gnuplot.sourceforge.net/demo/).

## Tarefa 3

A terceira e última tarefa deste trabalho consiste em estender o programa implementado até a realização da Tarefa 2 para que ele seja capaz de fornecer, na saída padrão, quais municípios apresentaram a maior e a menor *taxa de crescimento* no número de nascimentos nos últimos cinco anos da série histórica. A taxa de crescimento relativa $TC$ de um município $M$ pode ser calculada através da seguinte equação:

$$ TC(M) = \frac{N_{2020}(M)}{N_{2016}(M)} $$

em que $N_{2020}$ e $N_{2016}$ são respectivamente os números de nascidos vivos nos anos de 2020 e 2016. Para o caso em que o valor de $N_{2016}$ ou de $N_{2020}$ esteja faltante (o que implicaria em uma operação de divisão inválida), o município deverá ser desconsiderado em favor do próximo que apresentar número positivo de nascimentos tanto em 2016 quanto em 2020. Os municípios com menor e maior valores de taxa de crescimento deverão ser apresentados na saída padrão conforme o seguinte exemplo:

```bash
$ nascimentos Nascimentos_RN.csv
> Arquivo estatisticas.csv gerado
> Arquivo totais.dat gerado

> Lendo arquivo alvos.dat
Municípios definidos como alvo (5):
NATAL
EXTREMOZ
PARNAMIRIM
MACAIBA
SAO GONCALO DO AMARANTE

> Arquivo nascimentos-alvos.dat gerado

Município com maior taxa de crescimento 2016-2020: SAO MIGUEL DO GOSTOSO (1200.00%)
Município com maior taxa de queda 2016-2020: ALMINO AFONSO (-99.43%)
```

A porcentagem de crescimento (ou queda) $P$ apresentada na saída do programa, a qual deverá considerar estritamente duas casas decimais, pode ser calculada através da seguinte equação:

$$ P = 100 \times (TC - 1) $$

Note-se que, caso o percentual seja negativo, deve-se apresentar a taxa de crescimento como uma **taxa de queda** na saída padrão. 

## Requisitos

A implementação deste trabalho requer os seguintes elementos instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

Adicionalmente, deverão ser instalados no ambiente de desenvolvimento o(s) programa(s) e/ou biblioteca(s) escolhidos para dar suporte à geração automática de gráficos.

## Orientações gerais

Boas práticas de programação deverão ser constantemente aplicadas no desenvolvimento do programa. Assim, o programa deverá ser codificado de forma legível (com indentação de código fonte, nomes consistentes, etc.) e documentado adequadamente na forma de comentários.

O programa deverá ainda ser desenvolvido com qualidade, garantindo que o ele funcione de forma correta e eficiente. Deve-se também pensar nas possíveis entradas que poderão ser utilizadas para testar apropriadamente o programa, além de serem tratadas adequadamente possíveis entradas consideradas inválidas.

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre os integrantes da equipe. Após a implementação das soluções para os problemas propostos, o arquivo [`author.md`](author.md) presente no repositório deverá ser editado preenchendo as informações de identificação dos integrantes da equipe, na seção [Informações de Autoria](author.md#identificação-de-autoria).

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. **Trabalhos copiados no todo ou em parte de outros colegas ou da Internet serão anulados e receberão nota zero.**

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://git-scm.com) serão utilizados para possibilitar a entrega da implementação realizadas. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (*assignment*) no GitHub Classroom. Cada integrante de equipe deverá acessar este [*link*](https://classroom.github.com/a/NROUxqsa), aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelos integrantes da equipe e pelo docente, sob a organização [`ufrn-golang`](https://github.com/ufrn-golang). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo do arquivo executável gerado a partir da compilação do código fonte.

A implementação do programa objeto deste trabalho deverá ser realizada **até as 23:59 do dia 7 de maio de 2023** no respectivo repositório Git da equipe. Para fins de registro, o endereço do repositório também deverá ser **obrigatoriamente** enviado através da opção *Tarefas* na Turma Virtual do SIGAA, devendo **um único membro da equipe** realizar esse envio. Além disso, **não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.**

## Avaliação

A avaliação do programa implementado contabilizará nota de até 10,0 pontos. O programa implementado será avaliado de acordo com os seguintes critérios:

- utilização correta dos recursos providos pela linguagem de programação Go;
- corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação;
- aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte, e;
- correta utilização do repositório Git, no qual deverá ser registrado todo o histórico da implementação por meio de *commits*.

O não cumprimento de algum dos critérios de avaliação especificados poderá resultar nos seguintes decréscimos, aplicados sobre a nota obtida até então na avaliação:

| Falta | Decréscimo |
| :--- | ---: |
| Falta de comentários no código fonte | -10% |
| Uso inadequado de controle de versão com Git | -20% |
| Falta de especificação ou especificação incorreta da autoria do trabalho (arquivo [`author.md`](author.md)) | -20% |
| Código fonte com legibilidade prejudicada (por exemplo, com identação ou nomenclatura inadequada) | -30% |
| Implementação realizada em desacordo com as especificações postas para o trabalho | -50% |
| Programa apresenta erros de compilação, não executa ou apresenta saída incorreta | -70% |
| Plagiarismo | -100% |

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar *e-mail* para o endereço everton.cavalcante@ufrn.br;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.
