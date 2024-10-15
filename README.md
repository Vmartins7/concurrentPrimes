### Concurrent Primes

Projeto simples que visa explorar o uso de <i>goroutines</i> dividindo a faixa de divisores uteís de um número primo em n partes
e cada parte desta é testada em um sua própria <i>goroutine</i>, a mesmo função pode ser comparada com uma função não concorrente de teste de primalidade.

## Uso
No diretório do projeto:

```bash
go go run . <um ou mais números naturais> 
```

## Teste
Dentro do diretório [/cprime](./cprime) execute:
```bash
go test
```