# Gerador de chave JWT
Programinha que gera uma chave randômica para ser usada como secret e uma chave pública para ser utilizada com essa pública.

## Uso 
```$ go run main.go <nome do sub> <tempo de expiração (em dias)>```

## Resultado
O programa retornará no stdout um resultado como:

```
===================================== Saída =====================================
Chave secreta (campo JWT_SECRET_KEY do Vertc-Gestão):              ZCDIEsRYs29XY/gd92GAord2suVn2GI/cYuHgYayHEY=
Chave pública (campo JWT do sistema que acessará o Vertc-Gestão):  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjYwNjM2MTM0NTYsInN1YiI6InZlcnQtZGF0YSJ9.Low6I-p4QJ4Q-oOJqpku2Zxqsf6iVoarQmLBh1zVh6w
Nome do sistema:                                                   sistema-externo
Tempo de expiração (em dias):                                      2
=================================================================================
```