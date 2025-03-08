# go-datastructures

## Introdução

```bash
pkg/
└── graph/
├── graph.go (Estrutura básica do grafo)
├── adjacency_list.go (Representação com Lista de Adjacência)
├── adjacency_matrix.go (Representação com Matriz de Adjacência)
├── edge_list.go (Representação com Lista de Arestas)
├── algorithms/
│ ├── bfs.go (Busca em Largura - BFS)
│ ├── dfs.go (Busca em Profundidade - DFS)
│ ├── dijkstra.go (Caminho mínimo - Dijkstra)
│ ├── bellman_ford.go (Caminho mínimo - Bellman-Ford)
│ ├── floyd_warshall.go (Caminho mínimo para todos os pares)
│ ├── prim.go (Árvore Geradora Mínima - Prim)
│ ├── kruskal.go (Árvore Geradora Mínima - Kruskal)
│ ├── topological.go (Ordenação Topológica)
│ ├── strongly_connected.go (Componentes Fortemente Conectados)
│ ├── bridges.go (Detecção de Pontes)
│ ├── articulation_points.go (Pontos de Articulação)
│ ├── shortest_path.go (Caminho Mínimo)
│ ├── max_flow.go (Fluxo Máximo)
│ ├── min_cut.go (Corte Mínimo)
│ ├── traveling_salesman.go (Caixeiro Viajante)
│ ├── hamiltonian_cycle.go (Ciclo Hamiltoniano)
│ ├── eulerian_cycle.go (Ciclo Euleriano)
│ ├── is_bipartite.go (Bipartição)
│ ├── is_connected.go (Conectividade)
│ ├── is_cyclic.go (Ciclicidade)
│ ├── is_tree.go (Árvore)
│ ├── is_forest.go (Floresta)
│ ├── is_biconnected.go (Biconexidade)
│ ├── is_planar.go (Planaridade)
│ ├── is_eulerian.go (Eulerianidade)
│ ├── is_semi_eulerian.go (Semi-Eulerianidade)
│ ├── a-star.go (A*)
│ ├── ida_star.go (IDA*)
```

## Estrutura de Dados

### Grafo

Um grafo é uma estrutura de dados que consiste em um conjunto de vértices (ou nós) e um conjunto de arestas (ou arcos) que conectam pares de vértices.

Algoritmo A\* é um algoritmo de busca informada que utiliza uma heurística para encontrar o caminho mais curto entre dois vértices em um grafo.

```plaintext
ALGORITMO BuscaA*
INÍCIO
    Criar uma fila de prioridade aberta
    Adicionar o nó inicial à fila com custo 0

    ENQUANTO a fila não estiver vazia FAÇA
        Remover o nó com menor f(n) da fila
        SE nó for o objetivo ENTÃO
            Retornar o caminho encontrado
        FIM SE

        PARA cada vizinho do nó ATUAL FAÇA
            Calcular custo g(n) + h(n)
            SE vizinho ainda não foi visitado OU novo custo for menor ENTÃO
                Atualizar custo e adicionar à fila
            FIM SE
        FIM PARA
    FIM ENQUANTO
FIM

```
