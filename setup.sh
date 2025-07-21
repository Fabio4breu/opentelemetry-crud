#!/bin/bash

# Cria a pasta otel-traces se não existir
if [ ! -d "otel-traces" ]; then
  echo "Criando pasta otel-traces..."
  mkdir otel-traces
else
  echo "Pasta otel-traces já existe."
fi

# Dá permissão total para leitura, escrita e execução para todos os usuários
echo "Ajustando permissões da pasta otel-traces..."
chmod 777 otel-traces

# Altera o dono da pasta para o usuário atual
chown $(whoami):$(whoami) otel-traces

echo "Setup finalizado. A pasta otel-traces está pronta para uso."
