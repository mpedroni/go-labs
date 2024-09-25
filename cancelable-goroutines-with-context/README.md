# cancelable-goroutines-with-context

Esse é um exemplo de como cancelar a execução de goroutines utilizando o pacote [context](https://pkg.go.dev/context) da lib padrão do Go. A ideia é que, ao cancelar o context, todas as goroutines relacionadas a ele sejam canceladas, bem como os contexts criados a partir desse mesmo context.

## Criando um context cancelável

Podemos fazer isso de várias formas, dependendo de como o context deve ser cancelado. O método `context.WithCancel` retorna um novo context e uma função `cancel` que pode ser chamada para cancelar o context. O método `context.WithDeadline` retorna um novo context que será cancelado após o tempo limite especificado. O método `context.WithTimeout` retorna um novo context que será cancelado uma vez decorrido o tempo definido.

Vale ressaltar que todos os contexts são canceláveis via função `cancel` e via seu tipo específico, ou seja, todo context pode ser cancelado de duas maneiras: automaticamente, quando o tempo limite é atingido, ou manualmente, chamando a função `cancel`.

```go
// criando um context que será cancelado após 5 segundos
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

Note a declaração `defer cancel()` no final do bloco de código. Isso é importante para garantir que o contexto seja cancelado e que os recursos sejam liberados ao final da execução, além de também ser possível cancelar o context manualmente caso a operação seja concluída antes do tempo limite ou por qualquer outro motivo.

O método `context.Background()` retorna um context vazio, que não é cancelável e não possui valores associados. Normalmente é utilizado para derivar outros contexts a partir dele.

## Reagindo ao cancelamento do context

A maneira o context informa que foi cancelado é através de um channel chamado `Done`. O fechamento desse channel indica que o context foi cancelado. Para reagir ao cancelamento do context, basta escutar esse channel.

```go
select {
case <-ctx.Done():
    releaseResources()
    cleanup()
    // ...
    return
}
```

## O relacionamento hierárquico entre contexts

Durante a execução de uma operação, é comum criarmos vários contexts a partir de um context inicial. Quando criamos um context a partir de um outro context, cria-se uma relação hierárquica entre eles. Isso significa que, ao cancelar um context pai, todos os contexts filhos também serão cancelados.

## Por que encerrar goroutines ao cancelar um context ao invés de deixá-las rodando?

Se as goroutines não forem encerradas ao cancelar o context, elas continuarão rodando indefinidamente, consumindo recursos (CPU, memória, arquivos abertos, conexões de rede...) e possivelmente causando vazamento de goroutines. Além do consumo de recursos, algumas goroutines podem demandar algum tipo de graceful shutdown e, se não for realizado ao cancelar o context, a goroutine possivelmente só irá encerrar no encerramento do programa, resultando em um encerramento abrupto.
