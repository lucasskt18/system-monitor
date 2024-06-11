 ## System Monitor Tool

- System Monitor Tool é uma ferramenta de monitoramento de desempenho do sistema escrita em Golang. Ela coleta e reporta dados de uso de CPU, memória e disco. Esta ferramenta é útil para obter insights sobre o uso de recursos do sistema em tempo real e pode ser usada para diagnósticos e otimização de desempenho.

## Recursos

- Monitoramento do uso de CPU, memória e disco.
- Atualização das métricas a cada 10 segundos.
- Registro de métricas usando a biblioteca go-metrics.

## Dependências

- Golang (versão 1.18+)
- gopsutil para coleta de dados do sistema.
- go-metrics para gerenciamento de métricas.

## Instalação

- Pré-requisitos
Certifique-se de ter o Golang instalado em sua máquina. 

- Passos de Instalação
- Clone este repositório:

- git clone https://github.com/yourusername/system-monitor.git

- Instale as dependências:
- go mod tidy

## Uso
- Executando a Ferramenta
- Para executar a ferramenta de monitoramento:
- go run main.go

